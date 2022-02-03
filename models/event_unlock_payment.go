package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
	"strconv"

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

func GetEventUnlockPaymentsByPayloadCidUserWallet(sourceFileId int64, userWallet string) (*EventUnlockPayment, error) {
	srcFiles, err := GetSourceFilesIn1CarBySourceFileId(sourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(srcFiles) == 0 {
		return nil, nil
	}

	var eventUnlockPayments []*EventUnlockPayment
	sql := "select payload_cid,unlock_to_user_address, sum(unlock_to_user_amount) unlock_to_user_amount,max(unlock_time) unlock_time,count(*) unlock_times "
	sql = sql + "from event_unlock_payment "
	sql = sql + "where payload_cid=? and unlock_to_user_address=? "
	sql = sql + "group by payload_cid,unlock_to_user_address"
	err = database.GetDB().Raw(sql, srcFiles[0].DealFilePayloadCid, userWallet).Scan(&eventUnlockPayments).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(eventUnlockPayments) == 0 {
		return nil, nil
	}

	totalFileSize := 0.0
	fileSize := 0.0
	for _, srcFile := range srcFiles {
		totalFileSize = totalFileSize + float64(srcFile.FileSize)
		if srcFile.ID == sourceFileId {
			fileSize = float64(srcFile.FileSize)
		}
	}

	percent := fileSize / totalFileSize

	eventUnlockPayment := eventUnlockPayments[0]

	unlockAmt, err := strconv.ParseFloat(eventUnlockPayment.UnlockToUserAmount, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	unlockAmt = unlockAmt * percent

	eventUnlockPayment.UnlockToUserAmount = strconv.FormatFloat(unlockAmt, 'f', 0, 64)

	return eventUnlockPayment, nil
}
