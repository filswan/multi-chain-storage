package models

import (
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Wallet struct {
	ID       int64  `json:"id"`
	Address  string `json:"address"`
	CreateAt int64  `json:"create_at"`
}

func GetWalletByAddress(address string) (*Wallet, error) {
	var wallets []*Wallet
	err := database.GetDB().Where("address=?", address).Find(&wallets).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(wallets) > 0 {
		return wallets[0], nil
	}

	return nil, nil
}

func SaveWallet(address string) (*Wallet, error) {
	wallet := Wallet{
		Address:  address,
		CreateAt: utils.GetCurrentUtcSecond(),
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
