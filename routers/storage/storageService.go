package storage

import (
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
	"github.com/filswan/go-swan-lib/client/swan"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"os"
	"path/filepath"
	"payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"time"
)

func SaveFileAndCreateCarAndUploadToIPFSAndSaveDb(c *gin.Context, srcFile *multipart.FileHeader, duration int) (string, bool, error) {
	temDirDeal := config.GetConfig().SwanTask.DirDeal

	logs.GetLogger().Info("temp dir is ", temDirDeal)
	ifPayloadCidExist := false
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error("Cannot get home directory.")
		return "", ifPayloadCidExist, err
	}
	temDirDeal = filepath.Join(homedir, temDirDeal[2:])

	err = libutils.CreateDir(temDirDeal)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", ifPayloadCidExist, err
	}

	timeStr := time.Now().Format("20060102_150405")
	temDirDeal = filepath.Join(temDirDeal, timeStr)
	srcDir := filepath.Join(temDirDeal, "src")
	carDir := filepath.Join(temDirDeal, "car")
	err = libutils.CreateDir(srcDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", ifPayloadCidExist, err
	}

	err = libutils.CreateDir(carDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", ifPayloadCidExist, err
	}

	srcFilepath := filepath.Join(srcDir, srcFile.Filename)
	logs.GetLogger().Info("save your file to ", srcFilepath)
	err = c.SaveUploadedFile(srcFile, srcFilepath)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", ifPayloadCidExist, err
	}
	logs.GetLogger().Info("car files created in ", carDir)

	confCar := clientmodel.ConfCar{
		LotusClientApiUrl:      config.GetConfig().Lotus.ApiUrl,
		LotusClientAccessToken: config.GetConfig().Lotus.AccessToken,
		InputDir:               srcDir,
		OutputDir:              carDir,
	}
	fileList, err := subcommand.CreateCarFiles(&confCar)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", ifPayloadCidExist, err
	}
	logs.GetLogger().Info("car files created in ", carDir, "payload_cid=", fileList[0].DataCid)

	dealList, err := models.FindDealFileList(&models.DealFile{PayloadCid: fileList[0].DataCid}, "create_at desc", "10", "0")
	if len(dealList) > 0 {
		ifPayloadCidExist = true
		return fileList[0].DataCid, ifPayloadCidExist, nil
	} else {
		sourceFile := new(models.SourceFile)
		sourceFile.FileName = srcFile.Filename
		sourceFile.FileSize = strconv.FormatInt(srcFile.Size, 10)
		sourceFile.ResourceUri = srcFilepath
		sourceFile.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
		err = database.SaveOne(sourceFile)
		if err != nil {
			logs.GetLogger().Error(err)
			return "", ifPayloadCidExist, err
		}
		err = saveDealFileAndMapRelation(fileList, sourceFile, duration)
		if err != nil {
			logs.GetLogger().Error(err)
			return "", ifPayloadCidExist, err
		}
		return fileList[0].DataCid, ifPayloadCidExist, nil
	}
}

/*
func CreateTask(){
	taskDataset := config.GetConfig().SwanTask.CuratedDataset
	taskDescription := config.GetConfig().SwanTask.Description
	startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours
	startEpoch := libutils.GetCurrentEpoch() + (startEpochIntervalHours+1)*libconstants.EPOCH_PER_HOUR
	confTask := &clientmodel.ConfTask{
		SwanApiUrl:                 config.GetConfig().SwanApi.ApiUrl,
		SwanToken:                  jwtToken,
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
		//TaskName:                   taskName,
		MinerFid:                "",
		Dataset:                 taskDataset,
		Description:             taskDescription,
		StartEpochIntervalHours: startEpochIntervalHours,
		StartEpoch:              startEpoch,
		SourceId:                constants.SOURCE_ID_OF_PAYMENT,
		Duration:                duration,
	}

	_, fileInfoList, _, err := subcommand.CreateTask(confTask, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if len(fileInfoList) > 0 {
		err = saveDealFileAndMapRelation(fileInfoList, sourceFile)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}
	}

	logs.GetLogger().Info("task created")
	return fileInfoList, nil
}*/

func saveDealFileAndMapRelation(fileInfoList []*libmodel.FileDesc, sourceFile *models.SourceFile, duration int) error {
	dealFile := new(models.DealFile)
	dealList, err := models.FindDealFileList(&models.DealFile{LockPaymentStatus: constants.LOCK_PAYMENT_STATUS_WAITING}, "create_at desc", "10", "0")
	if len(dealList) > 0 {
		dealFile = dealList[0]
		dealFile.ID = 0
		dealFile.Duration = duration
		dealFile.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_FREE
	} else {
		dealFile.CarFileName = fileInfoList[0].CarFileName
		dealFile.CarFilePath = fileInfoList[0].CarFilePath
		dealFile.CarFileSize = fileInfoList[0].CarFileSize
		dealFile.CarMd5 = fileInfoList[0].CarFileMd5
		dealFile.PayloadCid = fileInfoList[0].DataCid
		dealFile.PieceCid = fileInfoList[0].PieceCid
		dealFile.SourceFilePath = sourceFile.ResourceUri
		dealFile.DealCid = fileInfoList[0].DealCid
		dealFile.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
		dealFile.Duration = duration
		dealFile.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_WAITING
	}
	err = database.SaveOne(dealFile)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	filepMap := new(models.SourceFileDealFileMap)
	filepMap.SourceFileId = sourceFile.ID
	filepMap.DealFileId = dealFile.ID
	filepMap.FileIndex = 0
	err = database.SaveOne(filepMap)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

func GetSourceFileAndDealFileInfo(limit, offset string) ([]*SourceFileAndDealFileInfo, error) {
	sql := "select s.file_name,s.file_size,s.create_at,df.miner_fid,df.payload_cid,df.deal_cid,df.piece_cid,df.deal_status,df.deal_status as status,df.pin_status from  source_file s " +
		" inner join source_file_deal_file_map sfdfm on s.id = sfdfm.source_file_id" +
		" inner join deal_file df on sfdfm.deal_file_id = df.id"
	var results []*SourceFileAndDealFileInfo
	err := database.GetDB().Raw(sql).Order("create_at desc").Scan(&results).Limit(limit).Offset(offset).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func GetSourceFileAndDealFileInfoCount() (int64, error) {
	sql := "select count(1) as total_record from  source_file s " +
		" inner join source_file_deal_file_map sfdfm on s.id = sfdfm.source_file_id" +
		" inner join deal_file df on sfdfm.deal_file_id = df.id"
	var recordCount common.RecordCount
	err := database.GetDB().Raw(sql).Scan(&recordCount).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return recordCount.TotalRecord, nil
}

func SendAutoBidDeals(confDeal *clientmodel.ConfDeal) ([]string, [][]*libmodel.FileDesc, error) {
	err := subcommand.CreateOutputDir(confDeal.OutputDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	logs.GetLogger().Info("output dir is:", confDeal.OutputDir)

	swanClient, err := swan.SwanGetClient(confDeal.SwanApiUrl, confDeal.SwanApiKey, confDeal.SwanAccessToken, confDeal.SwanToken)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	assignedTasks, err := swanClient.SwanGetAssignedTasks()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}
	logs.GetLogger().Info("autobid Swan task count:", len(assignedTasks))
	if len(assignedTasks) == 0 {
		logs.GetLogger().Info("no autobid task to be dealt with")
		return nil, nil, nil
	}

	var tasksDeals [][]*libmodel.FileDesc
	csvFilepaths := []string{}
	for _, assignedTask := range assignedTasks {
		assignedTaskInfo, err := swanClient.SwanGetOfflineDealsByTaskUuid(assignedTask.Uuid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		deals := assignedTaskInfo.Data.Deal
		task := assignedTaskInfo.Data.Task
		dealSentNum, csvFilePath, carFiles, err := subcommand.SendAutobidDeals4Task(confDeal, deals, task, confDeal.OutputDir)
		if err != nil {
			csvFilepaths = append(csvFilepaths, csvFilePath)
			logs.GetLogger().Error(err)
			continue
		}

		tasksDeals = append(tasksDeals, carFiles)

		if dealSentNum == 0 {
			logs.GetLogger().Info(dealSentNum, " deal(s) sent for task:", task.TaskName)
			continue
		}

		status := libconstants.TASK_STATUS_DEAL_SENT
		if dealSentNum != len(deals) {
			status = libconstants.TASK_STATUS_PROGRESS_WITH_FAILURE
		}

		response, err := swanClient.SwanUpdateAssignedTask(assignedTask.Uuid, status, csvFilePath)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		logs.GetLogger().Info(response.Message)
	}

	return csvFilepaths, tasksDeals, nil
}
