package models

type SourceFileMint struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	NftTxHash          string `json:"nft_tx_hash"`
	MintAddress        string `json:"mint_address"`
	TokenId            string `json:"token_id"`
	CreateAt           int64  `json:"create_at"`
	UpdateAt           int64  `json:"update_at"`
}
