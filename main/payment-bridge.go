package main

import (
	"context"
	"fmt"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	"os"
	"payment-bridge/blockchain/browsersync"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"payment-bridge/routers"
	"payment-bridge/routers/billing"
	"payment-bridge/routers/common"
	"payment-bridge/routers/storage"
	"payment-bridge/scheduler"
	"strconv"
	"time"
)

func main() {
	LoadEnv()
	// init database
	db := database.Init()

	initMethod()
	browsersync.Init()

	//TxLogs(polygon.WebConn.ConnWeb,"0x718f2d7fd893ec3e873631c386289911496fd87dacd3c718326bf609224eaba0")

	models.RunAllTheScan()

	scheduler.SendDealScheduler()

	//scheduler.DAOUnlockPaymentSchedule()

	//polygon.ScanDaoEventFromChainAndSaveEventLogData(20965958, 20966958)

	//factory := new(scanFactory.IEventScanFactory)

	//go factory.GenerateBlockChainNetwork(constants.NETWORK_TYPE_BSC).ScanEventFromChainAndSaveDataToDb()

	//go factory.GenerateBlockChainNetwork(constants.NETWORK_TYPE_GOERLI).ScanEventFromChainAndSaveDataToDb()

	//go factory.GenerateBlockChainNetwork(constants.NETWORK_TYPE_NBAI).ScanEventFromChainAndSaveDataToDb()

	//go factory.GenerateBlockChainNetwork(constants.NETWORK_TYPE_POLYGON).ScanEventFromChainAndSaveDataToDb()

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1 := r.Group("/api/v1")
	common.HostManager(v1.Group(constants.URL_HOST_GET_COMMON))
	routers.EventLogManager(v1.Group(constants.URL_EVENT_PREFIX))
	billing.BillingManager(v1.Group(constants.URL_BILLING_PREFIX))
	storage.SendDealManager(v1.Group(constants.URL_STORAGE_PREFIX))

	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}

func initMethod() string {
	config.InitConfig("")
	return ""
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		logs.GetLogger().Error(err)
	}
	fmt.Println("name: ", os.Getenv("privateKey"))
}

func TxLogs(client *ethclient.Client, txHsh string) {
	rp, _ := client.TransactionReceipt(context.Background(), common2.HexToHash(txHsh))
	abiFile, err := goBind.SwanPaymentMetaData.GetAbi()

	//SwanPayment contract function signature
	contractUnlockFunctionSignature := polygon.GetConfig().PolygonMainnetNode.ContractUnlockFunctionSignature
	//contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
	}
	for _, vLog := range rp.Logs {
		//if log have this contractor function signer
		if vLog.Topics[0].Hex() == contractUnlockFunctionSignature {
			eventList, err := models.FindEventUnlockPayments(&models.EventUnlockPayment{TxHash: vLog.TxHash.Hex(), BlockNo: strconv.FormatUint(vLog.BlockNumber, 10)}, "id desc", "10", "0")
			if err != nil {
				logs.GetLogger().Error(err)
				continue
			}

			if len(eventList) <= 0 {
				dataList, err := abiFile.Unpack("UnlockPayment", vLog.Data)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				fmt.Println(dataList)
			}
		}
	}
}
