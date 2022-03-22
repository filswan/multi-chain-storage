package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"
	"multi-chain-storage/logs"
)

type SourceFile struct {
	ID          int64  `json:"id"`
	FileName    string `json:"file_name"`
	ResourceUri string `json:"resource_uri"`
	FileSize    string `json:"file_size"`
	Uuid        string `json:"uuid"`
	Dataset     string `json:"dataset"`
	CreateAt    string `json:"create_at"`
	//UserId        int    `json:"user_id"`
	IpfsUrl       string `json:"ipfs_url"`
	PinStatus     string `json:"pin_status"`
	WalletAddress string `json:"wallet_address"`
	NftTxHash     string `json:"nft_tx_hash"`
	TokenId       string `json:"token_id"`
	MintAddress   string `json:"mint_address"`
	FileType      string `json:"file_type"`
}

func GetSourceFilesByDealFileId(dealFileId int64) ([]*SourceFile, error) {
	var sourceFiles []*SourceFile

	sql := "select a.* from source_file a, source_file_deal_file_map b, deal_file c where c.id=? and c.id=b.deal_file_id and b.source_file_id=a.id"

	err := database.GetDB().Raw(sql, dealFileId).First(&sourceFiles).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFiles, nil
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

func FindSourceFileByPayloadCid(payloadCid string) (*SourceFile, error) {
	db := database.GetDB()
	sql := "select s.id, s.file_name, s.resource_uri, s.file_size, s.uuid, s.dataset, s.create_at, s.ipfs_url, s.pin_status, s.wallet_address from source_file s " +
		" inner join source_file_deal_file_map sfdfm on s.id = sfdfm.source_file_id" +
		" inner join deal_file df on sfdfm.deal_file_id = df.id and df.payload_cid='" + payloadCid + "' "
	var results []*SourceFile
	err := db.Raw(sql).Order("create_at desc").Scan(&results).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if len(results) > 0 {
		return results[0], nil
	} else {
		return nil, nil
	}
}
