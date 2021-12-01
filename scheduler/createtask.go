package scheduler

import (
	"fmt"
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"github.com/shopspring/decimal"
	"path/filepath"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"strings"
	"time"
)

func CreateTaskScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.CreateTaskRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ create task  scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := DoCreateTask()
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

func DoCreateTask() error {
	whereCondition := "lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_SUCCESS + "') and task_uuid = '' "
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "50", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range dealList {
		//check if user have lock payment
		lockPaymentList, err := CheckIfHaveLockPayment(v.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		if len(lockPaymentList) > 0 {
			lockedFee, err := strconv.ParseInt(lockPaymentList[0].LockedFee, 10, 64)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if strings.Trim(v.TaskUuid, " ") == "" {
				confUpload := &clientmodel.ConfUpload{
					StorageServerType:           libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
					IpfsServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
					IpfsServerUploadUrl:         config.GetConfig().IpfsServer.UploadUrl,
					OutputDir:                   filepath.Dir(v.CarFilePath),
					InputDir:                    filepath.Dir(v.CarFilePath),
				}
				_, err = subcommand.UploadCarFiles(confUpload)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
				logs.GetLogger().Info("car files uploaded")

				taskDataset := config.GetConfig().SwanTask.CuratedDataset
				taskDescription := config.GetConfig().SwanTask.Description
				startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours
				startEpoch := libutils.GetCurrentEpoch() + (startEpochIntervalHours+1)*libconstants.EPOCH_PER_HOUR
				fmt.Println(filepath.Dir(v.CarFilePath))
				confTask := &clientmodel.ConfTask{
					SwanApiUrl:      config.GetConfig().SwanApi.ApiUrl,
					SwanToken:       "",
					PublicDeal:      true,
					SwanApiKey:      config.GetConfig().SwanApi.ApiKey,
					SwanAccessToken: config.GetConfig().SwanApi.AccessToken,
					BidMode:         libconstants.TASK_BID_MODE_AUTO,
					VerifiedDeal:    config.GetConfig().SwanTask.VerifiedDeal,
					OfflineMode:     false,
					FastRetrieval:   config.GetConfig().SwanTask.FastRetrieval,
					//MaxPrice:                 config.GetConfig().SwanTask.MaxPrice,
					MaxPrice:                   decimal.NewFromInt(lockedFee),
					StorageServerType:          libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
					WebServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
					ExpireDays:                 config.GetConfig().SwanTask.ExpireDays,
					OutputDir:                  filepath.Dir(v.CarFilePath),
					InputDir:                   filepath.Dir(v.CarFilePath),
					MinerFid:                   "",
					Dataset:                    taskDataset,
					Description:                taskDescription,
					StartEpochIntervalHours:    startEpochIntervalHours,
					StartEpoch:                 startEpoch,
					SourceId:                   constants.SOURCE_ID_OF_PAYMENT,
					Duration:                   v.Duration,
				}
				_, fileInfoList, _, err := subcommand.CreateTask(confTask, nil)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
				if len(fileInfoList) > 0 {
					logs.GetLogger().Info("task created, uuid=", fileInfoList[0].Uuid)
					err = updateTaskInfoToDB(fileInfoList, v)
					if err != nil {
						logs.GetLogger().Error(err)
						return err
					}
				}
			}
		}
	}
	return nil
}

func updateTaskInfoToDB(taskinfoList []*libmodel.FileDesc, dealFile *models.DealFile) error {
	dealFile.TaskUuid = taskinfoList[0].Uuid
	dealFile.DealCid = taskinfoList[0].DealCid
	dealFile.Cost = taskinfoList[0].Cost
	err := database.SaveOne(dealFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
