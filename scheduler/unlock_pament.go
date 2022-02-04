package scheduler

import (
	"fmt"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/robfig/cron"
)

func CreateUnlockScheduler() {
	Mutex := &sync.Mutex{}
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UnlockPaymentRule, func() {
		logs.GetLogger().Info("start")
		Mutex.Lock()
		err := UnlockPayment()
		Mutex.Unlock()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		logs.GetLogger().Info("end")
	})
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	c.Start()
}

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

	recipient, swanPaymentTransactor, err := client.GetSwanPaymentTransactor(ethClient)
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
		//if i > 0 {
		//	logs.GetLogger().Info(getLog(offlineDeal, "sleeping before unlock"))
		//	time.Sleep(unlockInterval * time.Minute)
		//}

		logs.GetLogger().Info(i)
		if offlineDeal.DealId != 87302 {
			continue
		}

		logs.GetLogger().Info(getLog(offlineDeal, "start to unlock"))
		unlockStatus, err := unlockDeal(filswanOracleSession, offlineDeal, ethClient, swanPaymentTransactor, tansactOpts, *recipient)
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		}

		if unlockStatus != nil {
			err = refund(offlineDeal, swanPaymentTransactor, tansactOpts)
			if err != nil {
				logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
				continue
			}
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

func unlockDeal(filswanOracleSession *goBind.FilswanOracleSession, offlineDeal *models.OfflineDeal, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, tansactOpts *bind.TransactOpts, recipient common.Address) (*string, error) {
	dealIdStr := strconv.FormatInt(offlineDeal.DealId, 10)
	isPaymentAvailable, err := filswanOracleSession.IsCarPaymentAvailable(dealIdStr, recipient)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return nil, err
	}

	if !isPaymentAvailable {
		logs.GetLogger().Info(getLog(offlineDeal, "payment is not available for recipient "+recipient.String()))
		return nil, nil
	}

	unlockStatusFailed := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCK_FAILED

	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, recipient)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return &unlockStatusFailed, err
	}

	logs.GetLogger().Info(tx.Hash().Hex())

	txReceipt, err := client.CheckTx(ethClient, tx)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, tx.Hash().Hex(), err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return &unlockStatusFailed, err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed! txHash=%s", tx.Hash().Hex())
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))

		err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusFailed, tx.Hash().Hex(), err.Error())
		if err != nil {
			logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
			return nil, err
		}

		return &unlockStatusFailed, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock success", "txHash="+tx.Hash().Hex()))

	unlockStatusUnlocked := constants.OFFLINE_DEAL_UNLOCK_STATUS_UNLOCKED
	err = models.UpdateOfflineDealUnlockStatus(offlineDeal.Id, unlockStatusUnlocked)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return &unlockStatusUnlocked, err
	}

	if len(txReceipt.Logs) == 0 {
		return &unlockStatusUnlocked, nil
	}

	err = client.SaveEventUnlockPaymentFromTxHash(txReceipt, recipient, offlineDeal)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return &unlockStatusUnlocked, err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "unlock successfully"))
	return &unlockStatusUnlocked, nil
}

func refund(offlineDeal *models.OfflineDeal, swanPaymentTransactor *goBind.SwanPaymentTransactor, tansactOpts *bind.TransactOpts) error {
	offlineDealsNotUnlocked, err := models.GetOfflineDealsNotUnlockedByDealFileId(offlineDeal.DealFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return err
	}

	if len(offlineDealsNotUnlocked) > 0 {
		msg := fmt.Sprintf("%d deals not unlocked, cannot refund for the deal file", len(offlineDealsNotUnlocked))
		logs.GetLogger().Info(getLog(offlineDeal, msg))
		return nil
	}

	var srcFilePayloadCids []string
	srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.DealFileId)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return err
	}

	for _, srcFile := range srcFiles {
		srcFilePayloadCids = append(srcFilePayloadCids, srcFile.PayloadCid)
	}

	refundStatus := constants.PROCESS_STATUS_UNLOCK_REFUNDED
	_, err = swanPaymentTransactor.Refund(tansactOpts, srcFilePayloadCids)
	if err != nil {
		refundStatus = constants.PROCESS_STATUS_UNLOCK_REFUNDFAILED
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
	}

	err = models.UpdateDealFileStatus(offlineDeal.DealFileId, refundStatus)
	if err != nil {
		logs.GetLogger().Error(getLog(offlineDeal, err.Error()))
		return err
	}

	logs.GetLogger().Info(getLog(offlineDeal, "refund with status:"+refundStatus))

	return nil
}
