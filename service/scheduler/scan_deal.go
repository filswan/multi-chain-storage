package scheduler

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"strconv"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-lib/client/lotus"
)

func ScanDeal() error {
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

		if offlineDeal.OnChainStatus == nil || *offlineDeal.OnChainStatus != dealInfo.Status || offlineDeal.DealId != dealInfo.DealId {
			offlineDeal.OnChainStatus = &dealInfo.Status
			offlineDeal.DealId = dealInfo.DealId
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

func GetExpiredDealInfoAndUpdateInfoToDB() error {
	eventLockPayment, err := models.FindExpiredLockPayment()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, v := range eventLockPayment {
		isLockedPaymentExists, err := client.IsLockedPaymentExists(v.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if !*isLockedPaymentExists {
				err = models.UpdateDealFileStatus(v.DealFileId, constants.PROCESS_STATUS_EXPIRE_REFUNDED)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
			continue
		}

		_dealFileId := v.DealFileId
		paymentStatus := constants.PROCESS_STATUS_EXPIRE_REFUNDING
		eventExpireList, err := models.FindEventExpirePayments(&models.EventExpirePayment{PayloadCid: v.PayloadCid}, "id desc", "10", "0")
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		for _, e := range eventExpireList {
			lockAmount, err := strconv.ParseInt(e.ExpireUserAmount, 10, 64)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			if lockAmount > 0 {
				paymentStatus = constants.PROCESS_STATUS_EXPIRE_REFUNDED
			}
		}
		err = models.UpdateDealFileStatus(_dealFileId, paymentStatus)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

	}
	return nil
}
