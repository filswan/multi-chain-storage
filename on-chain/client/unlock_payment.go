package client

import (
	"context"
	"payment-bridge/common/constants"
	"payment-bridge/database"
	"payment-bridge/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
)

func SaveEventUnlockPaymentFromTxHash(receipt *types.Receipt, recipient common.Address, oflineDeal *models.OfflineDeal) error {
	ethClient, rpcClient, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	txHash := receipt.TxHash.Hex()

	dealFile, err := models.GetDealFileByDealId(oflineDeal.DealId)
	if err != nil {
		logs.GetLogger().Info(err)
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
	unlockPayment.UnlockFromAddress = recipient.Hex()
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
	//block, err := ethClient.BlockByHash(context.Background(), *rpcTransaction.BlockHash)
	//if err != nil {
	//	logs.GetLogger().Error(err)
	//} else {
	//	//unlockPayment.UnlockTime = block.Time()
	//}

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
