package goerli

import (
	"math/big"
	"payment-bridge/scan/blockchain/goerliclient"
	"payment-bridge/scan/common/constants"
	common2 "payment-bridge/scan/common/utils"
	"payment-bridge/scan/config"
	"payment-bridge/scan/database"
	"payment-bridge/scan/logs"
	models2 "payment-bridge/scan/models"
	"strconv"
	"sync"
	"time"
)

func GoerliBlockBrowserSyncAndEventLogsSync() {
	lastCunrrentNumber := getStartBlockNo()

	for {
		var blockNoCurrent *big.Int
		var err error
		var getBlockFlag bool = true
		for getBlockFlag {
			blockNoCurrent, err = goerliclient.WebConn.GetBlockNumber()
			if err != nil {
				logs.GetLogger().Error(err)
				time.Sleep(5 * time.Second)
				continue
			}
			if err == nil {
				getBlockFlag = false
			}
		}

		scanStep := config.GetConfig().GoerliMainnetNode.ScanStep
		if (blockNoCurrent.Int64() < lastCunrrentNumber) || (blockNoCurrent.Int64()-lastCunrrentNumber) < scanStep {
			lastCunrrentNumber = blockNoCurrent.Int64() - scanStep
		}

		var mutex sync.Mutex
		mutex.Lock()

		err = ScanGoerliEventFromChainAndSaveEventLogData(lastCunrrentNumber, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		blockScanRecord := models2.BlockScanRecord{}
		whereCondition := "network_type='" + constants.NETWORK_TYPE_GOERLI + "'"
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			lastCunrrentNumber = 1
		}

		err = UpdateOrSaveBlockScanRecord(constants.NETWORK_TYPE_GOERLI, blockScanRecordList, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		lastCunrrentNumber = blockNoCurrent.Int64()
		mutex.Unlock()

		time.Sleep(time.Second * 5)
	}
}

func UpdateOrSaveBlockScanRecord(networkType string, blockScanRecordList []*models2.BlockScanRecord, lastCurrentBlockNumber int64) error {
	currentTime := strconv.FormatInt(common2.GetEpochInMillis(), 10)
	if len(blockScanRecordList) >= 1 {
		_, err := models2.UpdateBlockScanRecord(&models2.BlockScanRecord{NetworkType: networkType},
			map[string]interface{}{"update_at": currentTime, "last_current_block_number": lastCurrentBlockNumber})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	} else {
		newRecord := new(models2.BlockScanRecord)
		newRecord.NetworkType = networkType
		newRecord.LastCurrentBlockNumber = lastCurrentBlockNumber
		newRecord.UpdateAt = string(currentTime)
		err := database.SaveOne(newRecord)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}

func getStartBlockNo() int64 {
	var lastCunrrentNumber int64 = 1

	if config.GetConfig().PolygonMainnetNode.StartFromBlockNo > 0 {
		lastCunrrentNumber = config.GetConfig().GoerliMainnetNode.StartFromBlockNo
	}
	return lastCunrrentNumber
}
