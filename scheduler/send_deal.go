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

	"github.com/filswan/go-swan-lib/logs"

	libconstants "github.com/filswan/go-swan-lib/constants"
)

func SendDealScheduler() {
	Mutex := &sync.Mutex{}
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.CreateTaskRule, func() {
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
	whereCondition := "send_deal_status ='' and lock_payment_status='" + constants.LOCK_PAYMENT_STATUS_PROCESSING + "' and task_uuid != '' "
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "50", "0")
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

	for _, deal := range dealList {
		logs.GetLogger().Info("start to send deal for task:", deal.TaskUuid)
		cmdAutoBidDeal.OutputDir = filepath.Dir(deal.CarFilePath)

		_, fileDescs, err := cmdAutoBidDeal.SendAutoBidDealsByTaskUuid(deal.TaskUuid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		if len(fileDescs) == 0 {
			logs.GetLogger().Info("no deals sent")
			continue
		}

		deal.SendDealStatus = constants.SEND_DEAL_STATUS_SUCCESS
		deal.ClientWalletAddress = cmdAutoBidDeal.SenderWallet
		deal.DealCid = fileDescs[0].Deals[0].DealCid
		deal.MinerFid = fileDescs[0].Deals[0].MinerFid

		offlineDeal := models.OfflineDeal{
			DealFileId:   deal.ID,
			DealCid:      fileDescs[0].Deals[0].DealCid,
			MinerFid:     fileDescs[0].Deals[0].MinerFid,
			StartEpoch:   fileDescs[0].Deals[0].StartEpoch,
			SenderWallet: cmdAutoBidDeal.SenderWallet,
		}

		err = database.SaveOne(offlineDeal)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}
