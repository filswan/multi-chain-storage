package browsersync

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"math/big"
	"os"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/database"
	"payment-bridge/off-chain/logs"
	"payment-bridge/off-chain/models"
	"payment-bridge/off-chain/scan/eth"
	"strings"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

const SwanPaymentAbiJson = "on-chain/contracts/abi/SwanPayment.json"

// EventLogSave Find the event that executed the contract and save to db
func ScanEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	paymentAbiString, err := ReadContractAbiJsonFile()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(config.GetConfig().MainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := config.GetConfig().MainnetNode.ContractFunctionSignature

	//test block no. is : 5297224
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	logsInChain, err := eth.WebConn.ConnWeb.FilterLogs(context.Background(), query)
	if err != nil {
		logs.GetLogger().Fatal(err)
		return err
	}

	contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Fatal(err)
		return err
	}

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			fmt.Println(vLog.BlockHash.Hex())
			fmt.Println(vLog.BlockNumber)
			fmt.Println(vLog.TxHash.Hex())
			var event = new(models.Event)

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

			eventList, err := models.FindEvents(&models.Event{TxHash: event.TxHash, BlockNo: event.BlockNo}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				err = database.SaveOne(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}

func ReadContractAbiJsonFile() (string, error) {
	jsonFile, err := os.Open(SwanPaymentAbiJson)

	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	return string(byteValue), nil
}
