package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"net/http"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/gin-gonic/gin"
)

func SendDealManager(router *gin.RouterGroup) {
	router.POST("/ipfs/upload", UploadFile)
	router.GET("/tasks/deals", GetDealListFromLocal)
	router.GET("/deal/detail/:deal_id", GetDealListFromFilink)
	router.GET("/deal/file/:source_file_id", GetDeals4SourceFile)
	router.GET("/dao/signature/deals", GetDealListForDaoToSign)
	router.PUT("/dao/signature/deals", RecordDealListThatHaveBeenSignedByDao)
	router.POST("/mint/info", RecordMintInfo)
	router.POST("/deal/expire", RecordExpiredRefund)
}

type UpdateSourceFileParam struct {
	SourceFileId int64           `json:"source_file_id"`
	MaxPrice     decimal.Decimal `json:"max_price"`
}

type UploadResult struct {
	SourceFileId int64  `json:"source_file_id"`
	PayloadCid   string `json:"payload_cid"`
	IpfsUrl      string `json:"ipfs_url"`
	NeedPay      int    `json:"need_pay"`
	FileSize     int64  `json:"file_size"`
}

func GetDeals4SourceFile(c *gin.Context) {
	sourceFileIdStr := strings.Trim(c.Params.ByName("source_file_id"), " ")
	if sourceFileIdStr == "" {
		errMsg := "source file id can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}

	sourceFileId, err := strconv.ParseInt(sourceFileIdStr, 10, 64)
	if err != nil {
		errMsg := "source file id should be a valid number"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}

	offlineDeals, sourceFile, err := GetOfflineDealsBySourceFileId(sourceFileId)
	if err != nil {
		logs.GetLogger().Error(err.Error())
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{
		"source_file": sourceFile,
		"deals":       offlineDeals,
	}))
}

func RecordDealListThatHaveBeenSignedByDao(c *gin.Context) {
	var dealIdList []DealIdList
	err := c.BindJSON(&dealIdList)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE))
		return
	}

	daoSignRes := []daoBackendResponse{}

	for _, v := range dealIdList {
		daosignCount := 0
		deal_id, err := strconv.ParseInt(v.DealId, 10, 64)
		if err != nil {
			logs.GetLogger().Error(err)
		}

		if v.TxHash1 != "" {
			verification1, err := VerifyDaoSigOnContract(v.TxHash1)

			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				err = SaveDaoEventFromTxHash(v.TxHash1, v.PayloadCid, v.Recipent, deal_id, verification1)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				if verification1 {
					daosignCount++
				}
			}
		}

		if v.TxHash2 != "" {
			verification2, err := VerifyDaoSigOnContract(v.TxHash2)
			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				err = SaveDaoEventFromTxHash(v.TxHash2, v.PayloadCid, v.Recipent, deal_id, verification2)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				if verification2 {
					daosignCount++
				}
			}
		}

		if v.TxHash3 != "" {
			verification3, err := VerifyDaoSigOnContract(v.TxHash3)
			if err != nil {
				logs.GetLogger().Error(err)
			} else {
				err = SaveDaoEventFromTxHash(v.TxHash3, v.PayloadCid, v.Recipent, deal_id, verification3)
				if err != nil {
					logs.GetLogger().Error(err)
				}
				if verification3 {
					daosignCount++
				}
			}
		}

		events, err := models.GetEventDaoSignaturesByDealId(deal_id)

		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			if daosignCount >= 2 || len(events) >= 2 {
				daoFetchedDeal := new(models.DaoFetchedDeal)
				dealIdIntValue, err := strconv.ParseInt(v.DealId, 10, 64)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				daoFetchedDeal.DealId = dealIdIntValue
				daoFetchedDeal.CreateAt = utils.GetCurrentUtcMilliSecond()
				err = database.SaveOne(daoFetchedDeal)
				if err != nil {
					logs.GetLogger().Error(err)
					continue
				}
				var response daoBackendResponse
				response.DealId = v.DealId
				response.PayloadCid = v.PayloadCid
				response.SuccessDaoCount = len(events)
				daoSignRes = append(daoSignRes, response)
			}
		}
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(&daoSignRes))
}

func GetDealListForDaoToSign(c *gin.Context) {
	dealList, err := GetShoulBeSignDealListFromDB()
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dealList))
}

func GetDealListForDaoByDealId(c *gin.Context) {
	dealId := c.Params.ByName("deal_id")
	if strings.Trim(dealId, " ") == "" {
		errMsg := "deal id can not be null"
		err := errors.New("")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errMsg))
		return
	}
	dealIdIntValue, err := strconv.Atoi(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.TYPE_TRANSFER_ERROR_CODE))
		return
	}
	dealList, err := GetDealListThanGreaterDealID(int64(dealIdIntValue), 0, 100)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
		return
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(dealList))
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

	daoSignList, err := GetDaoSignEventByDealId(int64(dealIdIntValue))
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
	foundInfo, err := GetLockFoundInfoByPayloadCid(srcFilePayloadCid)
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
		if srcFile.RefundStatus != nil {
			unlockStatus = *srcFile.RefundStatus == constants.PROCESS_STATUS_UNLOCK_REFUNDED
		}
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
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.TYPE_TRANSFER_ERROR_CODE, "duration is not a number"))
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

	srcFileId, payloadCid, ipfsDownloadPath, needPay, srcFileSize, err := SaveFile(c, file, durationInt, fileTypeInt, walletAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.SAVE_FILE_ERROR))
		return
	}

	uploadResult := UploadResult{
		SourceFileId: *srcFileId,
		PayloadCid:   *payloadCid,
		NeedPay:      *needPay,
		IpfsUrl:      *ipfsDownloadPath,
		FileSize:     *srcFileSize,
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
	infoList, err := GetSourceFiles(pageSize, strconv.FormatInt(offset, 10), walletAddress, payloadCid, fileName, orderByColumn, ASCorDESC)
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

func RecordMintInfo(c *gin.Context) {
	var model mintInfoUpload
	c.BindJSON(&model)

	payloadCid := model.PayloadCid
	nftTxHash := model.TxHash
	tokenId := model.TokenId
	mintAddress := model.MintAddress

	if payloadCid == "" || nftTxHash == "" || tokenId == "" || mintAddress == "" {
		errMsg := "payload_cid, tx_hash, token_id and mint_address cannot be nil"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errMsg))
		return
	}

	sourceFile, err := models.GetSourceFileByPayloadCid(payloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.SAVE_DATA_TO_DB_ERROR_CODE))
		return
	} else {
		sourceFileMint := models.SourceFileMint{
			SourceFileId: sourceFile.ID,
			NftTxHash:    nftTxHash,
			TokenId:      tokenId,
			MintAddress:  mintAddress,
			CreateAt:     utils.GetCurrentUtcMilliSecond(),
		}

		database.SaveOne(sourceFileMint)
		c.JSON(http.StatusOK, common.CreateSuccessResponse(sourceFileMint))
	}
}

type mintInfoUpload struct {
	PayloadCid  string `json:"payload_cid"`
	TxHash      string `json:"tx_hash"`
	TokenId     string `json:"token_id"`
	MintAddress string `json:"mint_address"`
}

func RecordExpiredRefund(c *gin.Context) {
	URL := c.Request.URL.Query()
	tx_hash := URL.Get("tx_hash")
	if strings.Trim(tx_hash, " ") == "" {
		err := fmt.Errorf("transaction hash is required")
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, err.Error()))
		return
	}
	event, err := SaveExpirePaymentEvent(tx_hash)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.SAVE_DATA_TO_DB_ERROR_CODE))
		return
	} else {
		c.JSON(http.StatusOK, common.CreateSuccessResponse(event))
	}
}

type daoBackendResponse struct {
	PayloadCid      string `json:"payload_cid"`
	DealId          string `json:"deal_id"`
	SuccessDaoCount int    `json:"success_dao_count"`
}
