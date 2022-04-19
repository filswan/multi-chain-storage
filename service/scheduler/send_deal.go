package scheduler

import (
	"fmt"

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
	dealFiles, err := models.GetDeal2Send()
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

	currentUtcMilliSec := utils.GetCurrentUtcSecond()

	wallet, err := models.GetWalletByAddress(cmdAutoBidDeal.SenderWallet)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if wallet == nil {
		err = fmt.Errorf("wallet:%s not exists", cmdAutoBidDeal.SenderWallet)
		logs.GetLogger().Error(err)
		return err
	}

	for _, dealFile := range dealFiles {
		if currentUtcMilliSec-dealFile.CreateAt > 3*24*60*60*1000 {
			dealFile.Status = constants.PROCESS_STATUS_DEAL_SEND_CANCELLED
			err = database.SaveOne(dealFile)
			if err != nil {
				logs.GetLogger().Error(err)
			}

			continue
		}

		logs.GetLogger().Info("start to send deal for task:", dealFile.TaskUuid)
		cmdAutoBidDeal.OutputDir = filepath.Dir(dealFile.CarFilePath)

		_, fileDescs, err := cmdAutoBidDeal.SendAutoBidDealsByTaskUuid(dealFile.TaskUuid)
		if err != nil {
			logs.GetLogger().Error(err)
			dealFile.Status = constants.PROCESS_STATUS_DEAL_SENT_FAILED
			dealFile.Status = cmdAutoBidDeal.SenderWallet
			err = database.SaveOne(dealFile)
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

		dealFile.Status = constants.PROCESS_STATUS_DEAL_SENT
		dealFile.Status = cmdAutoBidDeal.SenderWallet
		dealFile.UpdateAt = currentUtcMilliSec

		db := database.GetDBTransaction()
		err = database.SaveOneInTransaction(db, dealFile)
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
				CarFileId:      dealFile.ID,
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