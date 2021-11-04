package polygon

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"strings"
	"time"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

// EventLogSave Find the event that executed the contract and save to db
func ScanPolygonEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("polygon blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.SwanPaymentMetaData.ABI)
	paymentAbiString := goBind.SwanPaymentMetaData.ABI

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().PolygonMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().PolygonMainnetNode.ContractFunctionSignature

	//test block no. is : 5297224
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
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

	contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			eventList, err := models.FindEventPolygons(&models.EventPolygon{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				var event = new(models.EventPolygon)
				dataList, err := contractAbi.Unpack(constants.EVENT_NAME_LOCAK_PAYMENT, vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}

				addrInfo, err := utils.GetFromAndToAddressByTxHash(WebConn.ConnWeb, big.NewInt(GetConfig().PolygonMainnetNode.ChainID), vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.AddressFrom = addrInfo.AddrFrom
					event.AddressTo = addrInfo.AddrTo
				}
				event.BlockNo = vLog.BlockNumber
				//todo
				event.CoinType = "USDC"
				event.PayloadCid = dataList[0].(string)
				event.TokenAddress = dataList[1].(common.Address).Hex()
				event.LockedFee = dataList[2].(*big.Int).String()
				event.MinPayment = dataList[3].(*big.Int).String()
				event.MinerAddress = dataList[4].(common.Address).Hex()
				deadLine := dataList[5].(*big.Int)
				event.EventName = constants.EVENT_NAME_LOCAK_PAYMENT
				event.Deadline = deadLine.String()
				event.TxHash = vLog.TxHash.Hex()
				event.ContractAddress = contractAddress.String()
				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.LockPaymentTime = strconv.FormatUint(block.Time(), 10)
				}
				//todo
				event.EventName = "USDC"
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
