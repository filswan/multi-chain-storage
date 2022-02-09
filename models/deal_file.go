package models

import (
	"fmt"
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
	PieceCid            string          `json:"piece_cid"`
	CarFileSize         int64           `json:"car_file_size"`
	PinStatus           string          `json:"pin_status"`
	CarFilePath         string          `json:"car_file_path"`
	CarMd5              string          `json:"car_md_5"`
	Duration            int             `json:"duration"`
	TaskUuid            string          `json:"task_uuid"`
	LockPaymentStatus   string          `json:"lock_payment_status"`
	ClientWalletAddress string          `json:"client_wallet_address"`
	MaxPrice            decimal.Decimal `json:"max_price"`
	CreateAt            int64           `json:"create_at"`
	UpdateAt            int64           `json:"update_at"`
}

func GetDealFileByDealId(dealId int64) (*DealFile, error) {
	sql := "select a.* from deal_file a, offline_deal b where a.id=b.deal_file_id and b.deal_id=?"

	var dealFiles []*DealFile

	err := database.GetDB().Raw(sql, dealId).Scan(&dealFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(dealFiles) > 0 {
		return dealFiles[0], nil
	}

	err = fmt.Errorf("deal with deal_id:%d not exists", dealId)
	logs.GetLogger().Error(err)
	return nil, err
}

func GetDealFilesByStatus(status string) ([]*DealFile, error) {
	sql := "select a.* from deal_file a where a.lock_payment_status=?"
	var dealFiles []*DealFile

	err := database.GetDB().Raw(sql, status).Scan(&dealFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return dealFiles, nil
}

func GetDeal2Send() ([]*DealFile, error) {
	var dealFiles []*DealFile

	err := database.GetDB().Where("lock_payment_status=? and task_uuid != ''", constants.PROCESS_STATUS_TASK_CREATED).Find(&dealFiles).Error

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

func UpdateDealFileStatus(id int64, status string) error {
	sql := "update deal_file set lock_payment_status=?,update_at=? where id=?"

	params := []interface{}{}
	params = append(params, status)
	params = append(params, utils.GetCurrentUtcMilliSecond())
	params = append(params, id)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
