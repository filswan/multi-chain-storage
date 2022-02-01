package models

import (
	"fmt"
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

func FindCoinByUuid(coinUuid string) (*Coin, error) {
	var coins []*Coin
	err := database.GetDB().Where("coin_address=?", coinUuid).Find(&coins).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(coins) > 0 {
		return coins[0], nil
	}

	err = fmt.Errorf("coin:%s not exists", coinUuid)
	logs.GetLogger().Error(err)
	return nil, err
}

func FindCoinIdById(id int64) (*Coin, error) {
	var coins []*Coin
	err := database.GetDB().Where("id=?", id).Find(&coins).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(coins) > 0 {
		return coins[0], nil
	}

	return nil, nil
}

func FindCoinByCoinAddress(coinAddress string) (*Coin, error) {
	var coins []*Coin
	err := database.GetDB().Where("coin_address=?", coinAddress).Find(&coins).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(coins) > 0 {
		return coins[0], nil
	}

	err = fmt.Errorf("coin:%s not exists", coinAddress)
	logs.GetLogger().Error(err)
	return nil, err
}
