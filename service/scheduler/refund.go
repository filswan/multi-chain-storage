package scheduler

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/on-chain/goBind"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
)

func Refund() error {
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
		err = refund(ethClient, carFile.ID, swanPaymentTransactor)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func refund(ethClient *ethclient.Client, carFileId int64, swanPaymentTransactor *goBind.SwanPaymentTransactor) error {
	offlineDealsNotUnlocked, err := models.GetOfflineDeals2BeUnlockedByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	if len(offlineDealsNotUnlocked) > 0 {
		msg := fmt.Sprintf("%d deals to be unlocked, cannot refund for the car file", len(offlineDealsNotUnlocked))
		logs.GetLogger().Info(msg)
		return nil
	}

	sourceFileUploads, err := models.GetSourceFileUploadsByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var srcFileUploadWCids []string
	for _, sourceFileUpload := range sourceFileUploads {
		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
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

	txReceipt, err := client.CheckTx(ethClient, tx)
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

	err = models.UpdateCarFileStatus(carFileId, constants.CAR_FILE_STATUS_UNLOCKED)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
