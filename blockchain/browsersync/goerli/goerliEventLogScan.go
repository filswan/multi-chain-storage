package goerli

import (
	"math/big"
	"payment-bridge/blockchain/initclient/goerliclient"
	"payment-bridge/blockchain/initclient/nbaiclient"
	"payment-bridge/common/constants"
	common2 "payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	models2 "payment-bridge/models"
	"strconv"
	"sync"
	"time"
)

func GoerliBlockBrowserSyncAndEventLogsSync() {
	startScanBlockNo := getStartBlockNo()

	for {
		var blockNoCurrent *big.Int
		var err error
		var getBlockFlag bool = true
		for getBlockFlag {
			blockNoCurrent, err = goerliclient.WebConn.GetBlockNumber()
			if err != nil {
				nbaiclient.ClientInit()
				logs.GetLogger().Error(err)
				time.Sleep(5 * time.Second)
				continue
			}
			if err == nil {
				getBlockFlag = false
			}
		}

		blockScanRecord := new(models2.BlockScanRecord)
		whereCondition := "network_type='" + constants.NETWORK_TYPE_GOERLI + "'"
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
			end := start + config.GetConfig().GoerliMainnetNode.ScanStep
			if startScanBlockNo > blockNoCurrent.Int64() {
				break
			}
			err = ScanGoerliEventFromChainAndSaveEventLogData(start, end)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			if end >= blockNoCurrent.Int64() {
				blockScanRecord.LastCurrentBlockNumber = blockNoCurrent.Int64()
			} else {
				blockScanRecord.LastCurrentBlockNumber = end
			}

			blockScanRecord.NetworkType = constants.NETWORK_TYPE_GOERLI
			blockScanRecord.UpdateAt = strconv.FormatInt(common2.GetEpochInMillis(), 10)

			err = database.SaveOne(blockScanRecord)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			startScanBlockNo = startScanBlockNo + config.GetConfig().GoerliMainnetNode.ScanStep
			if startScanBlockNo >= blockNoCurrent.Int64() {
				break
			}
		}

		getBlockFlag = true
		mutex.Unlock()

		time.Sleep(time.Second * config.GetConfig().GoerliMainnetNode.CycleTimeInterval)
	}
}

func getStartBlockNo() int64 {
	var lastCunrrentNumber int64 = 1

	if config.GetConfig().GoerliMainnetNode.StartFromBlockNo > 0 {
		lastCunrrentNumber = config.GetConfig().GoerliMainnetNode.StartFromBlockNo
	}
	return lastCunrrentNumber
}
