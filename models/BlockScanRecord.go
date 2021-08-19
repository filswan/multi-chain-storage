package models

import (
	"payment-bridge/common/utils"
	"payment-bridge/database"
)

type BlockScanRecord struct {
	ID                     int64  `json:"id"`
	NetworkType            string `json:"network_type"`
	LastCurrentBlockNumber int64  `json:"last_current_block_number"`
	UpdateAt               string `json:"update_at"`
}

func (self *BlockScanRecord) FindLastCurrentBlockNumber(whereCondition string) ([]*BlockScanRecord, error) {
	db := database.GetDB()
	var models []*BlockScanRecord
	err := db.Where(whereCondition).Find(&models).Error
	return models, err
}

//condition :&models.BlockScanRecord{"last_current_block_number": 18000}
//updateFields: map[string]interface{}{"update_at": taskT.ProcessingTime, "last_current_block_number": 18000}
func UpdateBlockScanRecord(whereCondition interface{}, updateFields interface{}) (BlockScanRecord, error) {
	db := database.GetDB()
	var record BlockScanRecord
	utils.GetEpochInMillis()
	err := db.Model(&record).Where(whereCondition).Update(updateFields).Error
	return record, err
}
