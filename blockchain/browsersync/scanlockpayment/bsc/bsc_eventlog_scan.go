package bsc

import (
	"math/big"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	models2 "payment-bridge/models"
	"strconv"
	"sync"
	"time"

	"github.com/filswan/go-swan-lib/logs"
)

func ScanEventFromChainAndSaveDataToDbForBsc() {
	startScanBlockNo := getStartBlockNo()

	for {
		var mutex sync.Mutex
		mutex.Lock()
		var blockNoCurrent *big.Int
		var err error
		var getBlockFlag bool = true
		for getBlockFlag {
			blockNoCurrent, err = WebConn.GetBlockNumber()
			if err != nil {
				ClientInit()
				logs.GetLogger().Error(err)
				time.Sleep(5 * time.Second)
				continue
			}
			if err == nil {
				getBlockFlag = false
			}
		}

		blockScanRecord := new(models2.BlockScanRecord)
		whereCondition := "network_type='" + constants.NETWORK_TYPE_BSC + "'"
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			startScanBlockNo = GetConfig().BscMainnetNode.StartFromBlockNo
		}
		if len(blockScanRecordList) > 0 {
			if blockScanRecordList[0].LastCurrentBlockNumber <= blockNoCurrent.Int64() {
				startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
			} else {
				startScanBlockNo = GetConfig().BscMainnetNode.StartFromBlockNo
			}
			blockScanRecord.ID = blockScanRecordList[0].ID
		}

		for {
			start := startScanBlockNo
			end := start + GetConfig().BscMainnetNode.ScanStep
			if startScanBlockNo > blockNoCurrent.Int64() {
				break
			}
			err = ScanBscEventFromChainAndSaveEventLogData(start, end)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			if end >= blockNoCurrent.Int64() {
				blockScanRecord.LastCurrentBlockNumber = blockNoCurrent.Int64()
			} else {
				blockScanRecord.LastCurrentBlockNumber = end
			}

			blockScanRecord.UpdateAt = strconv.FormatInt(utils.GetCurrentUtcMilliSecond(), 10)

			err = database.SaveOne(blockScanRecord)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			start = end
			startScanBlockNo = end
			if end >= blockNoCurrent.Int64() {
				break
			}
		}

		getBlockFlag = true
		mutex.Unlock()

		time.Sleep(time.Second * GetConfig().BscMainnetNode.CycleTimeInterval)
		logs.GetLogger().Info("-------------------------bsc----------------------------")
	}
}

func getStartBlockNo() int64 {
	var startScanBlockNo int64 = 1

	if GetConfig().BscMainnetNode.StartFromBlockNo > 0 {
		startScanBlockNo = GetConfig().BscMainnetNode.StartFromBlockNo
	}
	blockScanRecord := new(models2.BlockScanRecord)
	whereCondition := "network_type='" + constants.NETWORK_TYPE_BSC + "'"
	blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
	if err != nil {
		logs.GetLogger().Error(err)
		startScanBlockNo = GetConfig().BscMainnetNode.StartFromBlockNo
	}

	if len(blockScanRecordList) > 0 {
		if blockScanRecordList[0].LastCurrentBlockNumber > startScanBlockNo {
			startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
		}
	}
	return startScanBlockNo
}
