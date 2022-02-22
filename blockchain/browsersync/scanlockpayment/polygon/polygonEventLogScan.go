package polygon

import (
	"math/big"
	"payment-bridge/blockchain/browsersync/blockutils"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/logs"
	models2 "payment-bridge/models"
	"strconv"
	"sync"
	"time"
)

func ScanEventFromChainAndSaveDataToDbForPolygon() {
	startScanBlockNo := blockutils.GetStartBlockNo(constants.NETWORK_TYPE_POLYGON, GetConfig().PolygonMainnetNode.StartFromBlockNo)

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
		var whereCondition string
		networkId, err := models2.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
			startScanBlockNo = GetConfig().PolygonMainnetNode.StartFromBlockNo
		} else {
			whereCondition = "network_id=" + strconv.FormatInt(networkId, 10) + ""
		}

		blockScanRecord := new(models2.BlockScanRecord)
		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
		if err != nil {
			logs.GetLogger().Error(err)
			startScanBlockNo = GetConfig().PolygonMainnetNode.StartFromBlockNo
		}
		if len(blockScanRecordList) > 0 {
			if blockScanRecordList[0].LastCurrentBlockNumber <= blockNoCurrent.Int64() {
				startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
			} else {
				startScanBlockNo = GetConfig().PolygonMainnetNode.StartFromBlockNo
			}
			blockScanRecord.ID = blockScanRecordList[0].ID
		}

		for {
			start := startScanBlockNo
			end := start + GetConfig().PolygonMainnetNode.ScanStep
			if startScanBlockNo > blockNoCurrent.Int64() {
				break
			}
			err = ScanPolygonLockPaymentEventFromChainAndSaveEventLogData(start-1, end+1)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			err = ScanDaoEventFromChainAndSaveEventLogData(start-1, end+1)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			err = ScanPolygonExpirePaymentEventFromChainAndSaveEventLogData(start-1, end+1)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			if end >= blockNoCurrent.Int64() {
				blockScanRecord.LastCurrentBlockNumber = blockNoCurrent.Int64()
			} else {
				blockScanRecord.LastCurrentBlockNumber = end
			}
			networkId, err := models2.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				blockScanRecord.NetworkId = networkId
			}
			blockScanRecord.UpdateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)

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

		time.Sleep(time.Second * GetConfig().PolygonMainnetNode.CycleTimeInterval)
		logs.GetLogger().Info("-------------------------polygon----------------------------")
	}
}
