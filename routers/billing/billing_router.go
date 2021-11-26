package billing

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	common "payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
	"payment-bridge/common/utils"
	"payment-bridge/logs"
	"payment-bridge/models"
	"strconv"
	"strings"
)

func BillingManager(router *gin.RouterGroup) {
	router.GET("", GetUserBillingHistory)
	router.GET("/price/filecoin", GetFileCoinLastestPrice)
	router.POST("/deal/lockpayment/status", UpdateLockPaymentInfoByPayloadCid)
}
func UpdateLockPaymentInfoByPayloadCid(c *gin.Context) {
	logs.GetLogger().Info("~~~~~~~~~~~~~~~~~~~~~~~~ start to update lock payment status to db ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	authorization := c.Request.Header.Get("authorization")
	if len(authorization) == 0 {
		c.JSON(http.StatusUnauthorized, common.CreateErrorResponse(errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_CODE, errorinfo.NO_AUTHORIZATION_TOKEN_ERROR_MSG))
		return
	}

	payloadCid := c.PostForm("payload_cid")
	if strings.Trim(payloadCid, " ") == "" {
		errMsg := "payload_cid can not be null when updating lock payment info"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+" :"+errMsg))
		return
	}
	lockPaymentStatus := c.PostForm("lock_payment_status")
	if strings.Trim(lockPaymentStatus, " ") == "" {
		errMsg := "lock_payment_status can not be null when updating lock payment info"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+" :"+errMsg))
		return
	}
	networkName := c.PostForm("network_name")
	if strings.Trim(networkName, " ") == "" {
		errMsg := "network_name can not be null when updating lock payment info"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+" :"+errMsg))
		return
	}
	lockPaymentTx := c.PostForm("lock_payment_tx")
	if strings.Trim(lockPaymentTx, " ") == "" {
		errMsg := "lock_payment_tx can not be null when updating lock payment info"
		err := errors.New(errMsg)
		logs.GetLogger().Error(err)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+" :"+errMsg))
		return
	}
	dealList, err := models.FindDealFileList(&models.DealFile{PayloadCid: payloadCid}, "create_at desc", "10", "0")
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_MSG+": get deal info from db occurred error"))
		return
	}
	if len(dealList) > 0 {
		if dealList[0].LockPaymentTx != "" {
			err := models.UpdateDealFile(&models.DealFile{PayloadCid: payloadCid}, map[string]interface{}{"lock_payment_tx": lockPaymentTx, "lock_payment_status": lockPaymentStatus, "lock_payment_network": lockPaymentStatus})
			if err != nil {
				logs.GetLogger().Error(err)
				c.JSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.UPDATE_DATA_TO_DB_ERROR_CODE, errorinfo.UPDATE_DATA_TO_DB_ERROR_MSG+": update lock payment info to db occurred error"))
				return
			}
		}
	}
	c.JSON(http.StatusOK, common.CreateSuccessResponse(""))
}
func GetUserBillingHistory(c *gin.Context) {
	URL := c.Request.URL.Query()
	walletAddress := URL.Get("wallet_address")
	pageNumber := URL.Get("page_number")
	pageSize := URL.Get("page_size")

	if strings.Trim(pageNumber, " ") == "" {
		pageNumber = "1"
	}

	if strings.Trim(pageSize, " ") == "" {
		pageSize = constants.PAGE_SIZE_DEFAULT_VALUE
	}
	if strings.Trim(walletAddress, " ") == "" {
		errMsg := " :walletAddress can not be null"
		logs.GetLogger().Error("walletAddress")
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_MSG+errMsg))
		return
	}

	offset, err := utils.GetOffsetByPagenumber(pageNumber, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE, errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_MSG))
		return
	}

	recordCount, err := getBillingCount(walletAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_COUNT_ERROR_CODE, errorinfo.GET_RECORD_COUNT_ERROR_MSG))
		return
	}

	billingResultList, err := getBillHistoryList(walletAddress, pageSize, strconv.FormatInt(offset, 10))
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE, errorinfo.GET_RECORD_lIST_ERROR_MSG))
		return
	}

	page := new(common.PageInfo)
	page.PageNumber = pageNumber
	page.PageSize = pageSize
	page.TotalRecordCount = strconv.FormatInt(recordCount, 10)
	c.JSON(http.StatusOK, common.NewSuccessResponseWithPageInfo(billingResultList, page))
}

func getWhereCondition(txHash, walletAddress string) string {
	whereCondition := "where 1=1 "
	if strings.Trim(txHash, " ") != "" {
		whereCondition += " and tx_hash='" + txHash + "'"
	}
	if strings.Trim(walletAddress, " ") != "" {
		whereCondition += " and lower(address_from)='" + strings.ToLower(walletAddress) + "'"
	}
	logs.GetLogger().Error(whereCondition)
	return whereCondition
}

func GetFileCoinLastestPrice(c *gin.Context) {
	price, err := GetWfilPriceFromSushiPrice(polygon.WebConn.ConnWeb, "1")
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE, errorinfo.GET_LATEST_PRICE_OF_FILECOIN_ERROR_MSG))
		return
	}
	priceFloat, _ := new(big.Float).SetInt(price).Float64()
	c.JSON(http.StatusOK, common.CreateSuccessResponse(priceFloat))
}
