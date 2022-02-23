package models

import (
	"fmt"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Coin struct {
	ID          int64  `json:"id"`
	Name        string `json:"full_name"`
	Address     string `json:"coin_address"`
	NetworkId   int64  `json:"network_id"`
	GasPrice    int    `json:"gas_price"`
	GasLimit    int    `json:"gas_limit"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

func FindCoinByFullName(fullName string) (*Coin, error) {
	var coins []*Coin
	err := database.GetDB().Where("full_name=?", fullName).Find(&coins).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(coins) > 0 {
		return coins[0], nil
	}

	err = fmt.Errorf("coin:%s not exists", fullName)
	logs.GetLogger().Error(err)
	return nil, err
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
