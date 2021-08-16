package goerliclient

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"payment-bridge/scan/config"
	"payment-bridge/scan/logs"
	"time"
)

/**
 * created on 08/10/21.
 * author: nebula-ai-zhiqiang
 * Copyright defined in payment-bridge/LICENSE
 */

type ConnSetup struct {
	ConnWeb *ethclient.Client
}

//setup eth connection
var WebConn = new(ConnSetup)

func ClientInit() {

	for {
		rpcUrl := config.GetConfig().GoerliMainnetNode.RpcUrl
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
