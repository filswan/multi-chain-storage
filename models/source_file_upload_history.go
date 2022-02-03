package models

import (
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SourceFileUploadHistory struct {
	Id            int64  `json:"id"`
	SourceFileId  int64  `json:"source_file_id"`
	FileName      string `json:"file_name"`
	WalletAddress string `json:"wallet_address"`
	Status        string `json:"status"`
	CreateAt      int64  `json:"create_at"`
	UpdateAt      int64  `json:"update_at"`
}

func GetSourceFileUploadHistoryBySourceFileIdWallet(sourceFileId int64, walletAddress string) ([]*SourceFileUploadHistory, error) {
	var sourceFileUploadHistory []*SourceFileUploadHistory
	sql := "select * from source_file_upload_history where source_file_id=? and wallet_address=?"
	err := database.GetDB().Raw(sql, sourceFileId, walletAddress).Scan(&sourceFileUploadHistory).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sourceFileUploadHistory, nil
}
