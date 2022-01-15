package scheduler

import (
	"github.com/filswan/go-swan-client/command"

	"path/filepath"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"

	"github.com/filswan/go-swan-lib/logs"

	libconstants "github.com/filswan/go-swan-lib/constants"
)

func sendDeal() error {
	whereCondition := "send_deal_status ='' and lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_PROCESSING + "') and task_uuid != '' "
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

		err = database.SaveOne(deal)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}
