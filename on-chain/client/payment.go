package client

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"multi-chain-storage/config"
	"multi-chain-storage/on-chain/goBind"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
	decoder "github.com/mingjingc/abi-decoder"
	"github.com/shopspring/decimal"
)

type LockedPayment struct {
	TokenAddress string
	MinPayment   string
	LockedFee    decimal.Decimal
	AddressFrom  string
	AddressTo    string
	Deadline     int64
	Size         int64
	BlockNumber  int64
}

func IsLockedPaymentExists(srcFilePayloadCid string) (*bool, error) {
	swanPaymentSession, err := GetSwanPaymentSession()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	paymentInfo, err := swanPaymentSession.GetLockedPaymentInfo(srcFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &paymentInfo.IsExisted, nil
}

func GetLockedPaymentInfo(wCid string) (*LockedPayment, error) {
	swanPaymentSession, err := GetSwanPaymentSession()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	paymentInfo, err := swanPaymentSession.GetLockedPaymentInfo(wCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !paymentInfo.IsExisted {
		logs.GetLogger().Info("payment not exists for w_cid:", wCid)
		return nil, nil
	}

	lockedFee, err := decimal.NewFromString(paymentInfo.LockedFee.String())
	if err != nil {
		logs.GetLogger().Error(err)
		lockedFee = decimal.NewFromInt(-1)
	}

	lockedPayment := LockedPayment{
		TokenAddress: paymentInfo.Token.Hex(),
		MinPayment:   paymentInfo.MinPayment.String(),
		LockedFee:    lockedFee,
		AddressFrom:  paymentInfo.Owner.String(),
		AddressTo:    paymentInfo.Recipient.String(),
		Deadline:     paymentInfo.Deadline.Int64(),
		Size:         paymentInfo.Size.Int64(),
		BlockNumber:  paymentInfo.BlockNumber.Int64(),
	}

	return &lockedPayment, nil
}

func GetSwanPaymentSession() (*goBind.SwanPaymentSession, error) {
	ethClient, _, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	paymentContractAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)
	logs.GetLogger().Info("payment contract address:", paymentContractAddress.String())

	swanPayment, err := goBind.NewSwanPayment(paymentContractAddress, ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	swanPaymentSession := &goBind.SwanPaymentSession{
		Contract: swanPayment,
	}

	return swanPaymentSession, nil
}

func GetPaymentByBlockNumberWCid(blockNumber int64, wCid string) (*string, error) {
	ethClient, _, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	block, err := ethClient.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	contractRead, er := ioutil.ReadFile("./contracts/abi/FilswanOracle.json")
	if er != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	myContractAbi := string(contractRead)

	txDataDecoder := decoder.NewABIDecoder()
	txDataDecoder.SetABI(myContractAbi)

	for _, transaction := range block.Transactions() {
		txHash := transaction.Hash()
		transaction, _, err := ethClient.TransactionByHash(context.Background(), txHash)
		if er != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		inputHex := hex.EncodeToString(transaction.Data())
		if !strings.HasPrefix(inputHex, "4d043ef6") {
			continue
		}
		method, err1 := txDataDecoder.DecodeMethod(inputHex)
		if err1 != nil {
			continue
		}

		for _, param := range method.Params {
			if param.Name == "w_cid" {
				if strings.EqualFold(wCid, param.Value) {
					txHashStr := txHash.String()
					return &txHashStr, nil
				}
			}
		}
	}

	err = fmt.Errorf("cannot find tx hash in block%d, for w_cid:%s", blockNumber, wCid)
	logs.GetLogger().Error(err)

	return nil, err
}
