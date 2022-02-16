package scheduler

import (
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"
	"time"

	"github.com/filswan/go-swan-client/command"
	libconstants "github.com/filswan/go-swan-lib/constants"
	"github.com/filswan/go-swan-lib/logs"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/shopspring/decimal"
)

func CreateTask() error {
	err := CheckSourceFilesPaid()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	for {
		numSrcFiles, err := createTask()
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		if numSrcFiles == nil || *numSrcFiles == 0 {
			logs.GetLogger().Info("0 source file created to car file")
			return nil
		}

		logs.GetLogger().Info(*numSrcFiles, " source file(s) created to car file")
	}

}

func createTask() (*int, error) {
	srcFiles, err := models.GetSourceFilesNeed2Car()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(srcFiles) == 0 {
		logs.GetLogger().Info("0 source file to be created to car file")
		return nil, nil
	}

	currentTimeStr := time.Now().Format("2006-01-02T15:04:05")
	carSrcDir := filepath.Join(carDir, "src_"+currentTimeStr)
	carDestDir := filepath.Join(carDir, "car_"+currentTimeStr)

	err = libutils.CreateDir(carSrcDir)
	if err != nil {
		logs.GetLogger().Error("creating dir:", carSrcDir, " failed,", err)
		return nil, err
	}

	totalSize := int64(0)
	currentUtcMilliSec := utils.GetCurrentUtcMilliSecond()
	createdTimeMin := currentUtcMilliSec
	var maxPrice *decimal.Decimal

	fileCoinPriceInUsdc, err := client.GetWfilPriceFromSushiPrice("1")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	fileSizeMin := config.GetConfig().SwanTask.MinFileSize
	var srcFiles2Merged []*models.SourceFileExt
	for _, srcFile := range srcFiles {
		srcFilepathTemp := filepath.Join(carSrcDir, filepath.Base(srcFile.ResourceUri))

		bytesCopied, err := libutils.CopyFile(srcFile.ResourceUri, srcFilepathTemp)
		if err != nil {
			os.Remove(srcFilepathTemp)
			logs.GetLogger().Error(err)
			continue
		}

		maxPriceTemp, err := getMaxPrice(*srcFile, fileCoinPriceInUsdc)
		if err != nil {
			os.Remove(srcFilepathTemp)
			logs.GetLogger().Error(err)
			continue
		}

		totalSize = totalSize + bytesCopied

		if srcFile.CreateAt < createdTimeMin {
			createdTimeMin = srcFile.CreateAt
		}

		if maxPrice == nil {
			maxPrice = maxPriceTemp
		} else if maxPrice.Cmp(config.GetConfig().SwanTask.MaxPrice) < 0 {
			*maxPrice = config.GetConfig().SwanTask.MaxPrice
		}

		srcFiles2Merged = append(srcFiles2Merged, srcFile)

		if totalSize >= fileSizeMin {
			logs.GetLogger().Info("total size is:", totalSize, ", ", len(srcFiles2Merged), " files to be created to car file")
			break
		}
	}

	if totalSize == 0 {
		os.RemoveAll(carSrcDir)
		logs.GetLogger().Info("0 source file to be created to car file")
		return nil, nil
	}

	passedMilliSec := currentUtcMilliSec - createdTimeMin
	createAnyway := false
	if passedMilliSec >= 24*60*60*1000 {
		createAnyway = true
	}

	if !createAnyway && totalSize < fileSizeMin {
		os.RemoveAll(carSrcDir)
		err := fmt.Errorf("source file size is not enough")
		logs.GetLogger().Error("source file size is not enough")
		return nil, err
	}

	err = libutils.CreateDir(carDestDir)
	if err != nil {
		os.RemoveAll(carSrcDir)
		logs.GetLogger().Error("creating dir:", carDestDir, " failed,", err)
		return nil, err
	}

	fileDesc, err := createTask4SrcFiles(carSrcDir, carDestDir, *maxPrice, createAnyway, fileSizeMin)
	if err != nil {
		os.RemoveAll(carSrcDir)
		os.RemoveAll(carDestDir)
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = saveCarInfo2DB(fileDesc, srcFiles2Merged, *maxPrice)
	if err != nil {
		os.RemoveAll(carSrcDir)
		os.RemoveAll(carDestDir)
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = os.RemoveAll(carSrcDir)
	if err != nil {
		logs.GetLogger().Error(err)
	}

	numSrcFiles := len(srcFiles2Merged)
	return &numSrcFiles, nil
}

func getMaxPrice(srcFile models.SourceFileExt, rate *big.Int) (*decimal.Decimal, error) {
	_, sectorSize := libutils.CalculatePieceSize(srcFile.FileSize)

	lockedFeeInFileCoin := srcFile.LockedFee.Div(decimal.NewFromFloat(constants.LOTUS_PRICE_MULTIPLE_1E18)).Div(decimal.NewFromInt(rate.Int64()))

	durationEpoch := decimal.NewFromInt(constants.DURATION_DAYS_DEFAULT * constants.EPOCH_PER_DAY)
	sectorSizeGB := decimal.NewFromFloat(sectorSize).Div(decimal.NewFromInt(constants.BYTES_1GB))

	maxPrice := lockedFeeInFileCoin.Div(sectorSizeGB).Div(durationEpoch)

	confMaxPrice := config.GetConfig().SwanTask.MaxPrice

	if maxPrice.Cmp(confMaxPrice) > 0 {
		maxPrice = confMaxPrice
	}

	logs.GetLogger().Println("payload cid: ", srcFile.PayloadCid, " max price is ", maxPrice)

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

	durationEpoch := constants.DURATION_DAYS_DEFAULT * 24 * 60 * 2
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

func saveCarInfo2DB(fileDesc *libmodel.FileDesc, srcFiles []*models.SourceFileExt, maxPrice decimal.Decimal) error {
	db := database.GetDBTransaction()
	currentUtcMilliSecond := utils.GetCurrentUtcMilliSecond()
	dealFile := models.DealFile{
		CarFileName:       fileDesc.CarFileName,
		CarFilePath:       fileDesc.CarFilePath,
		CarFileSize:       fileDesc.CarFileSize,
		CarMd5:            fileDesc.CarFileMd5,
		PayloadCid:        fileDesc.PayloadCid,
		PieceCid:          fileDesc.PieceCid,
		CreateAt:          currentUtcMilliSecond,
		UpdateAt:          currentUtcMilliSecond,
		Duration:          constants.DURATION_DAYS_DEFAULT,
		LockPaymentStatus: constants.PROCESS_STATUS_TASK_CREATED,
		MaxPrice:          maxPrice,
		TaskUuid:          fileDesc.Uuid,
	}

	err := database.SaveOneInTransaction(db, &dealFile)
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcFiles {
		filepMap := models.SourceFileDealFileMap{
			SourceFileId: srcFile.ID,
			DealFileId:   dealFile.ID,
			FileIndex:    0,
			CreateAt:     currentUtcMilliSecond,
			UpdateAt:     currentUtcMilliSecond,
		}
		err = database.SaveOneInTransaction(db, filepMap)
		if err != nil {
			db.Rollback()
			logs.GetLogger().Error(err)
			return err
		}

		sql := "update source_file set status=?,update_at=? where id=?"

		params := []interface{}{}
		params = append(params, constants.SOURCE_FILE_STATUS_TASK_CREATED)
		params = append(params, currentUtcMilliSecond)
		params = append(params, srcFile.ID)

		err = db.Exec(sql, params...).Error
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

func CheckSourceFilesPaid() error {
	srcFiles, err := models.GetSourceFilesByStatus(constants.SOURCE_FILE_STATUS_CREATED)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcFiles {
		lockedPayment, err := client.GetLockedPaymentInfo(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		coin, err := models.FindCoinByCoinAddress(lockedPayment.TokenAddress)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		srcFile, err := models.GetSourceFileByPayloadCid(srcFile.PayloadCid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		currentUtcMilliSecond := utils.GetCurrentUtcMilliSecond()
		eventLockPayment := models.EventLockPayment{
			PayloadCid:      srcFile.PayloadCid,
			MinPayment:      lockedPayment.MinPayment,
			LockedFee:       lockedPayment.LockedFee,
			Deadline:        lockedPayment.Deadline,
			TokenAddress:    lockedPayment.TokenAddress,
			AddressFrom:     lockedPayment.AddressFrom,
			AddressTo:       lockedPayment.AddressTo,
			LockPaymentTime: currentUtcMilliSecond,
			CoinId:          coin.ID,
			NetworkId:       coin.NetworkId,
			SourceFileId:    srcFile.ID,
		}

		err = models.CreateEventLockPayment(eventLockPayment)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
	}

	return nil
}
