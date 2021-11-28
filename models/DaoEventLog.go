package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
	"payment-bridge/logs"
	"strconv"
)

type DaoEventLog struct {
	ID                    int64  `json:"id"`
	TxHash                string `json:"tx_hash"`
	Recipient             string `json:"recipient"`
	PayloadCid            string `json:"payload_cid"`
	OrderId               string `json:"order_id"`
	DealCid               string `json:"deal_cid"`
	Cost                  string `json:"cost"`
	DaoAddress            string `json:"dao_address"`
	DaoPassTime           string `json:"dao_pass_time"`
	BlockNo               uint64 `json:"block_no"`
	BlockTime             string `json:"block_time"`
	Network               string `json:"network"`
	Status                bool   `json:"status"`
	SignatureUnlockStatus string `json:"signature_unlock_status"`
}

type DaoSignatureResult struct {
	Recipient  string `json:"recipient"`
	PayloadCid string `json:"payload_cid"`
	OrderId    string `json:"order_id"`
	DealCid    string `json:"deal_cid"`
	Threshold  string `json:"threshold"`
}

// FindDaoEventLog (&DaoEventLog{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindDaoEventLog(whereCondition interface{}, orderCondition, limit, offset string) ([]*DaoEventLog, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*DaoEventLog
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

func GetDaoSignatureEventsSholdBeUnlock(threshHold uint8) ([]*DaoSignatureResult, error) {
	db := database.GetDB()
	var models []*DaoSignatureResult
	daoEventLog := DaoEventLog{}
	err := db.Model(daoEventLog).Select("payload_cid,order_id,deal_cid,count(*) as threshold").
		Where("signature_unlock_status = '0'").
		Group("payload_cid,order_id,deal_cid").
		Having("threshold >=" + strconv.Itoa(int(threshHold))).Scan(&models).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
}

//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateDaoEventLog(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	hardware := DaoEventLog{}
	err := db.Model(&hardware).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}
