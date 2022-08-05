package routers

import (
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/service"
	"net/http"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET("/host/info", GetHostInfo)
	router.GET("/system/params", GetSystemParams)
}

func GetHostInfo(c *gin.Context) {
	info := service.GetHostInfo()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
}

func GetSystemParams(c *gin.Context) {
	params, err := utils.GetSystemParam()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	params1 := make(map[string]interface{})
	params1["PAYMENT_CONTRACT_ADDRESS"] = params.PaymentContractAddress
	params1["PAYMENT_RECIPIENT_ADDRESS"] = params.PaymentRecipientAddress
	params1["MINT_CONTRACT_ADDRESS"] = params.MintContractAddress
	params1["USDC_ADDRESS"] = params.UsdcAddress
	params1["PAY_MULTIPLY_FACTOR"] = params.PayMultiplyFactor
	params1["LOCK_TIME"] = params.LockTime
	params1["GAS_LIMIT"] = params.GasLimit
	params1["USDC_ADDRESS"] = params.UsdcAddress

	c.JSON(http.StatusOK, common.CreateSuccessResponse(params))
}
