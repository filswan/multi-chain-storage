package goerli

import (
	"payment-bridge/models"
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
