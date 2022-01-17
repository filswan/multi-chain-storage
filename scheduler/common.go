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

const (
	DURATION = 500
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
		{Name: "create task", Rule: confScheduleRule.CreateTaskRule, Func: createTask},
		{Name: "send deal", Rule: confScheduleRule.SendDealRule, Func: sendDeal},
		{Name: "scan deal", Rule: confScheduleRule.ScanDealStatusRule, Func: scanDeal},
		{Name: "unlock payment", Rule: confScheduleRule.UnlockPaymentRule, Func: UnlockPayment},
	}

	for _, scheduleJob := range scheduleJobs {
		c := cron.New()

		err := c.AddFunc(scheduleJob.Rule, func() {
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

		c.Start()
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
