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

// EventLogSave Find the event that executed the contract and save to db
func ScanDaoEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("scan dao event on polygon : blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	daoEventAbiString := goBind.FilswanOralcleMetaData.ABI

	//SwanPayment contract address
	contractAddress := common.HexToAddress(GetConfig().PolygonMainnetNode.DaoSwanOracleAddress)
	//SwanPayment contract function signature
	contractFunctionSignature := GetConfig().PolygonMainnetNode.DaoEventFunctionSignature

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

	contractAbi, err := abi.JSON(strings.NewReader(daoEventAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractFunctionSignature {
			eventList, err := models.FindDaoEventLog(&models.DaoEventLog{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				var event = new(models.DaoEventLog)
				dataList, err := contractAbi.Unpack("SignTransaction", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}

				block, _ := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				event.BlockTime = strconv.FormatUint(block.Time(), 10)
				event.DaoPassTime = strconv.FormatUint(block.Time(), 10)

				addrInfo, err := utils.GetFromAndToAddressByTxHash(WebConn.ConnWeb, big.NewInt(GetConfig().PolygonMainnetNode.ChainID), vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.DaoAddress = addrInfo.AddrFrom
				}
				event.BlockNo = vLog.BlockNumber
				event.TxHash = vLog.TxHash.Hex()
				event.PayloadCid = dataList[0].(string)
				event.DealCid = dataList[1].(string)
				event.Recipient = dataList[3].(common.Address).String()
				event.Cost = dataList[4].(*big.Int).String()
				event.Terms = dataList[5].(*big.Int).String()
				event.Network = constants.NETWORK_TYPE_POLYGON
				event.Status = dataList[6].(bool)

				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
