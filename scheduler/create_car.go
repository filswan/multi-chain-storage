package scheduler

import (
	"fmt"
	"os"
	"path/filepath"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"sync"
	"time"

	"github.com/filswan/go-swan-client/command"
	"github.com/filswan/go-swan-lib/logs"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
)

const (
	SRC_FILE_SIZE_MIN = int64(1 * 1024) // * 1024 // * 1024
	CAR_FILE_SIZE_MIN = int64(1 * 1024) // * 1024 //* 1024
	DURATION          = 500
)

var carDir string
var srcDir string
var creatingCarMutex sync.Mutex

func init() {
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

func GetSrcDir() string {
	return srcDir
}

func CreateCarScheduler() {
	c := cron.New()

	err := c.AddFunc(config.GetConfig().ScheduleRule.CreateCarRule, func() {
		logs.GetLogger().Info("start")

		creatingCarMutex.Lock()
		createCar()
		creatingCarMutex.Unlock()

		logs.GetLogger().Info("end")
	})

	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	c.Start()
}

func createCar() error {
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

	totalSize := int64(0)
	createdTimeMin := utils.GetCurrentUtcMilliSecond()
	for _, srcFile := range srcFiles {
		if srcFile.CreateAt < createdTimeMin {
			createdTimeMin = srcFile.CreateAt
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

	if totalSize < SRC_FILE_SIZE_MIN {
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

	fileDesc, err := createCarFile(carSrcDir, carDestDir)
	if err != nil {
		os.RemoveAll(carSrcDir)
		os.RemoveAll(carDestDir)
		logs.GetLogger().Error(err)
		return err
	}

	if fileDesc.CarFileSize < CAR_FILE_SIZE_MIN {
		os.RemoveAll(carSrcDir)
		os.RemoveAll(carDestDir)
		err := fmt.Errorf("car file size is less than %d", CAR_FILE_SIZE_MIN)
		logs.GetLogger().Error(err)
		return err
	}

	err = saveCarInfo2DB(fileDesc, srcFiles)
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

func createCarFile(srcDir, carDir string) (*libmodel.FileDesc, error) {
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

	logs.GetLogger().Info("car files created in ", carDir, "payload_cid=", fileDesc.PayloadCid)

	return fileDesc, nil
}

func saveCarInfo2DB(fileDesc *libmodel.FileDesc, srcFiles []*models.SourceFile) error {
	db := database.GetDBTransaction()
	dealFile := new(models.DealFile)
	dealFile.CarFileName = fileDesc.CarFileName
	dealFile.CarFilePath = fileDesc.CarFilePath
	dealFile.CarFileSize = fileDesc.CarFileSize
	dealFile.CarMd5 = fileDesc.CarFileMd5
	dealFile.PayloadCid = fileDesc.PayloadCid
	dealFile.PieceCid = fileDesc.PieceCid
	dealFile.DealCid = fileDesc.PayloadCid
	dealFile.CreateAt = utils.GetCurrentUtcMilliSecond()
	dealFile.UpdateAt = dealFile.CreateAt
	dealFile.Duration = DURATION
	dealFile.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_WAITING
	dealFile.IsDeleted = utils.GetBoolPointer(false)
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
		err = database.SaveOne(filepMap)
		if err != nil {
			db.Rollback()
			logs.GetLogger().Error(err)
			return err
		}

		sql := "update source_file set status=? where id=?"

		params := []interface{}{}
		params = append(params, constants.SOURCE_FILE_STATUS_CAR_CREATED)
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
