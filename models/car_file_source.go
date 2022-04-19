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

func GetSourceFileDealFileMapBySourceFilePayloadCid(sourceFilePayloadCid string) ([]*CarFileSource, error) {
	var carFileSource []*CarFileSource
	sql := "select a.* from car_file_source a, source_file b where a.source_file_id=b.id and b.payload_cid=?"
	err := database.GetDB().Raw(sql, sourceFilePayloadCid).Scan(&carFileSource).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return carFileSource, nil
}
