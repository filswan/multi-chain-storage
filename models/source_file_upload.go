package models

import (
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
	PayloadCid            string `json:"payload_cid"`
	LockedFeeBeforeUnlock decimal.Decimal
	LockedFeeBeforeRefund decimal.Decimal
}

func GetSourceFileUploadsByCarFileId(carFileId int64) ([]*SourceFileUploadOut, error) {
	var sourceFileUploadOut []*SourceFileUploadOut
	sql := "select a.*,b.payload_cid from source_file_upload a\n" +
		"left join source_file b on a.source_file_id=b.id\n" +
		"where a.id in (select source_file_upload_id from car_file_source where car_file_id=?)"
	err := database.GetDB().Raw(sql, carFileId).Scan(&sourceFileUploadOut).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploadOut, nil
}

func GetSourceFileUploadsExpired() ([]*SourceFileUploadOut, error) {
	sql := "select a.*,b.payload_cid \n" +
		"from source_file_upload a, source_file b, transaction c\n" +
		"where a.file_type=? and a.source_file_id=b.id and a.id=c.source_file_upload_id and c.deadline<? and c.status=?"

	var models []*SourceFileUploadOut
	err := database.GetDB().Raw(sql).Scan(&models).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return models, nil
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
	Duration               int               `json:"duration"`
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

func GetSourceFileUploads(walletId int64, fileName, orderBy string, isAscend bool, limit, offset int) ([]*SourceFileUploadResult, *int, error) {
	sql := "select\n" +
		"a.id source_file_upload_id,d.id car_file_id,a.file_name,b.file_size,a.create_at upload_at,a.duration,\n" +
		"b.pin_status,d.payload_cid,a.status source_file_upload_status,d.status car_file_status,\n" +
		"e.token_id,e.mint_address,e.nft_tx_hash\n" +
		"from source_file_upload a\n" +
		"left join source_file b on a.source_file_id=b.id\n" +
		"left outer join car_file_source c on a.id=c.source_file_upload_id\n" +
		"left outer join car_file d on c.car_file_id=d.id\n" +
		"left outer join source_file_mint e on a.id=e.source_file_upload_id\n" +
		"where a.wallet_id=? and a.file_type=0"

	if !libutils.IsStrEmpty(&fileName) {
		sql = sql + " and a.file_name like '%" + fileName + "%' "
	}

	var sourceFileUploadResult []*SourceFileUploadResult

	err := database.GetDB().Raw(sql, walletId).Scan(&sourceFileUploadResult).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	switch strings.Trim(orderBy, " ") {
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
	start := (offset - 1) * limit
	end := start + limit
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
