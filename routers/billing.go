package routers

import (
	common "multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
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
		c.JSON(http.StatusOK, common.CreateErrorResponse(err.Error()))
		return
	}

	err = service.WriteLockPayment(lockPayment.SourceFileUploadId, lockPayment.TxHash)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(err.Error()))
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
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}
	lockPaymentList, err := models.GetEventLockPaymentBySrcPayloadCid(payloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
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
	walletAddress := URL.Get("wallet_address")
	pageNumber := URL.Get("page_number")
	pageSize := URL.Get("page_size")

	if strings.Trim(pageNumber, " ") == "" {
		pageNumber = "1"
	}

	if strings.Trim(pageSize, " ") == "" {
		pageSize = "constants.PAGE_SIZE_DEFAULT_VALUE"
	}
	if strings.Trim(walletAddress, " ") == "" {
		errMsg := " :walletAddress can not be null"
		logs.GetLogger().Error("walletAddress")
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errMsg))
		return
	}

	offset, err := utils.GetOffsetByPagenumber(pageNumber, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE))
		return
	}

	orderBy := URL.Get("order_by")
	orderByColumn, err := strconv.Atoi(orderBy)

	if err != nil {
		orderByColumn = 11
	}

	isAscending := URL.Get("is_ascending")
	ASCorDESC := "DESC"
	if strings.Trim(isAscending, " ") != "" {
		if strings.ToLower(strings.Trim(isAscending, " ")) == "y" {
			ASCorDESC = "ASC"
		}
	}

	txHash := strings.Trim(URL.Get("tx_hash"), " ")

	fileName := strings.Trim(URL.Get("file_name"), " ")

	/*
		totalRecords, err := getBillHistoriesByWalletAddress(walletAddress, fileName)
		if err != nil {
			logs.GetLogger().Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_COUNT_ERROR_CODE))
			return
		} */

	billingResultList, err := service.GetBillHistoryList(walletAddress, pageSize, strconv.FormatInt(offset, 10), txHash, fileName, orderByColumn, ASCorDESC)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
		return
	}

	page := new(common.PageInfo)
	page.PageNumber = pageNumber
	page.PageSize = pageSize
	page.TotalRecordCount = strconv.Itoa(len(billingResultList))
	c.JSON(http.StatusOK, common.NewSuccessResponseWithPageInfo(billingResultList, page))
}

func GetFileCoinLastestPrice(c *gin.Context) {
	latestPrice, err := client.GetFileCoinLastestPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(*latestPrice))
}
