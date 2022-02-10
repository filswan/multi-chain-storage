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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

func UnlockPayment() error {
	offlineDeals, err := models.GetOfflineDeals2BeUnlocked()
	//offlineDeals, err := models.GetOfflineDealByDealId(87416)
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

	mcpPaymentReceiverAddress := common.HexToAddress(config.GetConfig().Polygon.McpPaymentReceiverAddress)

	filswanOracleSession, err := client.GetFilswanOracleSession(ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	privateKey, publicKeyAddress, err := client.GetPrivateKeyPublicKey(constants.PRIVATE_KEY_ON_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, privateKey, *publicKeyAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockInterval := config.GetConfig().Polygon.UnlockIntervalMinute * time.Minute
	logs.GetLogger().Info("unlock interval is ", unlockInterval)

	unlockCnt := 0
	for _, offlineDeal := range offlineDeals {
		isUnlockable, err := checkUnlockable(offlineDeal, filswanOracleSession, mcpPaymentReceiverAddress)
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
		txHash, err := doUnlockDeal(offlineDeal, ethClient, swanPaymentTransactor, tansactOpts, mcpPaymentReceiverAddress)
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

func checkUnlockable(offlineDeal *models.OfflineDeal, filswanOracleSession *goBind.FilswanOracleSession, mcpPaymentReceiverAddress common.Address) (bool, error) {
	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	isPaymentAvailable, err := filswanOracleSession.IsCarPaymentAvailable(dealIdStr, mcpPaymentReceiverAddress)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return false, err
	}

	if !isPaymentAvailable {
		logs.GetLogger().Info(getLog(offlineDeal, "payment is not available for recipient "+mcpPaymentReceiverAddress.String()))
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

	blockTime := libutils.GetInt64FromStr(daoSignatures[0].BlockTime)
	if blockTime < 0 {
		err := fmt.Errorf(getLog(offlineDeal, "invalid block time:"+daoSignatures[0].BlockTime))
		logs.GetLogger().Error(err)
		return false, err
	}

	blockTime = blockTime * 1000
	curUtcMilliSec := utils.GetCurrentUtcMilliSecond()
	daoPassMilliSec := curUtcMilliSec - blockTime

	intervalDaoUnlockMinute := config.GetConfig().Polygon.IntervalDaoUnlockMinute
	if daoPassMilliSec < int64(intervalDaoUnlockMinute)*60*1000 {
		logs.GetLogger().Info(offlineDeal, "dao just signed, please wait")
		return false, nil
	}

	return true, nil
}

func updateUnlockPayment(offlineDeal *models.OfflineDeal, txHash string, rpcClient *rpc.Client) error {
	srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.DealFileId)
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

		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return err
		}

		unlockPayment.LockedFeeBeforeUnlock = lockedPayment.LockedFee

		unlockPayment.TokenAddress = lockedPayment.TokenAddress
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

func doUnlockDeal(offlineDeal *models.OfflineDeal, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, tansactOpts *bind.TransactOpts, mcpPaymentReceiverAddress common.Address) (*string, error) {
	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	unlockStatusFailed := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCK_FAILED

	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, mcpPaymentReceiverAddress)
	txHash := ""
	if tx != nil {
		txHash = tx.Hash().Hex()
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
	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "recipient.Hex()="+mcpPaymentReceiverAddress.Hex()))

	unlockStatusUnlocked := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCKED
	err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusUnlocked)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock successfully"))
	return &txHash, nil
}
