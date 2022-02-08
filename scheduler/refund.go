package scheduler

import (
	"fmt"
	"payment-bridge/common/constants"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

func Refund() error {
	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	_, swanPaymentTransactor, err := client.GetSwanPaymentTransactor(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	refund(int64(0), swanPaymentTransactor, tansactOpts)

	dealFiles, err := models.GetDealFiles2Refund()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, dealFile := range dealFiles {
		err = refund(dealFile.ID, swanPaymentTransactor, tansactOpts)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}

func refund(dealFileId int64, swanPaymentTransactor *goBind.SwanPaymentTransactor, tansactOpts *bind.TransactOpts) error {
	offlineDealsNotUnlocked, err := models.GetOfflineDealsNotUnlockedByDealFileId(dealFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	if len(offlineDealsNotUnlocked) > 0 {
		msg := fmt.Sprintf("%d deals not unlocked, cannot refund for the deal file", len(offlineDealsNotUnlocked))
		logs.GetLogger().Info(msg)
		return nil
	}

	srcFiles, err := models.GetSourceFilesByDealFileId(dealFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var srcFilePayloadCids []string
	for _, srcFile := range srcFiles {
		paymentInfo, err := client.GetPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		lockedFee, err := decimal.NewFromString(paymentInfo.LockedFee.String())
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		err = models.UpdateSourceFileRefundAmount(srcFile.ID, lockedFee)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		srcFilePayloadCids = append(srcFilePayloadCids, srcFile.PayloadCid)
	}

	refundStatus := constants.PROCESS_STATUS_UNLOCK_REFUNDED
	tx, err := swanPaymentTransactor.Refund(tansactOpts, srcFilePayloadCids)
	if err != nil {
		refundStatus = constants.PROCESS_STATUS_UNLOCK_REFUNDFAILED
		logs.GetLogger().Error(err.Error())
	}

	for _, srcFile := range srcFiles {
		txHash := ""
		if tx != nil {
			txHash = tx.Hash().Hex()
		}

		err = models.UpdateSourceFileRefundStatus(srcFile.ID, refundStatus, txHash)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			continue
		}
	}

	err = models.UpdateDealFileStatus(dealFileId, refundStatus)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
