package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SystemConf struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

func GetSystemConf() ([]*SystemConf, error) {
	var systemConf []*SystemConf
	err := database.GetDB().Find(&systemConf).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return systemConf, err
}
