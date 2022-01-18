package models

type OfflineDeal struct {
	Id           int64  `json:"id"`
	DealFileId   int64  `json:"deal_file_id"`
	DealCid      string `json:"deal_cid"`
	MinerFid     string `json:"miner_fid"`
	StartEpoch   int    `json:"start_epoch"`
	SenderWallet string `json:"sender_wallet"`
}
