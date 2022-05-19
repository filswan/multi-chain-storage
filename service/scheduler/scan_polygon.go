package scheduler

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	decoder "github.com/mingjingc/abi-decoder"
)

var txDataDecoder *decoder.ABIDecoder

func initTxDataDecoder() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.", err)
	}

	contractPath := filepath.Join(homedir, ".swan/mcs/SwanPayment.json")
	contractRead, err := ioutil.ReadFile(contractPath)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	myContractAbi := string(contractRead)

	txDataDecoder = decoder.NewABIDecoder()
	txDataDecoder.SetABI(myContractAbi)
}

func ScanPolygon() error {
	network, err := models.GetNetworkByName(constants.NETWORK_NAME_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	startBlockNumber := int64(0)
	if network.LastScanBlockNumber != nil {
		startBlockNumber = *network.LastScanBlockNumber + 1
	}

	endBlockNumberUint64, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	endBlockNumber := int64(endBlockNumberUint64)
	scanBlockStep := int64(config.GetConfig().Polygon.ScanPolygonBlockStep)
	for i := startBlockNumber; i <= endBlockNumber; {
		toBlockNumber := i + scanBlockStep - 1
		if toBlockNumber > endBlockNumber {
			toBlockNumber = endBlockNumber
		}
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(i),
			ToBlock:   big.NewInt(toBlockNumber),
			Addresses: []common.Address{
				common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress),
			},
		}

		vlogs, err := ethClient.FilterLogs(context.Background(), query)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		for _, vLog := range vlogs {
			transaction, isPending, err := ethClient.TransactionByHash(context.Background(), vLog.TxHash)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}

			if isPending {
				return nil
			}

			inputDataHex := hex.EncodeToString(transaction.Data())
			if strings.HasPrefix(inputDataHex, "f4d98717") {
				err = getPayment4Transaction(ethClient, inputDataHex, *transaction)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			} else if strings.HasPrefix(inputDataHex, "7d29985b") {
				//to do
			}
		}

		network.LastScanBlockNumber = &toBlockNumber
		err = database.GetDB().Save(network).Error
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		i = toBlockNumber + 1
	}

	return nil
}

func getPayment4Transaction(ethClient *ethclient.Client, inputDataHex string, transaction types.Transaction) error {
	txReceipt, err := client.CheckTx(ethClient, transaction.Hash())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if txReceipt.Status != uint64(1) {
		return nil
	}

	method, err := txDataDecoder.DecodeMethod(inputDataHex)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(method.Params) <= 0 {
		err = fmt.Errorf("method.Params is empty")
		return err
	}

	params := strings.Split(method.Params[0].Value, " ")
	if len(params) < 6 {
		err = fmt.Errorf("not enough params")
		return err
	}

	wCid := params[0]
	if len(wCid) <= 1 {
		err = fmt.Errorf("wCid is empty")
		return err
	}

	wCid = wCid[1:]
	lockTimeStr := params[3]
	lockTime, err := strconv.ParseInt(lockTimeStr, 10, 32)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	err = models.CreateTransaction4PayByWCid(wCid, transaction.Hash().String(), lockTime)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
