package billing

type BillingResult struct {
	TxHash      string `json:"tx_hash"`
	LockedFee   string `json:"locked_fee"`
	Deadline    string `json:"deadline"`
	PayloadCid  string `json:"payload_cid"`
	AddressFrom string `json:"address_from"`
	Network     string `json:"network"`
	CreateAt    string `json:"create_at"`
}

type BillingRequest struct {
	TxHash        string `json:"tx_hash"`
	WalletAddress string `json:"wallet_address"`
	PageNumber    string `json:"page_number"`
	PageSize      string `json:"page_size"`
}

type RecordCount struct {
	TotalRecord int64 `json:"total_record"`
}

type PriceResult struct {
	Filecoin struct {
		Usd float64 `json:"usd"`
	} `json:"filecoin"`
}
