package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type SystemConfigParam struct {
	ID         int64  `json:"id"`
	ParamKey   string `json:"param_key"`
	ParamValue string `json:"param_value"`
	Module     string `json:"module"`
	CreateAt   int64  `json:"create_at"`
}

// FindSystemConfigParam (&SystemConfigParam{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindSystemConfigParam(whereCondition interface{}, orderCondition, limit, offset string) ([]*SystemConfigParam, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.PAGE_SIZE_DEFAULT_VALUE
	}
	var models []*SystemConfigParam
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
