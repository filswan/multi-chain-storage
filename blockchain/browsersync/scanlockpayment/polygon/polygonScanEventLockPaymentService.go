package polygon

import (
	"context"
	"math/big"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

// EventLogSave Find the event that executed the contract and save to db
func ScanPolygonLockPaymentEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("polygon blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.SwanPaymentMetaData.ABI)
	paymentAbiString := goBind.SwanPaymentABI

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().PolygonMainnetNode.PaymentContractAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().PolygonMainnetNode.ContractLockFunctionSignature

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
			eventList, err := models.FindEventLockPayment(&models.EventLockPayment{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				var event = new(models.EventLockPayment)
				dataList, err := contractAbi.Unpack("LockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
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
					event.AddressFrom = addrInfo.AddrFrom
					event.AddressTo = addrInfo.AddrTo
				}
				event.BlockNo = vLog.BlockNumber
				usdcCoinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_USDC_ON_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.CoinId = usdcCoinId
				}
				networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.NetworkId = networkId
				}

				event.PayloadCid = dataList[0].(string)
				event.TokenAddress = dataList[1].(common.Address).Hex()
				event.LockedFee = dataList[2].(*big.Int).String()
				event.MinPayment = dataList[3].(*big.Int).String()
				event.MinerAddress = dataList[4].(common.Address).Hex()
				deadLine := dataList[5].(*big.Int)
				event.Deadline = deadLine.String()
				event.TxHash = vLog.TxHash.Hex()
				event.ContractAddress = contractAddress.String()
				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.LockPaymentTime = strconv.FormatUint(block.Time(), 10)
				}
				event.CreateAt = utils.GetCurrentUtcMilliSecond()
				currrentTime := strconv.FormatInt(event.CreateAt, 10)
				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				//condition :&models.EventPolygon{Ip: "192.168.88.80"}
				//updateFields: map[string]interface{}{"processing_time": taskT.ProcessingTime, "worker_reward": taskT.WorkerReward}
				err = models.UpdateDealFile(models.DealFile{PayloadCid: dataList[0].(string)},
					map[string]interface{}{"lock_payment_status": constants.LOCK_PAYMENT_STATUS_PROCESSING, "lock_payment_network": networkId, "lock_payment_tx": vLog.TxHash.Hex(), "update_at": currrentTime})
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
