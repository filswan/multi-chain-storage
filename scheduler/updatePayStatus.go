package scheduler

import (
	"github.com/robfig/cron"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"time"
)

func UpdatePayStatusScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UpdatePayStatusRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ create task  scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := DoUpdatePayStatus()
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

func DoUpdatePayStatus() error {
	list, err := FindUnpayDealFileList()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range list {
		err := models.UpdateDealFile(&models.DealFile{ID: v.ID}, map[string]interface{}{"lock_payment_status": constants.LOCK_PAYMENT_STATUS_SUCCESS, "lock_payment_tx": v.TxHash, "update_at": strconv.FormatInt(utils.GetEpochInMillis(), 10)})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return err
}

// find data that has not been paid within half an hour
func FindUnpayDealFileList() ([]*UnpaiedDealFile, error) {
	sql := " select a.*,b.tx_hash,b.locked_fee from ( " +
		" 	select * from deal_file " +
		"	where(unix_timestamp()*1000 - CONVERT(create_at, UNSIGNED INTEGER)) < " + strconv.Itoa(constants.TIME_HALF_AN_HOUR) +
		"    and lock_payment_status = '" + constants.LOCK_PAYMENT_STATUS_WAITING + "') a left join event_lock_payment b  " +
		" on a.payload_cid = b.payload_cid  " +
		" where b.tx_hash is not null "
	db := database.GetDB()
	var models []*UnpaiedDealFile
	err := db.Raw(sql).Scan(&models).Limit(constants.DEFAULT_SELECT_LIMIT).Offset(0).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
}

type UnpaiedDealFile struct {
	ID             int64  `json:"id"`
	PayloadCid     string `json:"payload_cid"`
	DealCid        string `json:"deal_cid"`
	DealId         int64  `json:"deal_id"`
	PieceCid       string `json:"piece_cid"`
	MinerFid       string `json:"miner_fid"`
	DealStatus     string `json:"deal_status"`
	PinStatus      string `json:"pin_status"`
	SourceFilePath string `json:"source_file_path"`
	TaskUuid       string `json:"task_uuid"`
	Cost           string `json:"cost"`
	CreateAt       string `json:"create_at"`
	TxHash         string `json:"tx_hash"`
	LockedFee      string `json:"locked_fee"`
}
