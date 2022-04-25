package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Token struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	NetworkId   int64  `json:"network_id"`
	Description string `json:"description"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}

func GetTokenById(id int64) (*Token, error) {
	var tokens []*Token
	err := database.GetDB().Where("id=?", id).Find(&tokens).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(tokens) > 0 {
		return tokens[0], nil
	}

	return nil, nil
}

func GetTokenByName(name string) (*Token, error) {
	var tokens []*Token
	err := database.GetDB().Where("name=?", name).Find(&tokens).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(tokens) > 0 {
		return tokens[0], nil
	}

	return nil, nil
}

func GetCoinByAddress(address string) (*Token, error) {
	var tokens []*Token
	err := database.GetDB().Where("address=?", address).Find(&tokens).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(tokens) > 0 {
		return tokens[0], nil
	}

	return nil, nil
}
