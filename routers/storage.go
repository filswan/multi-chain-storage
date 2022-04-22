package routers

import (
	"encoding/json"
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/config"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"multi-chain-storage/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

func Storage(router *gin.RouterGroup) {
	router.POST("/ipfs/upload", UploadFile)
	router.GET("/tasks/deals", GetDeals)
	router.GET("/deal/detail/:deal_id", GetDealFromFlink)
	router.GET("/deal/file/:source_file_id", GetDeals4SourceFile)
	router.POST("/deal/expire", RecordExpiredRefund)
}

func UploadFile(c *gin.Context) {
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

type DealOnChainResult struct {
	JobRunID int `json:"jobRunID"`
	Data     struct {
		Status string `json:"status"`
		Data   struct {
			Deal struct {
				DealID                   int    `json:"deal_id"`
				DealCid                  string `json:"deal_cid"`
				MessageCid               string `json:"message_cid"`
				Height                   int    `json:"height"`
				PieceCid                 string `json:"piece_cid"`
				VerifiedDeal             bool   `json:"verified_deal"`
				StoragePricePerEpoch     int    `json:"storage_price_per_epoch"`
				Signature                string `json:"signature"`
				SignatureType            string `json:"signature_type"`
				CreatedAt                int    `json:"created_at"`
				PieceSizeFormat          string `json:"piece_size_format"`
				StartHeight              int    `json:"start_height"`
				EndHeight                int    `json:"end_height"`
				Client                   string `json:"client"`
				ClientCollateralFormat   string `json:"client_collateral_format"`
				Provider                 string `json:"provider"`
				ProviderTag              string `json:"provider_tag"`
				VerifiedProvider         int    `json:"verified_provider"`
				ProviderCollateralFormat string `json:"provider_collateral_format"`
				Status                   int    `json:"status"`
				NetworkName              string `json:"network_name"`
				StoragePrice             int    `json:"storage_price"`
				IpfsUrl                  string `json:"ipfs_url"`
				FileName                 string `json:"file_name"`
			} `json:"deal"`
		} `json:"data"`
		Result struct {
		} `json:"result"`
	} `json:"data"`
	Result struct {
	} `json:"result"`
	StatusCode int `json:"statusCode"`
}

type flinkParams struct {
	ID   int `json:"id"`
	Data struct {
		Deal    int    `json:"deal"`
		Network string `json:"network"`
	} `json:"data"`
}

func GetDealFromFlink(c *gin.Context) {
	dealId := strings.Trim(c.Params.ByName("deal_id"), " ")
	if dealId == "" {
		errMsg := "deal_id is required"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, errMsg))
		return
	}
	dealIdInt, err := strconv.Atoi(dealId)
	if err != nil {
		err := fmt.Errorf("deal_id must be a number")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, err.Error()))
		return
	}

	if dealIdInt <= 0 {
		err := fmt.Errorf("deal_id must be greater than 0")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_INVALID_VALUE, err.Error()))
		return
	}

	URL := c.Request.URL.Query()
	var srcFilePayloadCid = strings.Trim(URL.Get("payload_cid"), " ")
	if srcFilePayloadCid == "" {
		err := fmt.Errorf("payload_cid is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	var walletAddress = strings.Trim(URL.Get("wallet_address"), " ")
	if walletAddress == "" {
		err := fmt.Errorf("wallet_address is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, err.Error()))
		return
	}

	result := DealOnChainResult{}
	url := config.GetConfig().FLinkUrl
	parameter := new(flinkParams)
	parameter.Data.Deal = dealIdInt
	parameter.Data.Network = config.GetConfig().FilecoinNetwork

	response, err := web.HttpGetNoToken(url, parameter)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		err = json.Unmarshal(response, &result)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}

	daoSignList, err := service.GetDaoSignEventByDealId(int64(dealIdInt))
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}
	signedDaoCount := 0
	for _, v := range daoSignList {
		if strings.Trim(v.PayloadCid, " ") != "" {
			signedDaoCount++
		}
	}
	foundInfo, err := service.GetLockFoundInfoByPayloadCid(srcFilePayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}
	srcFile, err := models.GetSourceFileExtByPayloadCid(srcFilePayloadCid, walletAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	unlockStatus := false
	if srcFile != nil {
		result.Data.Data.Deal.IpfsUrl = srcFile.IpfsUrl
		result.Data.Data.Deal.FileName = srcFile.FileName
		//if srcFile.RefundStatus != nil {
		//	unlockStatus = *srcFile.RefundStatus == constants.PROCESS_STATUS_UNLOCK_REFUNDED
		//}
	}
	threshHold, err := client.GetThreshHold()
	if err != nil {
		logs.GetLogger().Error(err)
	}

	result.Data.Data.Deal.CreatedAt = result.Data.Data.Deal.CreatedAt * 1000
	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"unlock_status":    unlockStatus,
		"dao_thresh_hold":  threshHold,
		"signed_dao_count": signedDaoCount,
		"dao_total_count":  len(daoSignList),
		"deal":             result.Data.Data.Deal,
		"found":            foundInfo,
		"dao":              daoSignList,
	}))
}
func GetDeals4SourceFile(c *gin.Context) {
	sourceFileIdStr := strings.Trim(c.Params.ByName("source_file_id"), " ")
	if sourceFileIdStr == "" {
		errMsg := "source file id can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_NULL, errMsg))
		return
	}

	sourceFileId, err := strconv.ParseInt(sourceFileIdStr, 10, 64)
	if err != nil {
		errMsg := "source file id should be a valid number"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_PARAM_WRONG_TYPE, errMsg))
		return
	}

	offlineDeals, sourceFile, err := service.GetOfflineDealsBySourceFileId(sourceFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.ERROR_INTERNAL, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"source_file": sourceFile,
		"deals":       offlineDeals,
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
