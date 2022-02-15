package scheduler

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"

	"github.com/filswan/go-swan-lib/logs"
)

func CheckSourceFilePaid() error {
	srcFiles, err := models.GetSourceFilesByStatus(constants.SOURCE_FILE_STATUS_CREATED)
	if err != nil {
		logs.GetLogger().Error(err)
	}

	for _, srcFile := range srcFiles {
		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		coin, err := models.FindCoinByCoinAddress(lockedPayment.TokenAddress)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		srcFile, err := models.GetSourceFileByPayloadCid(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		currentUtcMilliSecond := utils.GetCurrentUtcMilliSecond()
		eventLockPayment := models.EventLockPayment{
			MinPayment:      lockedPayment.MinPayment,
			LockedFee:       lockedPayment.LockedFee,
			Deadline:        lockedPayment.Deadline,
			TokenAddress:    lockedPayment.TokenAddress,
			AddressFrom:     lockedPayment.AddressFrom,
			AddressTo:       lockedPayment.AddressTo,
			LockPaymentTime: currentUtcMilliSecond,
			CoinId:          coin.ID,
			NetworkId:       coin.NetworkId,
			SourceFileId:    srcFile.ID,
		}

		err = models.CreateEventLockPayment(eventLockPayment)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
