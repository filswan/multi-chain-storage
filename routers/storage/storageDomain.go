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
