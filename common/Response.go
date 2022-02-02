package common

import (
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
)

type BasicResponse struct {
	Status   string      `json:"status"`
	Code     string      `json:"code"`
	Data     interface{} `json:"data,omitempty"`
	Message  string      `json:"message,omitempty"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
}

type PageInfo struct {
	PageNumber       string `json:"page_number"`
	PageSize         string `json:"page_size"`
	TotalRecordCount string `json:"total_record_count"`
}

type MixedResponse struct {
	BasicResponse
	MixData struct {
		Success interface{} `json:"success"`
		Fail    interface{} `json:"fail"`
	} `json:"mix_data"`
}

func CreateSuccessResponse(_data interface{}) BasicResponse {
	return BasicResponse{
		Status: constants.HTTP_STATUS_SUCCESS,
		Code:   constants.HTTP_CODE_200_OK,
		Data:   _data,
	}
}

func CreateErrorResponse(errCode string, errMsg ...string) BasicResponse {
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

func NewSuccessResponseWithPageInfo(_data interface{}, _page *PageInfo) BasicResponse {
	return BasicResponse{
		Status:   constants.HTTP_STATUS_SUCCESS,
		Code:     constants.HTTP_CODE_200_OK,
		Data:     _data,
		PageInfo: _page,
	}
}

type RecordCount struct {
	TotalRecord int64 `json:"total_record"`
}
