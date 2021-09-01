package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type ChildChainTransaction struct {
	ID           int64  `json:"id"`
	TxHashInBsc  string `json:"tx_hash_in_bsc"`
	TxHashInNbai string `json:"tx_hash_in_nbai""`
	Status       string `json:"status"`
	BlockNo      uint64 `json:"block_no"` //in nbai
	RedoTimes    int8   `json:"redo_times"`
	CreateAt     string `json:"create_at"`
	UpdateAt     string `json:"create_at"`
}

// FindChildChainTransaction (&ChildChainTransaction{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindChildChainTransaction(whereCondition interface{}, orderCondition, limit, offset string) ([]*ChildChainTransaction, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*ChildChainTransaction
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
