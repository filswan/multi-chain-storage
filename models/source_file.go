package models

import (
	"fmt"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/shopspring/decimal"
)

type SourceFile struct {
	ID          int64  `json:"id"`
	PayloadCid  string `json:"payload_cid"`
	ResourceUri string `json:"resource_uri"`
	IpfsUrl     string `json:"ipfs_url"`
	FileSize    int64  `json:"file_size"`
	Dataset     string `json:"dataset"`
	PinStatus   string `json:"pin_status"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

type SourceFileExt struct {
	SourceFile
	DealFilePayloadCid string           `json:"deal_file_payload_cid"`
	FileName           string           `json:"file_name"`
	DealFileId         int64            `json:"deal_file_id"`
	Duration           int              `json:"duration"`
	LockedFee          *decimal.Decimal `json:"locked_fee"`
	OfflineDeals       []*OfflineDeal   `json:"offline_deals"`
}

func GetSourceFileById(id int64) (*SourceFile, error) {
	var sourceFile SourceFile

	err := database.GetDB().Where("id=?", id).First(&sourceFile).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &sourceFile, nil
}

func GetSourceFileByPayloadCid(payloadCid string) (*SourceFile, error) {
	var sourceFiles []*SourceFile

	err := database.GetDB().Where("payload_cid=?", payloadCid).Find(&sourceFiles).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(sourceFiles) > 0 {
		return sourceFiles[0], nil
	}

	msg := fmt.Sprintf("source file with payload_cid:%s not exists", payloadCid)
	logs.GetLogger().Info(msg)

	return nil, nil
}

func CreateSourceFile(sourceFile *SourceFile) (*SourceFile, error) {
	value, err := database.SaveOneWithResult(sourceFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	sourceFileCreated := value.(*SourceFile)

	return sourceFileCreated, nil
}

func UpdateSourceFilePinStatus(sourceFileId int64, pinStatus string) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["pin_status"] = pinStatus
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(SourceFile{}).Where("id=?", sourceFileId).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
