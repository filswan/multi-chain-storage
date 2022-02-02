package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type DealFile struct {
	ID                  int64           `json:"id"`
	CarFileName         string          `json:"car_file_name"`
	PayloadCid          string          `json:"payload_cid"`
	DealCid             string          `json:"deal_cid"`
	DealId              int64           `json:"deal_id"`
	PieceCid            string          `json:"piece_cid"`
	CarFileSize         int64           `json:"car_file_size"`
	MinerFid            string          `json:"miner_fid"`
	DealStatus          string          `json:"deal_status"`
	PinStatus           string          `json:"pin_status"`
	SourceFilePath      string          `json:"source_file_path"`
	CarFilePath         string          `json:"car_file_path"`
	CarMd5              string          `json:"car_md_5"`
	Duration            int             `json:"duration"`
	TaskUuid            string          `json:"task_uuid"`
	Cost                string          `json:"cost"`
	CreateAt            int64           `json:"create_at"`
	UpdateAt            int64           `json:"update_at"`
	DeleteAt            int64           `json:"delete_at"`
	LockPaymentTx       string          `json:"lock_payment_tx"`
	LockPaymentStatus   string          `json:"lock_payment_status"`
	LockPaymentNetwork  int64           `json:"lock_payment_network"`
	DaoSignStatus       string          `json:"dao_sign_status"`
	SendDealStatus      string          `json:"send_deal_status"`
	Verified            bool            `json:"verified"`
	ClientWalletAddress string          `json:"client_wallet_address"`
	IsDeleted           *bool           `json:"is_deleted"`
	MaxPrice            decimal.Decimal `json:"max_price"`
}

func GetDeal2Send() ([]*DealFile, error) {
	var dealFiles []*DealFile

	err := database.GetDB().Where("send_deal_status=? and lock_payment_status=? and task_uuid != ''", constants.DEAL_FILE_STATUS_CREATED, constants.LOCK_PAYMENT_STATUS_PROCESSING).Find(&dealFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return dealFiles, nil
}

func GetDealFileBySourceFilePayloadCid(srcFilePayloadCid string) ([]*DealFile, error) {
	sql := "select a.* from deal_file a, source_file_deal_file_map b, source_file c where c.payload_cid=? and c.id=b.source_file_id and b.deal_file_id=a.id"

	var dealFiles []*DealFile

	err := database.GetDB().Raw(sql, srcFilePayloadCid).Scan(&dealFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return dealFiles, nil
}

func UpdateDealFileLockPaymentStatus(id int64, lockPaymentStatus string) error {
	sql := "update deal_file set lock_payment_status=?,update_at=? where id=?"

	params := []interface{}{}
	params = append(params, lockPaymentStatus)
	params = append(params, utils.GetCurrentUtcMilliSecond())
	params = append(params, id)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
