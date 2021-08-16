package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"payment-bridge/scan/blockchain/browsersync/goerli"
	"payment-bridge/scan/blockchain/browsersync/polygon"
	"payment-bridge/scan/blockchain/goerliclient"
	polygonclient "payment-bridge/scan/blockchain/polygonclient"
	"payment-bridge/scan/common/constants"
	"payment-bridge/scan/config"
	"payment-bridge/scan/database"
	"payment-bridge/scan/logs"
	"payment-bridge/scan/routers"
	"time"
)

func main() {

	config.InitConfig("")

	goerliclient.ClientInit()
	polygonclient.ClientInit()

	// init database
	db := database.Init()
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
