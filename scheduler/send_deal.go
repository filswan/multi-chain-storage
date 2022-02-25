package scheduler

import (
	"sync"

	"github.com/filswan/go-swan-client/command"
	"github.com/robfig/cron"

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

func CreateScheduler4SendDeal() {
	c := cron.New()
	name := "send deal"
	rule := config.GetConfig().ScheduleRule.SendDealRule
	mutex := &sync.Mutex{}

	err := c.AddFunc(rule, func() {
		logs.GetLogger().Info(name, " start")

		mutex.Lock()
		logs.GetLogger().Info(name, " running")
		err := SendDeal()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		mutex.Unlock()
		logs.GetLogger().Info(name, " end")
	})

	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	c.Start()
}

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

	currentUtcMilliSec := utils.GetCurrentUtcMilliSecond()
	for _, dealFile := range dealFiles {
		if currentUtcMilliSec-dealFile.CreateAt > 3*24*60*60*1000 {
			dealFile.LockPaymentStatus = constants.PROCESS_STATUS_DEAL_SEND_CANCELLED
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
			dealFile.LockPaymentStatus = constants.PROCESS_STATUS_DEAL_SENT_FAILED
			dealFile.ClientWalletAddress = cmdAutoBidDeal.SenderWallet
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

		dealFile.LockPaymentStatus = constants.PROCESS_STATUS_DEAL_SENT
		dealFile.ClientWalletAddress = cmdAutoBidDeal.SenderWallet
		dealFile.UpdateAt = currentUtcMilliSec

		db := database.GetDBTransaction()
		err = database.SaveOneInTransaction(db, dealFile)
		if err != nil {
			logs.GetLogger().Error(err)
			db.Rollback()
			return err
		}

		currentUtcMilliSec := utils.GetCurrentUtcMilliSecond()
		for _, deal := range fileDescs[0].Deals {
			dealInfo, err := lotusClient.LotusClientGetDealInfo(deal.DealCid)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			offlineDeal := models.OfflineDeal{
				DealFileId:   dealFile.ID,
				DealCid:      deal.DealCid,
				MinerFid:     deal.MinerFid,
				StartEpoch:   deal.StartEpoch,
				SenderWallet: cmdAutoBidDeal.SenderWallet,
				Status:       dealInfo.Status,
				DealId:       dealInfo.DealId,
				UnlockStatus: constants.OFFLINE_DEAL_UNLOCK_STATUS_NOT_UNLOCKED,
				CreateAt:     currentUtcMilliSec,
				UpdateAt:     currentUtcMilliSec,
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
