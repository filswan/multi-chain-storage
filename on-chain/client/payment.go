package client

import (
	"fmt"
	"multi-chain-storage/config"
	"multi-chain-storage/on-chain/goBind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

type LockedPayment struct {
	TokenAddress string
	MinPayment   string
	LockedFee    decimal.Decimal
	AddressFrom  string
	AddressTo    string
	Deadline     string
	Size         int64
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

func GetLockedPaymentInfo(srcFilePayloadCid string) (*LockedPayment, error) {
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
		Deadline:     paymentInfo.Deadline.String(),
		Size:         paymentInfo.Size.Int64(),
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
