package routers

import (
	"fmt"
	common "multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/gin-gonic/gin"
)

func BillingManager(router *gin.RouterGroup) {
	router.GET("", GetUserBillingHistory)
	router.GET("/price/filecoin", GetFileCoinLastestPrice)
	router.GET("/deal/lockpayment/info", GetLockPaymentInfoByPayloadCid)
	router.POST("/deal/lockpayment", WriteLockPayment)
}

type LockPayment struct {
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	TxHash             string `json:"tx_hash"`
}

func WriteLockPayment(c *gin.Context) {
	var lockPayment LockPayment
	err := c.BindJSON(&lockPayment)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.ERROR_PARAM_PARSE_TO_STRUCT, err.Error()))
		return
	}

	err = service.WriteLockPayment(lockPayment.SourceFileUploadId, lockPayment.TxHash)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(""))
}

func GetLockPaymentInfoByPayloadCid(c *gin.Context) {
	URL := c.Request.URL.Query()
	var payloadCid = strings.Trim(URL.Get("payload_cid"), " ")
	if payloadCid == "" {
		errMsg := "payload_cid can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, errMsg))
		return
	}
	lockPaymentList, err := models.GetEventLockPaymentBySrcPayloadCid(payloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL))
		return
	}
	if len(lockPaymentList) > 0 {
		c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{"tx_hash": lockPaymentList[0].TxHash, "payload_cid": lockPaymentList[0].PayloadCid, "locked_fee": lockPaymentList[0].LockedFee, "token_type": lockPaymentList[0].TokenAddress}))
	} else {
		c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{"tx_hash": "", "payload_cid": payloadCid, "locked_fee": "", "token_type": ""}))
	}

}

func GetUserBillingHistory(c *gin.Context) {
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

func GetFileCoinLastestPrice(c *gin.Context) {
	latestPrice, err := client.GetFileCoinLastestPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(*latestPrice))
}
