package storage

type SourceFileAndDealFileInfo struct {
	FileName   string `json:"file_name"`
	FileSize   string `json:"file_size"`
	CreateAt   string `json:"create_at"`
	MinerFid   string `json:"miner_fid"`
	DealStatus string `json:"deal_status"`
	Status     string `json:"status"`
	PinStatus  string `json:"pin_status"`
	PayloadCid string `json:"payload_cid"`
	DealCid    string `json:"deal_cid"`
	PieceCid   string `json:"piece_cid"`
}

type filinkParams struct {
	ID   int `json:"id"`
	Data struct {
		Deal int `json:"deal"`
	} `json:"data"`
}

type DealOnChainResult struct {
	Status string `json:"status"`
	Data   struct {
		Deal DealOnChainIno `json:"deal"`
	} `json:"data"`
}

type DealOnChainIno struct {
	DealID                   int    `json:"deal_id"`
	DealCid                  string `json:"deal_cid"`
	MessageCid               string `json:"message_cid"`
	Height                   int    `json:"height"`
	PieceCid                 string `json:"piece_cid"`
	VerifiedDeal             bool   `json:"verified_deal"`
	StoragePrice             int64  `json:"storage_price"`
	StoragePricePerEpoch     int    `json:"storage_price_per_epoch"`
	Signature                string `json:"signature"`
	SignatureType            string `json:"signature_type"`
	CreatedAt                int    `json:"created_at"`
	PieceSizeFormat          string `json:"piece_size_format"`
	StartHeight              int    `json:"start_height"`
	EndHeight                int    `json:"end_height"`
	Client                   string `json:"client"`
	ClientCollateralFormat   string `json:"client_collateral_format"`
	Provider                 string `json:"provider"`
	ProviderTag              string `json:"provider_tag"`
	VerifiedProvider         int    `json:"verified_provider"`
	ProviderCollateralFormat string `json:"provider_collateral_format"`
	Status                   int    `json:"status"`
	NetworkName              string `json:"network_name"`
}
