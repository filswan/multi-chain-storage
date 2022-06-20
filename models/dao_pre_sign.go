package models

type DaoPreSign struct {
	Id                       int64  `json:"id"`
	OfflineDealId            int64  `json:"offline_deal_id"`
	TxHash                   string `json:"tx_hash"`
	BatchNumber              int    `json:"batch_number"`
	BatchSizeMax             int    `json:"batch_size_max"`
	SourceFileUploadCntTotal int    `json:"source_file_upload_cnt_total"`
	SourceFileUploadCntSign  int    `json:"source_file_upload_cnt_sign"`
	CreateAt                 int64  `json:"create_at"`
	UpdateAt                 int64  `json:"update_at"`
}
