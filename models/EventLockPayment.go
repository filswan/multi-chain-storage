package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type EventLockPayment struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
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
	CoinId          int64  `json:"coin_id"`
	NetworkId       int64  `json:"network_id"`
	LockPaymentTime string `json:"lock_payment_time"`
	CreateAt        int64  `json:"create_at"`
	UnlockTxHash    string `json:"unlock_tx_hash"`
	UnlockTxStatus  string `json:"unlock_tx_status"`
	UnlockTime      string `json:"unlock_time"`
	SourceFileId    *int64 `json:"source_file_id"`
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventLockPayment(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventLockPayment, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventLockPayment
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

//condition :&models.EventPolygon{Ip: "192.168.88.80"}
//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateEventLockPayment(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	hardware := EventLockPayment{}
	err := db.Model(&hardware).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
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

func GetTotalLockFeeBySrcPayloadCid(carFilePayloadCid string) (*decimal.Decimal, error) {
	eventLockPayments, err := GetEventLockPaymentBySrcPayloadCid(carFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	totalLockFee := decimal.NewFromFloat(0.0)
	for _, localPayment := range eventLockPayments {
		lockFee, err := decimal.NewFromString(localPayment.LockedFee)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		totalLockFee = totalLockFee.Add(lockFee)
	}

	return &totalLockFee, nil
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
