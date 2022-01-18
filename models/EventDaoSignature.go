package models

import (
	"payment-bridge/common/constants"
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

// FindDaoEventLog (&DaoEventLog{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindDaoEventLog(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventDaoSignature, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventDaoSignature
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

func GetDeal2BeUnlocked(threshHold uint8) ([]*DealUnlockable, error) {
	sql := " SELECT payload_cid,deal_id,count(*) as sign_count FROM event_dao_signature " +
		" WHERE signature_unlock_status = '0' and deal_id > 0 " +
		" GROUP BY payload_cid,deal_id HAVING (sign_count>=?) "

	var models []*DealUnlockable
	err := database.GetDB().Raw(sql, threshHold).Scan(&models).Limit(constants.DEFAULT_SELECT_LIMIT).Offset(0).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
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
