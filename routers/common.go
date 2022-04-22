package routers

import (
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/service"
	"net/http"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constants.URL_HOST_GET_HOST_INFO, GetHostInfo)
	router.GET(constants.URL_SYSTEM_CONFIG_PARAMS, GetSystemParams)
}

func GetHostInfo(c *gin.Context) {
	info := service.GetHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

func GetSystemParams(c *gin.Context) {
	params, err := service.GetSystemParams()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(params))
}
