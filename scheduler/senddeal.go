package scheduler

import (
	"encoding/json"
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"payment-bridge/common/httpClient"
	"payment-bridge/config"
	"payment-bridge/logs"
	"payment-bridge/models"
	"time"
)

func SendDealScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.UnlockPaymentRule, func() {
		log.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^dao signature unlock payment schedule is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := doSendDealScheduler()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	})
	if err != nil {
		logs.GetLogger().Error(err)
	}
	c.Start()
}
func doSendDealScheduler() error {
	dealList, err := GetTaskListShouldBeSigService()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range dealList.Data.Deals {
		if v.Status == "Assigned" {
			err = sendDeal(v.UUID)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}
	return nil
}

func GetTaskListShouldBeSigService() (*models.OfflineDealResult, error) {
	url := config.GetConfig().SwanApi.ApiUrl + config.GetConfig().SwanApi.GetTaskApiUrlSuffix
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	var results *models.OfflineDealResult
	err = json.Unmarshal(response, &results)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func sendDeal(taskUuid string) error {

	task_uuid := taskUuid

	startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours
	startEpoch := libutils.GetCurrentEpoch() + (startEpochIntervalHours+1)*libconstants.EPOCH_PER_HOUR

	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	temDirDeal := config.GetConfig().Temp.DirDeal
	temDirDeal = filepath.Join(homedir, temDirDeal[2:])
	err = libutils.CreateDir(temDirDeal)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	timeStr := time.Now().Format("20060102_150405")
	temDirDeal = filepath.Join(temDirDeal, timeStr)
	carDir := filepath.Join(temDirDeal, "car")
	confDeal := &clientmodel.ConfDeal{
		SwanApiUrl:              config.GetConfig().SwanApi.ApiUrl,
		SwanApiKey:              config.GetConfig().SwanApi.ApiKey,
		SwanAccessToken:         config.GetConfig().SwanApi.ApiKey,
		SenderWallet:            config.GetConfig().FileCoinWallet,
		VerifiedDeal:            config.GetConfig().SwanTask.VerifiedDeal,
		FastRetrieval:           config.GetConfig().SwanTask.FastRetrieval,
		SkipConfirmation:        true,
		StartEpochIntervalHours: startEpochIntervalHours,
		StartEpoch:              startEpoch,
		OutputDir:               carDir,
	}

	dealSentNum, csvFilePath, carFiles, err := subcommand.SendAutoBidDealsByTaskUuid(confDeal, task_uuid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("------------------------------send deal success---------------------------------")
	logs.GetLogger().Info("dealSentNum = ", dealSentNum)
	logs.GetLogger().Info("csvFilePath = ", csvFilePath)
	logs.GetLogger().Info("carFiles = ", carFiles)
	return nil
}
