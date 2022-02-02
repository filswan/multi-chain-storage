package models

import (
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type EventDaoSignature struct {
	ID                    int64  `json:"id"`
	TxHash                string `json:"tx_hash"`
	Recipient             string `json:"recipient"`
	PayloadCid            string `json:"payload_cid"`
	DealId                int64  `json:"deal_id"`
	DaoPassTime           string `json:"dao_pass_time"`
	BlockNo               uint64 `json:"block_no"`
	BlockTime             string `json:"block_time"`
	Status                bool   `json:"status"`
	NetworkId             int64  `json:"network_id"`
	CoinId                int64  `json:"coin_id"`
	DaoAddress            string `json:"dao_address"`
	SignatureUnlockStatus string `json:"signature_unlock_status"`
	TxHashUnlock          string `json:"tx_hash_unlock"`
}

type DealUnlockable struct {
	PayloadCid string `json:"payload_cid"`
	DealId     int64  `json:"deal_id"`
	SignCount  string `json:"sign_count"`
}

//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateDaoEventLog(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	hardware := EventDaoSignature{}
	err := db.Model(&hardware).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}
