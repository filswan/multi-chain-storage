package scheduler

import (
	"crypto/ecdsa"
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/on-chain/client"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type Schedule struct {
	Name  string
	Rule  string
	Func  func() error
	Mutex *sync.Mutex
}

var carDir string
var srcDir string

var adminWalletPrivateKey *ecdsa.PrivateKey = nil
var adminWalletPublicKey *common.Address = nil

func GetSrcDir() string {
	return srcDir
}

func InitScheduler() {
	createDir()
	setAdminWallet()

	go runJob(CreateTask, config.GetConfig().ScheduleRule.CreateTaskIntervalSecond)
	go runJob(SendDeal, config.GetConfig().ScheduleRule.SendDealIntervalSecond)
	go runJob(ScanDeal, config.GetConfig().ScheduleRule.ScanDealStatusIntervalSecond)
	go runJob(UnlockPayment, config.GetConfig().ScheduleRule.UnlockIntervalSecond)
	go runJob(Refund, config.GetConfig().ScheduleRule.RefundIntervalSecond)
	go runJob(CheckPayment, config.GetConfig().ScheduleRule.CheckPaymentIntervalSecond)
}

func runJob(func2Run func() error, intervalSecond time.Duration) {
	for {
		funcName := runtime.FuncForPC(reflect.ValueOf(func2Run).Pointer()).Name()
		logs.GetLogger().Info(funcName, " start")
		err := func2Run()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		logs.GetLogger().Info(funcName, " end")

		time.Sleep(intervalSecond * time.Second)
	}
}

func createDir() {
	dealDir := config.GetConfig().SwanTask.DirDeal
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	if len(dealDir) < 2 {
		err := fmt.Errorf("deal directory config error")
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

func setAdminWallet() {
	privateKey, publicKey, err := client.GetPrivateKeyPublicKey(constants.PRIVATE_KEY_ON_POLYGON)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	adminWalletPrivateKey = privateKey
	adminWalletPublicKey = publicKey
}
