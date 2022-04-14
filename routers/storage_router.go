package routers

import (
	"encoding/json"
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
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
	router.GET("/tasks/deals", GetDealListFromLocal)
	router.GET("/deal/detail/:deal_id", GetDealListFromFilink)
}

func UploadFile(c *gin.Context) {
	walletAddress := c.PostForm("wallet_address")
	if strings.Trim(walletAddress, " ") == "" {
		err := fmt.Errorf("wallet_address can not be null")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, err.Error()))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, "get file from user occurred error,please try again"))
		return
	}

	duration := c.PostForm("duration")
	if strings.Trim(duration, " ") == "" {
		err = fmt.Errorf("duraion can not be null")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, err.Error()))
		return
	}

	durationInt, err := strconv.Atoi(duration)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_TYPE_ERROR_CODE, "duration should be a number"))
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
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.SAVE_FILE_ERROR))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(uploadResult))
}

func GetDealListFromLocal(c *gin.Context) {
	URL := c.Request.URL.Query()
	pageNumber := URL.Get("page_number")
	pageSize := URL.Get("page_size")
	walletAddress := URL.Get("wallet_address")
	if strings.Trim(walletAddress, " ") == "" {
		err := fmt.Errorf("wallet_address is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, err.Error()))
		return
	}

	if (strings.Trim(pageNumber, " ") == "") || (strings.Trim(pageNumber, " ") == "0") {
		pageNumber = "1"
	} else {
		tmpPageNumber, err := strconv.Atoi(pageNumber)
		if err != nil {
			pageNumber = "1"
		} else {
			pageNumber = strconv.Itoa(tmpPageNumber)
		}
	}

	if strings.Trim(pageSize, " ") == "" {
		pageSize = constants.PAGE_SIZE_DEFAULT_VALUE
	}

	orderBy := URL.Get("order_by")
	orderByColumn, err := strconv.Atoi(orderBy)

	if err != nil {
		orderByColumn = 5
	}

	isAscending := URL.Get("is_ascending")
	ASCorDESC := "DESC"
	if strings.Trim(isAscending, " ") != "" {
		if strings.ToLower(strings.Trim(isAscending, " ")) == "y" {
			ASCorDESC = "ASC"
		}
	}

	offset, err := utils.GetOffsetByPagenumber(pageNumber, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE))
		return
	}
	payloadCid := strings.Trim(URL.Get("payload_cid"), " ")
	fileName := strings.Trim(URL.Get("file_name"), " ")
	infoList, err := service.GetSourceFiles(pageSize, strconv.FormatInt(offset, 10), walletAddress, payloadCid, fileName, orderByColumn, ASCorDESC)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))

		return
	}
	pageInfo := new(common.PageInfo)
	pageInfo.PageSize = pageSize
	pageInfo.PageNumber = pageNumber
	/*
		sourceFiles, err := models.GetSourceFilesByWalletAddress(walletAddress)
		if err != nil {
			logs.GetLogger().Error(err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_COUNT_ERROR_CODE, err.Error()))

			return
		} */
	pageInfo.TotalRecordCount = strconv.Itoa(len(infoList))
	c.JSON(http.StatusOK, common.NewSuccessResponseWithPageInfo(infoList, pageInfo))
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

type filinkParams struct {
	ID   int `json:"id"`
	Data struct {
		Deal    int    `json:"deal"`
		Network string `json:"network"`
	} `json:"data"`
}

func GetDealListFromFilink(c *gin.Context) {
	dealId := strings.Trim(c.Params.ByName("deal_id"), " ")
	if strings.Trim(dealId, " ") == "" {
		errMsg := "deal id can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}
	dealIdIntValue, err := strconv.Atoi(dealId)
	if err != nil {
		errMsg := "deal_id must be a number"
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}

	URL := c.Request.URL.Query()
	var srcFilePayloadCid = URL.Get("payload_cid")
	if strings.Trim(srcFilePayloadCid, " ") == "" {
		errMsg := "payload_cid can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}

	var walletAddress = URL.Get("wallet_address")
	if strings.Trim(srcFilePayloadCid, " ") == "" {
		errMsg := "wallet_address can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}

	result := DealOnChainResult{}
	url := config.GetConfig().FLinkUrl
	parameter := new(filinkParams)
	parameter.Data.Deal = dealIdIntValue
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

	daoSignList, err := service.GetDaoSignEventByDealId(int64(dealIdIntValue))
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_CODE+": get dao info from db occurred error"))
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
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_CODE+": get lock found info from db occurred error"))
		return
	}
	srcFile, err := models.GetSourceFileExtByPayloadCid(srcFilePayloadCid, walletAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_CODE+": get deal file info from db occurred error"))
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
