package goerli

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
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

const SwanPaymentAbiJson = "on-chain/contracts/abi/SwanPayment.json"

// EventLogSave Find the event that executed the contract and save to db
func ScanGoerliEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	logs.GetLogger().Println("goerli blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))

	//read contract api json file
	paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.SwanPaymentABI)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().GoerliMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().GoerliMainnetNode.ContractFunctionSignature

	//test block no. is : 5297224
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	var logsInChain []types.Log
	var flag bool = true
	for flag {
		//logs, err := client.FilterLogs(context.Background(), query)
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
			eventList, err := models.FindEventGoerlis(&models.EventGoerli{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				var event = new(models.EventGoerli)
				dataList, err := contractAbi.Unpack("LockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				addrInfo, err := utils.GetFromAndToAddressByTxHash(WebConn.ConnWeb, big.NewInt(GetConfig().GoerliMainnetNode.ChainID), vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.AddressFrom = addrInfo.AddrFrom
					event.AddressTo = addrInfo.AddrTo
				}
				event.BlockNo = vLog.BlockNumber
				event.TxHash = vLog.TxHash.Hex()
				event.ContractName = "SwanPayment"
				event.ContractAddress = contractAddress.String()
				event.PayloadCid = dataList[0].(string)
				event.MinerAddress = dataList[3].(common.Address).Hex()
				lockFee := dataList[1].(*big.Int)
				event.LockedFee = lockFee.String()
				deadLine := dataList[4].(*big.Int)
				event.Deadline = deadLine.String()
				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.LockPaymentTime = strconv.FormatUint(block.Time(), 10)
				}
				//todo
				event.EventName = "USDC"
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				err = database.SaveOne(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
