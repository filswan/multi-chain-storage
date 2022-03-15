package polygon

import (
	"multi-chain-storage/models"
)

func Init() {
	models.Register(&models.RegInfo{
		Network:                           "Polygon",
		CoinName:                          "MATIC",
		ScanEventFromChainAndSaveDataToDb: ScanEventFromChainAndSaveDataToDb,
	})
}

func ScanEventFromChainAndSaveDataToDb() {
	ScanEventFromChainAndSaveDataToDbForPolygon()
}
