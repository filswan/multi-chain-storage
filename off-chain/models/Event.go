package models

import (
	"payment-bridge/off-chain/common/constants"
	"payment-bridge/off-chain/database"
)

type Event struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
	IndexedData     string `json:"indexed_data"`
	ContractName    string `json:"contract_name"`
	ContractAddress string `json:"contract_address"`
	LockedFee       string `json:"locked_fee"`
	Deadline        string `json:"deadline"`
	PayloadCid      string `json:"payload_cid"`
	BlockNo         uint64 `json:"block_no"`
}

func (self *Event) FindOneEvent(condition interface{}) (*Event, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEvents(whereCondition interface{}, orderCondition string, limit, offset string) ([]*Event, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*Event
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
