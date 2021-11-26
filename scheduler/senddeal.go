package scheduler

import (
	"encoding/json"
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"net/http"
	"os"
	"path/filepath"
	"payment-bridge/common/constants"
	"payment-bridge/common/httpClient"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"time"
)

func SendDealScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.SendDealRule, func() {
		logs.GetLogger().Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ send deal scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := DoSendDealScheduler()
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
	})
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	c.Start()
}
func DoSendDealScheduler() error {
	dealList, err := GetTaskListShouldBeSigServiceFromLocal()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range dealList {
		if v.DealStatus == constants.TASK_STATUS_ASSIGNED {
			logs.GetLogger().Println("################################## start to send deal ##################################")
			logs.GetLogger().Println(" task uuid : ", v.TaskUuid)
			err = sendDeal(v.TaskUuid, v)
			logs.GetLogger().Println("################################## end to send deal ##################################")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			v.DealStatus = "DealSent"
			err = database.SaveOne(v)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
		}
	}
	return nil
}

func GetTaskListShouldBeSigServiceFromLocal() ([]*models.DealFile, error) {
	whereCondition := "lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_SUCCESS + "' and task_uuid != null and task_uuid != '' "
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "10", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return dealList, nil
}

func GetTaskListShouldBeSigServiceFromSwan() (*models.OfflineDealResult, error) {
	url := config.GetConfig().SwanApi.ApiUrl + config.GetConfig().SwanApi.GetShouldSendTaskUrlSuffix
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

func sendDeal(taskUuid string, file *models.DealFile) error {
	startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours
	startEpoch := libutils.GetCurrentEpoch() + (startEpochIntervalHours+1)*libconstants.EPOCH_PER_HOUR

	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	temDirDeal := config.GetConfig().SwanTask.DirDeal
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
		SwanApiUrl:                   config.GetConfig().SwanApi.ApiUrl,
		SwanApiKey:                   config.GetConfig().SwanApi.ApiKey,
		SwanAccessToken:              config.GetConfig().SwanApi.AccessToken,
		SenderWallet:                 config.GetConfig().FileCoinWallet,
		VerifiedDeal:                 config.GetConfig().SwanTask.VerifiedDeal,
		FastRetrieval:                config.GetConfig().SwanTask.FastRetrieval,
		SkipConfirmation:             true,
		StartEpochIntervalHours:      startEpochIntervalHours,
		StartEpoch:                   startEpoch,
		OutputDir:                    carDir,
		LotusClientApiUrl:            config.GetConfig().Lotus.ApiUrl,
		LotusClientAccessToken:       config.GetConfig().Lotus.AccessToken,
		Duration:                     file.Duration,
		RelativeEpochFromMainNetwork: config.GetConfig().SwanTask.RelativeEpochFromMainNetwork,
	}
	confDeal.DealSourceIds = append(confDeal.DealSourceIds, libconstants.TASK_SOURCE_ID_SWAN_PAYMENT)

	dealSentNum, csvFilePath, carFiles, err := subcommand.SendAutoBidDealsByTaskUuid(confDeal, taskUuid)
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

func checkIfHaveLockPayment(payloadCid string) (bool, error) {
	polygonEventList, err := models.FindEventLockPayment(&models.EventLockPayment{PayloadCid: payloadCid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return false, err
	}
	if len(polygonEventList) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
