package scheduler

import (
	"context"
	"log"
	"math/big"
	"os"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robfig/cron"
)

func RefundUnlockPaymentSchedule() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UnlockPaymentRule, func() {
		log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^refund unlock payment scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := UnlockPaymentAsRefund()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	})
	if err != nil {
		logs.GetLogger().Error(err)
	}
	c.Start()
}

func UnlockPaymentAsRefund() error {
	lockpaymentDeal, err := models.FindExpiredLockPayment()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range lockpaymentDeal {
		parm := goBind.IPaymentMinimalunlockPaymentParam{}
		parm.Id = v.PayloadCid
		parm.DealId = v.DealId
		parm.Amount = big.NewInt(0)
		parm.Recipient = common.HexToAddress(v.Recipient)
		parm.OrderId = ""
		/*parm := goBind.IPaymentMinimalunlockPaymentParam{}
		parm.Id = "12321434534lkdlfg"
		parm.DealId ="1234"
		parm.Amount = big.NewInt(0)
		parm.Recipient = common.HexToAddress(daoEventLogList[0].Recipient)
		parm.OrderId = ""*/
		err = doUnlockPaymentforRefund(v, parm)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return err
}

func doUnlockPaymentforRefund(paymentEvent *models.EventLockPaymentQuery, unlockParams goBind.IPaymentMinimalunlockPaymentParam) error {
	pk := os.Getenv("privateKeyOnPolygon")
	adminAddress := common.HexToAddress(config.GetConfig().AdminWalletOnPolygon) //pay for gas
	client := polygon.WebConn.ConnWeb
	nonce, err := client.PendingNonceAt(context.Background(), adminAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if strings.HasPrefix(strings.ToLower(pk), "0x") {
		pk = pk[2:]
	}

	privateKey, _ := crypto.HexToECDSA(pk)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	callOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	//callOpts := new(bind.TransactOpts)
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = uint64(polygon.GetConfig().PolygonMainnetNode.GasLimit)
	callOpts.Context = context.Background()

	swanPaymentContractAddress := common.HexToAddress(polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress) //payment gateway on polygon
	swanPaymentContractInstance, err := goBind.NewSwanPayment(swanPaymentContractAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tx, err := swanPaymentContractInstance.UnlockTokenPayment(callOpts, unlockParams)
	unlockTxStatus := ""
	if err != nil {
		logs.GetLogger().Error(err)
		unlockTxStatus = constants.TRANSACTION_STATUS_FAIL
	} else {
		txReceipt, err := utils.CheckTx(client, tx)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if txReceipt.Status == uint64(1) {
				unlockTxStatus = constants.TRANSACTION_STATUS_SUCCESS
				logs.GetLogger().Println("unlock success! txHash=" + tx.Hash().Hex())
				if len(txReceipt.Logs) > 0 {
					eventLogs := txReceipt.Logs
					err = saveUnlockEventLogToDB(eventLogs, paymentEvent.Recipient, unlockTxStatus)
					if err != nil {
						logs.GetLogger().Error(err)
						return err
					}
					err = models.UpdateLockPaymentStatusByPayloadCid(paymentEvent.PayloadCid, constants.LOCK_PAYMENT_STATUS_REFUNDED)
					if err != nil {
						logs.GetLogger().Error(err)
						return err
					}
				}
			} else {
				unlockTxStatus = constants.TRANSACTION_STATUS_FAIL
				logs.GetLogger().Println("unlock failed! txHash=" + tx.Hash().Hex())
			}
		}
	}
	if err != nil {
		logs.GetLogger().Error(err)
	}
	logs.GetLogger().Info("unlock tx hash=", tx.Hash().Hex(), " for payloadCid=", paymentEvent.PayloadCid, " and deal_id=", paymentEvent.DealId)
	return nil
}
