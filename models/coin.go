package models

import (
	"errors"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Coin struct {
	ID          int64  `json:"id"`
	ShortName   string `json:"network_name"`
	FullName    string `json:"full_name"`
	CnName      string `json:"cn_name"`
	CoinAddress string `json:"coin_address"`
	NetworkId   int64  `json:"network_id"`
	GasPrice    int    `json:"gas_price"`
	GasLimit    int    `json:"gas_limit"`
	Description string `json:"description"`
}

func FindCoinIdByUUID(uuid string) (int64, error) {
	db := database.GetDB()
	var models []*Coin
	err := db.Where("uuid='" + uuid + "'").Find(&models).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, nil
	}
	if len(models) > 0 {
		return models[0].ID, nil
	} else {
		err = errors.New("There is no this coin info by uuid :" + uuid)
		return 0, err
	}
}
