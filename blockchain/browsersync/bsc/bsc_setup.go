package bsc

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"payment-bridge/logs"
	"time"
)

type ConnSetup struct {
	ConnWeb *ethclient.Client
}

var WebConn = new(ConnSetup)

func ClientInit() {
	for {
		rpcUrl := GetConfig().BscMainnetNode.RpcUrl
		client, err := ethclient.Dial(rpcUrl)
		if err != nil {
			logs.GetLogger().Error("Try to reconnect block chain node" + rpcUrl + " ...")
			time.Sleep(time.Second * 5)
		} else {
			WebConn.ConnWeb = client
			break
		}
	}
}

func (conn *ConnSetup) GetBlockNumber() (*big.Int, error) {
	//return conn.ConnWeb.Eth.GetBlockNumber()
	block, err := conn.ConnWeb.HeaderByNumber(context.Background(), nil)
	if block != nil {
		return block.Number, err
	}
	return nil, err
}

func init() {
	ClientInit()
}
