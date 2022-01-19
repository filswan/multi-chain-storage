package scheduler

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/routers/billing"
	"sync"
	"time"

	"github.com/filswan/go-swan-client/command"
	libconstants "github.com/filswan/go-swan-lib/constants"
	"github.com/filswan/go-swan-lib/logs"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"github.com/shopspring/decimal"
)

func CreateTaskScheduler() {
	Mutex := &sync.Mutex{}
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.CreateTaskRule, func() {
		logs.GetLogger().Info("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ create task  scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		Mutex.Lock()
		err := CreateTask()
		Mutex.Unlock()
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

func CreateTask() error {
	currentTimeStr := time.Now().Format("2006-01-02T15:04:05")
	carSrcDir := filepath.Join(carDir, "src_"+currentTimeStr)
	carDestDir := filepath.Join(carDir, "car_"+currentTimeStr)

	err := libutils.CreateDir(carSrcDir)
	if err != nil {
		logs.GetLogger().Error("creating dir:", carSrcDir, " failed,", err)
		return err
	}

	srcFiles, err := models.GetSourceFilesNeed2Car()
	if err != nil {
		os.RemoveAll(carSrcDir)
		logs.GetLogger().Error(err)
		return err
	}

	if len(srcFiles) == 0 {
		logs.GetLogger().Info("no source file to be created to car file")
		return nil
	}

	totalSize := int64(0)
	currentUtcMilliSec := utils.GetCurrentUtcMilliSecond()
	createdTimeMin := currentUtcMilliSec
	var maxPrice *decimal.Decimal
	for _, srcFile := range srcFiles {
		if srcFile.CreateAt < createdTimeMin {
			createdTimeMin = srcFile.CreateAt
		}

		maxPriceTemp, err := getMaxPrice(*srcFile)
		if err != nil {
			os.RemoveAll(carSrcDir)
			logs.GetLogger().Error(err)
			return err
		}

		if maxPrice == nil {
			maxPrice = maxPriceTemp
		} else if maxPrice.Cmp(config.GetConfig().SwanTask.MaxPrice) < 0 {
			*maxPrice = config.GetConfig().SwanTask.MaxPrice
		}

		srcFilepathTemp := filepath.Join(carSrcDir, srcFile.FileName)

		bytesCopied, err := libutils.CopyFile(srcFile.ResourceUri, srcFilepathTemp)
		if err != nil {
			os.RemoveAll(carSrcDir)
			logs.GetLogger().Error(err)
			return err
		}

		totalSize = totalSize + bytesCopied
	}

	passedMilliSec := currentUtcMilliSec - createdTimeMin
	createAnyway := false
	if passedMilliSec >= 24*60*60*1000 {
		createAnyway = true
	}

	fileSizeMin := config.GetConfig().SwanTask.MinFileSizeMb * 1024 * 1024

	if !createAnyway && totalSize < fileSizeMin {
		os.RemoveAll(carSrcDir)
		err := fmt.Errorf("source file size is not enough")
		logs.GetLogger().Error("source file size is not enough")
		return err
	}

	err = libutils.CreateDir(carDestDir)
	if err != nil {
		os.RemoveAll(carSrcDir)
		logs.GetLogger().Error("creating dir:", carDestDir, " failed,", err)
		return err
	}

	fileDesc, err := createTask4SrcFiles(carSrcDir, carDestDir, *maxPrice, createAnyway, fileSizeMin)
	if err != nil {
		os.RemoveAll(carSrcDir)
		os.RemoveAll(carDestDir)
		logs.GetLogger().Error(err)
		return err
	}

	err = saveCarInfo2DB(fileDesc, srcFiles, *maxPrice)
	if err != nil {
		os.RemoveAll(carSrcDir)
		os.RemoveAll(carDestDir)
		logs.GetLogger().Error(err)
		return err
	}

	err = os.RemoveAll(carSrcDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func getMaxPrice(srcFile models.SourceFile) (*decimal.Decimal, error) {
	totalLockFee, err := models.GetTotalLockFeeBySrcPayloadCid(srcFile.PayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	lockedFee := totalLockFee.IntPart()

	fileCoinPriceInUsdc, err := billing.GetWfilPriceFromSushiPrice(polygon.WebConn.ConnWeb, "1")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	maxPrice, err := GetMaxPriceForCreateTask(fileCoinPriceInUsdc, lockedFee, DURATION_DAYS, srcFile.FileSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	logs.GetLogger().Println("payload cid ", srcFile.PayloadCid, " max price is ", maxPrice)

	if maxPrice.Cmp(config.GetConfig().SwanTask.MaxPrice) < 0 {
		*maxPrice = config.GetConfig().SwanTask.MaxPrice
	}

	return maxPrice, nil
}

func GetMaxPriceForCreateTask(rate *big.Int, lockedFee int64, duration int, carFileSize int64) (*decimal.Decimal, error) {
	_, sectorSize := libutils.CalculatePieceSize(carFileSize)
	lockedFeeDecimal := decimal.NewFromInt(lockedFee)
	lockedFeeInFileCoin := lockedFeeDecimal.Div(decimal.NewFromFloat(constants.LOTUS_PRICE_MULTIPLE_1E18)).Div(decimal.NewFromInt(rate.Int64()))
	maxPrice := lockedFeeInFileCoin.Div(decimal.NewFromFloat(sectorSize).Div(decimal.NewFromInt(10204 * 1024 * 1024))).Div(decimal.NewFromInt(int64(duration)))
	return &maxPrice, nil
}

func createTask4SrcFiles(srcDir, carDir string, maxPrice decimal.Decimal, createAnyway bool, fileSizeMin int64) (*libmodel.FileDesc, error) {
	cmdIpfsCar := &command.CmdIpfsCar{
		LotusClientApiUrl:         config.GetConfig().Lotus.ClientApiUrl,
		LotusClientAccessToken:    config.GetConfig().Lotus.ClientAccessToken,
		InputDir:                  srcDir,
		OutputDir:                 carDir,
		GenerateMd5:               false,
		IpfsServerUploadUrlPrefix: config.GetConfig().IpfsServer.UploadUrlPrefix,
	}

	fileDescs, err := cmdIpfsCar.CreateIpfsCarFiles()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	fileDesc := fileDescs[0]
	if !createAnyway && fileDesc.CarFileSize < fileSizeMin {
		err := fmt.Errorf("car file size is less than %d", fileSizeMin)
		logs.GetLogger().Error(err)
		return nil, err
	}

	cmdUpload := command.CmdUpload{
		StorageServerType:           libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
		IpfsServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
		IpfsServerUploadUrlPrefix:   config.GetConfig().IpfsServer.UploadUrlPrefix,
		InputDir:                    carDir,
	}

	_, err = cmdUpload.UploadCarFiles()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	logs.GetLogger().Info("car files uploaded")

	taskDataset := config.GetConfig().SwanTask.CuratedDataset
	taskDescription := config.GetConfig().SwanTask.Description
	startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours

	durationEpoch := DURATION_DAYS * 24 * 60 * 2
	cmdTask := command.CmdTask{
		SwanApiUrl:                 config.GetConfig().SwanApi.ApiUrl,
		SwanToken:                  "",
		SwanApiKey:                 config.GetConfig().SwanApi.ApiKey,
		SwanAccessToken:            config.GetConfig().SwanApi.AccessToken,
		LotusClientApiUrl:          config.GetConfig().Lotus.ClientApiUrl,
		BidMode:                    libconstants.TASK_BID_MODE_AUTO,
		VerifiedDeal:               config.GetConfig().SwanTask.VerifiedDeal,
		OfflineMode:                false,
		FastRetrieval:              config.GetConfig().SwanTask.FastRetrieval,
		MaxPrice:                   maxPrice,
		StorageServerType:          libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
		WebServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
		ExpireDays:                 config.GetConfig().SwanTask.ExpireDays,
		InputDir:                   carDir,
		OutputDir:                  carDir,
		Dataset:                    taskDataset,
		Description:                taskDescription,
		StartEpochHours:            startEpochIntervalHours,
		SourceId:                   constants.SOURCE_ID_OF_PAYMENT,
		Duration:                   durationEpoch,
		MaxAutoBidCopyNumber:       5,
	}

	_, fileDescs, _, err = cmdTask.CreateTask(nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	fileDesc = fileDescs[0]

	logs.GetLogger().Info("car files created in ", carDir, "payload_cid=", fileDesc.PayloadCid)

	return fileDesc, nil
}

func saveCarInfo2DB(fileDesc *libmodel.FileDesc, srcFiles []*models.SourceFile, maxPrice decimal.Decimal) error {
	db := database.GetDBTransaction()
	dealFile := new(models.DealFile)
	dealFile.CarFileName = fileDesc.CarFileName
	dealFile.CarFilePath = fileDesc.CarFilePath
	dealFile.CarFileSize = fileDesc.CarFileSize
	dealFile.CarMd5 = fileDesc.CarFileMd5
	dealFile.PayloadCid = fileDesc.PayloadCid
	dealFile.PieceCid = fileDesc.PieceCid
	dealFile.CreateAt = utils.GetCurrentUtcMilliSecond()
	dealFile.UpdateAt = dealFile.CreateAt
	dealFile.Duration = DURATION_DAYS
	dealFile.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_PROCESSING
	dealFile.IsDeleted = utils.GetBoolPointer(false)
	dealFile.MaxPrice = maxPrice
	dealFile.TaskUuid = fileDesc.Uuid
	err := database.SaveOneInTransaction(db, dealFile)
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcFiles {
		filepMap := new(models.SourceFileDealFileMap)
		filepMap.SourceFileId = srcFile.ID
		filepMap.DealFileId = dealFile.ID
		filepMap.FileIndex = 0
		filepMap.CreateAt = dealFile.CreateAt
		filepMap.UpdateAt = dealFile.CreateAt
		err = database.SaveOneInTransaction(db, filepMap)
		if err != nil {
			db.Rollback()
			logs.GetLogger().Error(err)
			return err
		}

		sql := "update source_file set status=? where id=?"

		params := []interface{}{}
		params = append(params, constants.SOURCE_FILE_STATUS_TASK_CREATED)
		params = append(params, srcFile.ID)

		err := db.Exec(sql, params...).Error
		if err != nil {
			db.Rollback()
			logs.GetLogger().Error(err)
			return err
		}
	}

	err = db.Commit().Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
