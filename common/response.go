package common

import (
	"multi-chain-storage/common/constants"
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

func CreateErrorResponse(errMsg ...string) Response {
	return Response{
		Status:  constants.HTTP_STATUS_ERROR,
		Message: strings.Join(errMsg, ","),
	}
}
