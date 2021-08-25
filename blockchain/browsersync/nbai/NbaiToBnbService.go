package nbai

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"payment-bridge/blockchain/nbaiclient"
	"payment-bridge/logs"
	"strings"
)

func changeNbaiToBnb() {
	fromAddress := common.HexToAddress("0x678AC296Af90001A1bEeC5e4632ED60E64E090aB")
	client := nbaiclient.WebConn.ConnWeb
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logs.GetLogger().Error(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
	}
	// callOpts.NoSend = true
	callOpts := new(bind.TransactOpts)
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = 800000
	callOpts.Context = context.Background()

	childManagerAddress := common.HexToAddress("0xa1A58C82bA4a0D8e856D8fad76bB72b04fd9736b")
	childInstance, _ := NewChildChainManagerContract(childManagerAddress, client)

	contractAbi, err := abi.JSON(strings.NewReader(NbaiAbiJson))
}
