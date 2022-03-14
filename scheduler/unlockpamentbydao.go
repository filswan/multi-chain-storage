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
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robfig/cron"
)

func DAOUnlockPaymentSchedule() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UnlockPaymentRule, func() {
		log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^unlock payment scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
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
	threshHold, err := GetThreshHold()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	daoSigResult, err := models.GetDaoSignatureEventsSholdBeUnlock(threshHold)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range daoSigResult {
		daoEventLogList, err := models.FindDaoEventLog(&models.EventDaoSignature{PayloadCid: v.PayloadCid, DealId: v.DealId, SignatureUnlockStatus: constants.SIGNATURE_DEFAULT_VALUE}, "id desc", "10", "0")
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		if len(daoEventLogList) > 0 {
			parm := goBind.IPaymentMinimalunlockPaymentParam{}
			parm.Id = v.PayloadCid
			parm.DealId = strconv.FormatInt(v.DealId, 10)
			parm.Amount = big.NewInt(0)
			parm.Recipient = common.HexToAddress(daoEventLogList[0].Recipient)
			parm.OrderId = ""
			/*parm := goBind.IPaymentMinimalunlockPaymentParam{}
			parm.Id = "12321434534lkdlfg"
			parm.DealId ="1234"
			parm.Amount = big.NewInt(0)
			parm.Recipient = common.HexToAddress(daoEventLogList[0].Recipient)
			parm.OrderId = ""*/
			err = doUnlockPaymentOnContract(daoEventLogList[0], parm)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}
	}
	return err
}

func doUnlockPaymentOnContract(daoEvent *models.EventDaoSignature, unlockParams goBind.IPaymentMinimalunlockPaymentParam) error {
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
					err = saveUnlockEventLogToDB(eventLogs, daoEvent.Recipient, unlockTxStatus)
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
	//update unlock status in event_dao_signature
	err = models.UpdateDaoEventLog(&models.EventDaoSignature{PayloadCid: daoEvent.PayloadCid, DealId: daoEvent.DealId},
		map[string]interface{}{"tx_hash_unlock": tx.Hash().Hex(), "signature_unlock_status": unlockTxStatus})
	if err != nil {
		logs.GetLogger().Error(err)
	}
	logs.GetLogger().Info("unlock tx hash=", tx.Hash().Hex(), " for payloadCid=", daoEvent.PayloadCid, " and deal_id=", daoEvent.DealId)
	return nil
}

func saveUnlockEventLogToDB(logsInChain []*types.Log, recipient string, unlockStatus string) error {
	paymentAbiString := goBind.SwanPaymentABI

	//SwanPayment contract function signature
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
				event.PayloadCid = dataList[0].(string)
				event.TxHash = vLog.TxHash.Hex()
				neoworkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.NetworkId = neoworkId
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
				event.UnlockTime = strconv.FormatInt(utils.GetEpochInMillis(), 10)
				event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
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

func GetThreshHold() (uint8, error) {
	daoAddress := common.HexToAddress(polygon.GetConfig().PolygonMainnetNode.DaoSwanOracleAddress)
	client := polygon.WebConn.ConnWeb

	pk := os.Getenv("privateKeyOnPolygon")
	if strings.HasPrefix(strings.ToLower(pk), "0x") {
		pk = pk[2:]
	}

	callOpts := new(bind.CallOpts)
	callOpts.From = daoAddress
	callOpts.Context = context.Background()

	daoOracleContractInstance, err := goBind.NewFilswanOracle(daoAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}

	threshHold, err := daoOracleContractInstance.GetThreshold(callOpts)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	logs.GetLogger().Info("dao threshHold is : ", threshHold)
	return threshHold, nil
}
