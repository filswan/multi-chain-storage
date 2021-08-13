package polygon

import (
	"payment-bridge/off-chain/common/constants"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/logs"
	models2 "payment-bridge/off-chain/models"
	"payment-bridge/off-chain/scan/browsersync/goerli"
	"payment-bridge/off-chain/scan/goerliclient"
	polygonclient "payment-bridge/off-chain/scan/polygonclient"
	"sync"
	"time"
)

func PolygonBlockBrowserSyncAndEventLogsSync() {
	lastCunrrentNumber := getStartBlockNo()

	for {

		blockNoCurrent, err := polygonclient.WebConn.GetBlockNumber()
		if err != nil {
			goerliclient.ClientInit()
			logs.GetLogger().Error(err)
			continue
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
