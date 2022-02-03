package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type EventGoerli struct {
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
	MinerAddress    string `json:"miner_address"`
	AddressFrom     string `json:"address_from"`
	AddressTo       string `json:"address_to"`
	CoinType        string `json:"coin_type"`
	LockPaymentTime string `json:"lock_payment_time"`
	CreateAt        int64  `json:"create_at"`
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventGoerlis(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventGoerli, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventGoerli
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
