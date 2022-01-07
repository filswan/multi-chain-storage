package bsc

import (
	"context"
	"math/big"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/models"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

// EventLogSave Find the event that executed the contract and save to db
func ScanBscEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("bsc blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	paymentAbiString, err := utils.ReadContractAbiJsonFile(constants.SWAN_PAYMENT_ABI_JSON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().BscMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().BscMainnetNode.ContractFunctionSignature

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

	contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			eventList, err := models.FindEventBsc(&models.EventBsc{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				var event = new(models.EventBsc)
				dataList, err := contractAbi.Unpack("LockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				addrInfo, err := utils.GetFromAndToAddressByTxHash(WebConn.ConnWeb, big.NewInt(GetConfig().BscMainnetNode.ChainID), vLog.TxHash)
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
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.LockPaymentTime = strconv.FormatUint(block.Time(), 10)
				}
				//todo
				event.EventName = "USDC"
				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
