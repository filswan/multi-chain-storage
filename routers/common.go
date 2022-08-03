package routers

import (
	"multi-chain-storage/common"
	"multi-chain-storage/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET("/host/info", GetHostInfo)
}

func GetHostInfo(c *gin.Context) {
	info := service.GetHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}
