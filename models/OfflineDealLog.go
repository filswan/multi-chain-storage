package models

import (
	"payment-bridge/database"
	"payment-bridge/logs"
)

type OfflineDealLog struct {
	Id       int64  `json:"id"`
	DealCid  string `json:"deal_cid"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	CreateAt string `json:"create_at"`
}

func GetOfflineDealLogsByDealCid(dealCid string) ([]*OfflineDealLog, error) {
	var offlineDealLogs []*OfflineDealLog
	err := database.GetDB().Where("deal_cid=?", dealCid).Order("id desc").Find(&offlineDealLogs).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDealLogs, nil
}
