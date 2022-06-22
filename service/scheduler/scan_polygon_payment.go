package scheduler

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
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

var txDataDecoderPayment *decoder.ABIDecoder

func initTxDataDecoderPayment() {
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

	txDataDecoderPayment = decoder.NewABIDecoder()
	txDataDecoderPayment.SetABI(myContractAbi)
}

func ScanPolygon4Payment() error {
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
	if network.LastScanBlockNumberPayment != nil {
		startBlockNumber = *network.LastScanBlockNumberPayment + 1
	}

	endBlockNumberUint64, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	endBlockNumber := int64(endBlockNumberUint64)
	scanBlockStep := int64(config.GetConfig().Polygon.ScanPolygonBlockStep)
	paymentContractAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)

	logs.GetLogger().Info("scan block [", startBlockNumber, ",", endBlockNumber, "] start, scan block step:", scanBlockStep)
	logs.GetLogger().Info("payment contract address:", paymentContractAddress.String())

	for i := startBlockNumber; i <= endBlockNumber; {
		toBlockNumber := i + scanBlockStep - 1
		if toBlockNumber > endBlockNumber {
			toBlockNumber = endBlockNumber
		}

		logs.GetLogger().Info("scan block [", i, ",", toBlockNumber, "] start")

		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(i),
			ToBlock:   big.NewInt(toBlockNumber),
			Addresses: []common.Address{paymentContractAddress},
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
			logs.GetLogger().Info(inputDataHex)
			if strings.HasPrefix(inputDataHex, "f4d98717") {
				err = getPayment(ethClient, inputDataHex, *transaction)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			} else if strings.HasPrefix(inputDataHex, "7d29985b") {
				err = getRefund(ethClient, inputDataHex, *transaction)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			} else if strings.HasPrefix(inputDataHex, "ee4128f6") {
				err = getUnlock(ethClient, inputDataHex, *transaction)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			}
		}

		err = models.UpdateNetworkLastScanBlockNumberPayment(network.ID, toBlockNumber)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		logs.GetLogger().Info("scan block [", i, ",", toBlockNumber, "] end")

		i = toBlockNumber + 1
	}

	logs.GetLogger().Info("scan block [", startBlockNumber, ",", endBlockNumber, "] end, scan block step:", scanBlockStep)

	return nil
}

func getPayment(ethClient *ethclient.Client, inputDataHex string, transaction types.Transaction) error {
	txReceipt, err := client.CheckTx(ethClient, transaction.Hash())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if txReceipt.Status != uint64(1) {
		return nil
	}

	method, err := txDataDecoderPayment.DecodeMethod(inputDataHex)
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

func getUnlock(ethClient *ethclient.Client, inputDataHex string, transaction types.Transaction) error {
	txReceipt, err := client.CheckTx(ethClient, transaction.Hash())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if txReceipt.Status != uint64(1) {
		return nil
	}

	method, err := txDataDecoderPayment.DecodeMethod(inputDataHex)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(method.Params) <= 0 {
		err = fmt.Errorf("method.Params is empty")
		return err
	}

	dealIdStr := method.Params[0].Value
	networkName := method.Params[1].Value
	//recipient := method.Params[2].Value

	filecoinNetwork := config.GetConfig().FilecoinNetwork
	if networkName != filecoinNetwork {
		return nil
	}

	dealId, err := strconv.ParseInt(dealIdStr, 10, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	block, err := ethClient.BlockByNumber(context.Background(), txReceipt.BlockNumber)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockAt := block.ReceivedAt.UTC().Unix()

	txHash := transaction.Hash().String()

	offlineDeal, err := models.GetOfflineDealByDealId(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	err = models.UpdateOfflineDealUnlockInfo(offlineDeal.Id, txHash, unlockAt)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func getRefund(ethClient *ethclient.Client, inputDataHex string, transaction types.Transaction) error {
	txReceipt, err := client.CheckTx(ethClient, transaction.Hash())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if txReceipt.Status != uint64(1) {
		return nil
	}

	method, err := txDataDecoderPayment.DecodeMethod(inputDataHex)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(method.Params) <= 0 {
		err = fmt.Errorf("method.Params is empty")
		return err
	}

	wCidsStr := strings.TrimRight(strings.TrimLeft(method.Params[0].Value, "["), "]")
	wCids := strings.Split(wCidsStr, " ")

	block, err := ethClient.BlockByNumber(context.Background(), txReceipt.BlockNumber)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	refundAt := block.ReceivedAt.UTC().Unix()

	for _, wCid := range wCids {
		err = models.UpdateTransactionRefundInfo(wCid, transaction.Hash().String(), refundAt)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
