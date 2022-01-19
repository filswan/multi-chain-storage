package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/logs"
	"strconv"
)

type EventLockPayment struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	PayloadCid      string `json:"payload_cid"`
	TokenAddress    string `json:"token_address"`
	MinPayment      string `json:"min_payment"`
	ContractAddress string `json:"contract_address"`
	LockedFee       string `json:"locked_fee"`
	Deadline        string `json:"deadline"`
	BlockNo         uint64 `json:"block_no"`
	MinerAddress    string `json:"miner_address"`
	AddressFrom     string `json:"address_from"`
	AddressTo       string `json:"address_to"`
	CoinId          int64  `json:"coin_id"`
	NetworkId       int64  `json:"network_id"`
	LockPaymentTime string `json:"lock_payment_time"`
	CreateAt        string `json:"create_at"`
	UnlockTxHash    string `json:"unlock_tx_hash"`
	UnlockTxStatus  string `json:"unlock_tx_status"`
	UnlockTime      string `json:"unlock_time"`
}

type EventLockPaymentQuery struct {
	PayloadCid string `json:"payload_cid"`
	DealId     string `json:"deal_id"`
	Recipient  string `json:"recipient"`
	DealFileId int64  `json:"deal_file_id"`
}

func (self *EventLockPayment) FindOneEventPolygon(condition interface{}) (*EventLockPayment, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventLockPayment(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventLockPayment, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventLockPayment
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

//condition :&models.EventPolygon{Ip: "192.168.88.80"}
//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateEventLockPayment(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	hardware := EventLockPayment{}
	err := db.Model(&hardware).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}

func FindExpiredLockPayment() ([]*EventLockPaymentQuery, error) {
	sql :=
		"SELECT b.id as deal_file_id, a.payload_cid, b.deal_id, a.address_from as recipient " +
			"FROM event_lock_payment a, deal_file b " +
			"WHERE a.payload_cid = b.payload_cid and lock_payment_status <> '" + constants.LOCK_PAYMENT_STATUS_REFUNDED +
			"' and a.deadline < " + strconv.FormatInt(utils.GetEpochInMillis(), 10)
	db := database.GetDB()
	var models []*EventLockPaymentQuery
	err := db.Raw(sql).Scan(&models).Limit(constants.DEFAULT_SELECT_LIMIT).Offset(0).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
}
