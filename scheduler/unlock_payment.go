package scheduler

import (
	"context"
	"fmt"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

func UnlockPayment() error {
	//offlineDeals, err := models.GetOfflineDeals2BeUnlocked()
	offlineDeals, err := models.GetOfflineDealByDealId(87329)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(offlineDeals) == 0 {
		logs.GetLogger().Info("no deal to be unlocked")
		return nil
	}

	ethClient, rpcClient, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	recieverAddress, swanPaymentTransactor, err := client.GetSwanPaymentTransactor(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	filswanOracleSession, err := client.GetFilswanOracleSession(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockInterval := config.GetConfig().Polygon.UnlockIntervalMinute * time.Minute
	logs.GetLogger().Info("unlock interval is ", unlockInterval)
	for i, offlineDeal := range offlineDeals {
		if i > 0 {
			logs.GetLogger().Info(getLog(offlineDeal, "sleeping "+unlockInterval.String()+" before unlock"))
			time.Sleep(unlockInterval)
		}

		logs.GetLogger().Info(getLog(offlineDeal, "start to unlock"))

		err = setUnlockPayment(offlineDeal)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		txReceipt, err := unlockDeal(filswanOracleSession, offlineDeal, ethClient, swanPaymentTransactor, tansactOpts, *recieverAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		err = updateUnlockPayment(offlineDeal, txReceipt, rpcClient)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		}
	}
	return nil
}

func updateUnlockPayment(offlineDeal *models.OfflineDeal, receipt *types.Receipt, rpcClient *rpc.Client) error {
	srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.DealFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return err
	}

	txHash := receipt.TxHash.Hex()

	var rpcTransaction *models.RpcTransaction
	err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcFiles {
		paymentInfo, err := client.GetPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		lockedFee, err := decimal.NewFromString(paymentInfo.LockedFee.String())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		blockNo := ""
		if rpcTransaction.BlockNumber != nil {
			blockNo = *rpcTransaction.BlockNumber
		}

		err = models.UpdateUnlockAmount(srcFile.ID, offlineDeal.DealId, txHash, blockNo, lockedFee)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}
	}

	return nil
}

func setUnlockPayment(offlineDeal *models.OfflineDeal) error {
	srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.DealFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return err
	}

	for _, srcFile := range srcFiles {
		unlockPayment := models.EventUnlockPayment{
			PayloadCid:   srcFile.PayloadCid,
			SourceFileId: &srcFile.ID,
			DealId:       offlineDeal.DealId,
		}

		paymentInfo, err := client.GetPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return err
		}

		lockedFee, err := decimal.NewFromString(paymentInfo.LockedFee.String())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return err
		} else {
			unlockPayment.LockedFeeBeforeUnlock = lockedFee
		}

		unlockPayment.TokenAddress = paymentInfo.Token.Hex()
		//unlockPayment.UnlockFromAddress = paymentInfo.Owner.String()
		//unlockPayment.UnlockToAdminAddress = paymentInfo.Recipient.String()
		unlockPayment.UnlockTime = utils.GetCurrentUtcMilliSecond()
		coin, err := models.FindCoinByCoinAddress(unlockPayment.TokenAddress)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			unlockPayment.CoinId = coin.ID
			unlockPayment.NetworkId = coin.NetworkId
		}

		err = database.SaveOne(&unlockPayment)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}
	}

	return nil
}

func getLog(offlineDeal *models.OfflineDeal, messages ...string) string {
	text := fmt.Sprintf("id:%d,deal id:%d, deal file id:%d", offlineDeal.Id, offlineDeal.DealId, offlineDeal.DealFileId)
	if messages == nil {
		return text
	}

	for _, msg := range messages {
		text = text + "," + msg
	}

	return text
}

func unlockDeal(filswanOracleSession *goBind.FilswanOracleSession, offlineDeal *models.OfflineDeal, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, tansactOpts *bind.TransactOpts, recieverAddress common.Address) (*types.Receipt, error) {
	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	isPaymentAvailable, err := filswanOracleSession.IsCarPaymentAvailable(dealIdStr, recieverAddress)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	if !isPaymentAvailable {
		logs.GetLogger().Info(getLog(offlineDeal, "payment is not available for recipient "+recieverAddress.String()))
		return nil, nil
	}

	unlockStatusFailed := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCK_FAILED

	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, recieverAddress)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	txReceipt, err := client.CheckTx(ethClient, tx)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, tx.Hash().Hex(), err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed! txHash=%s", tx.Hash().Hex())
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, tx.Hash().Hex(), err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "txHash="+tx.Hash().Hex()))
	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "recipient.Hex()="+recieverAddress.Hex()))

	unlockStatusUnlocked := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCKED
	err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusUnlocked)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock successfully"))
	return txReceipt, nil
}
