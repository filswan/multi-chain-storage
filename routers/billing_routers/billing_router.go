package billing_routers

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
)

func BillingManager(router *gin.RouterGroup) {
	router.GET("", GetUserBillingHistory)
}

func GetUserBillingHistory(c *gin.Context) {
	var billingRequestParam BillingRequest
	err := c.BindJSON(&billingRequestParam)
	addressType := billingRequestParam.AddressType
	addressValue := billingRequestParam.AddressValue
	pageNumber := billingRequestParam.PageNumber
	pageSize := billingRequestParam.PageSize

	offset, err := utils.GetOffsetByPagenumber(pageNumber, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE, errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_MSG))
		return
	}

	whereCondition := getWhereCondition(addressType, addressValue)

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

func getWhereCondition(addressType, addressValue string) string {
	whereCondition := "where 1=1 "
	switch addressType {
	case "tx":
		whereCondition += "and tx_hash='" + addressValue + "'"
	case "wallet":
		whereCondition += "and address_from='" + addressValue + "'"
	}
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
