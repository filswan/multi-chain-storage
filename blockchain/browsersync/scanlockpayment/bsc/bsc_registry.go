package bsc

import (
	"multi-chain-storage/models"
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
