package storage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-bridge/common"
	"payment-bridge/common/errorinfo"
	"payment-bridge/logs"
	"payment-bridge/routers/storage/storageService"
	"strings"
)

func SendDealManager(router *gin.RouterGroup) {
	router.POST("/ipfs/upload", UploadFileToIpfs)
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
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+":file"))
		return
	}
	taskName := c.PostForm("task_name")

	err = storageService.CreateTask(c, taskName, jwtToken, file)
	if err != nil {
		c.JSON(http.StatusOK, common.CreateErrorResponse(errorinfo.SENDING_DEAL_ERROR_CODE, errorinfo.SENDING_DEAL_ERROR_MSG+":file"))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse("succeeded to send deal"))
	return
}
