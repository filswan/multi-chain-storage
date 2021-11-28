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
	DealId     int64  `json:"deal_id"`
	PieceCid   string `json:"piece_cid"`
}

type filinkParams struct {
	ID   int `json:"id"`
	Data struct {
		Deal int `json:"deal"`
	} `json:"data"`
}
type DealOnChainResult struct {
	JobRunID int `json:"jobRunID"`
	Data     struct {
		Status string `json:"status"`
		Data   struct {
			Deal struct {
				DealID                   int    `json:"deal_id"`
				DealCid                  string `json:"deal_cid"`
				MessageCid               string `json:"message_cid"`
				Height                   int    `json:"height"`
				PieceCid                 string `json:"piece_cid"`
				VerifiedDeal             bool   `json:"verified_deal"`
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
				StoragePrice             int    `json:"storage_price"`
			} `json:"deal"`
		} `json:"data"`
		Result struct {
		} `json:"result"`
	} `json:"data"`
	Result struct {
	} `json:"result"`
	StatusCode int `json:"statusCode"`
}

type DaoDealResult struct {
	DealID     int64  `json:"deal_id"`
	DealCid    string `json:"deal_cid"`
	PayloadCid string `json:"payload_cid"`
}

type DaoSignResult struct {
	DealID     int64  `json:"deal_id"`
	TxHash     string `json:"tx_hash"`
	DaoAddress string `json:"dao_address"`
	DealCid    string `json:"deal_cid"`
	PayloadCid string `json:"payload_cid"`
	Status     bool   `json:"status"`
}
