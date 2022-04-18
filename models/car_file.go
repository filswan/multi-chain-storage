package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type CarFile struct {
	ID          int64           `json:"id"`
	CarFileName string          `json:"car_file_name"`
	PayloadCid  string          `json:"payload_cid"`
	PieceCid    string          `json:"piece_cid"`
	CarFileSize int64           `json:"car_file_size"`
	CarFilePath string          `json:"car_file_path"`
	Duration    int             `json:"duration"`
	TaskUuid    string          `json:"task_uuid"`
	MaxPrice    decimal.Decimal `json:"max_price"`
	Status      string          `json:"status"`
	CreateAt    int64           `json:"create_at"`
	UpdateAt    int64           `json:"update_at"`
}

func GetCarFileById(id int64) (*CarFile, error) {
	sql := "select a.* from car_file a where a.id=?"

	var carFiles []*CarFile

	err := database.GetDB().Raw(sql, id).Scan(&carFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(carFiles) > 0 {
		return carFiles[0], nil
	}

	return nil, nil
}

func GetCarFileByDealId(dealId int64) (*CarFile, error) {
	sql := "select a.* from car_file a, offline_deal b where a.id=b.deal_file_id and b.deal_id=?"

	var carFiles []*CarFile

	err := database.GetDB().Raw(sql, dealId).Scan(&carFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(carFiles) > 0 {
		return carFiles[0], nil
	}

	return nil, nil
}

func GetCarFilesByStatus(status string) ([]*CarFile, error) {
	sql := "select a.* from car_file a where a.status=?"
	var carFile []*CarFile

	err := database.GetDB().Raw(sql, status).Scan(&carFile).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return carFile, nil
}

func GetDeal2Send() ([]*CarFile, error) {
	var carFiles []*CarFile

	err := database.GetDB().Where("status=? and task_uuid != ''", constants.PROCESS_STATUS_TASK_CREATED).Find(&carFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return carFiles, nil
}

func GetCarFileBySourceFilePayloadCid(srcFilePayloadCid string) ([]*CarFile, error) {
	sql := "select a.* from deal_file a, source_file_deal_file_map b, source_file c where c.payload_cid=? and c.id=b.source_file_id and b.deal_file_id=a.id"

	var carFiles []*CarFile

	err := database.GetDB().Raw(sql, srcFilePayloadCid).Scan(&carFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return carFiles, nil
}

func UpdateDealFileStatus(id int64, status string) error {
	sql := "update deal_file set lock_payment_status=?,update_at=? where id=?"

	params := []interface{}{}
	params = append(params, status)
	params = append(params, utils.GetCurrentUtcSecond())
	params = append(params, id)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
