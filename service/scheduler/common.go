package scheduler

import (
	"fmt"
	"multi-chain-storage/config"
	"os"
	"path/filepath"
	"sync"

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

func GetSrcDir() string {
	return srcDir
}

func InitScheduler() {
	createDir()
	//createScheduleJob()
	CreateScheduler4CreateTask()
	CreateScheduler4Refund()
	CreateScheduler4ScanDeal()
	CreateScheduler4SendDeal()
	CreateScheduler4UnlockPayment()
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
