package storageService

import (
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"payment-bridge/config"
	"payment-bridge/logs"
	"time"
)

func CreateTask(c *gin.Context, taskName, jwtToken string, srcFile *multipart.FileHeader) error {
	temDirDeal := config.GetConfig().Temp.DirDeal

	logs.GetLogger().Info("temp dir is ", temDirDeal)

	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error("Cannot get home directory.")
		return err
	}
	temDirDeal = filepath.Join(homedir, temDirDeal[2:])

	err = libutils.CreateDir(temDirDeal)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	timeStr := time.Now().Format("20060102_150405")
	temDirDeal = filepath.Join(temDirDeal, timeStr)
	srcDir := filepath.Join(temDirDeal, "src")
	carDir := filepath.Join(temDirDeal, "car")
	err = libutils.CreateDir(srcDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	err = libutils.CreateDir(carDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	srcFilepath := filepath.Join(srcDir, srcFile.Filename)
	logs.GetLogger().Info("save your file to ", srcFilepath)
	err = c.SaveUploadedFile(srcFile, srcFilepath)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("car files created in ", carDir)

	confCar := clientmodel.ConfCar{
		LotusApiUrl:      config.GetConfig().Lotus.ApiUrl,
		LotusAccessToken: config.GetConfig().Lotus.AccessToken,
		InputDir:         srcDir,
		OutputDir:        carDir,
	}
	_, err = subcommand.CreateCarFiles(&confCar)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("car files created in ", carDir)

	confUpload := &clientmodel.ConfUpload{
		StorageServerType:           libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
		IpfsServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
		IpfsServerUploadUrl:         config.GetConfig().IpfsServer.UploadUrl,
		OutputDir:                   carDir,
		InputDir:                    carDir,
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
	confTask := &clientmodel.ConfTask{
		SwanApiUrl:                 config.GetConfig().SwanApi.ApiUrl,
		SwanJwtToken:               jwtToken,
		PublicDeal:                 true,
		BidMode:                    libconstants.TASK_BID_MODE_AUTO,
		VerifiedDeal:               config.GetConfig().SwanTask.VerifiedDeal,
		OfflineMode:                false,
		FastRetrieval:              config.GetConfig().SwanTask.FastRetrieval,
		MaxPrice:                   config.GetConfig().SwanTask.MaxPrice,
		StorageServerType:          libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
		WebServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
		ExpireDays:                 config.GetConfig().SwanTask.ExpireDays,
		OutputDir:                  carDir,
		InputDir:                   carDir,
		TaskName:                   &taskName,
		MinerFid:                   nil,
		Dataset:                    &taskDataset,
		Description:                &taskDescription,
		StartEpochIntervalHours:    startEpochIntervalHours,
		StartEpoch:                 startEpoch,
	}

	_, _, err = subcommand.CreateTask(confTask, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("task created")

	return nil
}
