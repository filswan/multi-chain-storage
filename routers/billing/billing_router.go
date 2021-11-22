package billing

import (
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	common "payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
	"payment-bridge/common/utils"
	"payment-bridge/logs"
	"strconv"
	"strings"
)

func BillingManager(router *gin.RouterGroup) {
	router.GET("", GetUserBillingHistory)
	router.GET("/price/filecoin", GetFileCoinLastestPrice)
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
	c.JSON(http.StatusOK, common.CreateSuccessResponse(new(big.Float).SetInt(price)))
}
