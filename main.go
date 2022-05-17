package main

import (
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/routers"
	"multi-chain-storage/service/scheduler"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
)

func main() {
	LoadEnv()

	db := database.Init()
	defer database.CloseDB(db)

	scheduler.InitScheduler()

	//scheduler.CreateTask()
	//scheduler.UnlockPayment()
	//scheduler.Refund()
	//scheduler.SendDeal()
	//scheduler.UpdateSourceFile2Refundable()

	createGinServer()
}

func createGinServer() {
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

func LoadEnv() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	envFile := filepath.Join(homedir, ".swan/mcs/.env")
	err = godotenv.Load(envFile)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	keyName := "privateKeyOnPolygon"
	logs.GetLogger().Info(keyName, ":", os.Getenv(keyName))
}
