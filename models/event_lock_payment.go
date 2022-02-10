package models

import (
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
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
	MinerAddress    string          `json:"miner_address"`
	AddressFrom     string          `json:"address_from"`
	AddressTo       string          `json:"address_to"`
	CoinId          int64           `json:"coin_id"`
	NetworkId       int64           `json:"network_id"`
	LockPaymentTime int64           `json:"lock_payment_time"`
	CreateAt        int64           `json:"create_at"`
	UnlockTxHash    string          `json:"unlock_tx_hash"`
	UnlockTxStatus  string          `json:"unlock_tx_status"`
	UnlockTime      string          `json:"unlock_time"`
	SourceFileId    int64           `json:"source_file_id"`
}

func GetEventLockPaymentBySrcPayloadCid(srcFilePayloadCid string) ([]*EventLockPayment, error) {
	var eventLockPayments []*EventLockPayment
	sql := "select * from event_lock_payment a where a.payload_cid=?"

	query := database.GetDB().Raw(sql, srcFilePayloadCid).Scan(&eventLockPayments)

	err := query.Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return eventLockPayments, nil
}

func GetEventLockPaymentByPayloadCidWallet(payloadCid, walletAddress string) ([]*EventLockPayment, error) {
	var eventLockPayment []*EventLockPayment
	err := database.GetDB().Where("payload_cid=? and address_from=?", payloadCid, walletAddress).Find((&eventLockPayment)).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return eventLockPayment, nil
}

func GetEventLockPaymentByPayloadCid(payloadCid string) ([]*EventLockPayment, error) {
	var eventLockPayment []*EventLockPayment
	err := database.GetDB().Where("payload_cid=?", payloadCid).Find((&eventLockPayment)).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return eventLockPayment, nil
}
