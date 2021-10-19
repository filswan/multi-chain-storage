package utils

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"os"
	"payment-bridge/common/constants"
	"payment-bridge/logs"
	models2 "payment-bridge/models"
	"strconv"
	"time"
)

// GetEpochInMillis get current timestamp
func GetEpochInMillis() (millis int64) {
	nanos := time.Now().UnixNano()
	millis = nanos / 1000000
	return
}

func ReadContractAbiJsonFile(aptpath string) (string, error) {
	jsonFile, err := os.Open(aptpath)

	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	return string(byteValue), nil
}

func GetRewardPerBlock() *big.Int {
	rewardBig, _ := new(big.Int).SetString("35000000000000000000", 10) // the unit is wei
	return rewardBig
}

func CheckTx(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
retry:
	rp, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		if err == ethereum.NotFound {
			logs.GetLogger().Error("tx %v not found, check it later", tx.Hash().String())
			time.Sleep(1 * time.Second)
			goto retry
		} else {
			logs.GetLogger().Error("TransactionReceipt fail: %s", err)
			return nil, err
		}
	}
	return rp, nil
}

func GetFromAndToAddressByTxHash(client *ethclient.Client, chainID *big.Int, txHash common.Hash) (*addressInfo, error) {
	addrInfo := new(addressInfo)
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	addrInfo.AddrTo = tx.To().Hex()
	txMsg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	addrInfo.AddrFrom = txMsg.From().Hex()
	return addrInfo, nil
}

type addressInfo struct {
	AddrFrom string
	AddrTo   string
}

func GetOffsetByPagenumber(pageNumber, pageSize string) (int64, error) {
	pageNumberInt, err := strconv.ParseInt(pageNumber, 10, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	pageSizeInt, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	offset := (pageNumberInt - 1) * pageSizeInt
	return offset, nil
}

func GetStartBlockNo(networkName string, startBlockNoInConfig int64) int64 {
	var startScanBlockNo int64 = 1
	var whereCondition string

	switch networkName {
	case constants.NETWORK_TYPE_POLYGON:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_POLYGON + "'"
	case constants.NETWORK_TYPE_NBAI:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_NBAI + "'"
	case constants.NETWORK_TYPE_BSC:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_BSC + "'"
	case constants.NETWORK_TYPE_GOERLI:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_GOERLI + "'"
	case constants.NETWORK_TYPE_ETH:
		whereCondition = "network_type='" + constants.NETWORK_TYPE_ETH + "'"
	}

	if startBlockNoInConfig > 0 {
		startScanBlockNo = startBlockNoInConfig
	}

	blockScanRecord := new(models2.BlockScanRecord)
	blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber(whereCondition)
	if err != nil {
		logs.GetLogger().Error(err)
		startScanBlockNo = startBlockNoInConfig
	}

	if len(blockScanRecordList) > 0 {
		if blockScanRecordList[0].LastCurrentBlockNumber > startScanBlockNo {
			startScanBlockNo = blockScanRecordList[0].LastCurrentBlockNumber
		}
	}
	return startScanBlockNo
}
