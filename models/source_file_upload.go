package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type SourceFileUpload struct {
	Id           int64  `json:"id"`
	SourceFileId int64  `json:"source_file_id"`
	FileType     int    `json:"file_type"`
	FileName     string `json:"file_name"`
	Uuid         string `json:"uuid"`
	WalletId     int64  `json:"wallet_id"`
	Status       string `json:"status"`
	CreateAt     int64  `json:"create_at"`
	UpdateAt     int64  `json:"update_at"`
}

func GetSourceFileUploadBySourceFileIdWallet(sourceFileId int64, walletId int64) ([]*SourceFileUpload, error) {
	var sourceFileUpload []*SourceFileUpload
	err := database.GetDB().Where("source_file_id=? and wallet_id=?", sourceFileId, walletId).Find(&sourceFileUpload).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUpload, nil
}

func GetSourceFileUploadById(id int64) (*SourceFileUpload, error) {
	var sourceFileUpload []*SourceFileUpload
	err := database.GetDB().Where("id=?", id).Find(&sourceFileUpload).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(sourceFileUpload) > 0 {
		return sourceFileUpload[0], nil
	}

	return nil, nil
}

func CreateSourceFileUpload(sourceFileUpload *SourceFileUpload) (*SourceFileUpload, error) {
	value, err := database.SaveOneWithResult(sourceFileUpload)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	sourceFileUploadCreated := value.(*SourceFileUpload)

	return sourceFileUploadCreated, nil
}

func GetSourceFileUploadsByFileTypeStatus(fileType int, status string) ([]*SourceFileUpload, error) {
	var sourceFileUploads []*SourceFileUpload

	err := database.GetDB().Where("file_type=? and status=?", fileType, status).Find(&sourceFileUploads).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploads, nil
}

type SourceFileUploadsNeed2Car struct {
	SourceFileUploadId int64           `json:"source_file_upload_id"`
	ResourceUri        string          `json:"resource_uri"`
	FileSize           int64           `json:"file_size"`
	CreateAt           int64           `json:"create_at"`
	LockedFee          decimal.Decimal `json:"locked_fee"`
}

func GetSourceFileUploadsNeed2Car() ([]*SourceFileUploadsNeed2Car, error) {
	var sourceFileUploadsNeed2Car []*SourceFileUploadsNeed2Car
	sql := "select a.id source_file_upload_id,b.resource_uri,b.file_size,a.create_at,c.amount locked_fee\n" +
		"from source_file_upload a, source_file b, transaction c\n" +
		"where a.file_type=? and a.status=? and a.source_file_id=b.id and a.id=c.source_file_upload_id"
	err := database.GetDB().Raw(sql, constants.SOURCE_FILE_TYPE_NORMAL, constants.SOURCE_FILE_UPLOAD_STATUS_PAID).Scan(&sourceFileUploadsNeed2Car).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploadsNeed2Car, nil
}

type SourceFileUploadResult struct {
	SourceFileUploadId     int64             `json:"source_file_upload_id"`
	CarFileId              int64             `json:"car_file_id"`
	FileName               string            `json:"file_name"`
	FileSize               int64             `json:"file_size"`
	UploadAt               int64             `json:"upload_at"`
	PinStatus              string            `json:"pin_status"`
	PayloadCid             string            `json:"payload_cid"`
	SourceFileUploadStatus string            `json:"source_file_upload_status"`
	CarFileStatus          string            `json:"car_file_status"`
	Status                 string            `json:"status"`
	TokenId                string            `json:"token_id"`
	MintAddress            string            `json:"mint_address"`
	NftTxHash              string            `json:"nft_tx_hash"`
	OfflineDeals           []*OfflineDealOut `json:"offline_deal"`
}

func GetSourceFileUploads(walletId int64, fileName *string, limit, offset int) ([]*SourceFileUploadResult, *int64, error) {
	offset = offset - 1
	filterOnSourceFileUpload := "wallet_id=? and file_type=0"
	if fileName != nil {
		filterOnSourceFileUpload = filterOnSourceFileUpload + " and file_name like '%" + *fileName + "%'"
	}
	sql := "select\n" +
		"a.id source_file_upload_id,d.id car_file_id,a.file_name,b.file_size,a.create_at upload_at,\n" +
		"b.pin_status,d.payload_cid,a.status source_file_upload_status,d.status car_file_status,\n" +
		"e.token_id,e.mint_address,e.nft_tx_hash\n" +
		"from\n" +
		"(\n" +
		"Select * from source_file_upload where " + filterOnSourceFileUpload + " order by create_at desc LIMIT ? OFFSET ?\n" +
		") a\n" +
		"left join source_file b on a.source_file_id=b.id\n" +
		"left outer join car_file_source c on a.id=c.source_file_upload_id\n" +
		"left outer join car_file d on c.car_file_id=d.id\n" +
		"left outer join source_file_mint e on a.id=e.source_file_upload_id\n"

	var sourceFileUploadResult []*SourceFileUploadResult

	err := database.GetDB().Raw(sql, walletId, limit, offset).Scan(&sourceFileUploadResult).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	var totalRecordCount int64
	err = database.GetDB().Table("source_file_upload").Where(filterOnSourceFileUpload, walletId).Count(&totalRecordCount).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return sourceFileUploadResult, &totalRecordCount, nil
}
