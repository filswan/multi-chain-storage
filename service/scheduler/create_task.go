package scheduler

import (
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"os"
	"path/filepath"
	"time"

	"github.com/filswan/go-swan-client/command"
	libconstants "github.com/filswan/go-swan-lib/constants"
	"github.com/filswan/go-swan-lib/logs"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/shopspring/decimal"
)

func CreateTask() error {
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
	srcFileUploads, err := models.GetSourceFileUploadsNeed2Car()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(srcFileUploads) == 0 {
		logs.GetLogger().Info("0 source file upload to be created to car file")
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
	currentUtcMilliSec := libutils.GetCurrentUtcSecond()
	createdTimeMin := currentUtcMilliSec
	var maxPrice *decimal.Decimal

	fileCoinPriceInUsdc, err := client.GetWfilPriceFromSushiPrice("1")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	fileSizeMin := config.GetConfig().SwanTask.MinFileSize
	var srcFiles2Merged []*models.SourceFileUploadNeed2Car
	for _, srcFileUpload := range srcFileUploads {
		srcFilepathTemp := filepath.Join(carSrcDir, filepath.Base(srcFileUpload.ResourceUri))
		bytesCopied, err := libutils.CopyFile(srcFileUpload.ResourceUri, srcFilepathTemp)
		if err != nil {
			logs.GetLogger().Info(err)
			os.Remove(srcFilepathTemp)
			logs.GetLogger().Info("downloading ", srcFileUpload.IpfsUrl, " to ", srcFilepathTemp)
			err = utils.DownloadFile(srcFileUpload.IpfsUrl, srcFilepathTemp)
			if err != nil {
				logs.GetLogger().Error(err)
				os.Remove(srcFilepathTemp)
				continue
			}
			logs.GetLogger().Info("downloaded ", srcFileUpload.IpfsUrl, " to ", srcFilepathTemp)
		}

		maxPriceTemp, err := getMaxPrice(srcFileUpload.FileSize, srcFileUpload.PayAmount, fileCoinPriceInUsdc)
		if err != nil {
			logs.GetLogger().Error(err)
			os.Remove(srcFilepathTemp)
			continue
		}

		if maxPrice == nil {
			maxPrice = maxPriceTemp
		} else if maxPrice.Cmp(*maxPriceTemp) > 0 {
			*maxPrice = *maxPriceTemp
		}

		totalSize = totalSize + bytesCopied

		if srcFileUpload.CreateAt < createdTimeMin {
			createdTimeMin = srcFileUpload.CreateAt
		}

		srcFiles2Merged = append(srcFiles2Merged, srcFileUpload)

		if totalSize >= fileSizeMin || len(srcFiles2Merged) >= config.GetConfig().SwanTask.MaxFileNumPerCar {
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
	maxFileNumPerCar := config.GetConfig().SwanTask.MaxFileNumPerCar
	createAnyway := false
	if passedMilliSec >= 24*60*60*1000 {
		logs.GetLogger().Info("earliest uploaded file pass one day, create car file")
		createAnyway = true
	} else if len(srcFiles2Merged) >= maxFileNumPerCar {
		logs.GetLogger().Info(len(srcFiles2Merged), " uploaded files >= max files number in a car:", maxFileNumPerCar, ", create car file")
		createAnyway = true
	} else if totalSize >= fileSizeMin {
		logs.GetLogger().Info("total file size:", totalSize, " >= min car file size:", fileSizeMin, ", create car file")
		createAnyway = true
	}

	if !createAnyway {
		logs.GetLogger().Info("cannot meet conditions to create car file, wait")
		os.RemoveAll(carSrcDir)
		return nil, nil
	}

	err = libutils.CreateDir(carDestDir)
	if err != nil {
		logs.GetLogger().Error("creating dir:", carDestDir, " failed,", err)
		os.RemoveAll(carSrcDir)
		return nil, err
	}

	fileDesc, err := createTask4SrcFiles(carSrcDir, carDestDir, *maxPrice)
	if err != nil {
		logs.GetLogger().Error(err)
		os.RemoveAll(carSrcDir)
		//os.RemoveAll(carDestDir)
		return nil, err
	}

	err = saveCarInfo2DB(fileDesc, srcFiles2Merged, *maxPrice)
	if err != nil {
		os.RemoveAll(carSrcDir)
		//os.RemoveAll(carDestDir)
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

func getMaxPrice(fileSize int64, lockedFee decimal.Decimal, rate *big.Int) (*decimal.Decimal, error) {
	_, sectorSize := libutils.CalculatePieceSize(fileSize)

	lockedFeeInFileCoin := lockedFee.Div(decimal.NewFromFloat(libconstants.LOTUS_PRICE_MULTIPLE_1E18)).Div(decimal.NewFromInt(rate.Int64()))

	durationEpoch := decimal.NewFromInt(constants.DURATION_DAYS_DEFAULT * constants.EPOCH_PER_DAY)
	sectorSizeGB := decimal.NewFromFloat(sectorSize).Div(decimal.NewFromInt(constants.BYTES_1GB))

	maxPrice := lockedFeeInFileCoin.Div(sectorSizeGB).Div(durationEpoch)

	confMaxPrice := config.GetConfig().SwanTask.MaxPrice

	if maxPrice.Cmp(confMaxPrice) > 0 {
		maxPrice = confMaxPrice
	}

	return &maxPrice, nil
}

func createTask4SrcFiles(srcDir, carDir string, maxPrice decimal.Decimal) (*libmodel.FileDesc, error) {
	cmdIpfsCar := &command.CmdIpfsCar{
		LotusClientApiUrl:         config.GetConfig().Lotus.ClientApiUrl,
		LotusClientAccessToken:    config.GetConfig().Lotus.ClientAccessToken,
		InputDir:                  srcDir,
		OutputDir:                 carDir,
		GenerateMd5:               false,
		IpfsServerUploadUrlPrefix: config.GetConfig().IpfsServer.UploadUrlPrefix,
	}

	_, err := cmdIpfsCar.CreateIpfsCarFiles()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	logs.GetLogger().Info("car files created to ", carDir, " from ", srcDir)

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
	logs.GetLogger().Info("car files uploaded to ipfs from ", carDir)

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
		Dataset:                    config.GetConfig().SwanTask.CuratedDataset,
		Description:                config.GetConfig().SwanTask.Description,
		StartEpochHours:            config.GetConfig().SwanTask.StartEpochHours,
		SourceId:                   constants.SOURCE_ID_MCS,
		Duration:                   constants.DURATION_DAYS_DEFAULT * 24 * 60 * 2,
		MaxAutoBidCopyNumber:       5,
	}

	_, fileDescs, _, err := cmdTask.CreateTask(nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	fileDesc := fileDescs[0]

	logs.GetLogger().Info("task created for car files in ", carDir, ",payload_cid=", fileDesc.PayloadCid)

	return fileDesc, nil
}

func saveCarInfo2DB(fileDesc *libmodel.FileDesc, srcFiles []*models.SourceFileUploadNeed2Car, maxPrice decimal.Decimal) error {
	db := database.GetDBTransaction()
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	carFile := models.CarFile{
		CarFileName: fileDesc.CarFileName,
		CarFilePath: fileDesc.CarFilePath,
		CarFileSize: fileDesc.CarFileSize,
		PayloadCid:  fileDesc.PayloadCid,
		PieceCid:    fileDesc.PieceCid,
		CreateAt:    currentUtcSecond,
		UpdateAt:    currentUtcSecond,
		Duration:    constants.DURATION_DAYS_DEFAULT,
		Status:      constants.CAR_FILE_STATUS_TASK_CREATED,
		MaxPrice:    maxPrice,
		TaskUuid:    fileDesc.Uuid,
	}

	err := database.SaveOneInTransaction(db, &carFile)
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcFiles {
		carFileSource := models.CarFileSource{
			CarFileId:          carFile.ID,
			SourceFileUploadId: srcFile.SourceFileUploadId,
			CreateAt:           currentUtcSecond,
		}
		err = database.SaveOneInTransaction(db, &carFileSource)
		if err != nil {
			db.Rollback()
			logs.GetLogger().Error(err)
			return err
		}

		currentUtcSecond := libutils.GetCurrentUtcSecond()
		fields2BeUpdated := make(map[string]interface{})
		fields2BeUpdated["status"] = constants.SOURCE_FILE_UPLOAD_STATUS_TASK_CREATED
		fields2BeUpdated["update_at"] = currentUtcSecond

		err := db.Model(models.SourceFileUpload{}).Where("id=?", srcFile.SourceFileUploadId).Update(fields2BeUpdated).Error
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
