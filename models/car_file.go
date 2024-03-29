package models

import (
	"multi-chain-storage/database"

	libutils "github.com/filswan/go-swan-lib/utils"

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
	IsFree      bool            `json:"is_free"`
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
	var carFiles []*CarFile

	err := database.GetDB().Where("status=?", status).Find(&carFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return carFiles, nil
}

type CarFile2Refund struct {
	CarFileId     int64 `json:"car_file_id"`
	DealSuccesCnt int   `json:"success_cnt"`
	DealFailCnt   int   `json:"fail_cnt"`
	DealTotalCnt  int   `json:"total_cnt"`
	LastUnlockAt  int64 `json:"last_unlock_at"`
}

func GetCarFileBySourceFileUploadId(srcFileUploadId int64) (*CarFile, error) {
	sql := "select a.* from car_file a, car_file_source b, source_file_upload c where c.id=? and c.id=b.source_file_upload_id and b.car_file_id=a.id"

	var carFiles []*CarFile

	err := database.GetDB().Raw(sql, srcFileUploadId).Scan(&carFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(carFiles) > 0 {
		return carFiles[0], nil
	}

	return nil, nil
}

func UpdateCarFileStatus(id int64, status string) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["status"] = status
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(CarFile{}).Where("id=?", id).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
