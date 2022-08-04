package scheduler

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"strings"

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

			if strings.Contains(err.Error(), "datastore: key not found") {
				note := err.Error()
				offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_FAILED
				offlineDeal.Note = &note
				offlineDeal.UpdateAt = libutils.GetCurrentUtcSecond()
				err = database.SaveOne(offlineDeal)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
			}
			continue
		}

		if offlineDeal.OnChainStatus == nil || *offlineDeal.OnChainStatus != dealInfo.Status || offlineDeal.DealId == nil || *offlineDeal.DealId != dealInfo.DealId {
			offlineDeal.OnChainStatus = &dealInfo.Status
			if dealInfo.DealId != 0 {
				offlineDeal.DealId = &dealInfo.DealId
			}

			switch dealInfo.Status {
			case constants.ON_CHAIN_DEAL_STATUS_ERROR:
				if !strings.Contains(dealInfo.Message, constants.ON_CHAIN_MESSAGE_NOT_COMPLETED) {
					offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_FAILED
				}
			case constants.ON_CHAIN_DEAL_STATUS_ACTIVE:
				offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_ACTIVE
			}

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
