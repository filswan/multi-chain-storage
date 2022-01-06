package polygon

import (
	"context"
	"math/big"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

//scan unlock payment event to db
func ScanPolygonUnLockPaymentEventFromChainAndSaveToDB(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("polygon scan unlock enven blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.SwanPaymentMetaData.ABI)
	//paymentAbiString := goBind.SwanPaymentMetaData.ABI

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().PolygonMainnetNode.PaymentUnlockContractAddress)
	//SwanPayment contract function signature
	contractUnlockFunctionSignature := GetConfig().PolygonMainnetNode.ContractUnlockFunctionSignature

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

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractUnlockFunctionSignature {
			eventList, err := models.FindEventUnlockPayments(&models.EventUnlockPayment{TxHash: vLog.TxHash.Hex(), BlockNo: strconv.FormatUint(vLog.BlockNumber, 10)}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			var event *models.EventUnlockPayment
			if len(eventList) <= 0 {
				event = new(models.EventUnlockPayment)
				event.TxHash = vLog.TxHash.Hex()
				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.UnlockTime = strconv.FormatUint(block.Time(), 10)
				}
				chainId, err := WebConn.ConnWeb.ChainID(context.Background())
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				addrInfo, err := utils.GetFromAndToAddressByTxHash(WebConn.ConnWeb, chainId, vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.UnlockFromAddress = addrInfo.AddrFrom
				}
				event.UnlockFromAddress = polygonConfig.PolygonMainnetNode.PaymentContractAddress
				event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
			} else {
				event = eventList[0]
			}

			quantity := new(big.Int)
			quantity.SetBytes(vLog.Data)
			address := common.HexToAddress(vLog.Topics[1].Hex()).String()
			if strings.EqualFold(address, polygonConfig.PolygonMainnetNode.PaymentContractAddress) {
				if strings.EqualFold(common.HexToAddress(vLog.Topics[2].Hex()).String(), config.GetConfig().AdminWalletOnPolygon) {
					event.UnlockToAdminAddress = common.HexToAddress(vLog.Topics[2].Hex()).String()
					event.UnlockToAdminAmount = quantity.String()
				} else {
					event.UnlockToUserAddress = common.HexToAddress(vLog.Topics[2].Hex()).String()
					event.UnlockToUserAmount = quantity.String()
				}
			}

			err = database.SaveOneWithTransaction(event)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}

	return nil
}
