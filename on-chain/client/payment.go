package client

import (
	"fmt"
	"payment-bridge/on-chain/goBind"

	"github.com/filswan/go-swan-lib/logs"
)

func GetPaymentInfo(srcFilePayloadCid string) (*goBind.IPaymentMinimalTxInfo, error) {
	swanPaymentSession, err := GetPaymentSession()
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
