package routers

import (
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"net/http"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/gin-gonic/gin"
)

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
