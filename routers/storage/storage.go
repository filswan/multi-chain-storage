package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
	"payment-bridge/common/httpClient"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/routers/storage/storageService"
	"strconv"
	"strings"
)

func SendDealManager(router *gin.RouterGroup) {
	router.POST("/ipfs/upload", UploadFileToIpfs)
	//router.GET("/lotus/deal/:task_uuid", SendDeal)
	router.GET("/tasks/deals", GetDealListFromSwan)
}

func UploadFileToIpfs(c *gin.Context) {
	authorization := c.Request.Header.Get("authorization")
	if len(authorization) == 0 {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_CODE, errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_MSG))
		return
	}
	logs.GetLogger().Info(authorization)
	jwtToken := strings.TrimPrefix(authorization, "Bearer ")

	file, err := c.FormFile("file")
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+":get file from user occurred error,please try again"))
		return
	}

	duration := c.PostForm("duration")
	if strings.Trim(duration, " ") == "" {
		errMsg := "duraion can not be null"
		err = errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+":"+errMsg))
		return
	}

	durationInt, err := strconv.Atoi(duration)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.TYPE_TRANSFER_ERROR_CODE, errorinfo.TYPE_TRANSFER_ERROR_MSG+": duration is not a number"))
		return
	}

	fileInfoList, err := storageService.CreateTask(c, "", jwtToken, file, durationInt)
	if err != nil {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.SENDING_DEAL_ERROR_CODE, errorinfo.SENDING_DEAL_ERROR_MSG))
		return
	}

	if len(fileInfoList) > 0 {
		logs.GetLogger().Info("----------------------------payload_cid: ", fileInfoList[0].DataCid, "-----------------------------")
		c.JSON(http.StatusOK, common.CreateSuccessResponse(fileInfoList[0].DataCid))
		return
	} else {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.SENDING_DEAL_GET_NULL_RETURN_VALUE_CODE, errorinfo.SENDING_DEAL_GET_NULL_RETURN_VALUE_MSG))
		return
	}
}

func GetDealListFromSwan(c *gin.Context) {
	URL := c.Request.URL.Query()
	pageNumber := URL.Get("page_number")
	pageSize := URL.Get("page_size")
	authorization := c.Request.Header.Get("authorization")
	if len(authorization) == 0 {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_CODE, errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_MSG))
		return
	}

	if (strings.Trim(pageNumber, " ") == "") || (strings.Trim(pageNumber, " ") == "0") {
		pageNumber = "1"
	} else {
		tmpPageNumber, err := strconv.Atoi(pageNumber)
		pageNumber = strconv.Itoa(tmpPageNumber)
		if err != nil {
			pageNumber = "1"
		}
	}

	if strings.Trim(pageSize, " ") == "" {
		pageSize = constants.PAGE_SIZE_DEFAULT_VALUE
	}

	offset, err := utils.GetOffsetByPagenumber(pageNumber, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE, errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_MSG))
		return
	}
	url := config.GetConfig().SwanApi.ApiUrl + "/paymentgateway/deals?source_id=" + strconv.Itoa(constants.SOURCE_ID_OF_PAYMENT) +
		"&limit=" + pageSize + "&offset=" + strconv.FormatInt(offset, 10)
	header := make(http.Header)
	header.Add(constants.HTTP_REQUEST_HEADER_AUTHRORIZATION, authorization)
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, header)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_CODE, errorinfo.HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_MSG))
		return
	}
	var results *models.OfflineDealResult
	err = json.Unmarshal(response, &results)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE, errorinfo.HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_MSG))
		return
	}
	whereCondition := ""
	for i, v := range results.Data.Deals {
		whereCondition += "'" + v.PayloadCid + "'"
		if i < len(results.Data.Deals)-1 {
			whereCondition += ","
		}
	}
	fmt.Println(whereCondition)
	if strings.Trim(whereCondition, " ") == "" {
		whereCondition = "1=1"
	} else {
		whereCondition = "1=1 and payload_cid in (" + whereCondition + ")"
	}
	eventList, err := models.FindEventPolygons(whereCondition, "", strconv.Itoa(results.PagingInfo.Limit), strconv.Itoa(results.PagingInfo.Offset))
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_MSG))
		return
	}
	paidList := ""
	for _, v := range eventList {
		paidList += v.PayloadCid
	}
	for _, v := range results.Data.Deals {
		if strings.Contains(paidList, v.PayloadCid) {
			v.PayStatus = "Success"
		} else {
			v.PayStatus = "Fail"
		}
	}
	if len(eventList) == 0 {
		for _, v := range results.Data.Deals {
			v.PayStatus = "Fail"
		}
	}
	pageInfo := new(common.PageInfo)
	pageInfo.PageSize = pageSize
	pageInfo.PageNumber = pageNumber
	pageInfo.TotalRecordCount = strconv.Itoa(results.PagingInfo.TotalItems)
	c.JSON(http.StatusOK, common.NewSuccessResponseWithPageInfo(results.Data.Deals, pageInfo))
}
