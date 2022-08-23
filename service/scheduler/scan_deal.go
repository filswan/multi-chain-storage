package scheduler

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-lib/client/lotus"
)

func ScanDeal() error {
	err := ScanDealBeforeActive()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	err = ScanDealAfterActive()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func ScanDealBeforeActive() error {
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
		offlineDealStatusChanged := false
		dealInfo, err := lotusClient.LotusClientGetDealInfo(offlineDeal.DealCid)

		if err != nil {
			logs.GetLogger().Error(err)
			if strings.Contains(err.Error(), "datastore: key not found") {
				offlineDealStatusChanged = true
				note := err.Error()
				offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_FAILED
				offlineDeal.Note = &note
			} else {
				continue
			}
		} else {
			if offlineDeal.OnChainStatus == nil || *offlineDeal.OnChainStatus != dealInfo.Status || offlineDeal.DealId == nil || *offlineDeal.DealId != dealInfo.DealId ||
				offlineDeal.Note == nil || *offlineDeal.Note != dealInfo.Message {
				offlineDealStatusChanged = true
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

				offlineDeal.Note = &dealInfo.Message
			}
		}

		if offlineDeal.DealId != nil && !strings.EqualFold(offlineDeal.Status, constants.OFFLINE_DEAL_STATUS_ACTIVE) {
			isDealActive, err := utils.IsDealActive(*offlineDeal.DealId)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			if *isDealActive {
				offlineDealStatusChanged = true
				offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_ACTIVE
			}
		}

		if offlineDealStatusChanged {
			onChainStatus := ""
			if offlineDeal.OnChainStatus != nil {
				onChainStatus = *offlineDeal.OnChainStatus
			}

			onChainMessage := ""
			if offlineDeal.Note != nil {
				onChainMessage = *offlineDeal.Note
			}

			if offlineDeal.Status == constants.OFFLINE_DEAL_STATUS_ACTIVE {
				carFile, err := models.GetCarFileById(offlineDeal.CarFileId)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				if carFile.IsFree {
					offlineDeal.Status = constants.OFFLINE_DEAL_STATUS_SUCCESS
				}
			}
			err = models.CreateOfflineDealLog(offlineDeal.Id, onChainStatus, onChainMessage)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			offlineDeal.UpdateAt = libutils.GetCurrentUtcSecond()
			err = database.SaveOne(offlineDeal)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}
	}

	return nil
}

func ScanDealAfterActive() error {
	offlineDeals, err := models.GetOfflineDeals2BeScannedAfterActive()
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
			return err
		}

		if offlineDeal.OnChainStatus == nil || *offlineDeal.OnChainStatus != dealInfo.Status {
			err = models.CreateOfflineDealLog(offlineDeal.Id, dealInfo.Status, dealInfo.Message)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			err = models.UpdateOfflineDealOnChainStatus(offlineDeal.Id, dealInfo.Status)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}
	}

	return nil
}
