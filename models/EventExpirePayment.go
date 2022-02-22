package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type EventExpirePayment struct {
	ID               int64  `json:"id"`
	TxHash           string `json:"tx_hash"`
	PayloadCid       string `json:"payload_cid"`
	BlockNo          string `json:"block_no"`
	TokenAddress     string `json:"token_address"`
	ContractAddress  string `json:"contract_address"`
	UserAddress      string `json:"user_address"`
	ExpireUserAmount string `json:"expire_user_amount"`
	BlockTime        string `json:"block_time"`
	CreateAt         string `json:"create_at"`
	NetworkId        int64  `json:"network_id"`
	CoinId           int64  `json:"coin_id"`
}

func FindEventExpirePayments(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventExpirePayment, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventExpirePayment
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
