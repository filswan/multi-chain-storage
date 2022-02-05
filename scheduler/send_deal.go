package scheduler

import (
	"sync"

	"github.com/filswan/go-swan-client/command"

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

var sendDealMutex *sync.Mutex = &sync.Mutex{}
var sendDealRunning bool = false

func CreateSendDealScheduler() {
	CreateScheduler(config.GetConfig().ScheduleRule.SendDealRule, SendDeal, sendDealMutex, &sendDealRunning)
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
				DealId:       dealInfo.DealId,
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
