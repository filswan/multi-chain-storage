package billing

type BillingResult struct {
	TxHash              string `json:"tx_hash"`
	LockedFee           string `json:"locked_fee"`
	Deadline            string `json:"deadline"`
	PayloadCid          string `json:"payload_cid"`
	AddressFrom         string `json:"address_from"`
	Network             string `json:"network"`
	CoinType            string `json:"coin_type"`
	LockPaymentTime     string `json:"lock_payment_time"`
	CreateAt            string `json:"create_at"`
	UnlockToUserAddress string `json:"unlock_to_user_address"`
	UnlockToUserAmount  string `json:"unlock_to_user_amount"`
	UnlockTime          string `json:"unlock_time"`
	FileName            string `json:"file_name"`
}

type BillingRequest struct {
	TxHash        string `json:"tx_hash"`
	WalletAddress string `json:"wallet_address"`
	PageNumber    string `json:"page_number"`
	PageSize      string `json:"page_size"`
}

type PriceResult struct {
	Filecoin struct {
		Usd float64 `json:"usd"`
	} `json:"filecoin"`
}
