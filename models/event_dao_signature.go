package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"

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

func GetEventDaoSignaturesByDealId(dealId int64) ([]*EventDaoSignature, error) {
	var eventDaoSignatures []*EventDaoSignature
	sql := "select * from event_dao_signature a where a.deal_id=? and signature_unlock_status =" + constants.SIGNATURE_SUCCESS_VALUE

	query := database.GetDB().Raw(sql, dealId).Order("block_time desc").Scan(&eventDaoSignatures)

	err := query.Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return eventDaoSignatures, nil
}

func GetEventDaoSignatures(whereCondition interface{}) (*EventDaoSignature, error) {
	db := database.GetDB()
	var eventDaoSignatures []*EventDaoSignature
	err := db.Where(whereCondition).Find(&eventDaoSignatures).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(eventDaoSignatures) >= 1 {
		return eventDaoSignatures[0], nil
	} else {
		return nil, nil
	}
}
