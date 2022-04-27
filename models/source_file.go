package models

import (
	"fmt"
	"multi-chain-storage/database"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"
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

func UpdateSourceFileRefundAmount(srcFileId int64, refundAmount decimal.Decimal) error {
	sql := "update source_file set refund_amount=?,update_at=? where id=?"

	curUtcMilliSec := libutils.GetCurrentUtcSecond()

	params := []interface{}{}
	params = append(params, refundAmount)
	params = append(params, curUtcMilliSec)
	params = append(params, srcFileId)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func UpdateSourceFileRefundStatus(srcFileId int64, refundStatus string, refundTxHash string) error {
	sql := "update source_file set refund_status=?,refund_tx_hash=?,refund_at=?,update_at=? where id=?"

	curUtcMilliSec := libutils.GetCurrentUtcSecond()

	params := []interface{}{}
	params = append(params, refundStatus)
	params = append(params, refundTxHash)
	params = append(params, curUtcMilliSec)
	params = append(params, curUtcMilliSec)
	params = append(params, srcFileId)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
