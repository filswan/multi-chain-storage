package services

import (
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
)

func GetPaymentInfo(srcFilePayloadCid string) (*models.EventLockPayment, error) {
	rpcUrl := "https://polygon-mumbai.g.alchemy.com/v2/XnUfDZGQd7ACRbi5G33aUBjQ5KmlsHm-" // config.GetConfig().PolygonMainnetNode.RpcUrl
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		logs.GetLogger().Error("Try to reconnect block chain node" + rpcUrl + " ...")
		return nil, err
	}

	swanPaymentSession := goBind.SwanPaymentSession{}
	paymentGatewayAddress := common.HexToAddress(polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress)

	paymentGatewayInstance, err := goBind.NewSwanPayment(paymentGatewayAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	swanPaymentSession.Contract = paymentGatewayInstance

	paymentInfo, err := swanPaymentSession.GetLockedPaymentInfo(srcFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	logs.GetLogger().Info(paymentInfo.LockedFee.String())
	logs.GetLogger().Info(paymentInfo.Token.String())

	var event *models.EventLockPayment
	if paymentInfo.IsExisted {
		event = new(models.EventLockPayment)
		event.AddressFrom = paymentInfo.Owner.String()
		event.AddressTo = paymentInfo.Recipient.String()
		event.LockedFee = paymentInfo.LockedFee.String()
		database.SaveOne(event)
	}

	return event, nil
}
