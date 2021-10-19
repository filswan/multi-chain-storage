package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type DaoEventLog struct {
	ID          int64  `json:"id"`
	TxHash      string `json:"tx_hash"`
	PayloadCid  string `json:"payload_cid"`
	OrderId     string `json:"order_id"`
	DealCid     string `json:"deal_cid"`
	Terms       string `json:"terms"`
	Cost        string `json:"cost"`
	DaoAddress  string `json:"dao_address"`
	DaoPassTime string `json:"dao_pass_time"`
	BlockNo     uint64 `json:"block_no"`
	Network     string `json:"network"`
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
