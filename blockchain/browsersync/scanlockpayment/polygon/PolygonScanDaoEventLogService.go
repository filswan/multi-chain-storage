package polygon

import (
	"context"
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/logs"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/goBind"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EventLogSave Find the event that executed the contract and save to db
func ScanDaoEventFromChainAndSaveEventLogData(blockNoFrom, blockNoTo int64) error {
	//read contract api json file
	logs.GetLogger().Println("polygon scan dao event  : blockNoFrom=" + strconv.FormatInt(blockNoFrom, 10) + "--------------blockNoTo=" + strconv.FormatInt(blockNoTo, 10))
	daoEventAbiString := goBind.SwanPaymentABI
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
		Topics: [][]common.Hash{{common.HexToHash(contractFunctionSignature)}},
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
			dataList, err := contractAbi.Unpack("SignTransaction", vLog.Data)
			if err != nil {
				logs.GetLogger().Error(err)
			}
			logs.GetLogger().Info(dataList)
			eventList, err := models.FindDaoEventLog(&models.EventDaoSignature{TxHash: vLog.TxHash.Hex(), BlockNo: vLog.BlockNumber}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				var event = new(models.EventDaoSignature)
				dataList, err := contractAbi.Unpack("SignTransaction", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err, ", tx hash:", vLog.TxHash.Hex(), ", vLog.Data:", vLog.Data)
					continue
				}

				block, err := WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.BlockTime = strconv.FormatUint(block.Time(), 10)
					event.DaoPassTime = strconv.FormatUint(block.Time(), 10)
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
					event.DaoAddress = addrInfo.AddrFrom
				}

				wfilCoinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_WFIL_ON_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.CoinId = wfilCoinId
				}
				networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.NetworkId = networkId
				}
				event.BlockNo = vLog.BlockNumber
				event.TxHash = vLog.TxHash.Hex()
				event.PayloadCid = dataList[0].(string)
				dealId, err := strconv.ParseInt(dataList[1].(string), 10, 64)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.DealId = dealId
				}
				event.Recipient = dataList[2].(common.Address).String()
				event.Status = true
				event.SignatureUnlockStatus = constants.SIGNATURE_DEFAULT_VALUE

				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
	}
	return nil
}
