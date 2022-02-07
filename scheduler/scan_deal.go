package scheduler

import (
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"

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

		if deal.Status != dealInfo.Status || deal.DealId != dealInfo.DealId {
			deal.Status = dealInfo.Status
			deal.DealId = dealInfo.DealId
			deal.UpdateAt = utils.GetCurrentUtcMilliSecond()
			err = database.SaveOne(deal)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}
	return nil
}
