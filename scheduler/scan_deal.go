package scheduler

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-lib/client/lotus"
	"github.com/robfig/cron"
)

func ScanDealInfoScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.ScanDealStatusRule, func() {
		logs.GetLogger().Info("scanning deal info from chain scheduler")
		err := scanDeals()
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

func scanDeals() error {
	whereCondition := "deal_cid != '' and task_uuid != '' and lower(lock_payment_status) not in (lower('" + constants.LOCK_PAYMENT_STATUS_SUCCESS + "'), lower('" + constants.LOCK_PAYMENT_STATUS_REFUNDED + "'))"
	//" and deal_status not in (" + inList + ")"
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "100", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ClientApiUrl, config.GetConfig().Lotus.ClientAccessToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, deal := range dealList {
		dealInfo, err := lotusClient.LotusClientGetDealInfo(deal.DealCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		switch dealInfo.Status {
		case constants.DEAL_STATUS_ACTIVE:
			deal.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_SUCCESS
		case constants.DEAL_STATUS_ERROR:
			deal.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_REFUNDED
		default:
			deal.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_PROCESSING
		}

		deal.Verified = dealInfo.Verified
		deal.DealStatus = dealInfo.Status
		deal.DealId = dealInfo.DealId
		deal.Cost = dealInfo.CostComputed
		deal.UpdateAt = utils.GetCurrentUtcMilliSecond()
		err = database.SaveOne(deal)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}
