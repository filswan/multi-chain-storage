package models

import (
	"payment-bridge/common/constants"
	"payment-bridge/database"
)

type DaoInfo struct {
	ID          int64  `json:"id"`
	DaoName     string `json:"dao_name"`
	DaoAddress  string `json:"dao_address"`
	OrderIndex  string `json:"order_index"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
}

// FindDaoInfoList (&DaoInfo{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindDaoInfoList(whereCondition interface{}, orderCondition, limit, offset string) ([]*DaoInfo, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.PAGE_SIZE_DEFAULT_VALUE
	}
	var models []*DaoInfo
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
