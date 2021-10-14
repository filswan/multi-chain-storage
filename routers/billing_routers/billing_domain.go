package billing_routers

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
	AddressType  string `json:"address_type"`
	AddressValue string `json:"address_value""`
	PageNumber   string `json:"page_number"`
	PageSize     string `json:"page_size"`
}

type RecordCount struct {
	TotalRecord int64 `json:"total_record"`
}
