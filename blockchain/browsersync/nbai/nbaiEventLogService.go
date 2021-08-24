package nbai

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"payment-bridge/blockchain/nbaiclient"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"strings"
	"time"
)

/**
 * created on 08/20/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

const NbaiAbiJson = "on-chain/contracts/abi/SwanPayment.json"

// EventLogSave Find the event that executed the contract and save to db
func ScanNbaiEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	paymentAbiString, err := utils.ReadContractAbiJsonFile(NbaiAbiJson)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(config.GetConfig().NbaiMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := config.GetConfig().NbaiMainnetNode.ContractFunctionSignature

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
	for flag {
		logsInChain, err = nbaiclient.WebConn.ConnWeb.FilterLogs(context.Background(), query)
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
			eventList, err := models.FindEventNbai(&models.EventNbai{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				fmt.Println(vLog.BlockHash.Hex())
				fmt.Println(vLog.BlockNumber)
				fmt.Println(vLog.TxHash.Hex())
				var event = new(models.EventNbai)

				dataList, err := contractAbi.Unpack("LockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				fmt.Println(len(dataList))
				fmt.Println(dataList[0])
				fmt.Println(dataList[1])
				fmt.Println(dataList[2])
				fmt.Println(dataList[3])
				fmt.Println(dataList[4])
				event.BlockNo = vLog.BlockNumber
				event.TxHash = vLog.TxHash.Hex()
				event.ContractName = "SwanPayment"
				event.ContractAddress = contractAddress.String()
				event.BytesData = dataList[0].(string)
				event.Address = fmt.Sprintf("%x", dataList[2].(*big.Int))
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
