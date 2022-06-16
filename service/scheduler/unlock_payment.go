package scheduler

import (
	"context"
	"fmt"
	"multi-chain-storage/config"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/on-chain/goBind"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
)

func UnlockPayment() error {
	offlineDeals, err := models.GetOfflineDeals2BeUnlocked()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(offlineDeals) == 0 {
		logs.GetLogger().Info("no deal to be unlocked")
		return nil
	}

	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	swanPaymentTransactor, err := client.GetSwanPaymentTransactor(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	paymentRecipientAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentRecipientAddress)

	filswanOracleSession, err := client.GetFilswanOracleSession(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockIntervalSecond := config.GetConfig().Polygon.UnlockIntervalSecond * time.Second
	logs.GetLogger().Info("unlock interval is ", unlockIntervalSecond)

	unlockCnt := 0
	for _, offlineDeal := range offlineDeals {
		isUnlockable, err := checkUnlockable(ethClient, offlineDeal, filswanOracleSession, paymentRecipientAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		if !isUnlockable {
			logs.GetLogger().Info(getLog(offlineDeal, "not unlockable, please wait"))
			continue
		}

		if unlockCnt > 0 {
			logs.GetLogger().Info(getLog(offlineDeal, "sleeping "+unlockIntervalSecond.String()+" before unlock"))
			time.Sleep(unlockIntervalSecond)
		}

		srcFileUploads, err := setUnlockPayment(offlineDeal)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		unlockCnt = unlockCnt + 1
		txHash, err := unlockDeal(offlineDeal, ethClient, swanPaymentTransactor, paymentRecipientAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		err = models.UpdateOfflineDealUnlockInfo(offlineDeal.Id, *txHash)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		err = models.UpdateCarFileDealSuccess(offlineDeal.CarFileId)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		err = updateUnlockPayment(offlineDeal, srcFileUploads)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		}
	}
	return nil
}

func checkUnlockable(ethClient *ethclient.Client, offlineDeal *models.OfflineDeal, filswanOracleSession *goBind.FilswanOracleSession, mcsPaymentReceiverAddress common.Address) (bool, error) {
	if offlineDeal.DealId == nil || *offlineDeal.DealId <= 0 {
		err := fmt.Errorf("valid deal id must be greater than 0")
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	dealIdStr := strconv.FormatInt(*offlineDeal.DealId, 10)
	filecoinNetwork := config.GetConfig().FilecoinNetwork
	isPaymentAvailable, err := filswanOracleSession.IsCarPaymentAvailable(dealIdStr, filecoinNetwork, mcsPaymentReceiverAddress)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	if !isPaymentAvailable {
		logs.GetLogger().Info(getLog(offlineDeal, "payment is not available for recipient "+mcsPaymentReceiverAddress.String(), "network:"+filecoinNetwork))
		return false, nil
	}

	filswanOracleTransactions, err := filswanOracleSession.GetSignatureList(dealIdStr, filecoinNetwork)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	if len(filswanOracleTransactions) == 0 {
		logs.GetLogger().Info(getLog(offlineDeal, "no dao sigatures yet"))
		return false, nil
	}

	currentBlockNo, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	for _, filswanOracleTransaction := range filswanOracleTransactions {
		daoBlockNo := filswanOracleTransaction.BlockNumber.Uint64()
		blockInterval := int64(currentBlockNo - daoBlockNo)

		if blockInterval < config.GetConfig().Polygon.DaoUnlockIntervalBlock {
			msg := fmt.Sprintf("current block number:%d - dao block number:%d is less than block interval:%d", currentBlockNo, daoBlockNo, blockInterval)
			logs.GetLogger().Info(offlineDeal, msg)
			return false, nil
		}
	}

	return true, nil
}

func setUnlockPayment(offlineDeal *models.OfflineDeal) ([]*models.SourceFileUploadOut, error) {
	srcFileUploads, err := models.GetSourceFileUploadOutsByCarFileId(offlineDeal.CarFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	for _, srcFileUpload := range srcFileUploads {
		wCid := srcFileUpload.Uuid + srcFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		if lockedPayment == nil {
			err := fmt.Errorf("payment not exists for w_cid:%s", wCid)
			logs.GetLogger().Info(err)
			return nil, err
		}
		srcFileUpload.LockedFeeBeforeUnlock = lockedPayment.LockedFee
	}

	return srcFileUploads, nil
}

func updateUnlockPayment(offlineDeal *models.OfflineDeal, srcFileUploads []*models.SourceFileUploadOut) error {
	for _, srcFileUpload := range srcFileUploads {
		wCid := srcFileUpload.Uuid + srcFileUpload.PayloadCid
		lockedPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		if lockedPayment == nil {
			logs.GetLogger().Error("payment not exists for w_cid:", wCid)
			continue
		}

		unlockAmount := srcFileUpload.LockedFeeBeforeUnlock.Sub(lockedPayment.LockedFee)
		err = models.UpdateTransactionUnlockInfo(srcFileUpload.Id, unlockAmount)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}
	}

	return nil
}

func getLog(offlineDeal *models.OfflineDeal, messages ...string) string {
	text := fmt.Sprintf("id:%d,deal id:%d, car file id:%d", offlineDeal.Id, *offlineDeal.DealId, offlineDeal.CarFileId)
	if messages == nil {
		return text
	}

	for _, msg := range messages {
		text = text + "," + msg
	}

	return text
}

func unlockDeal(offlineDeal *models.OfflineDeal, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, mcsPaymentReceiverAddress common.Address) (*string, error) {
	if offlineDeal.DealId == nil || *offlineDeal.DealId <= 0 {
		err := fmt.Errorf("valid deal id must be greater than 0")
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, adminWalletPrivateKey, *adminWalletPublicKey)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	dealIdStr := strconv.FormatInt(*offlineDeal.DealId, 10)
	filecoinNetwork := config.GetConfig().FilecoinNetwork
	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, filecoinNetwork, mcsPaymentReceiverAddress)
	txHash := ""
	if tx != nil {
		txHash = tx.Hash().Hex()
	}

	logs.GetLogger().Info(getLog(offlineDeal, txHash))
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, "unlock failed, tx hash:"+txHash, err.Error()))
		return nil, err
	}

	txReceipt, err := client.CheckTx(ethClient, tx.Hash())
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, "unlock failed, tx hash:"+txHash, err.Error()))
		return nil, err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed, tx hash:%s, status:%d", txHash, txReceipt.Status)
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock succeeded, tx hash:"+txHash))

	return &txHash, nil
}
