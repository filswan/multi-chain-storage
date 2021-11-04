package billing

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
	"payment-bridge/common/utils"
	"payment-bridge/database"
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
	txHash := URL.Get("tx_hash")
	walletAddress := URL.Get("wallet_address")
	pageNumber := URL.Get("page_number")
	pageSize := URL.Get("page_size")

	if strings.Trim(pageNumber, " ") == "" {
		pageNumber = "1"
	}

	if strings.Trim(pageSize, " ") == "" {
		pageNumber = constants.PAGE_SIZE_DEFAULT_VALUE
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

	whereCondition := getWhereCondition(txHash, walletAddress)

	recordCount, err := getBillingCount(whereCondition)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_COUNT_ERROR_CODE, errorinfo.GET_RECORD_COUNT_ERROR_MSG))
		return
	}

	billingResultList, err := getBillHistoryList(whereCondition, pageSize, strconv.FormatInt(offset, 10))
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

func getBillingCount(whereCondition string) (int64, error) {
	sql := `select sum(total_record) as total_record from (
		select count(*) as total_record  from event_bsc ` + whereCondition + `
		union
		select count(*)  as totalRecord  from event_goerli ` + whereCondition + `
		union
		select count(*) as totalRecord  from event_polygon  ` + whereCondition + `
	) t`
	var recordCount RecordCount
	err := database.GetDB().Raw(sql).Scan(&recordCount).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return recordCount.TotalRecord, nil
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

func getBillHistoryList(whereCondition, limit, offset string) ([]*BillingResult, error) {
	selectColumn := "id,tx_hash,address_from,locked_fee,deadline,payload_cid"
	sqlBsc := "select " + selectColumn + " ,'bsc' as network from " + constants.TABLE_NAME_EVENT_BSC + " " + whereCondition
	sqlGoerli := "select " + selectColumn + " ,'goerli' as network  from " + constants.TABLE_NAME_EVENT_GOERLI + " " + whereCondition
	sqlPolygon := "select " + selectColumn + " ,'polygon' as network  from " + constants.TABLE_NAME_EVENT_POLYGON + " " + whereCondition
	finalSql := sqlBsc + " union " + sqlGoerli + " union " + sqlPolygon
	var billingResultList []*BillingResult
	err := database.GetDB().Raw(finalSql).Scan(&billingResultList).Limit(limit).Offset(offset).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}
	return billingResultList, nil
}

func GetFileCoinLastestPrice(c *gin.Context) {
	price, err := GetFileCoinLastestPriceService()
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE, errorinfo.GET_LATEST_PRICE_OF_FILECOIN_ERROR_MSG))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(price.Filecoin.Usd))
}
