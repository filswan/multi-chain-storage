package polygon

import (
	"math/big"
	"payment-bridge/scan/blockchain/browsersync/goerli"
	"payment-bridge/scan/blockchain/goerliclient"
	polygonclient "payment-bridge/scan/blockchain/polygonclient"
	"payment-bridge/scan/common/constants"
	"payment-bridge/scan/config"
	"payment-bridge/scan/logs"
	models2 "payment-bridge/scan/models"
	"sync"
	"time"
)

func PolygonBlockBrowserSyncAndEventLogsSync() {
	lastCunrrentNumber := getStartBlockNo()

	for {
		var blockNoCurrent *big.Int
		var err error
		var getBlockFlag bool = true
		for getBlockFlag {
			blockNoCurrent, err = polygonclient.WebConn.GetBlockNumber()
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
		err = ScanPolygonEventFromChainAndSaveEventLogData(lastCunrrentNumber, blockNoCurrent.Int64())
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

	if config.GetConfig().PolygonMainnetNode.StartFromBlockNo > 0 {
		lastCunrrentNumber = config.GetConfig().PolygonMainnetNode.StartFromBlockNo
	}
	return lastCunrrentNumber
}
