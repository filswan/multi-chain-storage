package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"
	"strconv"
	"time"

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

func FindExpiredLockPayment() ([]*EventLockPaymentQuery, error) {
	sql :=
		"SELECT b.id as deal_file_id, a.payload_cid, b.deal_id, a.address_from as recipient " +
			"FROM event_lock_payment a, deal_file b " +
			"WHERE a.payload_cid = b.payload_cid and a.payload_cid not in (SELECT payload_cid FROM event_unlock_payment c) and lock_payment_status <> '" + constants.PROCESS_STATUS_EXPIRE_REFUNDED +
			"' and a.deadline < " + strconv.FormatInt(time.Now().Unix(), 10) +
			" and not exists (select 1 from offline_deal d where d.deal_file_id = b.id and unlock_status ='Unlocked') and b.deal_id <> 0"
	db := database.GetDB()
	var models []*EventLockPaymentQuery
	err := db.Raw(sql).Scan(&models).Limit(constants.DEFAULT_SELECT_LIMIT).Offset(0).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
}
