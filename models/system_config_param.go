package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SystemConfigParam struct {
	ID         int64  `json:"id"`
	ParamKey   string `json:"param_key"`
	ParamValue string `json:"param_value"`
	Module     string `json:"module"`
}

func GetSystemConfigParams() ([]*SystemConfigParam, error) {
	var systemConfigParams []*SystemConfigParam
	err := database.GetDB().Find(&systemConfigParams).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return systemConfigParams, err
}
