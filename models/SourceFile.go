package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SourceFile struct {
	ID            int64  `json:"id"`
	FileName      string `json:"file_name"`
	ResourceUri   string `json:"resource_uri"`
	FileSize      string `json:"file_size"`
	Dataset       string `json:"dataset"`
	CreateAt      string `json:"create_at"`
	IpfsUrl       string `json:"ipfs_url"`
	PinStatus     string `json:"pin_status"`
	WalletAddress string `json:"wallet_address"`
	PayloadCid    string `json:"payload_cid"`
}

// FindSourceFileList (&SourceFile{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindSourceFileList(whereCondition interface{}, orderCondition, limit, offset string) ([]*SourceFile, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.PAGE_SIZE_DEFAULT_VALUE
	}
	var models []*SourceFile
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

func GetSourceFilesByPayloadCid(payloadCid string) ([]*SourceFile, error) {
	var sourceFiles []*SourceFile

	err := database.GetDB().Where("payload_cid=?").Find(&sourceFiles).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFiles, nil
}
