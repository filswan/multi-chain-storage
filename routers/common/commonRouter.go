package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
	"payment-bridge/logs"
	"strings"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(constants.URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
	router.GET(constants.URL_SYSTEM_CONFIG_PARAMS, GetSystemConfigParams)
}

func GetSwanMinerVersion(c *gin.Context) {
	info := getSwanMinerHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

func GetSystemConfigParams(c *gin.Context) {
	authorization := c.Request.Header.Get("authorization")
	if len(strings.Trim(authorization, " ")) == 0 {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_CODE, errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_MSG))
		return
	}
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
