package storage

type SourceFileAndDealFileInfo struct {
	ID                int64  `json:"id"`
	WalletAddress     string `json:"wallet_address"`
	FileName          string `json:"file_name"`
	FileSize          string `json:"file_size"`
	CreateAt          string `json:"create_at"`
	MinerFid          string `json:"miner_fid"`
	DealStatus        string `json:"deal_status"`
	Status            string `json:"status"`
	PinStatus         string `json:"pin_status"`
	PayloadCid        string `json:"payload_cid"`
	DealCid           string `json:"deal_cid"`
	DealId            int64  `json:"deal_id"`
	IpfsUrl           string `json:"ipfs_url"`
	PieceCid          string `json:"piece_cid"`
	Duration          int    `json:"duration"`
	LockPaymentStatus string `json:"lock_payment_status"`
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
				IpfsUrl                  string `json:"ipfs_url"`
				FileName                 string `json:"file_name"`
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

type LockFound struct {
	PayloadCid          string `json:"payload_cid"`
	CreateAt            string `json:"create_at"`
	LockedFee           string `json:"locked_fee"`
}

type DealForDaoSignResult struct {
	PayloadCid          string `json:"payload_cid"`
	DealCid             string `json:"deal_cid"`
	DealId              int64  `json:"deal_id"`
	PieceCid            string `json:"piece_cid"`
	MinerFid            string `json:"miner_fid"`
	Duration            int    `json:"duration"`
	Cost                string `json:"cost"`
	CreateAt            string `json:"create_at"`
	Verified            bool   `json:"verified"`
	ClientWalletAddress string `json:"client_wallet_address"`
}

type DealIdList struct {
	DealIdList string `json:"deal_id_list"`
}

type IpfsReturn struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}

type DaoInfoResult struct {
	DaoName         string `json:"dao_name"`
	DaoAddress      string `json:"dao_address"`
	OrderIndex      string `json:"order_index"`
	DealId          int64  `json:"deal_id"`
	DaoPassTime     string `json:"dao_pass_time"`
	PayloadCid      string `json:"payload_cid"`
	DaoAddressEvent string `json:"dao_address_event"`
	TxHash          string `json:"tx_hash"`
	Status          string `json:"status"`
}
