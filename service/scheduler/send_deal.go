package scheduler

import (
	"github.com/filswan/go-swan-client/command"

	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"path/filepath"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"

	libconstants "github.com/filswan/go-swan-lib/constants"
)

func SendDeal() error {
	carFiles, err := models.GetCarFilesByStatus(constants.CAR_FILE_STATUS_TASK_CREATED)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	cmdAutoBidDeal := &command.CmdAutoBidDeal{
		SwanApiUrl:             config.GetConfig().SwanApi.ApiUrl,
		SwanApiKey:             config.GetConfig().SwanApi.ApiKey,
		SwanAccessToken:        config.GetConfig().SwanApi.AccessToken,
		LotusClientApiUrl:      config.GetConfig().Lotus.ClientApiUrl,
		LotusClientAccessToken: config.GetConfig().Lotus.ClientAccessToken,
		SenderWallet:           config.GetConfig().FilecoinWallet,
		DealSourceIds:          []int{libconstants.TASK_SOURCE_ID_SWAN_PAYMENT},
	}

	currentUtcSec := libutils.GetCurrentUtcSecond()

	wallet, err := models.GetWalletByAddress(cmdAutoBidDeal.SenderWallet, constants.WALLET_TYPE_FILE_COIN)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, carFile := range carFiles {
		if currentUtcSec-carFile.CreateAt > 3*24*60*60 {
			carFile.Status = constants.CAR_FILE_STATUS_DEAL_SEND_EXPIRED
			err = database.SaveOne(carFile)
			if err != nil {
				logs.GetLogger().Error(err)
			}

			continue
		}

		logs.GetLogger().Info("start to send deal for task:", carFile.TaskUuid)
		cmdAutoBidDeal.OutputDir = filepath.Dir(carFile.CarFilePath)

		_, fileDescs, err := cmdAutoBidDeal.SendAutoBidDealsByTaskUuid(carFile.TaskUuid)
		if err != nil {
			logs.GetLogger().Error(err)
			carFile.Status = constants.CAR_FILE_STATUS_DEAL_SENT_FAILED
			carFile.UpdateAt = currentUtcSec
			err = database.SaveOne(carFile)
			if err != nil {
				logs.GetLogger().Error(err)
			}
			continue
		}

		if len(fileDescs) == 0 {
			logs.GetLogger().Info("no deals sent")
			continue
		}

		for _, fileDesc := range fileDescs {
			for _, deal := range fileDesc.Deals {
				miner, err := models.GeMinerByFid(deal.MinerFid)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}

				offlineDeal := models.OfflineDeal{
					CarFileId:      carFile.ID,
					DealCid:        deal.DealCid,
					MinerId:        miner.ID,
					StartEpoch:     deal.StartEpoch,
					SenderWalletId: wallet.ID,
					Status:         constants.OFFLINE_DEAL_STATUS_CREATED,
					DealId:         nil,
					CreateAt:       currentUtcSec,
					UpdateAt:       currentUtcSec,
				}

				err = database.SaveOne(&offlineDeal)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
			}
		}

		carFile.Status = constants.CAR_FILE_STATUS_DEAL_SENT
		carFile.UpdateAt = currentUtcSec

		err = database.SaveOne(carFile)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}
