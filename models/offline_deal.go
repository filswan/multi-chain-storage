package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type OfflineDeal struct {
	Id           int64  `json:"id"`
	DealFileId   int64  `json:"deal_file_id"`
	DealCid      string `json:"deal_cid"`
	MinerFid     string `json:"miner_fid"`
	StartEpoch   int    `json:"start_epoch"`
	SenderWallet string `json:"sender_wallet"`
	Status       string `json:"status"`
	DealId       int64  `json:"deal_id"`
	CreateAt     int64  `json:"create_at"`
	UpdateAt     int64  `json:"update_at"`
}

func GetOfflineDealsBySourceFileId(sourceFileId int64) ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select a.* from offline_deal a, source_file_deal_file_map b where b.source_file_id=? and a.deal_file_id=b.deal_file_id"
	err := database.GetDB().Raw(sql, sourceFileId).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDeals2BeScanned() ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select a.* from offline_deal a where a.status !=? and a.status!=?"
	err := database.GetDB().Raw(sql, constants.DEAL_STATUS_ACTIVE, constants.DEAL_STATUS_ERROR).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}
