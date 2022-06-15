package scheduler

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/on-chain/goBind"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-lib/client/lotus"
)

func ScanDeal() error {
	err := updateOfflineDealStatusAndLog()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	err = UpdateSourceFiles2Refundable()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	err = refundCarFiles()
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

			switch dealInfo.Status {
			case constants.ON_CHAIN_DEAL_STATUS_ERROR:
				if !strings.Contains(dealInfo.Message, constants.ON_CHAIN_MESSAGE_NOT_COMPLETED) {
					offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_FAILED
				}
			case constants.ON_CHAIN_DEAL_STATUS_ACTIVE:
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

func UpdateSourceFiles2Refundable() error {
	sourceFileUploads, err := models.GetSourceFileUploads2BeRefundedAfterExpired()
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

func refundCarFiles() error {
	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	swanPaymentTransactor, err := client.GetSwanPaymentTransactor(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	carFiles, err := models.GetCarFilesByStatus(constants.CAR_FILE_STATUS_DEAL_SENT)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, carFile := range carFiles {
		err = refundCarFile(ethClient, carFile.ID, swanPaymentTransactor)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func refundCarFile(ethClient *ethclient.Client, carFileId int64, swanPaymentTransactor *goBind.SwanPaymentTransactor) error {
	offlineDeals, err := models.GetOfflineDealsByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var offlineDealsCntFailed int = 0
	var offlineDealsCntSuccess int = 0
	var offlineDealsCntOther int = 0

	for _, offlineDeal := range offlineDeals {
		switch offlineDeal.Status {
		case constants.OFFLINE_DEAL_STATUS_FAILED:
			offlineDealsCntFailed = offlineDealsCntFailed + 1
		case constants.OFFLINE_DEAL_STATUS_SUCCESS:
			offlineDealsCntSuccess = offlineDealsCntSuccess + 1
		default:
			offlineDealsCntOther = offlineDealsCntOther + 1
		}
	}

	if offlineDealsCntOther > 0 {
		msg := fmt.Sprintf("%d deals to be unlocked, cannot refund for car file:%d", offlineDealsCntOther, carFileId)
		logs.GetLogger().Info(msg)
		return nil
	}

	sourceFileUploads, err := models.GetSourceFileUploadOutsByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	if offlineDealsCntFailed > 0 {
		for _, sourceFileUpload := range sourceFileUploads {
			err = models.UpdateSourceFileUploadStatus(sourceFileUpload.Id, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE)
			if err != nil {
				logs.GetLogger().Error(err.Error())
				return err
			}
		}

		err = models.UpdateCarFileStatus(carFileId, constants.CAR_FILE_STATUS_COMPLETED)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		return nil
	}

	var srcFileUploadWCids []string
	for _, sourceFileUpload := range sourceFileUploads {
		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			continue
		}

		if lockedPayment == nil {
			logs.GetLogger().Error("payment not exists for w_cid:", wCid)
			continue
		}

		sourceFileUpload.LockedFeeBeforeRefund = lockedPayment.LockedFee

		srcFileUploadWCids = append(srcFileUploadWCids, wCid)
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, adminWalletPrivateKey, *adminWalletPublicKey)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tx, err := swanPaymentTransactor.Refund(tansactOpts, srcFileUploadWCids)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}
	txHash := ""
	if tx != nil {
		txHash = tx.Hash().Hex()
	}

	txReceipt, err := client.CheckTx(ethClient, tx.Hash())
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed! txHash=%s, status:%d", tx.Hash().Hex(), txReceipt.Status)
		logs.GetLogger().Error(err.Error())
		return err
	}

	for _, sourceFileUpload := range sourceFileUploads {
		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		isPaymentAvailable, err := client.IsLockedPaymentExists(wCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			continue
		}

		if *isPaymentAvailable {
			err := fmt.Errorf("source file upload id:%d,lock payment is still available for w_cid:%s after refund success with tx hash:%s", sourceFileUpload.Id, wCid, txHash)
			logs.GetLogger().Error(err)
			continue
		}

		err = models.UpdateTransactionRefundAfterUnlock(sourceFileUpload.Id, txHash, sourceFileUpload.LockedFeeBeforeRefund)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			continue
		}
	}

	err = models.UpdateCarFileStatus(carFileId, constants.CAR_FILE_STATUS_COMPLETED)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
