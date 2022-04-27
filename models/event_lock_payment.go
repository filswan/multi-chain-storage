package models

import (
	"github.com/shopspring/decimal"
)

type EventLockPayment struct {
	ID              int64           `json:"id"`
	TxHash          string          `json:"tx_hash"`
	PayloadCid      string          `json:"payload_cid"`
	TokenAddress    string          `json:"token_address"`
	MinPayment      string          `json:"min_payment"`
	ContractAddress string          `json:"contract_address"`
	LockedFee       decimal.Decimal `json:"locked_fee"`
	Deadline        string          `json:"deadline"`
	BlockNo         uint64          `json:"block_no"`
	AddressFrom     string          `json:"address_from"`
	AddressTo       string          `json:"address_to"`
	CoinId          int64           `json:"coin_id"`
	NetworkId       int64           `json:"network_id"`
	LockPaymentTime int64           `json:"lock_payment_time"`
	CreateAt        int64           `json:"create_at"`
	SourceFileId    int64           `json:"source_file_id"`
}

type EventLockPaymentQuery struct {
	PayloadCid string `json:"payload_cid"`
	DealId     string `json:"deal_id"`
	Recipient  string `json:"recipient"`
	DealFileId int64  `json:"deal_file_id"`
}
