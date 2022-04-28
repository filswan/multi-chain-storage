package scheduler

import (
	"context"
	"fmt"
	"multi-chain-storage/common/constants"
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
	//offlineDeals, err := models.GetOfflineDealByDealId(87843)
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

	mcsPaymentReceiverAddress := common.HexToAddress(config.GetConfig().Polygon.McsPaymentReceiverAddress)

	filswanOracleSession, err := client.GetFilswanOracleSession(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockInterval := config.GetConfig().Polygon.UnlockIntervalMinute * time.Minute
	logs.GetLogger().Info("unlock interval is ", unlockInterval)

	unlockCnt := 0
	for _, offlineDeal := range offlineDeals {
		isUnlockable, err := checkUnlockable(ethClient, offlineDeal, filswanOracleSession, mcsPaymentReceiverAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		if !isUnlockable {
			logs.GetLogger().Info(getLog(offlineDeal, "not unlockable, please wait"))
			continue
		}

		if unlockCnt > 0 {
			logs.GetLogger().Info(getLog(offlineDeal, "sleeping "+unlockInterval.String()+" before unlock"))
			time.Sleep(unlockInterval)
		}

		srcFileUploads, err := setUnlockPayment(offlineDeal)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		unlockCnt = unlockCnt + 1
		_, err = unlockDeal(offlineDeal, ethClient, swanPaymentTransactor, mcsPaymentReceiverAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
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
	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	filecoinNetwork := config.GetConfig().FilecoinNetwork
	isPaymentAvailable, err := filswanOracleSession.IsCarPaymentAvailable(dealIdStr, filecoinNetwork, mcsPaymentReceiverAddress)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	if !isPaymentAvailable {
		logs.GetLogger().Info(getLog(offlineDeal, "payment is not available for recipient "+mcsPaymentReceiverAddress.String()))
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

		if blockInterval < config.GetConfig().Polygon.IntervalDaoUnlockBlock {
			msg := fmt.Sprintf("current block number:%d - dao block number:%d is less than block interval:%d", currentBlockNo, daoBlockNo, blockInterval)
			logs.GetLogger().Info(offlineDeal, msg)
			return false, nil
		}
	}

	return true, nil
}

func setUnlockPayment(offlineDeal *models.OfflineDeal) ([]*models.SourceFileUploadOut, error) {
	srcFileUploads, err := models.GetSourceFileUploadsByCarFileId(offlineDeal.CarFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	for _, srcFileUpload := range srcFileUploads {
		wCid := srcFileUpload.Uuid + srcFileUpload.PayloadCid
		lockPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
		srcFileUpload.LockedFeeBeforeUnlock = lockPayment.LockedFee
	}

	return srcFileUploads, nil
}

func updateUnlockPayment(offlineDeal *models.OfflineDeal, srcFileUploads []*models.SourceFileUploadOut) error {
	for _, srcFileUpload := range srcFileUploads {
		wCid := srcFileUpload.Uuid + srcFileUpload.PayloadCid
		lockPayment, err := client.GetLockedPaymentInfo(wCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		unlockAmount := srcFileUpload.LockedFeeBeforeUnlock.Sub(lockPayment.LockedFee)
		err = models.UpdateTransactionUnlockInfo(srcFileUpload.Id, unlockAmount)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}
	}

	return nil
}

func getLog(offlineDeal *models.OfflineDeal, messages ...string) string {
	text := fmt.Sprintf("id:%d,deal id:%d, deal file id:%d", offlineDeal.Id, offlineDeal.DealId, offlineDeal.CarFileId)
	if messages == nil {
		return text
	}

	for _, msg := range messages {
		text = text + "," + msg
	}

	return text
}

func unlockDeal(offlineDeal *models.OfflineDeal, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, mcsPaymentReceiverAddress common.Address) (*string, error) {
	tansactOpts, err := client.GetTransactOpts(ethClient, adminWalletPrivateKey, *adminWalletPublicKey)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	unlockStatusFailed := constants.OFFLINE_DEAL_STATUS_UNLOCK_FAILED

	filecoinNetwork := config.GetConfig().FilecoinNetwork
	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, filecoinNetwork, mcsPaymentReceiverAddress)
	txHash := ""
	if tx != nil {
		txHash = tx.Hash().Hex()
	}

	logs.GetLogger().Info(getLog(offlineDeal, txHash))
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockInfo(offlineDeal.Id, unlockStatusFailed, txHash, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	txReceipt, err := client.CheckTx(ethClient, tx)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockInfo(offlineDeal.Id, unlockStatusFailed, txHash, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed! txHash=%s, status:%d", tx.Hash().Hex(), txReceipt.Status)
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockInfo(offlineDeal.Id, unlockStatusFailed, txHash, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "txHash="+tx.Hash().Hex()))

	unlockStatusUnlocked := constants.OFFLINE_DEAL_STATUS_UNLOCKED
	err = models.UpdateOfflineDealUnlockInfo(offlineDeal.Id, unlockStatusUnlocked, txHash)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock successfully"))
	return &txHash, nil
}
