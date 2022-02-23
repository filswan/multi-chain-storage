package scheduler

import (
	"fmt"
	"multi-chain-storage/config"
	"os"
	"path/filepath"
	"sync"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
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
	createScheduleJob()
}

func createScheduleJob() {
	confScheduleRule := config.GetConfig().ScheduleRule
	scheduleJobs := []Schedule{
		{Name: "create task", Rule: confScheduleRule.CreateTaskRule, Func: CreateTask, Mutex: &sync.Mutex{}},
		{Name: "send deal", Rule: confScheduleRule.SendDealRule, Func: SendDeal, Mutex: &sync.Mutex{}},
		{Name: "scan deal", Rule: confScheduleRule.ScanDealStatusRule, Func: ScanDeal, Mutex: &sync.Mutex{}},
		{Name: "unlock payment", Rule: confScheduleRule.UnlockPaymentRule, Func: UnlockPayment, Mutex: &sync.Mutex{}},
		{Name: "refund", Rule: confScheduleRule.RefundRule, Func: Refund, Mutex: &sync.Mutex{}},
	}

	for _, scheduleJob := range scheduleJobs {
		createScheduler(scheduleJob.Name, scheduleJob.Rule, scheduleJob.Func, scheduleJob.Mutex)
	}
}

func createScheduler(name, rule string, func2Run func() error, mutex *sync.Mutex) {
	c := cron.New()
	err := c.AddFunc(rule, func() {
		logs.GetLogger().Info(name, " start")

		mutex.Lock()
		logs.GetLogger().Info(name, " running")
		err := func2Run()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		mutex.Unlock()
		logs.GetLogger().Info(name, " end")
	})

	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	c.Start()
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
