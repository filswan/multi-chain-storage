package routers

import (
	"fmt"
	"multi-chain-storage/common"
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

func Storage(router *gin.RouterGroup) {
	router.POST("/ipfs/upload", UploadFile)
	router.GET("/tasks/deals", GetDealListFromLocal)
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
