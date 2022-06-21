package scheduler

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/on-chain/goBind"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
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

	carFiles, err := models.GetCarFiles2Refund()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, carFile := range carFiles {
		err = refundCarFile(ethClient, carFile, swanPaymentTransactor)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func refundCarFile(ethClient *ethclient.Client, carFile2Refund *models.CarFile2Refund, swanPaymentTransactor *goBind.SwanPaymentTransactor) error {
	sourceFileUploads, err := models.GetSourceFileUploadOutsByCarFileId(carFile2Refund.CarFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var srcFileUploadWCids []string
	for _, sourceFileUpload := range sourceFileUploads {
		if sourceFileUpload.Status == constants.SOURCE_FILE_UPLOAD_STATUS_DEAL_COMPLETED {
			continue
		}

		wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		if lockedPayment == nil {
			logs.GetLogger().Error("payment not exists for w_cid:", wCid)
			sourceFileUpload.LockedFeeBeforeRefund = decimal.NewFromInt(0)
		} else {
			sourceFileUpload.LockedFeeBeforeRefund = lockedPayment.LockedFee
		}

		srcFileUploadWCids = append(srcFileUploadWCids, wCid)

		err = models.UpdateTransactionUnlockInfo(sourceFileUpload.Id, sourceFileUpload.LockedFeeBeforeRefund, carFile2Refund.LastUnlockAt)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}
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

	wCids := strings.Join(srcFileUploadWCids, ",")
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

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed, wCids:%s, tx hash:%s, tx receipt status:%d", wCids, txHash, txReceipt.Status)
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
