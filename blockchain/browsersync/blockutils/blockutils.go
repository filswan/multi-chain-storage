package blockutils

import (
	"payment-bridge/common/constants"
	"payment-bridge/models"
	models2 "payment-bridge/models"
	"strconv"

	"github.com/filswan/go-swan-lib/logs"
)

func GetStartBlockNo(networkName string, startBlockNoInConfig int64) int64 {
	var startScanBlockNo int64 = 1
	var uuid string

	switch networkName {
	case constants.NETWORK_TYPE_POLYGON:
		uuid = constants.NETWORK_TYPE_POLYGON_UUID
	case constants.NETWORK_TYPE_NBAI:
		uuid = constants.NETWORK_TYPE_NBAI_UUID
	case constants.NETWORK_TYPE_BSC:
		uuid = constants.NETWORK_TYPE_BSC_UUID
	case constants.NETWORK_TYPE_GOERLI:
		uuid = constants.NETWORK_TYPE_GOERLI_UUID
	case constants.NETWORK_TYPE_ETH:
		uuid = constants.NETWORK_TYPE_ETH_UUID
	}
	var whereCondition string
	networkId, err := models2.FindNetworkIdByUUID(uuid)
	if err != nil {
		logs.GetLogger().Error(err)
		startScanBlockNo = startBlockNoInConfig
	} else {
		whereCondition = "network_id=" + strconv.FormatInt(networkId, 10) + ""
	}

	if startBlockNoInConfig > 0 {
		startScanBlockNo = startBlockNoInConfig
	}

	blockScanRecordList, err := models.FindLastCurrentBlockNumber(whereCondition)
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
