package main

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/routers"
	"multi-chain-storage/service/scheduler"
	"os"
	"strconv"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	paymentChainName := constants.PAYMENT_CHAIN_NAME_POLYGON
	if len(os.Args) > 1 {
		paymentChainName = os.Args[1]
	}
	config.InitConfig(paymentChainName)

	db := database.Init()
	defer database.CloseDB(db)

	scheduler.InitScheduler()

	createGinServer()
}

func createGinServer() {
	if config.GetConfig().Release {
		gin.SetMode(gin.ReleaseMode)
	}

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
	routers.HostManager(v1.Group("common"))
	routers.BillingManager(v1.Group("billing"))
	routers.Storage(v1.Group("storage"))
	routers.Dao(v1.Group("dao"))

	err := r.Run(":" + strconv.Itoa(config.GetConfig().Port))
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}
