package routers

import (
	"errors"
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
	router.GET("/deal/log/:offline_deal_id", GetDealLogs)
	router.POST("/mint/info", RecordMintInfo)
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

	durationStr := strings.Trim(c.PostForm("duration"), " ")
	if durationStr == "" {
		err = fmt.Errorf("duraion is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, "duration should be a number"))
		return
	}

	//if duration < 180 || duration > 530 {
	//	err := fmt.Errorf("duration must be in [180,530]")
	if duration != 525 {
		err := fmt.Errorf("duration must be 525")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE, err.Error()))
		return

	}

	fileTypeStr := strings.Trim(c.PostForm("file_type"), " ")
	if fileTypeStr == "" {
		fileTypeStr = "0"
	}

	fileType, err := strconv.Atoi(fileTypeStr)
	if err != nil {
		fileType = 0
	}

	uploadResult, err := service.SaveFile(c, file, duration, fileType, walletAddress)
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

	status := strings.Trim(URL.Get("status"), " ")
	fileName := URL.Get("file_name")
	is_minted := strings.Trim(URL.Get("is_minted"), " ")

	orderBy := strings.Trim(URL.Get("order_by"), " ")

	isAscend := strings.EqualFold(strings.Trim(URL.Get("is_ascend"), " "), "y")

	sourceFileUploads, totalRecordCount, err := service.GetSourceFileUploads(walletAddress, status, fileName, orderBy, is_minted, isAscend, limit, offset)
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

	dealId, err := strconv.ParseInt(dealIdStr, 10, 64)
	if err != nil {
		dealId = 0
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

	sourceFileUploadDeal, daoSignatures, err := service.GetSourceFileUploadDeal(sourceFileUploadId, dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"source_file_upload_deal": sourceFileUploadDeal,
		"dao_threshold":           threshold,
		"dao_signature":           daoSignatures,
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

type mintInfoUpload struct {
	SourceFileIploadId int64  `json:"source_file_upload_id"`
	TxHash             string `json:"tx_hash"`
	TokenId            string `json:"token_id"`
	MintAddress        string `json:"mint_address"`
}

func RecordMintInfo(c *gin.Context) {
	var model mintInfoUpload
	c.BindJSON(&model)

	sourceFileIploadId := model.SourceFileIploadId
	nftTxHash := model.TxHash
	tokenId := model.TokenId
	mintAddress := model.MintAddress

	if sourceFileIploadId <= 0 {
		err := fmt.Errorf("source_file_upload_id must be greater than 0")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE, err.Error()))
		return
	}

	if nftTxHash == "" || tokenId == "" || mintAddress == "" {
		errMsg := "tx_hash, token_id and mint_address cannot be empty"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, errMsg))
		return
	}

	sourceFileMint, err := service.RecordMintInfo(sourceFileIploadId, nftTxHash, tokenId, mintAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(sourceFileMint))
}
