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
	"strings"

	"github.com/ethereum/go-ethereum/common"
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

func GetPayments() error {
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

	endBlockNumber, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for i := startBlockNumber; i <= int64(endBlockNumber); i++ {
		err = GetPaymentsByBlockNumber(ethClient, i)
		if err != nil {
			logs.GetLogger().Error(err)
			break
		}

		network.LastScanBlockNumber = &i
		err = database.GetDB().Save(network).Error
		if err != nil {
			logs.GetLogger().Error(err)
			break
		}
	}

	return nil
}

func GetPaymentsByBlockNumber(ethClient *ethclient.Client, blockNumber int64) error {
	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, transaction := range block.Transactions() {
		inputDataHex := hex.EncodeToString(transaction.Data())
		if strings.HasPrefix(inputDataHex, "f4d98717") {
			err = getPayment4Transaction(ethClient, inputDataHex, transaction.Hash())
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}

	return err
}

func getPayment4Transaction(ethClient *ethclient.Client, inputDataHex string, txHash common.Hash) error {
	txReceipt, err := client.CheckTx(ethClient, txHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if txReceipt.Status != uint64(1) {
		return nil
	}

	if txReceipt.ContractAddress.String() != config.GetConfig().Polygon.PaymentContractAddress {
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
	if len(params) <= 0 {
		err = fmt.Errorf("params is empty")
		return err
	}

	wCid := params[0]
	if len(wCid) <= 1 {
		err = fmt.Errorf("wCid is empty")
		return err
	}

	wCid = wCid[1:]
	err = models.CreateTransaction4PayByWCid(wCid, txHash.String())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
