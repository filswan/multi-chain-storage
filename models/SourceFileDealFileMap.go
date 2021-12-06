package models

import (
	"payment-bridge/database"
	"payment-bridge/logs"
)

type SourceFileDealFileMap struct {
	SourceFileId int64  `json:"source_file_id"`
	FileIndex    int    `json:"file_index"`
	DealFileId   int64  `json:"source_file_id"`
	CreateAt     string `json:"create_at"`
	UpdateAt     string `json:"update_at"`
}

//condition :&models.SourceFileDealFileMap{DealCid: "123"}
//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
func UpdateSourceFileDealFileMap(whereCondition interface{}, updateFields interface{}) error {
	db := database.GetDB()
	sdfMap := SourceFileDealFileMap{}
	err := db.Model(&sdfMap).Where(whereCondition).Update(updateFields).Error
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return err
}
