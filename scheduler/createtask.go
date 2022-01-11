package scheduler

import (
	"fmt"
	"io/ioutil"

	//clientmodel "github.com/filswan/go-swan-client/model"
	"math/big"
	"path/filepath"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/models"
	"payment-bridge/routers/billing"
	"strconv"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/filswan/go-swan-client/command"
	libmodel "github.com/filswan/go-swan-lib/model"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"github.com/shopspring/decimal"
)

var SrcFilesDir string = ""

const (
	SRC_FILE_SIZE_MIN = 1 * 1024 // * 1024 // * 1024
	CAR_FILE_SIZE_MIN = 1 * 1024 // * 1024 //* 1024
)

var SrcDirs []SrcDirInfo

type SrcDirInfo struct {
	SrcDir    string
	TotalSize int64
	SrcFiles  []SrcFileInfo
}

type SrcFileInfo struct {
	Id         int64
	PayloadCid string
	VrfRandStr string
	Filepath   string
	FileUrl    string
}

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

func CreateTask() {
	for _, srcDir := range SrcDirs {
		if srcDir.TotalSize < SRC_FILE_SIZE_MIN {
			continue
		}
	}
}

func createCarFile() error {
	srcDir := SrcFilesDir
	srcFiles, err := ioutil.ReadDir(srcDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	srcFilesSize := int64(0)
	for _, srcFile := range srcFiles {
		srcFilesSize = srcFilesSize + srcFile.Size()
	}

	if srcFilesSize < SRC_FILE_SIZE_MIN {
		return nil
	}

	temDirDeal := filepath.Base(srcDir)

	carDir := filepath.Join(temDirDeal, "car")
	err = libutils.CreateDir(carDir)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	cmdIpfsCar := &command.CmdIpfsCar{
		LotusClientApiUrl:         config.GetConfig().Lotus.ClientApiUrl,
		LotusClientAccessToken:    config.GetConfig().Lotus.ClientAccessToken,
		InputDir:                  srcDir,
		OutputDir:                 carDir,
		GenerateMd5:               false,
		IpfsServerUploadUrlPrefix: config.GetConfig().IpfsServer.UploadUrlPrefix,
	}
	fileList, err := cmdIpfsCar.CreateIpfsCarFiles()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	carFileSize := fileList[0].CarFileSize
	if carFileSize < CAR_FILE_SIZE_MIN {
		return nil
	}

	payloadCid := fileList[0].PayloadCid

	logs.GetLogger().Info("car files created in ", carDir, "payload_cid=", payloadCid)

	srcDir = ""

	return nil
}

func saveDealFileAndMapRelation(fileDesc *libmodel.FileDesc, srcDir SrcDirInfo, duration int) error {
	db := database.GetDBTransaction()
	currentTime := utils.GetEpochInMillis()
	dealFile := new(models.DealFile)
	dealFile.CarFileName = fileDesc.CarFileName
	dealFile.CarFilePath = fileDesc.CarFilePath
	dealFile.CarFileSize = fileDesc.CarFileSize
	dealFile.CarMd5 = fileDesc.CarFileMd5
	dealFile.PayloadCid = fileDesc.PayloadCid
	dealFile.PieceCid = fileDesc.PieceCid
	dealFile.DealCid = fileDesc.PayloadCid
	dealFile.CreateAt = strconv.FormatInt(currentTime, 10)
	dealFile.UpdateAt = strconv.FormatInt(currentTime, 10)
	dealFile.Duration = duration
	dealFile.LockPaymentStatus = constants.LOCK_PAYMENT_STATUS_WAITING
	dealFile.IsDeleted = utils.GetBoolPointer(false)
	err := database.SaveOneInTransaction(db, dealFile)
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFile := range srcDir.SrcFiles {
		filepMap := new(models.SourceFileDealFileMap)
		filepMap.SourceFileId = srcFile.Id
		filepMap.DealFileId = dealFile.ID
		filepMap.FileIndex = 0
		filepMap.CreateAt = strconv.FormatInt(currentTime, 10)
		filepMap.UpdateAt = strconv.FormatInt(currentTime, 10)
		err = database.SaveOne(filepMap)
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

func DoCreateTask() error {
	fileCoinPriceInUsdc, err := billing.GetWfilPriceFromSushiPrice(polygon.WebConn.ConnWeb, "1")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	whereCondition := "lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_PROCESSING + "') and task_uuid = '' and is_deleted=false "
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "50", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range dealList {
		//duplicatedList, err := GetDuplicateTaskInfoByPayloadCid("10", "0", v.PayloadCid)
		whereCondition := "lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_PROCESSING + "') and task_uuid = '' and is_deleted=false " +
			" and payload_cid='" + v.PayloadCid + "'"
		duplicatedList, err := models.FindDealFileList(whereCondition, "create_at desc", "50", "0")
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		if len(duplicatedList) > 1 {
			dataIndex := 0
			uuid := ""
			for i, v := range duplicatedList {
				if v.TaskUuid != "" {
					uuid = v.TaskUuid
					dataIndex = i
				}
			}
			var taskInfo *models.DealFile = nil
			if uuid != "" {
				taskInfo = duplicatedList[dataIndex]
			} else {
				dataIndex = 0
				taskInfo = duplicatedList[0]
			}
			for i, v := range duplicatedList {
				if i != dataIndex {
					err = models.DeleteDealFile(&models.DealFile{ID: v.ID})
					if err != nil {
						logs.GetLogger().Error(err)
						continue
					}
					currentTime := strconv.FormatInt(utils.GetEpochInMillis(), 10)
					err = models.UpdateSourceFileDealFileMap(&models.SourceFileDealFileMap{DealFileId: v.ID}, map[string]interface{}{"deal_file_id": taskInfo.ID, "update_at": currentTime})
					if err != nil {
						logs.GetLogger().Error(err)
						continue
					}
				}
			}
		} else {
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
					/*
						confUpload := &clientmodel.ConfUpload{
							StorageServerType:           libconstants.STORAGE_SERVER_TYPE_IPFS_SERVER,
							IpfsServerDownloadUrlPrefix: config.GetConfig().IpfsServer.DownloadUrlPrefix,
							IpfsServerUploadUrl:         config.GetConfig().IpfsServer.UploadUrl,
							OutputDir:                   filepath.Dir(v.CarFilePath),
							InputDir:                    filepath.Dir(v.CarFilePath),
						}*/
					//_, err = subcommand.UploadCarFiles(confUpload)
					_, err = command.UploadCarFilesByConfig(filepath.Dir(v.CarFilePath))
					if err != nil {
						logs.GetLogger().Error(err)
						return err
					}
					logs.GetLogger().Info("car files uploaded")

					taskDataset := config.GetConfig().SwanTask.CuratedDataset
					taskDescription := config.GetConfig().SwanTask.Description
					//startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours
					//startEpoch := libutils.GetCurrentEpoch() + (startEpochIntervalHours+1)*libconstants.EPOCH_PER_HOUR
					fmt.Println(filepath.Dir(v.CarFilePath))
					/*
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
						} */
					//_, fileInfoList, _, err := subcommand.CreateTask(confTask, nil)
					//Adapt to new version of swan-client
					inoutputDir := filepath.Dir(v.CarFilePath)
					_, fileInfoList, _, err := command.CreateTaskByConfig(inoutputDir, &inoutputDir, "", "", taskDataset, taskDescription)
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
	//dealFile.Cost = taskinfoList[0].Cost
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
