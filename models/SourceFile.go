package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type SourceFile struct {
	ID            int64  `json:"id"`
	FileName      string `json:"file_name"`
	ResourceUri   string `json:"resource_uri"`
	Status        string `json:"status"`
	FileSize      int64  `json:"file_size"`
	Dataset       string `json:"dataset"`
	CreateAt      int64  `json:"create_at"`
	IpfsUrl       string `json:"ipfs_url"`
	PinStatus     string `json:"pin_status"`
	WalletAddress string `json:"wallet_address"`
	PayloadCid    string `json:"payload_cid"`
	NftTxHash     string `json:"nft_tx_hash"`
	TokenId       string `json:"token_id"`
	MintAddress   string `json:"mint_address"`
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

func GetSourceFileById(id int64) (*SourceFile, error) {
	var sourceFile SourceFile

	err := database.GetDB().Where("id=?", id).First(&sourceFile).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &sourceFile, nil
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

func CreateSourceFile(sourceFile SourceFile) (*SourceFile, error) {
	value, err := database.SaveOneWithResult(&sourceFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	sourceFileCreated := value.(*SourceFile)

	return sourceFileCreated, nil
}

func GetSourceFiles(limit, offset string, walletAddress, payloadCid string) ([]*SourceFile, error) {
	sql := "select s.id, s.file_name,s.file_size,s.pin_status,s.create_at,s.payload_cid,df.lock_payment_status as status,df.duration, evpm.locked_fee from source_file s "
	sql = sql + "left join source_file_deal_file_map sfdfm on s.id = sfdfm.source_file_id "
	sql = sql + "left join deal_file df on sfdfm.deal_file_id = df.id "

	params := []interface{}{}

	if strings.Trim(payloadCid, " ") != "" {
		sql = sql + " and s.payload_cid=?"
		params = append(params, payloadCid)
	}

	sql = sql + "left outer join event_lock_payment evpm on evpm.payload_cid = s.payload_cid "
	sql = sql + "where wallet_address=?"
	params = append(params, walletAddress)

	var results []*SourceFile

	err := database.GetDB().Raw(sql, params...).Order("create_at desc").Limit(limit).Offset(offset).Scan(&results).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func GetSourceFilesCount(walletAddress string) (int64, error) {
	var count int64
	err := database.GetDB().Table("source_file").Where("wallet_address = ?", walletAddress).Count(&count).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return count, nil
}
