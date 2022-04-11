package scheduler

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/logs"
	"multi-chain-storage/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/client/lotus"
	libutils "github.com/filswan/go-swan-lib/utils"
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
	for {
		err := GetExpiredDealInfoAndUpdateInfoToDB()
		if err != nil {
			logs.GetLogger().Error(err)
		}

		time.Sleep(30 * time.Minute)
	}
}

func GetDealInfoByLotusClientAndUpdateInfoToDB() error {
	inList := "'" + strings.Join(strings.Split(config.GetConfig().Lotus.FinalStatusList, ","), "', '") + "'"
	fmt.Println(inList)
	whereCondition := "deal_cid != '' and task_uuid != '' and lock_payment_status not in ('" + constants.LOCK_PAYMENT_STATUS_SUCCESS + "', '" + constants.LOCK_PAYMENT_STATUS_REFUNDED + "','" + constants.LOCK_PAYMENT_STATUS_REFUNDING + "')"
	//" and deal_status not in (" + inList + ")"
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "10000", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ApiUrl, config.GetConfig().Lotus.AccessToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, v := range dealList {
		dealInfo, err := lotusClient.LotusClientGetDealInfo(v.DealCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		paymentStatus := ""
		if strings.EqualFold(dealInfo.Status, constants.DEAL_STATUS_ACTIVE) {
			paymentStatus = constants.LOCK_PAYMENT_STATUS_SUCCESS
			err := deleteFilesInDealTemp(v.ID, v.CarFilePath)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		} else if strings.EqualFold(dealInfo.Status, constants.DEAL_STATUS_ERROR) {
			//logs.GetLogger().Error("deal:", v.DealCid, ", status:", dealInfo.Status)
			//paymentStatus = constants.LOCK_PAYMENT_STATUS_WAITING
			paymentStatus = constants.LOCK_PAYMENT_STATUS_PROCESSING
			err := deleteFilesInDealTemp(v.ID, v.CarFilePath)
			if err != nil {
				logs.GetLogger().Error(err)
			}
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
		}

		err = SaveOfflineDealStatus(v.DealCid, dealInfo.Status, dealInfo.Message)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}
	return nil
}

func deleteFilesInDealTemp(dealFileId int64, carFilepath string) error {
	carFileDir := filepath.Base(filepath.Base(carFilepath))

	if libutils.IsDirExists(carFileDir) {
		err := os.RemoveAll(carFileDir)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}

	sourceFiles, err := models.GetSourceFilesByDealFileId(dealFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range sourceFiles {
		srcFileDir := filepath.Base(filepath.Base(srcFile.ResourceUri))

		if libutils.IsDirExists(srcFileDir) {
			err := os.RemoveAll(srcFileDir)
			if err != nil {
				logs.GetLogger().Error(err)
			}
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

	//logs.GetLogger().Info(dealCid, " on chain status:", status, " on chain message:", message)
	if logExists {
		logs.GetLogger().Info(dealCid, " deal status not changed")
		return nil
	}

	err = database.SaveOne(&dealLog)
	if err != nil {
		logs.GetLogger().Error(err, " deal:", dealCid, " on chain status:", status, " on chain message:", message)
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
