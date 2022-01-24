package models

import (
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
