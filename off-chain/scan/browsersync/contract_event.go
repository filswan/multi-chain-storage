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
	constants "payment-bridge/off-chain/common"
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

func GetAbi() {
	paymentAbiString, err := ReadContractAbiJsonFile()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	contractAddress := common.HexToAddress(constants.PAYMENT_CONTRACT_ADDRESS)
	contractFunctionSignature := constants.CONTRACT_FUNCTION_SIGNATURE
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(5297224),
		ToBlock:   big.NewInt(5297224),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//logs, err := client.FilterLogs(context.Background(), query)
	logsInChain, err := eth.WebConn.ConnWeb.FilterLogs(context.Background(), query)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Fatal(err)
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
			event.DataCid = dataList[0].(string)
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

type event struct {
	ID         string
	lockedFee  *big.Int
	minPayment *big.Int
	recipient  *big.Int
	deadline   *big.Int
}
