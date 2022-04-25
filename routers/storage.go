package routers

import (
	"fmt"
	"multi-chain-storage/common"
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

func Storage(router *gin.RouterGroup) {
	router.POST("/ipfs/upload", UploadFile)
	router.GET("/tasks/deals", GetDeals)
	router.GET("/deal/detail/:deal_id", GetDealFromFlink)
	router.POST("/deal/expire", RecordExpiredRefund)
	router.GET("/deal/log/:offline_deal_id", GetDealLogs)
}

func UploadFile(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	walletAddress := c.PostForm("wallet_address")
	if strings.Trim(walletAddress, " ") == "" {
		err := fmt.Errorf("wallet_address can not be null")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, "get file from user occurred error,please try again"))
		return
	}

	duration := c.PostForm("duration")
	if strings.Trim(duration, " ") == "" {
		err = fmt.Errorf("duraion can not be null")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	durationInt, err := strconv.Atoi(duration)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, "duration should be a number"))
		return
	}
	durationInt = durationInt * 24 * 60 * 60 / 30

	fileType := c.PostForm("file_type")
	if strings.Trim(fileType, " ") == "" {
		fileType = "0"
	}

	fileTypeInt, err := strconv.Atoi(fileType)
	if err != nil {
		fileTypeInt = 0
	}

	uploadResult, err := service.SaveFile(c, file, durationInt, fileTypeInt, walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(uploadResult))
}

func GetDeals(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	URL := c.Request.URL.Query()
	pageNumber := strings.Trim(URL.Get("page_number"), " ")
	var offset int = 1
	if pageNumber != "" {
		pageNumberTemp, err := strconv.Atoi(pageNumber)
		if err != nil {
			logs.GetLogger().Error(err)
		} else if pageNumberTemp > 0 {
			offset = pageNumberTemp
		}
	}

	pageSize := strings.Trim(URL.Get("page_size"), " ")
	var limit int = constants.PAGE_SIZE_DEFAULT_VALUE
	if pageSize != "" {
		pageSizeTemp, err := strconv.Atoi(pageSize)
		if err != nil {
			logs.GetLogger().Error(err)
		} else if pageSizeTemp > 0 {
			limit = pageSizeTemp
		}
	}

	walletAddress := strings.Trim(URL.Get("wallet_address"), " ")
	if walletAddress == "" {
		err := fmt.Errorf("wallet_address is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	fileName := URL.Get("file_name")

	orderBy := strings.Trim(URL.Get("order_by"), " ")

	isAscend := strings.Trim(URL.Get("is_ascend"), " ") == "y"

	sourceFileUploads, totalRecordCount, err := service.GetSourceFileUploads(walletAddress, fileName, orderBy, isAscend, limit, offset)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"source_file_upload": sourceFileUploads,
		"total_record_count": *totalRecordCount,
	}))
}

func GetDealFromFlink(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	dealIdStr := strings.Trim(c.Params.ByName("deal_id"), " ")
	if dealIdStr == "" {
		errMsg := "deal_id is required"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, errMsg))
		return
	}
	dealId, err := strconv.Atoi(dealIdStr)
	if err != nil {
		err := fmt.Errorf("deal_id must be a valid number")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, err.Error()))
		return
	}

	URL := c.Request.URL.Query()
	var sourceFileUploadIdStr = strings.Trim(URL.Get("source_file_upload_id"), " ")
	if sourceFileUploadIdStr == "" {
		err := fmt.Errorf("source_file_upload_id is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	sourceFileUploadId, err := strconv.ParseInt(sourceFileUploadIdStr, 10, 32)
	if err != nil {
		err := fmt.Errorf("source_file_upload_id must be a valid number")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, err.Error()))
		return
	}

	if sourceFileUploadId <= 0 {
		err := fmt.Errorf("source_file_upload_id must be greater than 0")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE, err.Error()))
		return
	}

	threshold, err := client.GetThreshHold()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	sourceFileUploadDeal, err := service.GetSourceFileUploadDeal(sourceFileUploadId, dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"source_file_upload_deal": sourceFileUploadDeal,
		"dao_threshold":           threshold,
	}))
}

func GetDealLogs(c *gin.Context) {
	logs.GetLogger().Info("ip:", c.ClientIP(), ",port:", c.Request.URL.Port())
	offlineDealIdStr := strings.Trim(c.Params.ByName("offline_deal_id"), " ")
	if offlineDealIdStr == "" {
		err := fmt.Errorf("offline_deal_id is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	offlineDealId, err := strconv.ParseInt(offlineDealIdStr, 10, 64)
	if err != nil {
		err := fmt.Errorf("offline_deal_id must be a valid number")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, err.Error()))
		return
	}

	if offlineDealId <= 0 {
		err := fmt.Errorf("offline_deal_id must be greater than 0")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE, err.Error()))
		return
	}

	offlineDealLogs, err := models.GetOfflineDealLogsByOfflineDealId(offlineDealId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"offline_deal_log": offlineDealLogs,
	}))
}

func RecordExpiredRefund(c *gin.Context) {
	URL := c.Request.URL.Query()
	tx_hash := URL.Get("tx_hash")
	if strings.Trim(tx_hash, " ") == "" {
		err := fmt.Errorf("transaction hash is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}
	event, err := service.SaveExpirePaymentEvent(tx_hash)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	} else {
		c.JSON(http.StatusOK, common.CreateSuccessResponse(event))
	}
}
