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

func RefundCarFiles() error {
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

	carFiles2Refund, err := models.GetCarFiles2Refund()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, carFile2Refund := range carFiles2Refund {
		err = refundCarFile(ethClient, carFile2Refund, swanPaymentTransactor)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func refundCarFile(ethClient *ethclient.Client, carFile2Refund *models.CarFile2Refund, swanPaymentTransactor *goBind.SwanPaymentTransactor) error {
	sourceFileUploads, err := models.GetSourceFileUploads2RefundByCarFileId(carFile2Refund.CarFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var srcFileUploadWCids []string
	for i, sourceFileUpload := range sourceFileUploads {
		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		if lockedPayment == nil {
			logs.GetLogger().Info("payment not exists for wCid:", wCid)
			continue
		}

		err = models.UpdateTransactionUnlockInfo(sourceFileUpload.Id, lockedPayment.LockedFee, carFile2Refund.LastUnlockAt)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		srcFileUploadWCids = append(srcFileUploadWCids, wCid)

		if len(srcFileUploadWCids) >= constants.MAX_WCID_COUNT_IN_TRANSACTION || i >= len(sourceFileUploads)-1 {
			err = refundSourceFiles(ethClient, srcFileUploadWCids, swanPaymentTransactor)
			if err != nil {
				logs.GetLogger().Error(err.Error())
				return err
			}

			srcFileUploadWCids = []string{}
		}
	}

	return nil
}

func refundSourceFiles(ethClient *ethclient.Client, wCids []string, swanPaymentTransactor *goBind.SwanPaymentTransactor) error {
	tansactOpts, err := client.GetTransactOpts(ethClient, adminWalletPrivateKey, *adminWalletPublicKey)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tx, err := swanPaymentTransactor.Refund(tansactOpts, wCids)
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
		err := fmt.Errorf("unlock failed, wCids:%s, tx hash:%s, error:%s", wCids, txHash, err.Error())
		logs.GetLogger().Error()
		return err
	}

	if txReceipt == nil {
		err := fmt.Errorf("unlock failed, wCids:%s, tx hash:%s,", wCids, txHash)
		logs.GetLogger().Error(err.Error())
		return err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed, wCids:%s, tx hash:%s, tx receipt status:%d", wCids, txHash, txReceipt.Status)
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
