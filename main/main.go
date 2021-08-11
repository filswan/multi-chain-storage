package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/database"
	"payment-bridge/off-chain/logs"
	"payment-bridge/off-chain/scan/browsersync"
	"payment-bridge/off-chain/scan/eth"
	"time"
)

func main() {

	config.InitConfig("")

	eth.ClientInit()

	// init database
	db := database.Init()

	go browsersync.BlockBrowserSyncAndEventLogsSync()

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

	err := r.Run(":" + config.GetConfig().Port)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

}
