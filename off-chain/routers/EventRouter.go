package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "payment-bridge/off-chain/common"
	"payment-bridge/off-chain/common/constants"
	"payment-bridge/off-chain/logs"
	"payment-bridge/off-chain/models"
)

func EventLogManager(router *gin.RouterGroup) {
	router.GET("/logs/:contractAddress/:cid", GetEventLogData)
}

func GetEventLogData(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	cid := c.Param("cid")

	eventList, err := models.FindEvents(&models.Event{ContractAddress: contractAddress, PayloadCid: cid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(constants.GET_EVENT_FROM_DB_ERROR_CODE, constants.GET_EVENT_FROM_DB_ERROR_MSG))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(eventList))
}
