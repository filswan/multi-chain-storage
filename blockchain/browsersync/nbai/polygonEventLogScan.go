package nbai

import (
	"math/big"
	"payment-bridge/blockchain/browsersync/goerli"
	"payment-bridge/blockchain/goerliclient"
	nbaiclient "payment-bridge/blockchain/nbaiclient"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/logs"
	models2 "payment-bridge/models"
	"sync"
	"time"
)

func NbaiBlockBrowserSyncAndEventLogsSync() {
	lastCunrrentNumber := getStartBlockNo()

	for {
		var blockNoCurrent *big.Int
		var err error
		var getBlockFlag bool = true
		for getBlockFlag {
			blockNoCurrent, err = nbaiclient.WebConn.GetBlockNumber()
			if err != nil {
				goerliclient.ClientInit()
				logs.GetLogger().Error(err)
				time.Sleep(5 * time.Second)
				continue
			}
			if err == nil {
				getBlockFlag = false
			}
		}

		scanStep := config.GetConfig().PolygonMainnetNode.ScanStep
		if (blockNoCurrent.Int64() < lastCunrrentNumber) || (blockNoCurrent.Int64()-lastCunrrentNumber) < scanStep {
			lastCunrrentNumber = blockNoCurrent.Int64() - 10000
		}

		var mutex sync.Mutex
		mutex.Lock()
		err = ScanNbaiEventFromChainAndSaveEventLogData(lastCunrrentNumber, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		blockScanRecord := models2.BlockScanRecord{}
		whereCondition := "network_type='" + constants.NETWORK_TYPE_POLYGON + "'"
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			lastCunrrentNumber = 1
		}

		err = goerli.UpdateOrSaveBlockScanRecord(constants.NETWORK_TYPE_POLYGON, blockScanRecordList, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		lastCunrrentNumber = blockNoCurrent.Int64()
		mutex.Unlock()

		time.Sleep(time.Second * 5)
	}
}

func getStartBlockNo() int64 {
	var lastCunrrentNumber int64 = 1

	if config.GetConfig().NbaiMainnetNode.StartFromBlockNo > 0 {
		lastCunrrentNumber = config.GetConfig().NbaiMainnetNode.StartFromBlockNo
	}
	return lastCunrrentNumber
}
