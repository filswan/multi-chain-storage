package client

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/models"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

func SaveEventUnlockPayment(receipt *types.Receipt, unlockStatus string, oflineDeal *models.OfflineDeal, txHash string) error {
	//contractUnlockFunctionSignature := config.GetConfig().Polygon.ContractUnlockFunctionSignature
	//
	//contractAbi, err := GetContractAbi()
	//if err != nil {
	//	logs.GetLogger().Error(err)
	//	return err
	//}
	//
	//for _, vLog := range logsInChain {
	//	logs.GetLogger().Info("dealId:", oflineDeal.DealId, ",vLog.Topics[0].Hex():", vLog.Topics[0].Hex())
	//	if vLog.Topics[0].Hex() != contractUnlockFunctionSignature {
	//		continue
	//	}
	//
	//
	//	dataList, err := contractAbi.Unpack("unlockCarPayment", vLog.Data)
	//	if err != nil {
	//		logs.GetLogger().Error("dealId:", oflineDeal.DealId, err)
	//		continue
	//	}
	dealFile, err := models.GetDealFileByDealId(87323)
	if err == nil {
		logs.GetLogger().Info(dealFile)
		return err
	}

	event := new(models.EventUnlockPayment)
	event.DealId = oflineDeal.DealId
	event.TxHash = txHash
	event.TokenAddress = constants.COIN_ADDRESS_USDC
	event.PayloadCid = dealFile.PayloadCid
	//event.UnlockToAdminAmount = dataList[2].(*big.Int).String()
	//event.UnlockToUserAmount = dataList[3].(*big.Int).String()
	//event.UnlockToAdminAddress = dataList[4].(common.Address).Hex()
	//event.UnlockToUserAddress = dataList[5].(common.Address).Hex()
	event.UnlockTime = strconv.FormatInt(utils.GetCurrentUtcMilliSecond(), 10)
	event.BlockNo = receipt.BlockNumber.String()
	event.CreateAt = utils.GetCurrentUtcMilliSecond()
	event.UnlockStatus = unlockStatus

	network, err := models.GetNetworkByName(constants.NETWORK_NAME_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		event.NetworkId = network.ID
	}
	coin, err := models.FindCoinByCoinAddress(event.TokenAddress)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		event.CoinId = coin.ID
	}

	err = database.SaveOne(event)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	//}

	return nil
}
