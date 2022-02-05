package scheduler

import (
	"fmt"
	"os"
	"path/filepath"
	"payment-bridge/config"
	"sync"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
)

type Schedule struct {
	Name      string
	Rule      string
	Func      func() error
	Mutex     *sync.Mutex
	IsRunning bool
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

func CreateScheduler(name, rule string, func2Run func() error, mutex *sync.Mutex, isRunning *bool) {
	c := cron.New()
	err := c.AddFunc(rule, func() {
		logs.GetLogger().Info(name, " start")
		if *isRunning {
			logs.GetLogger().Info(name, " already running, exit")
			return
		}

		mutex.Lock()
		logs.GetLogger().Info(name, " running")
		*isRunning = true
		err := func2Run()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		*isRunning = false
		mutex.Unlock()
		logs.GetLogger().Info(name, " end")
	})

	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	c.Start()
}

func createScheduleJob() {
	confScheduleRule := config.GetConfig().ScheduleRule
	scheduleJobs := []Schedule{
		{Name: "create task", Rule: confScheduleRule.CreateTaskRule, Func: CreateTask, Mutex: &sync.Mutex{}, IsRunning: false},
		{Name: "send deal", Rule: confScheduleRule.SendDealRule, Func: SendDeal, Mutex: &sync.Mutex{}, IsRunning: false},
		{Name: "scan deal", Rule: confScheduleRule.ScanDealStatusRule, Func: ScanDeal, Mutex: &sync.Mutex{}, IsRunning: false},
		{Name: "unlock payment", Rule: confScheduleRule.UnlockPaymentRule, Func: UnlockPayment, Mutex: &sync.Mutex{}, IsRunning: false},
	}

	for _, scheduleJob := range scheduleJobs {
		CreateScheduler(scheduleJob.Name, scheduleJob.Rule, scheduleJob.Func, scheduleJob.Mutex, &scheduleJob.IsRunning)
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
