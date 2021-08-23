package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"payment-bridge/blockchain/browsersync/goerli"
	"payment-bridge/blockchain/browsersync/nbai"
	"payment-bridge/blockchain/browsersync/polygon"
	"payment-bridge/blockchain/goerliclient"
	polygonclient "payment-bridge/blockchain/polygonclient"
	"payment-bridge/common/constants"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/routers"
	"time"
)

func main() {

	config.InitConfig("")

	goerliclient.ClientInit()
	polygonclient.ClientInit()
	polygonclient.ClientInit()

	// init database
	db := database.Init()

	nbai.ScanNbaiEventFromChainAndSaveEventLogData(60000, 6899437)
	go polygon.PolygonBlockBrowserSyncAndEventLogsSync()

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
