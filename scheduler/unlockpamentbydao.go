package scheduler

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robfig/cron"
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
	threshHold, err := getThreshHold()
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
		txRecept, err := utils.CheckTx(client, tx)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if txRecept.Status == uint64(1) {
				unlockTxStatus = constants.TRANSACTION_STATUS_SUCCESS
				logs.GetLogger().Println("unlock success! txHash=" + tx.Hash().Hex())
			} else {
				unlockTxStatus = constants.TRANSACTION_STATUS_FAIL
				logs.GetLogger().Println("unlock failed! txHash=" + tx.Hash().Hex())
			}
		}
		if len(txRecept.Logs) > 0 {
			eventLogs := txRecept.Logs
			err = saveUnlockEventLogToDB(eventLogs, daoEvent.Recipient)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
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
func saveUnlockEventLogToDB(logsInChain []*types.Log, recipient string) error {
	//paymentAbiString, err := utils.ReadContractAbiJsonFile(goBind.SwanPaymentMetaData.ABI)
	paymentAbiString := goBind.SwanPaymentMetaData.ABI

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
			var event *models.EventUnlockPayment
			if len(eventList) <= 0 {
				dataList, err := contractAbi.Unpack("UnlockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				fmt.Println(dataList)
				event = new(models.EventUnlockPayment)
				event.TxHash = vLog.TxHash.Hex()
				event.Network = constants.NETWORK_TYPE_POLYGON
				block, err := polygon.WebConn.ConnWeb.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.UnlockTime = strconv.FormatUint(block.Time(), 10)
				}
				chainId, err := polygon.WebConn.ConnWeb.ChainID(context.Background())
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				addrInfo, err := utils.GetFromAndToAddressByTxHash(polygon.WebConn.ConnWeb, chainId, vLog.TxHash)
				if err != nil {
					logs.GetLogger().Error(err)
				} else {
					event.UnlockFromAddress = addrInfo.AddrFrom
				}
				event.UnlockFromAddress = polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress
				event.BlockNo = strconv.FormatUint(vLog.BlockNumber, 10)
				event.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
			} else {
				event = eventList[0]
			}

			quantity := new(big.Int)
			quantity.SetBytes(vLog.Data)
			address := common.HexToAddress(vLog.Topics[1].Hex()).String()
			if strings.ToLower(address) == strings.ToLower(polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress) {
				if strings.ToLower(common.HexToAddress(vLog.Topics[2].Hex()).String()) == strings.ToLower(recipient) {
					event.UnlockToAdminAddress = common.HexToAddress(vLog.Topics[2].Hex()).String()
					event.UnlockToAdminAmount = quantity.String()
				} else {
					event.UnlockToUserAddress = common.HexToAddress(vLog.Topics[2].Hex()).String()
					event.UnlockToUserAmount = quantity.String()
				}
			}

			err = database.SaveOneWithTransaction(event)
			if err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}
	return nil
}

func getThreshHold() (uint8, error) {
	daoAddress := common.HexToAddress(config.GetConfig().SwanDaoOracleAddress)
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
