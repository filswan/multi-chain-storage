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

type SourceFileAndDealFileInfoExtend struct {
	ID          int64  `json:"id"`
	FileName    string `json:"file_name"`
	FileSize    string `json:"file_size"`
	CreateAt    string `json:"create_at"`
	MinerFid    string `json:"miner_fid"`
	DealStatus  string `json:"deal_status"`
	Status      string `json:"status"`
	PinStatus   string `json:"pin_status"`
	PayloadCid  string `json:"payload_cid"`
	DealCid     string `json:"deal_cid"`
	DealId      int64  `json:"deal_id"`
	PieceCid    string `json:"piece_cid"`
	Duration    int    `json:"duration"`
	LockedFee   string `json:"locked_fee"`
	NftTxHash   string `json:"nft_tx_hash"`
	TokenId     string `json:"token_id"`
	MintAddress string `json:"mint_address"`
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

type DealForDaoSignResult struct {
	DealFileId            int64    `json:"deal_file_id"`
	PayloadCid            string   `json:"payload_cid"`
	DealCid               string   `json:"deal_cid"`
	DealId                int64    `json:"deal_id"`
	PieceCid              string   `json:"piece_cid"`
	MinerFid              string   `json:"miner_fid"`
	Duration              int      `json:"duration"`
	Cost                  string   `json:"cost"`
	CreateAt              string   `json:"create_at"`
	Verified              bool     `json:"verified"`
	ClientWalletAddress   string   `json:"client_wallet_address"`
	SourceFilePayloadCids []string `json:"payload_cids_source"`
}

type DealIdList struct {
	DealId     string `json:"deal_id"`
	PayloadCid string `json:"payload_cid"`
	Recipent   string `json:"recipent"`
	TxHash1    string `json:"tx_hash_1"`
	TxHash2    string `json:"tx_hash_2"`
	TxHash3    string `json:"tx_hash_3"`
}

type IpfsReturn struct {
	Name string `json:"Name"`
	Hash string `json:"Hash"`
	Size string `json:"Size"`
}
