package common

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
	"strings"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func CreateSuccessResponse(data interface{}) Response {
	return Response{
		Status: constants.HTTP_STATUS_SUCCESS,
		Data:   data,
	}
}

func CreateErrorResponse(errCode int, errMsg ...string) Response {
	message := errorinfo.GetErrMsg(errCode)
	if errCode == errorinfo.ERROR_INTERNAL {
		message = ""
	}
	if errMsg != nil {
		if message != "" {
			message = message + ":"
		}
		message = message + strings.Join(errMsg, ",")
	}
	return Response{
		Status:  constants.HTTP_STATUS_ERROR,
		Message: message,
	}
}
