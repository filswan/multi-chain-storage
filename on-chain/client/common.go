package client

import (
	"context"
	"multi-chain-storage/config"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetEthClient() (*ethclient.Client, *rpc.Client, error) {
	polygonRpcUrl := config.GetConfig().Polygon.PolygonRpcUrl
	rpcClient, err := rpc.DialContext(context.Background(), polygonRpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	ethClient, err := ethclient.Dial(polygonRpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, rpcClient, err
	}

	return ethClient, rpcClient, nil
}
