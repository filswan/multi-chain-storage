package polygon

import (
	"math/big"
	"payment-bridge/blockchain/initclient/goerliclient"
	"payment-bridge/blockchain/initclient/polygonclient"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	models2 "payment-bridge/models"
	"strconv"
	"sync"
	"time"
)

func PolygonBlockBrowserSyncAndEventLogsSync() {
	startScanBlockNo := getStartBlockNo()

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

		blockScanRecord := new(models2.BlockScanRecord)
		whereCondition := "network_type='" + constants.NETWORK_TYPE_POLYGON + "'"
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			startScanBlockNo = getStartBlockNo()
		}
		if len(blockScanRecordList) > 0 {
			startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
			blockScanRecord.ID = blockScanRecordList[0].ID
		}

		var mutex sync.Mutex
		mutex.Lock()
		for {
			start := startScanBlockNo
			end := start + config.GetConfig().PolygonMainnetNode.ScanStep
			if startScanBlockNo > blockNoCurrent.Int64() {
				break
			}
			err = ScanPolygonEventFromChainAndSaveEventLogData(start, end)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			if end >= blockNoCurrent.Int64() {
				blockScanRecord.LastCurrentBlockNumber = blockNoCurrent.Int64()
			} else {
				blockScanRecord.LastCurrentBlockNumber = end
			}

			blockScanRecord.NetworkType = constants.NETWORK_TYPE_POLYGON
			blockScanRecord.UpdateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)

			err = database.SaveOne(blockScanRecord)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			startScanBlockNo = startScanBlockNo + config.GetConfig().PolygonMainnetNode.ScanStep
			if startScanBlockNo >= blockNoCurrent.Int64() {
				break
			}
		}

		getBlockFlag = true

		mutex.Unlock()

		time.Sleep(time.Second * config.GetConfig().PolygonMainnetNode.CycleTimeInterval)
	}
}

func getStartBlockNo() int64 {
	var startScanBlockNo int64 = 1

	if config.GetConfig().PolygonMainnetNode.StartFromBlockNo > 0 {
		startScanBlockNo = config.GetConfig().PolygonMainnetNode.StartFromBlockNo
	}
	return startScanBlockNo
}
