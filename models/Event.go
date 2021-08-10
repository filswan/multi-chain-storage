package models

import (
	"errors"
	"payment-bridge/database"
	"strconv"
)

type Event struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
	IndexedData     string `json:"indexed_data"`
	Data            string `json:"data"`
	ContractName    string `json:"contract_name"`
	ContractAddress string `json:"contract_address"`
}

func (self *Event) FindOneEvent(condition interface{}) (*Event, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

func FindEventsByTxHash(txHash, limit, offset string) ([]Event, int, error, int) { //the last int returned represents error situation
	db := database.GetDB()
	var events []Event

	var hasOffset bool = len(offset) > 0
	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		return events, 0, err, 1
	}
	offset_int, err := strconv.Atoi(offset)
	if err != nil && hasOffset {
		return events, 0, err, 2
	}
	count := 0
	db.Where(&Event{TxHash: txHash}).Find(&events).Count(&count)
	if offset_int >= count && hasOffset {
		return events, 0, errors.New("offset exceeds count"), 3
	}
	db.Order("id desc").Where(&Event{TxHash: txHash}).Offset(offset_int).Limit(limit_int).Find(&events)

	return events, count, err, 4
}
