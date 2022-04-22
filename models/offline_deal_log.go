package models

import (
	"multi-chain-storage/database"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"
)

type OfflineDealLog struct {
	Id             int64  `json:"id"`
	OfflineDealId  int64  `json:"offline_deal_id"`
	OnChainStatus  string `json:"on_chain_status"`
	OnChainMessage string `json:"on_chain_message"`
	CreateAt       int64  `json:"create_at"`
}

func GetOfflineDealLogsByOfflineDealId(offlineDealId int64) ([]*OfflineDealLog, error) {
	var offlineDealLogs []*OfflineDealLog
	err := database.GetDB().Where("offline_deal_id=?", offlineDealId).Find(&offlineDealLogs).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDealLogs, nil
}

func CreateOfflineDealLog(offlineDealId int64, onChainStatus, onChainMessage string) error {
	var offlineDealLogs []*OfflineDealLog
	err := database.GetDB().Where("offline_deal_id=? and on_chain_status=? and on_chain_message=?", offlineDealId, onChainStatus, onChainMessage).Find(&offlineDealLogs).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(offlineDealLogs) > 0 {
		logs.GetLogger().Info("offline deal id:", offlineDealId, " deal status not changed")
		return nil
	}
	dealLog := OfflineDealLog{
		OfflineDealId:  offlineDealId,
		OnChainStatus:  onChainStatus,
		OnChainMessage: onChainMessage,
		CreateAt:       libutils.GetCurrentUtcSecond(),
	}

	err = database.SaveOne(&dealLog)
	if err != nil {
		logs.GetLogger().Error(err, ",offline deal id:", offlineDealId, " on chain status:", onChainStatus, " on chain message:", onChainMessage)
		return err
	}

	return nil
}
