package scheduler

import (

	//clientmodel "github.com/filswan/go-swan-client/model"
	"math/big"
	"path/filepath"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/routers/billing"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-client/command"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"github.com/shopspring/decimal"
)

func CreateTaskScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.CreateTaskRule, func() {
		logs.GetLogger().Info("create task scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
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

func CheckIfHaveLockPayment(payloadCid string) ([]*models.EventLockPayment, error) {
	polygonEventList, err := models.FindEventLockPayment(&models.EventLockPayment{PayloadCid: payloadCid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return polygonEventList, nil
}

func DoCreateTask() error {
	whereCondition := "lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_PROCESSING + "') and task_uuid = '' and is_deleted=false "
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "50", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, deal := range dealList {
		totalLockFee, err := models.GetTotalLockFeeByCarPayloadCid(deal.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		lockedFee := totalLockFee.IntPart()

		fileCoinPriceInUsdc, err := billing.GetWfilPriceFromSushiPrice(polygon.WebConn.ConnWeb, "1")
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		maxPrice, err := GetMaxPriceForCreateTask(fileCoinPriceInUsdc, lockedFee, deal.Duration, deal.CarFileSize)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		logs.GetLogger().Println("payload cid ", deal.PayloadCid, " max price is ", maxPrice)

		if maxPrice.Cmp(config.GetConfig().SwanTask.MaxPrice) < 0 {
			*maxPrice = config.GetConfig().SwanTask.MaxPrice
		}

		cmdUpload := command.CmdUpload{
			StorageServerType:           libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
			IpfsServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
			IpfsServerUploadUrlPrefix:   config.GetConfig().IpfsServer.UploadUrlPrefix,
			OutputDir:                   filepath.Dir(deal.CarFilePath),
			InputDir:                    filepath.Dir(deal.CarFilePath),
		}

		_, err = cmdUpload.UploadCarFiles()
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		logs.GetLogger().Info("car files uploaded")

		taskDataset := config.GetConfig().SwanTask.CuratedDataset
		taskDescription := config.GetConfig().SwanTask.Description
		startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours

		cmdTask := command.CmdTask{
			SwanApiUrl:                 config.GetConfig().SwanApi.ApiUrl,
			SwanToken:                  "",
			SwanApiKey:                 config.GetConfig().SwanApi.ApiKey,
			SwanAccessToken:            config.GetConfig().SwanApi.AccessToken,
			BidMode:                    libconstants.TASK_BID_MODE_AUTO,
			VerifiedDeal:               config.GetConfig().SwanTask.VerifiedDeal,
			OfflineMode:                false,
			FastRetrieval:              config.GetConfig().SwanTask.FastRetrieval,
			MaxPrice:                   *maxPrice,
			StorageServerType:          libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
			WebServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
			ExpireDays:                 config.GetConfig().SwanTask.ExpireDays,
			OutputDir:                  filepath.Dir(deal.CarFilePath),
			InputDir:                   filepath.Dir(deal.CarFilePath),
			Dataset:                    taskDataset,
			Description:                taskDescription,
			StartEpochHours:            startEpochIntervalHours,
			SourceId:                   constants.SOURCE_ID_OF_PAYMENT,
			Duration:                   deal.Duration,
		}

		_, fileDescs, _, err := cmdTask.CreateTask(nil)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		if len(fileDescs) > 0 {
			logs.GetLogger().Info("task created, uuid=", fileDescs[0].Uuid)
			err = updateTaskInfoToDB(fileDescs, deal)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
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

func GetDuplicateTaskInfoByPayloadCid(limit, offset, payloadCid string) ([]*DuplicatedTaskInfo, error) {
	sql := "select s.id as sid,df.id as did,df.miner_fid,df.payload_cid,df.deal_cid,df.task_uuid,df.piece_cid,df.deal_status,df.lock_payment_status as status,df.create_at from  source_file s " +
		" inner join source_file_deal_file_map sfdfm on s.id = sfdfm.source_file_id" +
		" inner join deal_file df on sfdfm.deal_file_id = df.id and df.is_deleted = false "
	if strings.Trim(payloadCid, " ") != "" {
		sql = sql + " and df.payload_cid='" + payloadCid + "'"
	}

	var results []*DuplicatedTaskInfo
	err := database.GetDB().Raw(sql).Order("create_at desc").Limit(limit).Offset(offset).Scan(&results).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func updateTaskInfoToDB(taskinfoList []*libmodel.FileDesc, dealFile *models.DealFile) error {
	dealFile.TaskUuid = taskinfoList[0].Uuid
	dealFile.DealCid = taskinfoList[0].Deals[0].DealCid
	err := database.SaveOne(dealFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

type DuplicatedTaskInfo struct {
	Sid               int64  `json:"sid"`
	Did               int64  `json:"did"`
	UserId            int    `json:"user_id"`
	MinerFid          string `json:"miner_fid"`
	DealStatus        string `json:"deal_status"`
	Status            string `json:"status"`
	PayloadCid        string `json:"payload_cid"`
	DealCid           string `json:"deal_cid"`
	TaskUuid          string `json:"task_uuid"`
	IpfsUrl           string `json:"ipfs_url"`
	PieceCid          string `json:"piece_cid"`
	LockPaymentStatus string `json:"lock_payment_status"`
}
