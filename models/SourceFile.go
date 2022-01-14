package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type SourceFile struct {
	ID            int64           `json:"id"`
	FileName      string          `json:"file_name"`
	ResourceUri   string          `json:"resource_uri"`
	Status        string          `json:"status"`
	FileSize      string          `json:"file_size"`
	Dataset       string          `json:"dataset"`
	CreateAt      int64           `json:"create_at"`
	IpfsUrl       string          `json:"ipfs_url"`
	PinStatus     string          `json:"pin_status"`
	WalletAddress string          `json:"wallet_address"`
	PayloadCid    string          `json:"payload_cid"`
	MaxPrice      decimal.Decimal `json:"max_price"`
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

	err := database.GetDB().Where("payload_cid=?", payloadCid).Order("create_at").Find(&sourceFiles).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFiles, nil
}

func GetSourceFilesNeed2Car() ([]*SourceFile, error) {
	var sourceFiles []*SourceFile
	sql := "select a.* from source_file a,event_lock_payment b where a.status =? and a.payload_cid=b.payload_cid and a.wallet_address=b.address_from"
	err := database.GetDB().Raw(sql, constants.SOURCE_FILE_STATUS_CREATED).Scan(&sourceFiles).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFiles, nil
}

func UpdateSourceFileMaxPrice(id int64, maxPrice decimal.Decimal) error {
	sql := "update source_file set max_price=? where id=?"

	params := []interface{}{}
	params = append(params, maxPrice)
	params = append(params, id)

	err := database.GetDB().Exec(sql, params...).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
