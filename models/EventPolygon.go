package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
	"payment-bridge/logs"
)

type EventPolygon struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
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
	CoinType        string `json:"coin_type"`
	LockPaymentTime string `json:"lock_payment_time"`
	CreateAt        string `json:"create_at"`
	UnlockTxHash    string `json:"unlock_tx_hash"`
	UnlockTxStatus  string `json:"unlock_tx_hash"`
	UnlockTime      string `json:"unlock_time"`
}

func (self *EventPolygon) FindOneEventPolygon(condition interface{}) (*EventPolygon, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventPolygons(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventPolygon, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventPolygon
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

//condition :&models.EventPolygon{Ip: "192.168.88.80"}
//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateEventPolygon(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	hardware := EventPolygon{}
	err := db.Model(&hardware).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}
