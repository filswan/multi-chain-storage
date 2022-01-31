package client

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/models"

	"github.com/filswan/go-swan-lib/logs"
)

func GetPaymentInfo(srcFilePayloadCid string) (*models.EventLockPayment, error) {
	swanPaymentSession, err := GetPaymentSession()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	paymentInfo, err := swanPaymentSession.GetLockedPaymentInfo(srcFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var event *models.EventLockPayment
	if paymentInfo.IsExisted {
		event = new(models.EventLockPayment)
		event.TokenAddress = paymentInfo.Token.Hex()
		event.AddressFrom = paymentInfo.Owner.String()
		event.AddressTo = paymentInfo.Recipient.String()
		event.LockedFee = paymentInfo.LockedFee.String()
		event.PayloadCid = srcFilePayloadCid
		event.LockPaymentTime = utils.GetCurrentUtcMilliSecond()
		usdcCoinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_USDC_ON_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.CoinId = usdcCoinId
		}

		networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.NetworkId = networkId
		}

		err = database.SaveOne(event)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		dealFiles, err := models.GetDealFileBySourceFilePayloadCid(srcFilePayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		if len(dealFiles) > 0 {
			dealFileId := dealFiles[0].ID

			err := models.UpdateDealFileLockPaymentStatus(dealFileId, constants.LOCK_PAYMENT_STATUS_PROCESSING)
			if err != nil {
				logs.GetLogger().Error(err)
				return nil, err
			}
		}
	}

	return event, nil
}
