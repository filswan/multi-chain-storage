package goerli

import (
	"multi-chain-storage/models"
)

func Init() {
	models.Register(&models.RegInfo{
		Network:                           "Goerli",
		CoinName:                          "ETH(Goerli network)",
		ScanEventFromChainAndSaveDataToDb: ScanEventFromChainAndSaveDataToDb,
	})
}

func ScanEventFromChainAndSaveDataToDb() {
	ScanEventFromChainAndSaveDataToDbForGoerli()
}
