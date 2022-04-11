package main

import (
	"fmt"
	"multi-chain-storage/blockchain/browsersync"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/logs"
	"multi-chain-storage/models"
	"multi-chain-storage/routers"
	"multi-chain-storage/routers/billing"
	"multi-chain-storage/routers/common"
	"multi-chain-storage/routers/storage"
	"multi-chain-storage/scheduler"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()
	// init database
	db := database.Init()

	initMethod()
	browsersync.Init()

	models.RunAllTheScan()

	scheduler.CreateTaskScheduler()

	scheduler.SendDealScheduler()

	scheduler.DAOUnlockPaymentSchedule()
	//scheduler.RefundUnlockPaymentSchedule()
	scheduler.ScanDealInfoScheduler()

	go scheduler.ScanExpiredDealInfoScheduler()

	defer func() {
		err := db.Close()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}()

	r := gin.Default()
	r.MaxMultipartMemory = config.GetConfig().MaxMultipartMemory << 20
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
