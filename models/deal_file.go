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
	PieceCid            string          `json:"piece_cid"`
	CarFileSize         int64           `json:"car_file_size"`
	PinStatus           string          `json:"pin_status"`
	CarFilePath         string          `json:"car_file_path"`
	CarMd5              string          `json:"car_md_5"`
	Duration            int             `json:"duration"`
	TaskUuid            string          `json:"task_uuid"`
	Status              string          `json:"status"`
	ClientWalletAddress string          `json:"client_wallet_address"`
	MaxPrice            decimal.Decimal `json:"max_price"`
	CreateAt            int64           `json:"create_at"`
	UpdateAt            int64           `json:"update_at"`
}

func GetDeal2Send() ([]*DealFile, error) {
	var dealFiles []*DealFile

	err := database.GetDB().Where("status=? and task_uuid != ''", constants.PROCESS_STATUS_TASK_CREATED).Find(&dealFiles).Error

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
	sql := "update deal_file set status=?,update_at=? where id=?"

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
