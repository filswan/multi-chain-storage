package routers

import (
	"fmt"
	common "multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/gin-gonic/gin"
)

func BillingManager(router *gin.RouterGroup) {
	router.GET("", GetUserBillingHistory)
	router.GET("/deal/lockpayment/info", GetLockPaymentInfo)
	router.GET("/price/filecoin", GetFilecoinPrice)
}

func GetUserBillingHistory(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	URL := c.Request.URL.Query()
	pageNumber := strings.Trim(URL.Get("page_number"), " ")
	var offset int = 1
	if pageNumber != "" {
		pageNumberTemp, err := strconv.Atoi(pageNumber)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if pageNumberTemp > 0 {
				offset = pageNumberTemp
			}
		}
	}

	pageSize := strings.Trim(URL.Get("page_size"), " ")
	var limit int = constants.PAGE_SIZE_DEFAULT_VALUE
	if pageSize != "" {
		pageSizeTemp, err := strconv.Atoi(pageSize)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if pageSizeTemp > 0 {
				limit = pageSizeTemp
			}
		}
	}

	walletAddress := strings.Trim(URL.Get("wallet_address"), " ")
	if walletAddress == "" {
		err := fmt.Errorf("wallet_address is required")
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	orderBy := strings.Trim(URL.Get("order_by"), " ")

	isAscend := strings.Trim(URL.Get("is_ascend"), " ") == "y"

	txHash := strings.Trim(URL.Get("tx_hash"), " ")

	fileName := URL.Get("file_name")

	billings, totalRecordCount, err := service.GetTransactions(walletAddress, txHash, fileName, orderBy, isAscend, limit, offset)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"billing":            billings,
		"total_record_count": *totalRecordCount,
	}))
}

type LockPayment struct {
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	TxHash             string `json:"tx_hash"`
}

func GetLockPaymentInfo(c *gin.Context) {
	URL := c.Request.URL.Query()
	sourceFileUploadIdStr := strings.Trim(URL.Get("source_file_upload_id"), " ")
	if sourceFileUploadIdStr == "" {
		err := fmt.Errorf("source_file_upload_id is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	sourceFileUploadId, err := strconv.ParseInt(sourceFileUploadIdStr, 10, 64)
	if err != nil {
		err := fmt.Errorf("source_file_upload_id must be a valid number")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, err.Error()))
		return
	}

	sourceFileUploadInfo, err := service.GetLockPaymentInfo(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(sourceFileUploadInfo))
}

func GetFilecoinPrice(c *gin.Context) {
	params, err := utils.GetSystemParam()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(params.FilecoinPrice))
}
