package client

import (
	"math/big"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

func SaveEventUnlockPayment(logsInChain []*types.Log, unlockStatus string, dealId int64) error {
	contractUnlockFunctionSignature := config.GetConfig().Polygon.ContractUnlockFunctionSignature

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
		event.TokenAddress = dataList[1].(common.Address).Hex()
		event.UnlockToAdminAmount = dataList[2].(*big.Int).String()
		event.UnlockToUserAmount = dataList[3].(*big.Int).String()
		event.UnlockToAdminAddress = dataList[4].(common.Address).Hex()
		event.UnlockToUserAddress = dataList[5].(common.Address).Hex()
		event.UnlockTime = strconv.FormatInt(utils.GetCurrentUtcMilliSecond(), 10)
		event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
		event.CreateAt = utils.GetCurrentUtcMilliSecond()
		event.UnlockStatus = unlockStatus

		network, err := models.GetNetworkByName(constants.NETWORK_NAME_POLYGON)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.NetworkId = network.ID
		}
		coin, err := models.FindCoinByCoinAddress(event.TokenAddress)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.CoinId = coin.ID
		}

		err = database.SaveOne(event)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
