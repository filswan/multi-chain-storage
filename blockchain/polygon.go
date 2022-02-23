package blockchain

import (
	"context"
	"math/big"
	"multi-chain-storage/config"
	"multi-chain-storage/on-chain/client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
)

func ScanEventOnPolygon(blockNoFrom, blockNoTo int64) error {
	contractAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)
	refundFunctionSignature := ""

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{contractAddress},
	}

	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	logsInChain, err := ethClient.FilterLogs(context.Background(), query)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	//swanPaymentFilter, err := client.GetSwanPaymentFilterer()
	//if err != nil {
	//	logs.GetLogger().Error(err)
	//	return err
	//}

	for _, vLog := range logsInChain {
		if vLog.Topics[0].Hex() != refundFunctionSignature {
			continue
		}
	}
	return nil
}
