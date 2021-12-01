package scheduler

import (
	"encoding/json"
	"fmt"
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
	"github.com/filswan/go-swan-lib/client/lotus"
	libconstants "github.com/filswan/go-swan-lib/constants"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/robfig/cron"
	"github.com/shopspring/decimal"
	"net/http"
	"os"
	"path/filepath"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/common/httpClient"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/routers/billing"
	"payment-bridge/routers/storage"
	"strconv"
	"strings"
	"time"
)

func SendDealScheduler() {
	c := cron.New()
	err := c.AddFunc(config.GetConfig().ScheduleRule.SendDealRule, func() {
		logs.GetLogger().Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ send deal scheduler is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := DoSendDealScheduler()
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
func DoSendDealScheduler() error {
	fmt.Println(config.GetConfig().SwanTask.RelativeEpochFromMainNetwork)
	dealList, err := GetTaskListShouldBeSendDealFromLocal()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	for _, v := range dealList {
		taskInfo, err := GetTaskStatusByUuid(v.TaskUuid)
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}
		if taskInfo.Data.Task.Status == constants.TASK_STATUS_ASSIGNED {
			//filecoin's unit price in usdc
			filPriceInUsdc, err := billing.GetWfilPriceFromSushiPrice(polygon.WebConn.ConnWeb, "1")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			//miner's unit price in usdc
			unitPriceOfMinerInUsdcForPerEpoach, err := GetMinerPerEpoachPriceInOtherCoin(taskInfo.Data.Miner.MinerID, filPriceInUsdc.Int64(), taskInfo.Data.Task.Type, v.CarFileSize)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			//estimated total price of this deal
			finalPreDealPrice := unitPriceOfMinerInUsdcForPerEpoach * int64(v.Duration)
			lockFound, err := storage.GetLockFoundInfoByPayloadCid(v.PayloadCid)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			//locked fee of this deal
			lockedFee, err := strconv.ParseInt(lockFound.LockedFee, 10, 64)
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}
			if lockedFee >= finalPreDealPrice {
				logs.GetLogger().Println("################################## start to send deal ##################################")
				logs.GetLogger().Println(" task uuid : ", v.TaskUuid)
				v.SendDealStatus = constants.SEND_DEAL_STATUS_SUCCESS
				v.MinerFid = taskInfo.Data.Miner.MinerID
				v.ClientWalletAddress = config.GetConfig().FileCoinWallet
				dealCid, err := sendDeal(v.TaskUuid, v)
				if err != nil {
					logs.GetLogger().Error(err)
					v.SendDealStatus = constants.SEND_DEAL_STATUS_FAIL
					err = database.SaveOne(v)
					if err != nil {
						logs.GetLogger().Error(err)
						continue
					}
					continue
				}
				logs.GetLogger().Println("################################## end to send deal ##################################")
				v.DealCid = dealCid
				err = database.SaveOne(v)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
			}
		}
	}
	return nil
}

func GetTaskListShouldBeSendDealFromLocal() ([]*models.DealFile, error) {
	whereCondition := "send_deal_status ='' and lower(lock_payment_status)=lower('" + constants.LOCK_PAYMENT_STATUS_SUCCESS + "') and task_uuid != '' "
	dealList, err := models.FindDealFileList(whereCondition, "create_at desc", "50", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return dealList, nil
}

func GetTaskListShouldBeSigServiceFromSwan() (*models.OfflineDealResult, error) {
	url := config.GetConfig().SwanApi.ApiUrl + config.GetConfig().SwanApi.GetShouldSendTaskUrlSuffix
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	var results *models.OfflineDealResult
	err = json.Unmarshal(response, &results)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func sendDeal(taskUuid string, file *models.DealFile) (string, error) {
	startEpochIntervalHours := config.GetConfig().SwanTask.StartEpochHours
	startEpoch := libutils.GetCurrentEpoch() + (startEpochIntervalHours+1)*libconstants.EPOCH_PER_HOUR

	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	temDirDeal := config.GetConfig().SwanTask.DirDeal
	temDirDeal = filepath.Join(homedir, temDirDeal[2:])
	err = libutils.CreateDir(temDirDeal)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}

	timeStr := time.Now().Format("20060102_150405")
	temDirDeal = filepath.Join(temDirDeal, timeStr)
	carDir := filepath.Join(temDirDeal, "car")
	confDeal := &clientmodel.ConfDeal{
		SwanApiUrl:                   config.GetConfig().SwanApi.ApiUrl,
		SwanApiKey:                   config.GetConfig().SwanApi.ApiKey,
		SwanAccessToken:              config.GetConfig().SwanApi.AccessToken,
		SenderWallet:                 config.GetConfig().FileCoinWallet,
		VerifiedDeal:                 config.GetConfig().SwanTask.VerifiedDeal,
		FastRetrieval:                config.GetConfig().SwanTask.FastRetrieval,
		SkipConfirmation:             true,
		StartEpochIntervalHours:      startEpochIntervalHours,
		StartEpoch:                   startEpoch,
		OutputDir:                    carDir,
		LotusClientApiUrl:            config.GetConfig().Lotus.ApiUrl,
		LotusClientAccessToken:       config.GetConfig().Lotus.AccessToken,
		Duration:                     file.Duration,
		RelativeEpochFromMainNetwork: config.GetConfig().SwanTask.RelativeEpochFromMainNetwork,
	}
	confDeal.DealSourceIds = append(confDeal.DealSourceIds, libconstants.TASK_SOURCE_ID_SWAN_PAYMENT)

	dealSentNum, csvFilePath, carFiles, err := subcommand.SendAutoBidDealsByTaskUuid(confDeal, taskUuid)
	if err != nil {
		logs.GetLogger().Error(err)
		return "", err
	}
	logs.GetLogger().Info("------------------------------send deal success---------------------------------")
	logs.GetLogger().Info("dealSentNum = ", dealSentNum)
	logs.GetLogger().Info("csvFilePath = ", csvFilePath)
	logs.GetLogger().Info("carFiles = ", carFiles)
	return carFiles[0].DealCid, nil
}

func CheckIfHaveLockPayment(payloadCid string) (bool, error) {
	polygonEventList, err := models.FindEventLockPayment(&models.EventLockPayment{PayloadCid: payloadCid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return false, err
	}
	if len(polygonEventList) > 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetTaskStatusByUuid(taskUuid string) (*TaskDetailResult, error) {
	url := config.GetConfig().SwanApi.ApiUrl + "/tasks/" + taskUuid
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var taskInfo *TaskDetailResult
	err = json.Unmarshal(response, &taskInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return taskInfo, nil
}

func GetMinerPerEpoachPriceInOtherCoin(minerFid string, rate int64, verifiedType string, carFileSize int64) (int64, error) {
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
	finalPrice := decimal.NewFromFloat(float64(rate)).Mul(unitPriceMinerWithFileSize)
	return finalPrice.IntPart(), nil

}

type TaskDetailResult struct {
	Data struct {
		AverageBid       string        `json:"average_bid"`
		Bid              []interface{} `json:"bid"`
		BidCount         int           `json:"bid_count"`
		Deal             []Deal        `json:"deal"`
		DealCompleteRate string        `json:"deal_complete_rate"`
		Miner            Miner         `json:"miner"`
		Poster           Poster        `json:"poster"`
		Task             Task          `json:"task"`
		TotalDealCount   int           `json:"total_deal_count"`
		TotalItems       int           `json:"total_items"`
	} `json:"data"`
	Status string `json:"status"`
}

type Miner struct {
	AddressBalance       string      `json:"address_balance"`
	AdjustedPower        string      `json:"adjusted_power"`
	AutoBidTaskCnt       int         `json:"auto_bid_task_cnt"`
	AutoBidTaskPerDay    int         `json:"auto_bid_task_per_day"`
	BidMode              int         `json:"bid_mode"`
	LastAutoBidAt        int64       `json:"last_auto_bid_at"`
	Location             string      `json:"location"`
	MaxPieceSize         string      `json:"max_piece_size"`
	MinPieceSize         string      `json:"min_piece_size"`
	MinerID              string      `json:"miner_id"`
	OfflineDealAvailable bool        `json:"offline_deal_available"`
	Price                interface{} `json:"price"`
	Score                int         `json:"score"`
	StartEpoch           int         `json:"start_epoch"`
	Status               string      `json:"status"`
	UpdateTimeStr        string      `json:"update_time_str"`
	VerifiedPrice        interface{} `json:"verified_price"`
	YearlyPrice          interface{} `json:"yearly_price"`
	YearlyVerifiedPrice  interface{} `json:"yearly_verified_price"`
}

type Task struct {
	BidMode        int         `json:"bid_mode"`
	CreatedOn      string      `json:"created_on"`
	CuratedDataset interface{} `json:"curated_dataset"`
	Description    interface{} `json:"description"`
	Duration       int         `json:"duration"`
	ExpireDays     int         `json:"expire_days"`
	FastRetrieval  int         `json:"fast_retrieval"`
	IsPublic       int         `json:"is_public"`
	MaxPrice       string      `json:"max_price"`
	MinPrice       interface{} `json:"min_price"`
	MinerID        interface{} `json:"miner_id"`
	SourceID       int         `json:"source_id"`
	Status         string      `json:"status"`
	Tags           interface{} `json:"tags"`
	TaskFileName   string      `json:"task_file_name"`
	TaskID         int         `json:"task_id"`
	TaskName       string      `json:"task_name"`
	Type           string      `json:"type"`
	UpdatedOn      string      `json:"updated_on"`
	UUID           string      `json:"uuid"`
}

type Deal struct {
	ContractID    string      `json:"contract_id"`
	Cost          interface{} `json:"cost"`
	CreatedAt     string      `json:"created_at"`
	DealCid       interface{} `json:"deal_cid"`
	FileName      string      `json:"file_name"`
	FilePath      interface{} `json:"file_path"`
	FileSize      string      `json:"file_size"`
	FileSourceURL string      `json:"file_source_url"`
	ID            int         `json:"id"`
	Md5Origin     string      `json:"md5_origin"`
	MinerID       interface{} `json:"miner_id"`
	Note          interface{} `json:"note"`
	PayloadCid    string      `json:"payload_cid"`
	PieceCid      string      `json:"piece_cid"`
	PinStatus     string      `json:"pin_status"`
	StartEpoch    int         `json:"start_epoch"`
	Status        string      `json:"status"`
	TaskID        int         `json:"task_id"`
	UpdatedAt     string      `json:"updated_at"`
	UserID        int         `json:"user_id"`
}

type Poster struct {
	AvatarURL         string      `json:"avatar_url"`
	CompleteTaskCount int         `json:"complete_task_count"`
	ContactInfo       interface{} `json:"contact_info"`
	MemberSince       string      `json:"member_since"`
}
