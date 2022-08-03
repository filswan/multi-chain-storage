package scheduler

import (
	"fmt"
	"multi-chain-storage/config"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"time"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

var carDir string
var srcDir string

func GetSrcDir() string {
	return srcDir
}

func InitScheduler() {
	createDir()

	go runJob(CreateTask, config.GetConfig().ScheduleRule.CreateTaskIntervalSecond)
	go runJob(SendDeal, config.GetConfig().ScheduleRule.SendDealIntervalSecond)
	go runJob(ScanDeal, config.GetConfig().ScheduleRule.ScanDealStatusIntervalSecond)
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
