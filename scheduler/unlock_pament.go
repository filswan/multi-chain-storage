package scheduler

import (
	"context"
	"math/big"
	"os"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/robfig/cron"
)

func CreateUnlockScheduler() {
	Mutex := &sync.Mutex{}
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UnlockPaymentRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ create task  scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		Mutex.Lock()
		err := UnlockPayment()
		Mutex.Unlock()
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
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
	for _, offlineDeal := range offlineDeals {
		err := doUnlockPaymentOnContract(offlineDeal.DealId)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}
	return err
}

func doUnlockPaymentOnContract(dealId int64) error {
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
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = uint64(polygon.GetConfig().PolygonMainnetNode.GasLimit)
	callOpts.Context = context.Background()

	swanPaymentContractAddress := common.HexToAddress(polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress) //payment gateway on polygon
	swanPaymentContractInstance, err := goBind.NewSwanPaymentTransactor(swanPaymentContractAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	dealIdStr := strconv.FormatInt(dealId, 10)

	tx, err := swanPaymentContractInstance.UnlockCarPayment(callOpts, dealIdStr, swanPaymentContractAddress)
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
					err = saveUnlockEventLogToDB(eventLogs, unlockTxStatus, dealId)
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

	offlineDeals, err := models.GetOfflineDealByDealId(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(offlineDeals) > 0 {
		offlineDeal := offlineDeals[0]
		offlineDealsNotUnlocked, err := models.GetOfflineDealsNotUnlockedByDealFileId(offlineDeal.DealFileId)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		if len(offlineDealsNotUnlocked) == 0 {
			var srcFilePayloadCids []string
			srcFiles, err := models.GetSourceFilesByDealFileId(offlineDeal.DealFileId)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}

			for _, srcFile := range srcFiles {
				srcFilePayloadCids = append(srcFilePayloadCids, srcFile.PayloadCid)
			}

			refundStatusAfterUnlock := constants.LOCK_PAYMENT_STATUS_UNLOCK_REFUNDED
			tx, err := swanPaymentContractInstance.Refund(callOpts, srcFilePayloadCids)
			if err != nil {
				refundStatusAfterUnlock = constants.LOCK_PAYMENT_STATUS_UNLOCK_REFUNDFAILED
				logs.GetLogger().Error(err)
			}

			logs.GetLogger().Info(tx)

			currrentTime := utils.GetCurrentUtcMilliSecond()
			err = models.UpdateDealFile(models.DealFile{ID: offlineDeal.DealFileId},
				map[string]interface{}{"lock_payment_status": refundStatusAfterUnlock, "update_at": currrentTime})
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}

	return nil
}

func saveUnlockEventLogToDB(logsInChain []*types.Log, unlockStatus string, dealId int64) error {
	paymentAbiString := goBind.SwanPaymentABI

	contractUnlockFunctionSignature := polygon.GetConfig().PolygonMainnetNode.ContractUnlockFunctionSignature
	contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, vLog := range logsInChain {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractUnlockFunctionSignature {
			eventList, err := models.FindEventUnlockPayments(&models.EventUnlockPayment{TxHash: vLog.TxHash.Hex(), BlockNo: strconv.FormatUint(vLog.BlockNumber, 10)}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if len(eventList) <= 0 {
				event := new(models.EventUnlockPayment)
				dataList, err := contractAbi.Unpack("UnlockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				event.DealId = dealId
				event.TxHash = vLog.TxHash.Hex()
				networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.NetworkId = networkId
				}
				coinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_USDC_ON_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.CoinId = coinId
				}
				event.TokenAddress = dataList[1].(common.Address).Hex()
				event.UnlockToAdminAmount = dataList[2].(*big.Int).String()
				event.UnlockToUserAmount = dataList[3].(*big.Int).String()
				event.UnlockToAdminAddress = dataList[4].(common.Address).Hex()
				event.UnlockToUserAddress = dataList[5].(common.Address).Hex()
				event.UnlockTime = strconv.FormatInt(utils.GetCurrentUtcMilliSecond(), 10)
				event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
				event.CreateAt = utils.GetCurrentUtcMilliSecond()
				event.UnlockStatus = unlockStatus
				err = database.SaveOneWithTransaction(event)
				if err != nil {
					logs.GetLogger().Error(err)
				}

			}
		}
	}
	return nil
}
