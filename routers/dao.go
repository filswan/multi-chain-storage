package routers

import (
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/service"
	"net/http"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func Dao(router *gin.RouterGroup) {
	router.GET("/dao/signature/deals", GetDealListForDaoToSign)
	router.POST("/dao/signature", WriteDaoSignature)
}

func GetDealListForDaoToSign(c *gin.Context) {
	dealList, err := service.GetShoulBeSignDealListFromDB()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dealList))
}

type DaoSignature struct {
	DealId   int64  `json:"deal_id"`
	Recipent string `json:"recipent"`
	TxHash   string `json:"tx_hash"`
}

func WriteDaoSignature(c *gin.Context) {
	var daoSignature DaoSignature
	err := c.BindJSON(&daoSignature)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_PARSE_TO_STRUCT))
		return
	}

	if daoSignature.DealId <= 0 {
		err := fmt.Errorf("deal_id should be greater than 0")
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE))
		return
	}

	if daoSignature.TxHash == "" || !strings.HasPrefix(daoSignature.TxHash, "0x") {
		err := fmt.Errorf("tx_hash is invalid")
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE))
		return
	}

	if daoSignature.Recipent == "" {
		err := fmt.Errorf("recipent is required")
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL))
		return
	}

	err = service.WriteDaoSignature(daoSignature.TxHash, daoSignature.Recipent, daoSignature.DealId)
	if err != nil {
		logs.GetLogger().Error(err)
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(""))
}
