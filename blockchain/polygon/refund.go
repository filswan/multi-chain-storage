package polygon

import (
	"context"
	"math/big"
	"payment-bridge/config"
	"payment-bridge/on-chain/client"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
)

func ScanRefundEventOnPolygon(blockNoFrom, blockNoTo int64) error {
	contractAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)
	refundFunctionSignature := ""

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNoFrom),
		ToBlock:   big.NewInt(blockNoTo),
		Addresses: []common.Address{
			contractAddress,
		},
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

	for _, vLog := range logsInChain {
		if vLog.Topics[0].Hex() != refundFunctionSignature {
			continue
		}

	}
	return nil
}
