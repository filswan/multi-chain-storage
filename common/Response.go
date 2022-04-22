package common

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/errorinfo"
)

type BasicResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func CreateSuccessResponse(_data interface{}) BasicResponse {
	return BasicResponse{
		Status: constants.HTTP_STATUS_SUCCESS,
		Code:   constants.HTTP_CODE_200_OK,
		Data:   _data,
	}
}

func CreateErrorResponse(errCode int, errMsg ...string) BasicResponse {
	message := errorinfo.GetErrMsg(errCode)
	if errMsg != nil {
		message = message + ":" + errMsg[0]
	}
	return BasicResponse{
		Status:  constants.HTTP_STATUS_ERROR,
		Code:    errCode,
		Message: message,
	}
}
