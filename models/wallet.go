package models

import (
	"multi-chain-storage/database"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"
)

type Wallet struct {
	ID       int64  `json:"id"`
	Type     int    `json:"type"`
	Address  string `json:"address"`
	IsDao    *bool  `json:"is_dao"`
	CreateAt int64  `json:"create_at"`
	UpdateAt int64  `json:"update_at"`
}

func GetWalletByAddress(address string, walletType int) (*Wallet, error) {
	var wallets []*Wallet
	err := database.GetDB().Where("address=? and type=?", address, walletType).Find(&wallets).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(wallets) > 0 {
		return wallets[0], nil
	}

	wallet, err := SaveWallet(address, walletType)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return wallet, nil
}

func SaveWallet(address string, walletType int) (*Wallet, error) {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	wallet := Wallet{
		Type:     walletType,
		Address:  address,
		CreateAt: currentUtcSecond,
		UpdateAt: currentUtcSecond,
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

func SetWalletAsDao(walletId int64) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["is_dao"] = true
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(Wallet{}).Where("id=?", walletId).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
