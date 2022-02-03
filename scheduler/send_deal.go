package scheduler

import (
	"sync"
	"time"

	"github.com/filswan/go-swan-client/command"
	"github.com/robfig/cron"

	"path/filepath"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"

	"payment-bridge/common/utils"

	"github.com/filswan/go-swan-lib/client/lotus"
	"github.com/filswan/go-swan-lib/logs"

	libconstants "github.com/filswan/go-swan-lib/constants"
)

func SendDealScheduler() {
	Mutex := &sync.Mutex{}
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.SendDealRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ create task  scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		Mutex.Lock()
		err := SendDeal()
		Mutex.Unlock()
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
	})
	if err != nil {
		logs.GetLogger().Error(err)
		return
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
		SenderWallet:           config.GetConfig().FileCoinWallet,
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
			dealFile.SendDealStatus = constants.DEAL_FILE_STATUS_CANCELLED
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
			dealFile.SendDealStatus = constants.DEAL_FILE_STATUS_FAILED
			dealFile.ClientWalletAddress = cmdAutoBidDeal.SenderWallet
			dealFile.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_DEAL_SENT_FAILED
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

		dealFile.SendDealStatus = constants.DEAL_FILE_STATUS_DEAL_SENT
		dealFile.ClientWalletAddress = cmdAutoBidDeal.SenderWallet
		dealFile.UpdateAt = currentUtcMilliSec
		err = database.SaveOne(dealFile)
		if err != nil {
			logs.GetLogger().Error(err)
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
				DealId:       dealFile.DealId,
				UnlockStatus: constants.OFFLINE_DEAL_UNLOCK_STATUS_NOT_UNLOCKED,
				CreateAt:     currentUtcMilliSec,
				UpdateAt:     currentUtcMilliSec,
			}

			err = database.SaveOne(&offlineDeal)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}

	return nil
}
