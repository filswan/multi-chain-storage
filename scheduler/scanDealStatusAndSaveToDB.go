package scheduler

import (
	"fmt"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/client/lotus"
	"github.com/robfig/cron"
)

func ScanDealInfoScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.ScanDealStatusRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ scan deal info from chain scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := GetDealInfoByLotusClientAndUpdateInfoToDB()
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

func ScanExpiredDealInfoScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UpdatePayStatusRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ scan expired deal info scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := GetExpiredDealInfoAndUpdateInfoToDB()
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

func GetDealInfoByLotusClientAndUpdateInfoToDB() error {
	inList := "'" + strings.Join(strings.Split(config.GetConfig().Lotus.FinalStatusList, ","), "', '") + "'"
	fmt.Println(inList)
	whereCondition := "deal_cid != '' and task_uuid != '' and lower(lock_payment_status) not in (lower('" + constants.LOCK_PAYMENT_STATUS_SUCCESS + "'), lower('" + constants.LOCK_PAYMENT_STATUS_REFUNDED + "'))"
	//" and deal_status not in (" + inList + ")"
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "100", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, v := range dealList {
		lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ApiUrl, config.GetConfig().Lotus.AccessToken)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		dealInfo, err := lotusClient.LotusClientGetDealInfo(v.DealCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		paymentStatus := ""
		if strings.ToLower(dealInfo.Status) == strings.ToLower(constants.DEAL_STATUS_ACTIVE) {
			paymentStatus = constants.LOCK_PAYMENT_STATUS_SUCCESS
		} else if strings.ToLower(dealInfo.Status) == strings.ToLower(constants.DEAL_STATUS_ERROR) {
			logs.GetLogger().Error("deal:", v.DealCid, ", status:", dealInfo.Status)
			//paymentStatus = constants.LOCK_PAYMENT_STATUS_WAITING
			paymentStatus = constants.LOCK_PAYMENT_STATUS_PROCESSING
		} else {
			paymentStatus = constants.LOCK_PAYMENT_STATUS_PROCESSING
		}
		v.Verified = dealInfo.Verified
		v.DealStatus = dealInfo.Status
		v.DealId = dealInfo.DealId
		v.Cost = dealInfo.CostComputed
		v.LockPaymentStatus = paymentStatus
		v.UpdateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
		err = database.SaveOne(v)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		err = SaveOfflineDealStatus(v.DealCid, dealInfo.Status, dealInfo.Message)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}

func SaveOfflineDealStatus(dealCid, status, message string) error {
	dealLog := models.OfflineDealLog{
		DealCid:  dealCid,
		Status:   status,
		Message:  message,
		CreateAt: strconv.FormatInt(utils.GetEpochInMillis(), 10),
	}

	offlineDealLogs, err := models.GetOfflineDealLogsByDealCid(dealCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	logExists := false
	for _, offlineDealLog := range offlineDealLogs {
		if strings.EqualFold(offlineDealLog.Status, status) && strings.EqualFold(offlineDealLog.Message, message) {
			logExists = true
		}
	}

	logs.GetLogger().Info(dealCid, " on chain status:", status, " on chain message:", message)
	if logExists {
		logs.GetLogger().Info(dealCid, " deal status not changed")
		return nil
	}

	err = database.SaveOne(&dealLog)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func GetExpiredDealInfoAndUpdateInfoToDB() error {
	eventLockPayment, err := models.FindExpiredLockPayment()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range eventLockPayment {
		_dealFileId := v.DealFileId
		paymentStatus := constants.LOCK_PAYMENT_STATUS_REFUNDING
		eventExpireList, err := models.FindEventExpirePayments(&models.EventExpirePayment{PayloadCid: v.PayloadCid}, "id desc", "10", "0")
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		for _, e := range eventExpireList {
			lockAmount, err := strconv.ParseInt(e.ExpireUserAmount, 10, 64)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			if lockAmount > 0 {
				paymentStatus = constants.LOCK_PAYMENT_STATUS_REFUNDED
			}
		}
		err = models.UpdateDealFile(&models.DealFile{ID: _dealFileId}, &models.DealFile{LockPaymentStatus: paymentStatus, UpdateAt: strconv.FormatInt(utils.GetEpochInMillis(), 10)})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}
