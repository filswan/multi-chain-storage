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

	err = updateExpiredSourceFileUploadStatus()
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

func updateExpiredSourceFileUploadStatus() error {
	sourceFileUploads, err := models.GetSourceFileUploadsExpired()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, sourceFileUpload := range sourceFileUploads {
		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		isLockedPaymentExists, err := client.IsLockedPaymentExists(wCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		sourceFileUploadStatus := constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED
		if *isLockedPaymentExists {
			sourceFileUploadStatus = constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE
		}

		err = models.UpdateSourceFileUploadStatus(sourceFileUpload.Id, sourceFileUploadStatus)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}

	return nil
}
