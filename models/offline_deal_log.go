package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type OfflineDealLog struct {
	Id             int64  `json:"id"`
	OfflineDealId  int64  `json:"offline_deal_id"`
	OnChainStatus  string `json:"on_chain_status"`
	OnChainMessage string `json:"on_chain_message"`
	CreateAt       int64  `json:"create_at"`
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
