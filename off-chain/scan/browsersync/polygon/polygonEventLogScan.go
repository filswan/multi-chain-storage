package polygon

import (
	"fmt"
	"payment-bridge/off-chain/common/constants"
	"payment-bridge/off-chain/logs"
	models2 "payment-bridge/off-chain/models"
	"payment-bridge/off-chain/scan/browsersync/goerli"
	"payment-bridge/off-chain/scan/goerliclient"
	polygonclient "payment-bridge/off-chain/scan/polygonclient"
	"sync"
	"time"
)

func PolygonBlockBrowserSyncAndEventLogsSync() {

	for {
		var lastCunrrentNumber int64 = 1
		blockScanRecord := models2.BlockScanRecord{}
		whereCondition := "network_type='" + constants.NETWORK_TYPE_POLYGON + "'"
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			lastCunrrentNumber = 1
		}

		if len(blockScanRecordList) >= 1 {
			lastCunrrentNumber = blockScanRecordList[0].LastCurrentBlockNumber
		}
		fmt.Println(lastCunrrentNumber)
		blockNoCurrent, err := polygonclient.WebConn.GetBlockNumber()
		if err != nil {
			goerliclient.ClientInit()
			logs.GetLogger().Error(err)
			continue
		}

		var mutex sync.Mutex
		mutex.Lock()
		err = ScanPolygonEventFromChainAndSaveEventLogData(lastCunrrentNumber, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		err = goerli.UpdateOrSaveBlockScanRecord(constants.NETWORK_TYPE_POLYGON, blockScanRecordList, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		mutex.Unlock()

		time.Sleep(time.Second * 5)
	}
}
