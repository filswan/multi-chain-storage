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
	"math/big"
	"path/filepath"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/routers/billing"
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
	fileCoinPriceInUsdc, err := billing.GetWfilPriceFromSushiPrice(polygon.WebConn.ConnWeb, "1")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	whereCondition := "lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_PROCESSING + "') and task_uuid = '' "
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
			maxPrice, err := GetMaxPriceForCreateTask(fileCoinPriceInUsdc, lockedFee, v.Duration, v.CarFileSize)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			logs.GetLogger().Println("payload cid ", v.PayloadCid, " max price is ", maxPrice)
			if maxPrice.Cmp(config.GetConfig().SwanTask.MaxPrice) < 0 {
				*maxPrice = config.GetConfig().SwanTask.MaxPrice
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
					//MaxPrice:        config.GetConfig().SwanTask.MaxPrice,
					MaxPrice:                   *maxPrice,
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

func GetMaxPriceForCreateTask(rate *big.Int, lockedFee int64, duration int, carFileSize int64) (*decimal.Decimal, error) {
	_, sectorSize := libutils.CalculatePieceSize(carFileSize)
	lockedFeeDecimal := decimal.NewFromInt(lockedFee)
	lockedFeeInFileCoin := lockedFeeDecimal.Div(decimal.NewFromFloat(constants.LOTUS_PRICE_MULTIPLE_1E18)).Div(decimal.NewFromInt(rate.Int64()))
	maxPrice := lockedFeeInFileCoin.Div(decimal.NewFromFloat(sectorSize).Div(decimal.NewFromInt(10204 * 1024 * 1024))).Div(decimal.NewFromInt(int64(duration)))
	return &maxPrice, nil
}

/*func GetMinerPerEpoachPriceInOtherCoin(minerFid string, rate *big.Int, verifiedType string, carFileSize int64) (int64, error) {
	lotusClient, err := lotus.LotusGetClient(config.GetConfig().Lotus.ApiUrl, config.GetConfig().Lotus.AccessToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}

	minerConfig, err := lotusClient.LotusClientQueryAsk(minerFid)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	//minerPrice, minerVerifiedPrice, _, _ := lotusClient.LotusGetMinerConfig(minerFid)
	var unitPriceMiner decimal.Decimal
	if strings.Trim(verifiedType, " ") == constants.LOTUS_TASK_TYPE_VERIFIED {
		unitPriceMiner = minerConfig.VerifiedPrice
	} else {
		unitPriceMiner = minerConfig.Price
	}
	_, sectorSize := libutils.CalculatePieceSize(carFileSize)

	unitPriceMinerWithFileSize := libutils.CalculateRealCost(sectorSize, unitPriceMiner)
	finalPrice := decimal.NewFromInt(rate.Int64()).Mul(unitPriceMinerWithFileSize)
	return finalPrice.IntPart(), nil

}*/

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
