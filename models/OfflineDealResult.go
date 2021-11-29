package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type OfflineDealResult struct {
	Data       DataStruct `json:"data"`
	PagingInfo PagingInfo `json:"paging_info"`
	Status     string     `json:"status"`
}

type DataStruct struct {
	Deals []*OfflineDeal `json:"deals"`
}

type PagingInfo struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalItems int `json:"total_items"`
}

type OfflineDeal struct {
	ID            int64       `json:"id"`
	UUID          string      `json:"uuid"`
	UserID        int64       `json:"user_id"`
	TaskName      string      `json:"task_name"`
	TaskId        int64       `json:"task_id"`
	DealCid       string      `json:"deal_cid"`
	FileName      interface{} `json:"file_name"`
	FilePath      string      `json:"file_path"`
	FileSize      string      `json:"file_size"`
	FileSourceURL string      `json:"file_source_url"`
	Md5Origin     string      `json:"md5_origin"`
	MinerID       int         `json:"miner_id"`
	MinerFid      string      `json:"miner_fid"`
	Note          string      `json:"note"`
	PayloadCid    string      `json:"payload_cid"`
	PieceCid      string      `json:"piece_cid"`
	PinStatus     string      `json:"pin_status"`
	StartEpoch    int64       `json:"start_epoch"`
	TaskStatus    string      `json:"task_status"`
	Status        string      `json:"status"`
	Cost          string      `json:"cost"`
	PayStatus     string      `json:"pay_status"`
	UpdatedAt     string      `json:"updated_at"`
	CreatedAt     string      `json:"created_at"`
}

// FindOfflineDeals (&OfflineDeal{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindOfflineDeals(whereCondition interface{}, orderCondition, limit, offset string) ([]*OfflineDeal, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*OfflineDeal
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
