package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type SystemConfigParam struct {
	ID         int64  `json:"id"`
	ParamKey   string `json:"param_key"`
	ParamValue string `json:"param_value"`
	Module     string `json:"module"`
}

func GetSystemConfigParams(limit, offset string) ([]*SystemConfigParam, error) {
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.PAGE_SIZE_DEFAULT_VALUE
	}

	var systemConfigParams []*SystemConfigParam
	err := database.GetDB().Offset(offset).Limit(limit).Order("id desc").Find(&systemConfigParams).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return systemConfigParams, err
}
