package nbai

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
 * created on 08/20/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

// EventLogSave Find the event that executed the contract and save to db
func ScanNbaiEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("nbai blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.StateSenderABI)
	paymentAbiString, err := abi.JSON(strings.NewReader(string(goBind.SwanPaymentMetaData.ABI)))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().NbaiMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().NbaiMainnetNode.ContractFunctionSignature

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

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			eventList, err := models.FindEventNbai(&models.EventNbai{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				receiveMap := map[string]interface{}{}
				err = paymentAbiString.UnpackIntoMap(receiveMap, "StateSynced", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				var event = new(models.EventNbai)
				addrInfo, err := utils.GetFromAndToAddressByTxHash(WebConn.ConnWeb, big.NewInt(GetConfig().NbaiMainnetNode.ChainID), vLog.TxHash)
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
				event.BytesData = receiveMap["data"].([]byte)
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
			}
		}
	}
	return nil
}
