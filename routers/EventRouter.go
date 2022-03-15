package routers

import (
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/goerli"
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/polygon"
	common "multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/logs"
	"multi-chain-storage/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EventLogManager(router *gin.RouterGroup) {
	router.GET("/logs/lock/:cid", GetLockPaymentEventLogData)
}

func GetLockPaymentEventLogData(c *gin.Context) {
	cid := c.Param("cid")
	eventListMap := make(map[string]interface{})
	goerliEventList, err := models.FindEventGoerlis(&models.EventGoerli{ContractAddress: goerli.GetConfig().GoerliMainnetNode.PaymentContractAddress, PayloadCid: cid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_EVENT_FROM_DB_ERROR_CODE, errorinfo.GET_EVENT_FROM_DB_ERROR_MSG))
		return
	}
	eventListMap[constants.NETWORK_TYPE_GOERLI] = goerliEventList
	polygonEventList, err := models.FindEventLockPayment(&models.EventLockPayment{ContractAddress: polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress, PayloadCid: cid}, "id desc", "", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_EVENT_FROM_DB_ERROR_CODE, errorinfo.GET_EVENT_FROM_DB_ERROR_MSG))
		return
	}
	eventListMap[constants.NETWORK_TYPE_POLYGON] = polygonEventList
	c.JSON(http.StatusOK, common.CreateSuccessResponse(eventListMap))
}
