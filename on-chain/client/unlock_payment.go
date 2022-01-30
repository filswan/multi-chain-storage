package client

import (
	"math/big"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/models"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

func SaveEventUnlockPayment(logsInChain []*types.Log, unlockStatus string, dealId int64) error {
	contractUnlockFunctionSignature := polygon.GetConfig().PolygonMainnetNode.ContractUnlockFunctionSignature

	contractAbi, err := GetContractAbi()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, vLog := range logsInChain {
		if vLog.Topics[0].Hex() != contractUnlockFunctionSignature {
			continue
		}

		event := new(models.EventUnlockPayment)
		dataList, err := contractAbi.Unpack("UnlockPayment", vLog.Data)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		event.DealId = dealId
		event.TxHash = vLog.TxHash.Hex()
		networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.NetworkId = networkId
		}
		coinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_USDC_ON_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.CoinId = coinId
		}
		event.TokenAddress = dataList[1].(common.Address).Hex()
		event.UnlockToAdminAmount = dataList[2].(*big.Int).String()
		event.UnlockToUserAmount = dataList[3].(*big.Int).String()
		event.UnlockToAdminAddress = dataList[4].(common.Address).Hex()
		event.UnlockToUserAddress = dataList[5].(common.Address).Hex()
		event.UnlockTime = strconv.FormatInt(utils.GetCurrentUtcMilliSecond(), 10)
		event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
		event.CreateAt = utils.GetCurrentUtcMilliSecond()
		event.UnlockStatus = unlockStatus

		err = database.SaveOneWithTransaction(event)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
