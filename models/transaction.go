package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Transaction struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	Type               int    `json:"type"`
	TxHash             string `json:"tx_hash"`
	WalletIdFrom       int64  `json:"wallet_id_from"`
	WalletIdTo         int64  `json:"wallet_id_to"`
	CoinId             int64  `json:"coin_id"`
	Amount             string `json:"amount"`
	BlockNumber        int64  `json:"block_number"`
	TransactionAt      int64  `json:"transaction_at"`
	CreateAt           int64  `json:"create_at"`
}

func CreateTransaction(eventLockPayment EventLockPayment) error {
	currentUtcMilliSecond := utils.GetCurrentUtcSecond()
	eventLockPayment.CreateAt = currentUtcMilliSecond

	eventLockPayments, err := GetEventLockPaymentByPayloadCid(eventLockPayment.PayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(eventLockPayments) > 0 {
		eventLockPayment.ID = eventLockPayments[0].ID
	}

	db := database.GetDBTransaction()
	err = database.SaveOneInTransaction(db, &eventLockPayment)
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	sql := "update source_file set status=?,update_at=? where id=?"

	params := []interface{}{}
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_PAID)
	params = append(params, currentUtcMilliSecond)
	params = append(params, eventLockPayment.SourceFileId)

	err = db.Exec(sql, params...).Error
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	err = db.Commit().Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
