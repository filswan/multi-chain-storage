package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
	"payment-bridge/logs"
)

type DealFile struct {
	ID                 int64  `json:"id"`
	CarFileName        string `json:"car_file_name"`
	PayloadCid         string `json:"payload_cid"`
	DealCid            string `json:"deal_cid"`
	DealId             int64  `json:"deal_id"`
	PieceCid           string `json:"piece_cid"`
	CarFileSize        int64  `json:"car_file_size"`
	MinerFid           string `json:"miner_fid"`
	DealStatus         string `json:"deal_status"`
	PinStatus          string `json:"pin_status"`
	SourceFilePath     string `json:"source_file_path"`
	CarFilePath        string `json:"car_file_path"`
	CarMd5             string `json:"car_md_5"`
	Cost               string `json:"cost"`
	CreateAt           string `json:"create_at"`
	UpdateAt           string `json:"update_at"`
	DeleteAt           string `json:"delete_at"`
	LockPaymentTx      string `json:"lock_payment_tx"`
	LockPaymentStatus  string `json:"lock_payment_status"`
	LockPaymentNetwork string `json:"lock_payment_network"`
	DaoSignStatus      string `json:"dao_sign_status"`
}

// FindDealFileList (&DealFile{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindDealFileList(whereCondition interface{}, orderCondition, limit, offset string) ([]*DealFile, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.PAGE_SIZE_DEFAULT_VALUE
	}
	var models []*DealFile
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}

//condition :&models.DealFile{DealCid: "123"}
//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateDealFile(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	hardware := DealFile{}
	err := db.Model(&hardware).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}
