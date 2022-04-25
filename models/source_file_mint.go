package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SourceFileMint struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	NftTxHash          string `json:"nft_tx_hash"`
	MintAddress        string `json:"mint_address"`
	TokenId            string `json:"token_id"`
	CreateAt           int64  `json:"create_at"`
	UpdateAt           int64  `json:"update_at"`
}

func GetSourceFileMintBySourceFileUploadId(sourceFileUploadId int64) (*SourceFileMint, error) {
	var sSourceFileMints []*SourceFileMint
	err := database.GetDB().Where("source_file_upload_id=?", sourceFileUploadId).Find(&sSourceFileMints).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(sSourceFileMints) > 0 {
		return sSourceFileMints[0], nil
	}

	return nil, nil
}
