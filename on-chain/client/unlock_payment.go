package client

import (
	"context"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/models"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

func SaveEventUnlockPayment(receipt *types.Receipt, oflineDeal *models.OfflineDeal, txHash string) error {
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
	dealFile, err := models.GetDealFileByDealId(oflineDeal.DealId)
	if err == nil {
		logs.GetLogger().Info(dealFile)
		return err
	}

	receipt.TxHash.Hex()
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
	event.UnlockStatus = constants.TRANSACTION_STATUS_SUCCESS

	coin, err := models.FindCoinByCoinAddress(event.TokenAddress)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		event.CoinId = coin.ID
		event.NetworkId = coin.NetworkId
	}

	err = database.SaveOne(event)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	//}

	return nil
}

func SaveEventUnlockPaymentFromTxHash(receipt *types.Receipt, recipent string, oflineDeal *models.OfflineDeal) error {
	ethClient, rpcClient, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	txHash := receipt.TxHash.Hex()

	dealFile, err := models.GetDealFileByDealId(oflineDeal.DealId)
	if err == nil {
		logs.GetLogger().Info(dealFile)
		return err
	}

	var rpcTransaction *models.RpcTransaction
	err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	transaction, _, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	var unlockPayment models.EventUnlockPayment
	unlockPayment.TxHash = txHash
	unlockPayment.UnlockFromAddress = recipent
	unlockPayment.PayloadCid = dealFile.PayloadCid
	wfilCoinId, err := models.FindCoinByFullName(constants.COIN_NAME_USDC)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		unlockPayment.CoinId = wfilCoinId.ID
		unlockPayment.NetworkId = wfilCoinId.NetworkId
		unlockPayment.TokenAddress = wfilCoinId.Address
	}
	unlockPayment.DealId = oflineDeal.DealId
	block, err := ethClient.BlockByHash(context.Background(), *rpcTransaction.BlockHash)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		unlockPayment.UnlockTime = strconv.FormatUint(block.Time(), 10)
	}

	logs.GetLogger().Info("*rpcTransaction.BlockNumber:", *rpcTransaction.BlockNumber)
	logs.GetLogger().Info("receipt.BlockNumber:", receipt.BlockNumber)

	unlockPayment.BlockNo = receipt.BlockNumber.String()
	unlockPayment.UnlockStatus = constants.TRANSACTION_STATUS_SUCCESS
	addrInfo, err := GetFromAndToAddressByTxHash(ethClient, transaction.ChainId(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		unlockPayment.UnlockFromAddress = addrInfo.AddrFrom
	}
	err = database.SaveOneWithTransaction(unlockPayment)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
