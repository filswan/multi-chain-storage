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
	deals2Unlock, err := models.GetDeals2Unlock()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(deals2Unlock) == 0 {
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

	for i, deal2Unlock := range deals2Unlock {
		isUnlockable, err := checkUnlockable(ethClient, deal2Unlock.DealId, filswanOracleSession, paymentRecipientAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(deal2Unlock.DealId, err.Error()))
			continue
		}

		if !isUnlockable {
			logs.GetLogger().Info(getLog(deal2Unlock.DealId, "not unlockable, please wait"))
			continue
		}

		if i > 0 {
			logs.GetLogger().Info(getLog(deal2Unlock.DealId, "sleeping "+unlockIntervalSecond.String()+" before unlock"))
			time.Sleep(unlockIntervalSecond)
		}

		_, err = unlockDeal(deal2Unlock.DealId, ethClient, swanPaymentTransactor, paymentRecipientAddress)
		if err != nil {
			logs.GetLogger().Error(getLog(deal2Unlock.DealId, err.Error()))
			continue
		}
	}
	return nil
}

func checkUnlockable(ethClient *ethclient.Client, dealId int64, filswanOracleSession *goBind.FilswanOracleSession, recipient common.Address) (bool, error) {
	if dealId <= 0 {
		err := fmt.Errorf("valid deal id must be greater than 0")
		logs.GetLogger().Error(getLog(dealId, err.Error()))
		return false, err
	}

	dealIdStr := strconv.FormatInt(dealId, 10)
	filecoinNetwork := config.GetConfig().FilecoinNetwork
	isPaymentAvailable, err := filswanOracleSession.IsCarPaymentAvailable(dealIdStr, filecoinNetwork, recipient)
	if err != nil {
		logs.GetLogger().Error(getLog(dealId, err.Error()))
		return false, err
	}

	if !isPaymentAvailable {
		logs.GetLogger().Info(getLog(dealId, "payment is not available for recipient:"+recipient.String(), ", network:"+filecoinNetwork))
		return false, nil
	}

	filswanOracleTransactions, err := filswanOracleSession.GetSignatureList(dealIdStr, filecoinNetwork)
	if err != nil {
		logs.GetLogger().Error(getLog(dealId, err.Error()))
		return false, err
	}

	if len(filswanOracleTransactions) == 0 {
		logs.GetLogger().Info(getLog(dealId, "no dao sigatures yet"))
		return false, nil
	}

	currentBlockNo, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		logs.GetLogger().Error(getLog(dealId, err.Error()))
		return false, err
	}

	for _, filswanOracleTransaction := range filswanOracleTransactions {
		daoBlockNo := filswanOracleTransaction.BlockNumber.Uint64()
		blockInterval := int64(currentBlockNo - daoBlockNo)

		daoUnlockIntervalBlock := config.GetConfig().Polygon.DaoUnlockIntervalBlock
		if blockInterval < config.GetConfig().Polygon.DaoUnlockIntervalBlock {
			msg := fmt.Sprintf("current block number:%d - dao block number:%d is less than configed dao-unlock block interval:%d", currentBlockNo, daoBlockNo, daoUnlockIntervalBlock)
			logs.GetLogger().Info(dealId, msg)
			return false, nil
		}
	}

	return true, nil
}

func getLog(dealId int64, messages ...string) string {
	text := fmt.Sprintf("deal id:%d", dealId)
	if messages == nil {
		return text
	}

	for _, msg := range messages {
		text = text + "," + msg
	}

	return text
}

func unlockDeal(dealId int64, ethClient *ethclient.Client, swanPaymentTransactor *goBind.SwanPaymentTransactor, recipient common.Address) (*string, error) {
	if dealId <= 0 {
		err := fmt.Errorf("deal id must be greater than 0")
		logs.GetLogger().Error(getLog(dealId, err.Error()))
		return nil, err
	}

	tansactOpts, err := client.GetTransactOpts(ethClient, adminWalletPrivateKey, *adminWalletPublicKey)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	dealIdStr := strconv.FormatInt(dealId, 10)
	filecoinNetwork := config.GetConfig().FilecoinNetwork
	tx, err := swanPaymentTransactor.UnlockCarPayment(tansactOpts, dealIdStr, filecoinNetwork, recipient)
	txHash := ""
	if tx != nil {
		txHash = tx.Hash().Hex()
	}

	logs.GetLogger().Info(getLog(dealId, txHash))
	if err != nil {
		logs.GetLogger().Error(getLog(dealId, "unlock failed, tx hash:"+txHash, err.Error()))
		return nil, err
	}

	txReceipt, err := client.CheckTx(ethClient, tx.Hash())
	if err != nil {
		logs.GetLogger().Error(getLog(dealId, "unlock failed, tx hash:"+txHash, err.Error()))
		return nil, err
	}

	if txReceipt.Status != uint64(1) {
		err := fmt.Errorf("unlock failed, tx hash:%s, status:%d", txHash, txReceipt.Status)
		logs.GetLogger().Error(getLog(dealId, err.Error()))
		return nil, err
	}

	logs.GetLogger().Info(getLog(dealId, "unlock succeeded, tx hash:"+txHash))

	return &txHash, nil
}
