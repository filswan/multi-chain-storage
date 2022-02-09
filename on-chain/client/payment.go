package client

import (
	"fmt"
	"payment-bridge/config"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
)

func GetPaymentInfo(srcFilePayloadCid string) (*goBind.IPaymentMinimalTxInfo, error) {
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

	if !paymentInfo.IsExisted {
		err := fmt.Errorf("payment for source file with payload_cid:%s not exists", srcFilePayloadCid)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &paymentInfo, nil
}

func GetSwanPaymentSession() (*goBind.SwanPaymentSession, error) {
	ethClient, _, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	paymentContractAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)

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
