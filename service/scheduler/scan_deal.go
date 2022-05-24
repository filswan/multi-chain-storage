package scheduler

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-lib/client/lotus"
)

func ScanDeal() error {
	err := updateOfflineDealStatusAndLog()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	err = UpdateSourceFile2Refundable()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	return nil
}

func updateOfflineDealStatusAndLog() error {
	offlineDeals, err := models.GetOfflineDeals2BeScanned()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ClientApiUrl, config.GetConfig().Lotus.ClientAccessToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, offlineDeal := range offlineDeals {
		dealInfo, err := lotusClient.LotusClientGetDealInfo(offlineDeal.DealCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		if offlineDeal.OnChainStatus == nil || *offlineDeal.OnChainStatus != dealInfo.Status || offlineDeal.DealId == nil || *offlineDeal.DealId != dealInfo.DealId {
			offlineDeal.OnChainStatus = &dealInfo.Status
			if dealInfo.DealId != 0 {
				offlineDeal.DealId = &dealInfo.DealId
			}

			if dealInfo.Status == constants.ON_CHAIN_DEAL_STATUS_ERROR {
				offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_FAILED
				updateSourceFile2RefundableAfterDealFailed(offlineDeal.CarFileId)
			} else if dealInfo.Status == constants.ON_CHAIN_DEAL_STATUS_ACTIVE {
				offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_ACTIVE
			}

			offlineDeal.UpdateAt = libutils.GetCurrentUtcSecond()
			err = database.SaveOne(offlineDeal)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}

		err = models.CreateOfflineDealLog(offlineDeal.Id, dealInfo.Status, dealInfo.Message)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func updateSourceFile2RefundableAfterDealFailed(carFileId int64) error {
	offlineDeals, err := models.GetOfflineDealNotFailedByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(offlineDeals) > 0 {
		return nil
	}

	sourceFileUploads, err := models.GetSourceFileUploadsByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, sourceFileUpload := range sourceFileUploads {
		sourceFileUpload.Status = constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE
		sourceFileUpload.UpdateAt = libutils.GetCurrentUtcSecond()
		err = database.SaveOne(sourceFileUpload)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func UpdateSourceFile2Refundable() error {
	sourceFileUploads, err := models.GetSourceFileUploadsExpired()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	setSourceFiles2Refundable(sourceFileUploads)

	sourceFileUploads, err = models.GetSourceFileUploads2BeRefundedByCarFileStatus(constants.CAR_FILE_STATUS_DEAL_SEND_EXPIRED)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	setSourceFiles2Refundable(sourceFileUploads)

	sourceFileUploads, err = models.GetSourceFileUploads2BeRefundedByCarFileStatus(constants.CAR_FILE_STATUS_DEAL_SENT_FAILED)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	setSourceFiles2Refundable(sourceFileUploads)

	return nil
}

func setSourceFiles2Refundable(sourceFileUploads []*models.SourceFileUploadOut) {
	for _, sourceFileUpload := range sourceFileUploads {
		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		sourceFileUploadStatus := constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED
		if lockedPayment != nil {
			sourceFileUploadStatus = constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE
		} else {
			logs.GetLogger().Info("payment not exists for ", wCid, ", source file upload id:", sourceFileUpload.Id)
		}

		err = models.UpdateSourceFileUploadStatus(sourceFileUpload.Id, sourceFileUploadStatus)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}
}
