package bsc

import (
	"payment-bridge/models"
)

func Init() {
	models.Register(&models.RegInfo{
		Network:                           "BSC",
		CoinName:                          "BNB",
		ScanEventFromChainAndSaveDataToDb: ScanEventFromChainAndSaveDataToDb,
	})
}

func ScanEventFromChainAndSaveDataToDb() {
	ScanEventFromChainAndSaveDataToDbForBsc()
}
