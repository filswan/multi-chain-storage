package scheduler

import (
	"github.com/filswan/go-swan-client/command"

	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"path/filepath"

	"multi-chain-storage/common/utils"

	"github.com/filswan/go-swan-lib/client/lotus"
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
		SenderWallet:           config.GetConfig().SwanPlatformFilWallet,
		DealSourceIds:          []int{libconstants.TASK_SOURCE_ID_SWAN_PAYMENT},
	}

	lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ClientApiUrl, config.GetConfig().Lotus.ClientAccessToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSec := utils.GetCurrentUtcSecond()

	wallet, err := models.GetWalletByAddress(cmdAutoBidDeal.SenderWallet, constants.WALLET_TYPE_FILE_COIN)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, carFile := range carFiles {
		if currentUtcSec-carFile.CreateAt > 3*24*60*60*1000 {
			carFile.Status = constants.CAR_FILE_STATUS_DEAL_SEND_CANCELLED
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
			carFile.Status = constants.PROCESS_STATUS_DEAL_SENT_FAILED
			carFile.Status = cmdAutoBidDeal.SenderWallet
			err = database.SaveOne(carFile)
			if err != nil {
				logs.GetLogger().Error(err)
			}
			logs.GetLogger().Error(err)
			continue
		}

		if len(fileDescs) == 0 {
			logs.GetLogger().Info("no deals sent")
			continue
		}

		carFile.Status = constants.PROCESS_STATUS_DEAL_SENT
		carFile.Status = cmdAutoBidDeal.SenderWallet
		carFile.UpdateAt = currentUtcSec

		db := database.GetDBTransaction()
		err = database.SaveOneInTransaction(db, carFile)
		if err != nil {
			logs.GetLogger().Error(err)
			db.Rollback()
			return err
		}

		currentUtcMilliSec := utils.GetCurrentUtcSecond()
		for _, deal := range fileDescs[0].Deals {
			dealInfo, err := lotusClient.LotusClientGetDealInfo(deal.DealCid)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			offlineDeal := models.OfflineDeal{
				CarFileId:      carFile.ID,
				DealCid:        deal.DealCid,
				MinerId:        deal.MinerFid,
				StartEpoch:     deal.StartEpoch,
				SenderWalletId: wallet.ID,
				Status:         dealInfo.Status,
				DealId:         dealInfo.DealId,
				CreateAt:       currentUtcMilliSec,
				UpdateAt:       currentUtcMilliSec,
			}

			err = database.SaveOneInTransaction(db, &offlineDeal)
			if err != nil {
				logs.GetLogger().Error(err)
				db.Rollback()
				return err
			}
		}

		err = db.Commit().Error
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
