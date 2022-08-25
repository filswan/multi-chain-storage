package routers

import (
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/service"
	"net/http"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET("/host/info", GetHostInfo)
	router.GET("/system/params", GetSystemParams)
	router.GET("/system/params_all_chain", GetSystemParams4AllChains)
}

func GetHostInfo(c *gin.Context) {
	info := service.GetHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

func GetSystemParams(c *gin.Context) {
	params, err := utils.GetSystemParam("")
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(params))
}

func GetSystemParams4AllChains(c *gin.Context) {
	paramsPolygonMumbai, err := utils.GetSystemParam(config.GetConfig().Web3ApiUrlPolygonMumbai)
	if err != nil {
		logs.GetLogger().Error(err)
	}

	paramsBscTestnet, err := utils.GetSystemParam(config.GetConfig().Web3ApiUrlBscTestnet)
	if err != nil {
		logs.GetLogger().Error(err)
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"polygon_mumbai": paramsPolygonMumbai,
		"bsc_testnet":    paramsBscTestnet,
	}))
}
