package nbai

import (
	"multi-chain-storage/models"
)

func Init() {
	models.Register(&models.RegInfo{
		Network:                           "Nbai",
		CoinName:                          "NBAI",
		ScanEventFromChainAndSaveDataToDb: ScanEventFromChainAndSaveDataToDb,
	})
}

func ScanEventFromChainAndSaveDataToDb() {
	ScanEventFromChainAndSaveDataToDbForNBAI()
}
