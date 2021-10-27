package blockutils

import (
	"payment-bridge/common/constants"
	"payment-bridge/logs"
	models2 "payment-bridge/models"
)

func GetStartBlockNo(networkName string, startBlockNoInConfig int64) int64 {
	var startScanBlockNo int64 = 1
	var whereCondition string

	switch networkName {
	case constants.NETWORK_TYPE_POLYGON:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_POLYGON + "'"
	case constants.NETWORK_TYPE_NBAI:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_NBAI + "'"
	case constants.NETWORK_TYPE_BSC:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_BSC + "'"
	case constants.NETWORK_TYPE_GOERLI:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_GOERLI + "'"
	case constants.NETWORK_TYPE_ETH:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_ETH + "'"
	}

	if startBlockNoInConfig > 0 {
		startScanBlockNo = startBlockNoInConfig
	}

	blockScanRecord := new(models2.BlockScanRecord)
	blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
	if err != nil {
		logs.GetLogger().Error(err)
		startScanBlockNo = startBlockNoInConfig
	}

	if len(blockScanRecordList) > 0 {
		if blockScanRecordList[0].LastCurrentBlockNumber > startScanBlockNo {
			startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
		}
	}
	return startScanBlockNo
}
