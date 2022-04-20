package scheduler

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"strconv"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-lib/client/lotus"
)

func ScanDeal() error {
	dealList, err := models.GetOfflineDeals2BeScanned()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ClientApiUrl, config.GetConfig().Lotus.ClientAccessToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, deal := range dealList {
		dealInfo, err := lotusClient.LotusClientGetDealInfo(deal.DealCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		if deal.OnChainStatus == nil || *deal.OnChainStatus != dealInfo.Status || deal.DealId != dealInfo.DealId {
			deal.OnChainStatus = &dealInfo.Status
			deal.DealId = dealInfo.DealId
			deal.UpdateAt = utils.GetCurrentUtcSecond()
			err = database.SaveOne(deal)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
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
