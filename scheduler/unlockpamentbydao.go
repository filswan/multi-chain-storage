package scheduler

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robfig/cron"
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
	"strconv"
	"strings"
	"time"
)

func DAOUnlockPaymentSchedule() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UnlockPaymentRule, func() {
		log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^dao signature unlock payment schedule is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := UnlockPaymentByDao()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	})
	if err != nil {
		logs.GetLogger().Error(err)
	}
	c.Start()
}
func UnlockPaymentByDao() error {
	daoSigResult, err := models.GetThresholdForDao()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range daoSigResult {
		daoEventLogList, err := models.FindDaoEventLog(&models.DaoEventLog{PayloadCid: v.PayloadCid, DealCid: v.DealCid, SignatureUnlockStatus: constants.SIGNATURE_DEFAULT_VALUE}, "id desc", "10", "0")
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		if len(daoEventLogList) > 0 {
			parm := goBind.IPaymentMinimalunlockPaymentParam{}
			parm.Id = v.PayloadCid
			parm.DealId = v.DealCid
			n := new(big.Int)
			n, _ = n.SetString(daoEventLogList[0].Cost, 10)
			parm.Amount = n
			parm.Recipient = common.HexToAddress(daoEventLogList[0].Recipient)
			parm.OrderId = v.OrderId
			err = doUnlockPaymentOnContract(daoEventLogList[0], parm)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			//update signature unlock action success
			err = models.UpdateDaoEventLog(&models.DaoEventLog{PayloadCid: v.PayloadCid, DealCid: v.DealCid}, map[string]interface{}{"SignatureUnlockStatus": constants.SIGNATURE_SUCCESS_VALUE})
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}
	}
	return err
}

func doUnlockPaymentOnContract(daoEvent *models.DaoEventLog, unlockParams goBind.IPaymentMinimalunlockPaymentParam) error {
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

	swanPaymentContractAddress := common.HexToAddress(config.GetConfig().SwanPaymentAddressOnPolygon) //payment gateway on polygon
	swanPaymentContractInstance, err := goBind.NewSwanPayment(swanPaymentContractAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	tx, err := swanPaymentContractInstance.UnlockTokenPayment(callOpts, unlockParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	unlockTxStatus := ""
	if err != nil {
		logs.GetLogger().Error(err)
		unlockTxStatus = constants.TRANSACTION_STATUS_FAIL
	} else {
		unlockTxStatus = constants.TRANSACTION_STATUS_SUCCESS
	}
	err = updateUnlockPaymentStatus(daoEvent.PayloadCid, daoEvent.DealCid, unlockTxStatus, tx.Hash().Hex())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("unlock tx hash=", tx.Hash().Hex(), " for payloadCid=", daoEvent.PayloadCid)
	return nil
}

func updateUnlockPaymentStatus(payloadCid, dealCid, unLockTxStatus, unlockTxHash string) error {
	updateTime := strconv.FormatInt(utils.GetEpochInMillis(), 10)
	err := models.UpdateEventPolygon(&models.EventPolygon{PayloadCid: payloadCid}, map[string]interface{}{"unlock_time": updateTime, "unlock_tx_status": unLockTxStatus, "unlock_tx_hash": unlockTxHash})
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
