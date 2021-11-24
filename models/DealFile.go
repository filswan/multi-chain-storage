package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type DealFile struct {
	ID             int64  `json:"id"`
	CarFileName    string `json:"car_file_name"`
	PayloadCid     string `json:"payload_cid"`
	DealCid        string `json:"deal_cid"`
	PieceCid       string `json:"piece_cid"`
	CarFileSize    int64  `json:"car_file_size"`
	DealStatus     string `json:"deal_status"`
	SourceFilePath string `json:"source_file_path"`
	CarFilePath    string `json:"car_file_path"`
	CarMd5         string `json:"car_md_5"`
	CreateAt       string `json:"create_at"`
}

// FindDealFileList (&DealFile{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindDealFileList(whereCondition interface{}, orderCondition, limit, offset string) ([]*DealFile, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.PAGE_SIZE_DEFAULT_VALUE
	}
	var models []*DealFile
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
