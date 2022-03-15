package polygon

import (
	"context"
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/logs"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/goBind"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-dannyng
 * Copyright defined in multi-chain-storage/LICENSE
 */

// EventLogSave Find the event that executed the contract and save to db
func ScanPolygonExpirePaymentEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("polygon expireEvent scan blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	expirePaymentAbiString := goBind.SwanPaymentABI

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().PolygonMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().PolygonMainnetNode.ExpireEventFunctionSignature

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
		Topics: [][]common.Hash{{common.HexToHash(contractFunctionSignature)}},
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	var logsInChain []types.Log
	var flag bool = true
	var err error
	for flag {
		logsInChain, err = WebConn.ConnWeb.FilterLogs(context.Background(), query)
		if err != nil {
			logs.GetLogger().Error(err)
			time.Sleep(5 * time.Second)
			continue
		}
		if err == nil {
			flag = false
		}
	}

	contractAbi, err := abi.JSON(strings.NewReader(expirePaymentAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			dataList, err := contractAbi.Unpack("ExpirePayment", vLog.Data)
			if err != nil {
				logs.GetLogger().Error(err)
			}
			logs.GetLogger().Info(dataList)
			eventList, err := models.FindEventExpirePayments(&models.EventExpirePayment{TxHash: vLog.TxHash.Hex(), BlockNo: strconv.FormatUint(vLog.BlockNumber, 10)}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			var event *models.EventExpirePayment
			if len(eventList) <= 0 {
				event = new(models.EventExpirePayment)
				event.TxHash = vLog.TxHash.Hex()
				event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.BlockTime = strconv.FormatUint(block.Time(), 10)
				}
				wfilCoinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_WFIL_ON_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.CoinId = wfilCoinId
				}
				networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.NetworkId = networkId
				}
				event.PayloadCid = dataList[0].(string)
				event.TokenAddress = dataList[1].(common.Address).Hex()
				event.ExpireUserAmount = dataList[2].(*big.Int).String()
				event.UserAddress = dataList[3].(common.Address).Hex()
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				event.ContractAddress = contractAddress.String()
				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
