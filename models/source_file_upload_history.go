package models

type SourceFileUploadHistory struct {
	Id            int64  `json:"id"`
	SourceFileId  int64  `json:"source_file_id"`
	FileName      string `json:"file_name"`
	WalletAddress string `json:"wallet_address"`
	Status        string `json:"status"`
	CreateAt      int64  `json:"create_at"`
	UpdateAt      int64  `json:"update_at"`
}
