package storage

import (
	"errors"
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/errorinfo"
	"net/http"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"

	"github.com/gin-gonic/gin"
)

func SendDealManager(router *gin.RouterGroup) {
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
