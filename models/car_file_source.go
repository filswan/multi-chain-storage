package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type CarFileSource struct {
	Id                 int   `json:"id"`
	CarFileId          int64 `json:"car_file_id"`
	SourceFileUploadId int64 `json:"source_file_upload_id"`
	CreateAt           int64 `json:"create_at"`
}

func GetCarFileSourceBySourceFileUploadId(sourceFileUploadId int64) (*CarFileSource, error) {
	var carFileSources []*CarFileSource
	err := database.GetDB().Where("source_file_upload_id=?", sourceFileUploadId).Scan(&carFileSources).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(carFileSources) > 0 {
		return carFileSources[0], nil
	}

	return nil, nil
}
