package models

type Transaction struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	Type               int    `json:"type"`
	TxHash             string `json:"tx_hash"`
	WalletIdFrom       int64  `json:"wallet_id_from"`
	WalletIdTo         int64  `json:"wallet_id_to"`
	CoinId             int64  `json:"coin_id"`
	Amount             string `json:"amount"`
	TransactionAt      int64  `json:"transaction_at"`
	BlockNumber        int64  `json:"block_number"`
	CreateAt           int64  `json:"create_at"`
}
