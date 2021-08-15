package goerli

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"payment-bridge/off-chain/common/utils"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/database"
	"payment-bridge/off-chain/logs"
	"payment-bridge/off-chain/models"
	"payment-bridge/off-chain/scan/goerliclient"
	"strconv"
	"strings"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

const SwanPaymentAbiJson = "on-chain/contracts/abi/SwanPayment.json"

// EventLogSave Find the event that executed the contract and save to db
func ScanGoerliEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	logs.GetLogger().Println("blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))

	//read contract api json file
	paymentAbiString, err := utils.ReadContractAbiJsonFile(SwanPaymentAbiJson)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(config.GetConfig().GoerliMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := config.GetConfig().GoerliMainnetNode.ContractFunctionSignature

	//test block no. is : 5297224
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	logsInChain, err := goerliclient.WebConn.ConnWeb.FilterLogs(context.Background(), query)
	if err != nil {
		goerliclient.ClientInit()
		logs.GetLogger().Error(err)
		return err
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
				fmt.Println(vLog.BlockHash.Hex())
				fmt.Println(vLog.BlockNumber)
				fmt.Println(vLog.TxHash.Hex())
				var event = new(models.EventGoerli)
				dataList, err := contractAbi.Unpack("LockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
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
				err = database.SaveOne(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
