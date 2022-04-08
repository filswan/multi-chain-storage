package common

import (
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/logs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constants.URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
	router.GET(constants.URL_SYSTEM_CONFIG_PARAMS, GetSystemConfigParams)
}

func GetSwanMinerVersion(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	info := getSwanMinerHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

func GetSystemConfigParams(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	limit := strings.Trim(c.Query("limit"), " ")
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}

	params, err := getSystemConfigParams(limit)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_MSG+":getting system config from db occurred error"))
		return
	}
	config := map[string]string{}
	for _, v := range params {
		config[v.ParamKey] = v.ParamValue
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(config))
}
