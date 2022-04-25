package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Coin struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	NetworkId   int64  `json:"network_id"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

func GetCoinById(id int64) (*Coin, error) {
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

func GetCoinByName(name string) (*Coin, error) {
	var coins []*Coin
	err := database.GetDB().Where("name=?", name).Find(&coins).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(coins) > 0 {
		return coins[0], nil
	}

	return nil, nil
}

func GetCoinByAddress(address string) (*Coin, error) {
	var coins []*Coin
	err := database.GetDB().Where("address=?", address).Find(&coins).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(coins) > 0 {
		return coins[0], nil
	}

	return nil, nil
}
