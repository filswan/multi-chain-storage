package billing

import (
	"net/http"
	common "payment-bridge/common"
	"payment-bridge/common/constants"
	"payment-bridge/common/errorinfo"
	"payment-bridge/common/utils"
	"payment-bridge/models"
	"payment-bridge/on-chain/client"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/gin-gonic/gin"
)

func BillingManager(router *gin.RouterGroup) {
	router.GET("", GetUserBillingHistory)
	router.GET("/price/filecoin", GetFileCoinLastestPrice)
	router.GET("/deal/lockpayment/info", GetLockPaymentInfoByPayloadCid)
	router.POST("/deal/lockpayment", WriteLockPayment)
}

func WriteLockPayment(c *gin.Context) {
	var eventLockPayment models.EventLockPayment
	err := c.BindJSON(&eventLockPayment)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(err.Error()))
		return
	}

	lockedPayment, err := client.GetLockedPaymentInfo(eventLockPayment.PayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		usdcCoin, err := models.FindCoinByFullName(constants.COIN_NAME_USDC)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			eventLockPayment.CoinId = usdcCoin.ID
			eventLockPayment.NetworkId = usdcCoin.NetworkId
			eventLockPayment.TokenAddress = usdcCoin.Address
		}
	} else {
		eventLockPayment.MinPayment = lockedPayment.MinPayment
		eventLockPayment.LockedFee = lockedPayment.LockedFee
		eventLockPayment.Deadline = lockedPayment.Deadline
		eventLockPayment.TokenAddress = lockedPayment.TokenAddress
		eventLockPayment.AddressFrom = lockedPayment.AddressFrom
		eventLockPayment.AddressTo = lockedPayment.AddressTo
		eventLockPayment.LockPaymentTime = utils.GetCurrentUtcMilliSecond()
		coin, err := models.FindCoinByCoinAddress(lockedPayment.TokenAddress)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			eventLockPayment.CoinId = coin.ID
			eventLockPayment.NetworkId = coin.NetworkId
		}
	}
	srcFile, err := models.GetSourceFileByPayloadCid(eventLockPayment.PayloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		eventLockPayment.SourceFileId = srcFile.ID
	}

	err = models.CreateEventLockPayment(eventLockPayment)
	if err != nil {
		logs.GetLogger().Error(err)
		c.JSON(http.StatusOK, common.CreateErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(""))
}

func GetLockPaymentInfoByPayloadCid(c *gin.Context) {
	URL := c.Request.URL.Query()
	var payloadCid = strings.Trim(URL.Get("payload_cid"), " ")
	if payloadCid == "" {
		errMsg := "payload_cid can not be null"
		logs.GetLogger().Error(errMsg)
		c.JSON(http.StatusBadRequest, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAM_TYPE_ERROR_CODE, errMsg))
		return
	}
	lockPaymentList, err := models.GetEventLockPaymentBySrcPayloadCid(payloadCid)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
		return
	}
	if len(lockPaymentList) > 0 {
		c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{"tx_hash": lockPaymentList[0].TxHash, "payload_cid": lockPaymentList[0].PayloadCid, "locked_fee": lockPaymentList[0].LockedFee, "token_type": lockPaymentList[0].TokenAddress}))
	} else {
		c.JSON(http.StatusOK, common.CreateSuccessResponse(gin.H{"tx_hash": "", "payload_cid": payloadCid, "locked_fee": "", "token_type": ""}))
	}

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
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.HTTP_REQUEST_PARAMS_NULL_ERROR_CODE, errMsg))
		return
	}

	offset, err := utils.GetOffsetByPagenumber(pageNumber, pageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE))
		return
	}

	orderBy := URL.Get("order_by")
	orderByColumn, err := strconv.Atoi(orderBy)

	if err != nil {
		orderByColumn = 11
	}

	isAscending := URL.Get("is_ascending")
	ASCorDESC := "DESC"
	if strings.Trim(isAscending, " ") != "" {
		if strings.ToLower(strings.Trim(isAscending, " ")) == "y" {
			ASCorDESC = "ASC"
		}
	}

	txHash := strings.Trim(URL.Get("tx_hash"), " ")

	fileName := strings.Trim(URL.Get("file_name"), " ")

	totalRecords, err := getBillHistoriesByWalletAddress(walletAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_COUNT_ERROR_CODE))
		return
	}

	billingResultList, err := getBillHistoryList(walletAddress, pageSize, strconv.FormatInt(offset, 10), txHash, fileName, orderByColumn, ASCorDESC)
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_RECORD_lIST_ERROR_CODE))
		return
	}

	page := new(common.PageInfo)
	page.PageNumber = pageNumber
	page.PageSize = pageSize
	page.TotalRecordCount = strconv.Itoa(len(totalRecords))
	c.JSON(http.StatusOK, common.NewSuccessResponseWithPageInfo(billingResultList, page))
}

func GetFileCoinLastestPrice(c *gin.Context) {
	latestPrice, err := client.GetFileCoinLastestPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.CreateErrorResponse(errorinfo.GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE))
		return
	}

	c.JSON(http.StatusOK, common.CreateSuccessResponse(*latestPrice))
}
