package bsc

import "payment-bridge/blockchain/browsersync"

func init() {
	browsersync.Register(&browsersync.RegInfo{
		Network:                           "BSC",
		CoinName:                          "BNB",
		ScanEventFromChainAndSaveDataToDb: ScanEventFromChainAndSaveDataToDb,
	})
}

func ScanEventFromChainAndSaveDataToDb() {
	ScanEventFromChainAndSaveDataToDbForBsc()
}
