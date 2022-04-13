package scheduler

import (
	"context"
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/on-chain/goBind"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/robfig/cron"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
)

func CreateScheduler4UnlockPayment() {
	c := cron.New()
	name := "unlock payment"
	rule := config.GetConfig().ScheduleRule.UnlockPaymentRule
	mutex := &sync.Mutex{}

	err := c.AddFunc(rule, func() {
		logs.GetLogger().Info(name, " start")

		mutex.Lock()
		logs.GetLogger().Info(name, " running")
		err := UnlockPayment()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		mutex.Unlock()
		logs.GetLogger().Info(name, " end")
	})

	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	c.Start()
}

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

	ethClient, rpcClient, err := client.GetEthClient()
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

		logs.GetLogger().Info(getLog(offlineDeal, "start to unlock"))

		err = setUnlockPayment(offlineDeal)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		unlockCnt = unlockCnt + 1
		txHash, err := doUnlockDeal(offlineDeal, ethClient, swanPaymentTransactor, mcsPaymentReceiverAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		if txHash == nil {
			logs.GetLogger().Info(getLog(offlineDeal, " no tx hash returned"))
			continue
		}

		err = updateUnlockPayment(offlineDeal, *txHash, rpcClient)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		}
	}
	return nil
}

func getDaoSignatures(ethClient *ethclient.Client, offlineDeal *models.OfflineDeal) error {
	dealFile, err := models.GetCarFileById(offlineDeal.CarFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	filswanOracleSession, err := client.GetFilswanOracleSession(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	filecoinNetwork := config.GetConfig().FilecoinNetwork
	daoSignatures, err := filswanOracleSession.GetSignatureList(dealIdStr, filecoinNetwork)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, daoSingature := range daoSignatures {
		eventDaoSignature := models.EventDaoSignature{
			Recipient:   daoSingature.Recipient.Hex(),
			PayloadCid:  dealFile.PayloadCid,
			DealId:      offlineDeal.DealId,
			DaoAddress:  daoSingature.Signer.Hex(),
			BlockNo:     daoSingature.BlockNumber.Uint64(),
			BlockTime:   daoSingature.Timestamp.String(),
			DaoPassTime: daoSingature.Timestamp.String(),
			Status:      daoSingature.Status,
		}
		err = database.SaveOne(eventDaoSignature)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		logs.GetLogger().Info(daoSingature.Paid)
		logs.GetLogger().Info(daoSingature.Terms)
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

	daoSignatures, err := models.GetEventDaoSignaturesByDealId(offlineDeal.DealId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	if len(daoSignatures) == 0 {
		logs.GetLogger().Info("no dao sigatures yet")
		return false, nil
	}

	currentBlockNo, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	blockInterval := int64(currentBlockNo - daoSignatures[0].BlockNo)

	if blockInterval < config.GetConfig().Polygon.IntervalDaoUnlockBlock {
		msg := fmt.Sprintf("current block number:%d minus last dao block number:%d is less than block interval:%d", currentBlockNo, daoSignatures[0].BlockNo, blockInterval)
		logs.GetLogger().Info(offlineDeal, msg)
		return false, nil
	}

	return true, nil
}

func updateUnlockPayment(offlineDeal *models.OfflineDeal, txHash string, rpcClient *rpc.Client) error {
	srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.CarFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return err
	}

	var rpcTransaction *models.RpcTransaction
	err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcFiles {
		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}

		blockNo := ""
		if rpcTransaction.BlockNumber != nil {
			blockNo = *rpcTransaction.BlockNumber
		}

		err = models.UpdateUnlockAmount(srcFile.ID, offlineDeal.DealId, txHash, blockNo, lockedPayment.LockedFee)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			continue
		}
	}

	return nil
}

func setUnlockPayment(offlineDeal *models.OfflineDeal) error {
	srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.CarFileId)
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

		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return err
		}

		unlockPayment.LockedFeeBeforeUnlock = lockedPayment.LockedFee

		unlockPayment.TokenAddress = lockedPayment.TokenAddress
		unlockPayment.UnlockTime = utils.GetCurrentUtcMilliSecond()
		coin, err := models.GetCoinByAddress(unlockPayment.TokenAddress)
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
	text := fmt.Sprintf("id:%d,deal id:%d, deal file id:%d", offlineDeal.Id, offlineDeal.DealId, offlineDeal.CarFileId)
	if messages == nil {
		return text
	}

	for _, msg := range messages {
		text = text + "," + msg
	}

	return text
}

func doUnlockDeal(offlineDeal *models.OfflineDeal, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, mcsPaymentReceiverAddress common.Address) (*string, error) {
	privateKey, publicKeyAddress, err := client.GetPrivateKeyPublicKey(constants.PRIVATE_KEY_ON_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, privateKey, *publicKeyAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	unlockStatusFailed := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCK_FAILED

	filecoinNetwork := config.GetConfig().FilecoinNetwork
	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, filecoinNetwork, mcsPaymentReceiverAddress)
	txHash := ""
	if tx != nil {
		txHash = tx.Hash().Hex()
	} else {
		logs.GetLogger().Info("tx hash is nil")
	}

	logs.GetLogger().Info(getLog(offlineDeal, txHash))
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, "txHash="+txHash, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	if tx == nil {
		err := fmt.Errorf("tx hash is nil")
		logs.GetLogger().Error(err)
		return nil, err
	}

	txReceipt, err := client.CheckTx(ethClient, tx)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, "txHash="+txHash, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed! txHash=%s", tx.Hash().Hex())
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, "txHash="+txHash, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "txHash="+tx.Hash().Hex()))
	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "recipient.Hex()="+mcsPaymentReceiverAddress.Hex()))

	unlockStatusUnlocked := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCKED
	err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusUnlocked)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock successfully"))
	return &txHash, nil
}
