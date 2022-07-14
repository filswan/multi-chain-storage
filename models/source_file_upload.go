package models

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"
	"sort"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
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
	Duration     int    `json:"duration"`
	CreateAt     int64  `json:"create_at"`
	UpdateAt     int64  `json:"update_at"`
}

type SourceFileUploadOut struct {
	SourceFileUpload
	PayloadCid string `json:"payload_cid"`
}

func GetSourceFileUploads2RefundByCarFileId(carFileId int64) ([]*SourceFileUploadOut, error) {
	var sourceFileUploadOut []*SourceFileUploadOut
	sql := "select a.*,b.payload_cid from source_file_upload a\n" +
		"left join source_file b on a.source_file_id=b.id\n" +
		"where a.id in (select source_file_upload_id from car_file_source where car_file_id=?)\n" +
		" and a.status!=? and a.status!=?"
	err := database.GetDB().Raw(sql, carFileId, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED, constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS).Scan(&sourceFileUploadOut).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploadOut, nil
}

func GetSourceFileUploadsByCarFileId(carFileId int64, batchNo int) ([]*SourceFileUploadOut, error) {
	var sourceFileUploads []*SourceFileUploadOut
	sql := "select b.*,c.payload_cid from car_file_source a, source_file_upload b, source_file c\n" +
		"where a.car_file_id=? and a.source_file_upload_id=b.id and b.source_file_id=c.id\n" +
		"order by b. id limit ? offset ?"

	offset := constants.MAX_WCID_COUNT_IN_TRANSACTION * batchNo
	err := database.GetDB().Raw(sql, carFileId, constants.MAX_WCID_COUNT_IN_TRANSACTION, offset).Scan(&sourceFileUploads).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploads, nil
}

func GetSourceFileUploadsNotCompletedByCarFileId(carFileId int64) ([]*SourceFileUpload, error) {
	var sourceFileUploads []*SourceFileUpload
	sql := "select b.* from car_file_source a, source_file_upload b\n" +
		"where a.car_file_id=? and a.source_file_upload_id=b.id\n" +
		"  and b.status!=? and b.status!=?"
	err := database.GetDB().Raw(sql, carFileId, constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED).Scan(&sourceFileUploads).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploads, nil
}

func GetSourceFileUploads2BeRefundedAfterExpired() ([]*SourceFileUploadOut, error) {
	sql := "select a.*,b.payload_cid \n" +
		"from source_file_upload a, source_file b, transaction c\n" +
		"where a.file_type=? and a.source_file_id=b.id and a.id=c.source_file_upload_id and c.deadline<? and a.status not in (?,?,?,?)"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	var models []*SourceFileUploadOut
	params := []interface{}{}
	params = append(params, constants.SOURCE_FILE_TYPE_NORMAL)
	params = append(params, currentUtcSecond)
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS)
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE)
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED)
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_PENDING)
	err := database.GetDB().Raw(sql, params...).Scan(&models).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
}

func GetSourceFileUploads2BeRefundedByCarFileStatus(carFileStatus string) ([]*SourceFileUploadOut, error) {
	sql := "select a.*,d.payload_cid from source_file_upload a, car_file_source b, car_file c, source_file d\n" +
		"where a.id=b.source_file_upload_id and b.car_file_id=c.id and a.source_file_id=d.id\n" +
		"  and a.file_type=? and a.status not in (?,?) and c.status=?"
	var models []*SourceFileUploadOut
	params := []interface{}{}
	params = append(params, constants.SOURCE_FILE_TYPE_NORMAL)
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE)
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED)
	params = append(params, carFileStatus)
	err := database.GetDB().Raw(sql, params...).Scan(&models).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
}

func GetSourceFileUploadBySourceFileIdUuid(sourceFileId int64, uuid string) (*SourceFileUpload, error) {
	var sourceFileUploads []*SourceFileUpload
	err := database.GetDB().Where("source_file_id=? and uuid=?", sourceFileId, uuid).Find(&sourceFileUploads).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(sourceFileUploads) > 0 {
		return sourceFileUploads[0], nil
	}

	return nil, nil
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

type SourceFileUploadNeed2Car struct {
	SourceFileUploadId int64           `json:"source_file_upload_id"`
	ResourceUri        string          `json:"resource_uri"`
	IpfsUrl            string          `json:"ipfs_url"`
	FileSize           int64           `json:"file_size"`
	CreateAt           int64           `json:"create_at"`
	PayAmount          decimal.Decimal `json:"pay_amount"`
}

func GetSourceFileUploadsNeed2Car() ([]*SourceFileUploadNeed2Car, error) {
	var sourceFileUploadsNeed2Car []*SourceFileUploadNeed2Car
	sql := "select a.id source_file_upload_id,b.resource_uri,b.ipfs_url,b.file_size,a.create_at,c.pay_amount\n" +
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
	SourceFileUploadId int64             `json:"source_file_upload_id"`
	FileName           string            `json:"file_name"`
	FileSize           int64             `json:"file_size"`
	UploadAt           int64             `json:"upload_at"`
	Duration           int               `json:"duration"`
	IpfsUrl            string            `json:"ipfs_url"`
	PinStatus          string            `json:"pin_status"`
	PayAmount          string            `json:"pay_amount"`
	Status             string            `json:"status"`
	IsMinted           bool              `json:"is_minted"`
	TokenId            *string           `json:"token_id"`
	MintAddress        *string           `json:"mint_address"`
	NftTxHash          *string           `json:"nft_tx_hash"`
	OfflineDeals       []*OfflineDealOut `json:"offline_deal"`
}
type SourceFileUploadResultByFileName []*SourceFileUploadResult

func (a SourceFileUploadResultByFileName) Len() int           { return len(a) }
func (a SourceFileUploadResultByFileName) Less(i, j int) bool { return a[i].FileName < a[j].FileName }
func (a SourceFileUploadResultByFileName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type SourceFileUploadResultByFileSize []*SourceFileUploadResult

func (a SourceFileUploadResultByFileSize) Len() int           { return len(a) }
func (a SourceFileUploadResultByFileSize) Less(i, j int) bool { return a[i].FileSize < a[j].FileSize }
func (a SourceFileUploadResultByFileSize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type SourceFileUploadResultByUploadAt []*SourceFileUploadResult

func (a SourceFileUploadResultByUploadAt) Len() int           { return len(a) }
func (a SourceFileUploadResultByUploadAt) Less(i, j int) bool { return a[i].UploadAt < a[j].UploadAt }
func (a SourceFileUploadResultByUploadAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetSourceFileUploads(walletId int64, status, fileName, orderBy, is_minted *string, isAscend bool, limit, offset *int) ([]*SourceFileUploadResult, *int, error) {
	sql := "select\n" +
		"a.id source_file_upload_id,a.file_name,b.file_size,a.create_at upload_at,a.duration,\n" +
		"b.ipfs_url,b.pin_status,f.pay_amount,a.status,\n" +
		"e.id is not null is_minted,e.token_id,e.mint_address,e.nft_tx_hash\n" +
		"from source_file_upload a\n" +
		"left join source_file b on a.source_file_id=b.id\n" +
		"left outer join source_file_mint e on a.id=e.source_file_upload_id\n" +
		"left outer join transaction f on a.id=f.source_file_upload_id\n" +
		"where a.wallet_id=? and a.file_type=0"

	if fileName != nil {
		sql = sql + " and a.file_name like '%" + *fileName + "%'\n"
	}

	params := []interface{}{}
	params = append(params, walletId)

	if !libutils.IsStrEmpty(status) {
		switch strings.Trim(*status, " ") {
		case constants.SOURCE_FILE_UPLOAD_STATUS_PENDING,
			constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE,
			constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED,
			constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS:
			sql = sql + " and a.status=?"
			params = append(params, status)
		case constants.SOURCE_FILE_UPLOAD_STATUS_PROCESSING:
			sql = sql + " and a.status not in (?,?,?,?)"
			params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_PENDING)
			params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDABLE)
			params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED)
			params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS)
		default:
			logs.GetLogger().Info("input status:", status, ", get records with all kinds of statuses")
		}
	}

	if is_minted != nil {
		if strings.EqualFold(*is_minted, "y") {
			sql = sql + " and e.id is not null"
		} else if strings.EqualFold(*is_minted, "n") {
			sql = sql + " and e.id is null"
		}
	}

	var sourceFileUploadResult []*SourceFileUploadResult

	err := database.GetDB().Raw(sql, params...).Scan(&sourceFileUploadResult).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if orderBy == nil {
		orderByDefault := "upload_at"
		orderBy = &orderByDefault
	}

	switch strings.Trim(*orderBy, " ") {
	case "file_name":
		if isAscend {
			sort.Sort(SourceFileUploadResultByFileName(sourceFileUploadResult))
		} else {
			sort.Sort(sort.Reverse(SourceFileUploadResultByFileName(sourceFileUploadResult)))
		}
	case "file_size":
		if isAscend {
			sort.Sort(SourceFileUploadResultByFileSize(sourceFileUploadResult))
		} else {
			sort.Sort(sort.Reverse(SourceFileUploadResultByFileSize(sourceFileUploadResult)))
		}
	case "upload_at":
		if isAscend {
			sort.Sort(SourceFileUploadResultByUploadAt(sourceFileUploadResult))
		} else {
			sort.Sort(sort.Reverse(SourceFileUploadResultByUploadAt(sourceFileUploadResult)))
		}
	default:
		if isAscend {
			sort.Sort(SourceFileUploadResultByUploadAt(sourceFileUploadResult))
		} else {
			sort.Sort(sort.Reverse(SourceFileUploadResultByUploadAt(sourceFileUploadResult)))
		}
	}

	totalRecordCount := len(sourceFileUploadResult)

	if limit == nil || offset == nil {
		return sourceFileUploadResult, &totalRecordCount, nil
	}

	start := (*offset - 1) * *limit
	end := start + *limit
	if start >= totalRecordCount {
		return nil, &totalRecordCount, nil
	}

	if end >= totalRecordCount {
		end = totalRecordCount
	}

	result := sourceFileUploadResult[start:end]
	return result, &totalRecordCount, nil
}

func UpdateSourceFileUploadStatus(id int64, status string) error {
	sql := "update source_file_upload set status=?,update_at=? where id=?"

	params := []interface{}{}
	params = append(params, status)
	params = append(params, libutils.GetCurrentUtcSecond())
	params = append(params, id)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func GetSourceFileUploadByWCid(wCid string) (*SourceFile, *SourceFileUpload, error) {
	sourceFilePayloadCidIndex := strings.Index(wCid, "Qm")
	if sourceFilePayloadCidIndex < 0 {
		logs.GetLogger().Info("invalid wCid: ", wCid)
		return nil, nil, nil
	}

	sourceFileUploadUuid := wCid[0:sourceFilePayloadCidIndex]
	sourceFilePayloadCid := wCid[sourceFilePayloadCidIndex:]

	sourceFile, err := GetSourceFileByPayloadCid(sourceFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if sourceFile == nil {
		msg := fmt.Sprintf("no source file with payload cid:%s", sourceFilePayloadCid)
		logs.GetLogger().Info(msg)
		return nil, nil, nil
	}

	sourceFileUpload, err := GetSourceFileUploadBySourceFileIdUuid(sourceFile.ID, sourceFileUploadUuid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if sourceFileUpload == nil {
		msg := fmt.Sprintf("no source file upload with source file id:%d, uuid:%s", sourceFile.ID, sourceFileUploadUuid)
		logs.GetLogger().Info(msg)
		return nil, nil, nil
	}

	return sourceFile, sourceFileUpload, nil
}
