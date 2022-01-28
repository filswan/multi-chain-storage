package scheduler

import (
	"fmt"
	"os"
	"path/filepath"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
)

type Schedule struct {
	Cron  *cron.Cron
	Name  string
	Rule  string
	Func  func() error
	Mutex *sync.Mutex
}

var carDir string
var srcDir string

var ethClient *ethclient.Client
var privateKeyOnPolygon string

func GetSrcDir() string {
	return srcDir
}

func InitScheduler() {
	createDir()
	//createScheduleJob()

	privateKeyOnPolygon = os.Getenv("privateKeyOnPolygon")
	if len(privateKeyOnPolygon) <= 0 {
		err := fmt.Errorf("env variable privateKeyOnPolygon is not defined")
		logs.GetLogger().Fatal(err)
	}

	if strings.HasPrefix(strings.ToLower(privateKeyOnPolygon), "0x") {
		privateKeyOnPolygon = privateKeyOnPolygon[2:]
	}
}

func createScheduleJob() {
	confScheduleRule := config.GetConfig().ScheduleRule
	scheduleJobs := []Schedule{
		{Cron: cron.New(), Name: "create task", Rule: confScheduleRule.CreateTaskRule, Func: CreateTask, Mutex: &sync.Mutex{}},
		{Cron: cron.New(), Name: "send deal", Rule: confScheduleRule.SendDealRule, Func: SendDeal},
		{Cron: cron.New(), Name: "scan deal", Rule: confScheduleRule.ScanDealStatusRule, Func: ScanDeal},
		{Cron: cron.New(), Name: "unlock payment", Rule: confScheduleRule.UnlockPaymentRule, Func: UnlockPayment},
	}

	for _, scheduleJob := range scheduleJobs {
		err := scheduleJob.Cron.AddFunc(scheduleJob.Rule, func() {
			logs.GetLogger().Info(scheduleJob.Name + " start")

			if scheduleJob.Mutex != nil {
				scheduleJob.Mutex.Lock()
			}

			scheduleJob.Func()

			if scheduleJob.Mutex != nil {
				scheduleJob.Mutex.Unlock()
			}

			logs.GetLogger().Info(scheduleJob.Name + " end")
		})

		if err != nil {
			logs.GetLogger().Fatal(err)
		}

		scheduleJob.Cron.Start()
	}
}

func createDir() {
	dealDir := config.GetConfig().SwanTask.DirDeal
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	if len(dealDir) < 2 {
		err := fmt.Errorf("deal directory config error, please contact administrator")
		logs.GetLogger().Fatal(err)
	}

	dealDir = filepath.Join(homedir, dealDir[2:])

	err = libutils.CreateDir(dealDir)
	if err != nil {
		logs.GetLogger().Error(err)
		logs.GetLogger().Fatal("creating dir:", dealDir, " failed")
	}

	srcDir = filepath.Join(dealDir, "src")
	err = libutils.CreateDir(srcDir)
	if err != nil {
		logs.GetLogger().Error(err)
		logs.GetLogger().Fatal("creating dir:", srcDir, " failed")
	}

	carDir = filepath.Join(dealDir, "car")
	err = libutils.CreateDir(srcDir)
	if err != nil {
		logs.GetLogger().Error(err)
		logs.GetLogger().Fatal("creating dir:", carDir, " failed")
	}
}

func DialEthClient() error {
	if ethClient != nil {
		return nil
	}

	polygonRpcUrl := config.GetConfig().PolygonRpcUrl
	client, err := ethclient.Dial(polygonRpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	ethClient = client

	return nil
}

func GetPaymentSession() (*goBind.SwanPaymentSession, error) {
	err := DialEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	swanPaymentSession := &goBind.SwanPaymentSession{}

	paymentGatewayAddress := common.HexToAddress(polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress)

	paymentGatewayInstance, err := goBind.NewSwanPayment(paymentGatewayAddress, ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	swanPaymentSession.Contract = paymentGatewayInstance

	return swanPaymentSession, nil
}

func GetPaymentInfo(srcFilePayloadCid string) (*models.EventLockPayment, error) {
	swanPaymentSession, err := GetPaymentSession()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	paymentInfo, err := swanPaymentSession.GetLockedPaymentInfo(srcFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var event *models.EventLockPayment
	if paymentInfo.IsExisted {
		event = new(models.EventLockPayment)
		event.AddressFrom = paymentInfo.Owner.String()
		event.AddressTo = paymentInfo.Recipient.String()
		event.LockedFee = paymentInfo.LockedFee.String()
		event.PayloadCid = srcFilePayloadCid
		usdcCoinId, err := models.FindCoinIdByUUID(constants.COIN_TYPE_USDC_ON_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.CoinId = usdcCoinId
		}

		networkId, err := models.FindNetworkIdByUUID(constants.NETWORK_TYPE_POLYGON_UUID)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			event.NetworkId = networkId
		}

		err = database.SaveOne(event)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		dealFiles, err := models.GetDealFileBySourceFilePayloadCid(srcFilePayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		if len(dealFiles) > 0 {
			dealFileId := dealFiles[0].ID

			err := models.UpdateDealFileLockPaymentStatus(dealFileId, constants.LOCK_PAYMENT_STATUS_PROCESSING)
			if err != nil {
				logs.GetLogger().Error(err)
				return nil, err
			}
		}
	}

	return event, nil
}
