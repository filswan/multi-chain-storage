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

	//refund(int64(903), swanPaymentTransactor, tansactOpts)

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
	offlineDealsNotUnlocked, err := models.GetOfflineDealsNotUnlockedByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	if len(offlineDealsNotUnlocked) > 0 {
		msg := fmt.Sprintf("%d deals not unlocked or unlock failed, cannot refund for the car file", len(offlineDealsNotUnlocked))
		logs.GetLogger().Info(msg)
		return nil
	}

	srcFiles, err := models.GetSourceFileUploadsByCarFileId(carFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	var srcFilePayloadCids []string
	for _, srcFile := range srcFiles {
		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		err = models.UpdateSourceFileRefundAmount(srcFile.Id, lockedPayment.LockedFee)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			return err
		}

		srcFilePayloadCids = append(srcFilePayloadCids, srcFile.PayloadCid)
	}

	privateKey, publicKeyAddress, err := client.GetPrivateKeyPublicKey(constants.PRIVATE_KEY_ON_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, privateKey, *publicKeyAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
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

		logs.GetLogger().Info("refund stats:", refundStatus, " tx hash:", txHash)

		err = models.UpdateSourceFileRefundStatus(srcFile.Id, refundStatus, txHash)
		if err != nil {
			logs.GetLogger().Error(err.Error())
			continue
		}
	}

	err = models.UpdateDealFileStatus(carFileId, refundStatus)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		return err
	}

	return nil
}
