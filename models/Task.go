package models

import (
	"payment-bridge/database"
)

type Task struct {
	ID             int64  `json:"id"`
	Address        string `json:"address"`
	CompleteStatus int    `json:"complete_status"`
}

func (self *Task) FindOneTask(condition interface{}) (*Task, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

func UpdateCompleteStatus(contractAddress string, completeStatus int) error {
	db := database.GetDB()
	task := Task{}
	err := db.Model(&task).Where(&Task{Address: contractAddress}).Update("CompleteStatus", completeStatus).Error
	return err
}

func FindTasksByCompleteStatus(completeStatus int) ([]Task, error) {
	db := database.GetDB()
	var tasks []Task
	err := db.Order("id desc").Where(&Task{CompleteStatus: completeStatus}).Find(&tasks).Error
	return tasks, err
	//completestatus: 1 = task created; 2 = task completed; 3 = task cancelled
}
