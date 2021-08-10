package models

import (
	"github.com/nebulaai/nbai-node/common"
	"math/big"
	"payment-bridge/database"
	"regexp"
)

type Wallet struct {
	ID            int64  `json:"id"`
	WalletAddress string `json:"wallet_address"`
	FirstTxID     string `json:"first_tx_id"`
}

func (self *Wallet) FindOneWallet(condition interface{}) (*Wallet, error) {
	db := database.GetDB()
	err := db.Where(condition).First(&self).Error
	return self, err
}

func FindWalletAddresseses(wallets *[]Wallet) (*[]Wallet, error) {
	db := database.GetDB()
	err := db.Find(&wallets).Error
	return wallets, err
}

type WalletWithBalance struct {
	WalletAddress string   `json:"wallet_address"`
	Balance       *big.Int `json:"balance"`
	Percentage    string   `json:"percentage"`
}

type WalletContainer struct {
	WalletWithBalances []*WalletWithBalance
}

type ByBalance []*WalletWithBalance

func (s ByBalance) Len() int {
	return len(s)
}

func (s ByBalance) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByBalance) Less(i, j int) bool {
	return s[i].Balance.Cmp(s[j].Balance) == 1
}

var FirstSlice WalletContainer = WalletContainer{[]*WalletWithBalance{}}
var SecondSlice WalletContainer = WalletContainer{[]*WalletWithBalance{}}

func GetFirstSlice() *WalletContainer {
	return &FirstSlice
}

func GetSecondSlice() *WalletContainer {
	return &SecondSlice
}

func (walletContainer *WalletContainer) StoreArray(walletWithBalance WalletWithBalance) {
	walletContainer.WalletWithBalances = append(walletContainer.WalletWithBalances, &walletWithBalance)
}

func IsValidAddress(iaddress interface{}) bool {

	//	todo
	//	1. common.IsHexAddress()
	//	2. start with "0x"
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
