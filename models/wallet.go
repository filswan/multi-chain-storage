package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/filswan/go-swan-lib/utils"
)

type Wallet struct {
	ID       int64  `json:"id"`
	Address  string `json:"address"`
	Type     int    `json:"type"`
	CreateAt int64  `json:"create_at"`
}

func GetWalletByAddressType(address string, walletType int) (*Wallet, error) {
	var wallets []*Wallet
	err := database.GetDB().Where("address=? and type=?", address, walletType).Find(&wallets).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(wallets) > 0 {
		return wallets[0], nil
	}

	return nil, nil
}

func GetWalletsByType(walletType int) ([]*Wallet, error) {
	var wallets []*Wallet
	err := database.GetDB().Where("type=?", walletType).Find(&wallets).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return wallets, nil
}

func SaveWallet(address string, walletType int) (*Wallet, error) {
	wallet := Wallet{
		Address:  address,
		Type:     walletType,
		CreateAt: utils.GetEpochInMillis(),
	}

	walletResult := database.GetDB().Create(&wallet)
	err := walletResult.Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	walletCreated := walletResult.Value.(*Wallet)

	return walletCreated, nil
}
