package scheduler

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"multi-chain-storage/config"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	decoder "github.com/mingjingc/abi-decoder"
)

var txDataDecoderDao *decoder.ABIDecoder

type RunParam struct {
	LastScannedBlockNo int64
}

func ReadRunParam() (*RunParam, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error("cannot get home directory.", err)
		return nil, err
	}

	runParamFile := filepath.Join(homedir, ".swan/dao-runner/.run_param")
	contents, err := ioutil.ReadFile(runParamFile)
	if err != nil {
		logs.GetLogger().Error(err)
		runParam := RunParam{}
		runParam.LastScannedBlockNo = 0
		err = WriteRunParam(runParam.LastScannedBlockNo)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		return &runParam, nil
	}

	runParam := RunParam{}
	err = json.Unmarshal(contents, &runParam)
	if err != nil {
		logs.GetLogger().Error(err)
		runParam.LastScannedBlockNo = 0
		err = WriteRunParam(runParam.LastScannedBlockNo)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
	}

	return &runParam, nil
}

func WriteRunParam(lastScannedBlockNo int64) error {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error("cannot get home directory.", err)
		return err
	}

	runParamFile := filepath.Join(homedir, ".swan/dao-runner/.run_param")

	runParam := RunParam{LastScannedBlockNo: lastScannedBlockNo}

	content, err := json.MarshalIndent(runParam, "", " ")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	err = ioutil.WriteFile(runParamFile, content, 0644)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func initTxDataDecoderDao() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("cannot get home directory.", err)
	}

	contractPath := filepath.Join(homedir, ".swan/dao-runner/FilswanOracle.json")
	contractRead, err := ioutil.ReadFile(contractPath)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	myContractAbi := string(contractRead)

	txDataDecoder = decoder.NewABIDecoder()
	txDataDecoder.SetABI(myContractAbi)
}

func ScanPolygon4Dao() {
	for {
		runParam, err := ReadRunParam()
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		startBlockNumber := runParam.LastScannedBlockNo + 1
		err = scanPolygon(startBlockNumber)
		if err != nil {
			logs.GetLogger().Error(err)
		}

		time.Sleep(config.GetConfig().ScheduleRule.ScanPolygonIntervalSecond * time.Second)
	}
}

func scanPolygon(startBlockNumber int64) error {
	if txDataDecoder == nil {
		initTxDataDecoder()
	}

	ethClient, err := ethclient.Dial(config.GetConfig().Polygon.PolygonRpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	endBlockNumberUint64, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	endBlockNumber := int64(endBlockNumberUint64)
	scanBlockStep := config.GetConfig().Polygon.ScanPolygonBlockStep
	logs.GetLogger().Info("scan from block:", startBlockNumber, ", to block:", endBlockNumber, ", scan block step:", scanBlockStep)

	daoContractAddress := common.HexToAddress(config.GetConfig().Polygon.DaoContractAddress)
	logs.GetLogger().Info("dao contract address:", daoContractAddress.String())

	for i := startBlockNumber; i <= endBlockNumber; {
		toBlockNumber := i + scanBlockStep - 1
		if toBlockNumber > endBlockNumber {
			toBlockNumber = endBlockNumber
		}
		logs.GetLogger().Info("scan block [", i, ",", toBlockNumber, "] start")
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(i),
			ToBlock:   big.NewInt(toBlockNumber),
			Addresses: []common.Address{daoContractAddress},
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
			if strings.HasPrefix(inputDataHex, "4d043ef6") {
				err = getDaoSignature(ethClient, inputDataHex, transaction)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
			}
		}

		err = WriteRunParam(toBlockNumber)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		logs.GetLogger().Info("scan block [", i, ",", toBlockNumber, "] end")
		i = toBlockNumber + 1
	}

	return nil
}

func getDaoSignature(ethClient *ethclient.Client, inputDataHex string, transaction *types.Transaction) error {
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

	if len(method.Params) < 4 {
		err = fmt.Errorf("method.Params has not enough parameters,len(method.Params)=%d", len(method.Params))
		logs.GetLogger().Error(err)
		return err
	}

	dealIdStr := method.Params[1].Value
	networkName := method.Params[2].Value
	recipient := method.Params[3].Value

	filecoinNetwork := config.GetConfig().FilecoinNetwork
	if networkName != filecoinNetwork {
		return nil
	}

	dealId, err := strconv.ParseInt(dealIdStr, 10, 32)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	txHash := transaction.Hash().String()
	err = models.WriteDaoSignature(txHash, recipient, dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
