package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type EventUnlockPayment struct {
	ID                   int64  `json:"id"`
	TxHash               string `json:"tx_hash"`
	PayloadCid           string `json:"payload_cid"`
	DealId               int64  `json:"deal_id"`
	BlockNo              string `json:"block_no"`
	NetworkId            int64  `json:"network"`
	TokenAddress         string `json:"token_address"`
	UnlockFromAddress    string `json:"unlock_from_address"`
	UnlockToUserAddress  string `json:"unlock_to_user_address"`
	UnlockToUserAmount   string `json:"unlock_to_user_amount"`
	UnlockToAdminAddress string `json:"unlock_to_admin_address"`
	UnlockToAdminAmount  string `json:"unlock_to_admin_amount"`
	UnlockTime           string `json:"unlock_time"`
	CreateAt             int64  `json:"create_at"`
	CoinId               int64  `json:"coin_id"`
	UnlockStatus         string `json:"unlock_status"`
	SourceFileId         *int64 `json:"source_file_id"`
}

func GetEventUnlockPaymentsByPayloadCid(payloadCid string, limit, offset string) ([]*EventUnlockPayment, error) {
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}

	var dealFiles []*EventUnlockPayment

	err := database.GetDB().Where("payload_cid=?", payloadCid).Offset(offset).Limit(limit).Find(&dealFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return dealFiles, nil
}
