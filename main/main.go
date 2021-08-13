package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"payment-bridge/off-chain/common/constants"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/database"
	"payment-bridge/off-chain/logs"
	"payment-bridge/off-chain/routers"
	"payment-bridge/off-chain/scan/browsersync/goerli"
	"payment-bridge/off-chain/scan/browsersync/polygon"
	"payment-bridge/off-chain/scan/goerliclient"
	polygonclient "payment-bridge/off-chain/scan/polygonclient"
	"time"
)

func main() {

	config.InitConfig("")

	goerliclient.ClientInit()
	polygonclient.ClientInit()

	// init database
	db := database.Init()
	polygon.PolygonBlockBrowserSyncAndEventLogsSync()

	go goerli.GoerliBlockBrowserSyncAndEventLogsSync()

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
	routers.EventLogManager(v1.Group(constants.URL_EVENT_PREFIX))

	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}
