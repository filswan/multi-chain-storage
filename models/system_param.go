package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SystemParam struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

func GetSystemParams() ([]*SystemParam, error) {
	var systemParams []*SystemParam
	err := database.GetDB().Find(&systemParams).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return systemParams, err
}
