package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SourceFileDealFileMap struct {
	SourceFileId int64 `json:"source_file_id"`
	FileIndex    int   `json:"file_index"`
	DealFileId   int64 `json:"deal_file_id"`
	CreateAt     int64 `json:"create_at"`
	UpdateAt     int64 `json:"update_at"`
}

func GetSourceFileDealFileMapBySourceFilePayloadCid(sourceFilePayloadCid string) ([]*SourceFileDealFileMap, error) {
	var sourceFileDealFileMap []*SourceFileDealFileMap
	sql := "select a.* from source_file_deal_file_map a, source_file b where a.source_file_id=b.id and b.payload_cid=?"
	err := database.GetDB().Raw(sql, sourceFilePayloadCid).Scan(&sourceFileDealFileMap).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileDealFileMap, nil
}
