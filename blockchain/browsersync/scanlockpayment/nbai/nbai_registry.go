package nbai

import (
	"payment-bridge/models"
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
