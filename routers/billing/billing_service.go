package billing

import (
	"encoding/json"
	"net/http"
	"payment-bridge/common/constants"
	"payment-bridge/common/httpClient"
	"payment-bridge/database"
	"payment-bridge/logs"
)

func GetFileCoinLastestPriceService() (*PriceResult, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=filecoin&vs_currencies=usd"
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	var price *PriceResult
	err = json.Unmarshal(response, &price)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return price, nil
}

func getBillHistoryList(whereCondition, limit, offset string) ([]*BillingResult, error) {
	selectColumn := "id,tx_hash,address_from,locked_fee,deadline,payload_cid,lock_payment_time,coin_type "
	sqlBsc := "select " + selectColumn + " ,'bsc' as network from " + constants.TABLE_NAME_EVENT_BSC + " " + whereCondition
	sqlGoerli := "select " + selectColumn + " ,'goerli' as network  from " + constants.TABLE_NAME_EVENT_GOERLI + " " + whereCondition
	sqlPolygon := "select " + selectColumn + " ,'polygon' as network  from " + constants.TABLE_NAME_EVENT_POLYGON + " " + whereCondition
	finalSql := sqlBsc + " union " + sqlGoerli + " union " + sqlPolygon + " order by lock_payment_time desc"
	var billingResultList []*BillingResult
	err := database.GetDB().Raw(finalSql).Scan(&billingResultList).Limit(limit).Offset(offset).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}
	return billingResultList, nil
}

func GetTaskDealsService(url string) (*PriceResult, error) {
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	var price *PriceResult
	err = json.Unmarshal(response, &price)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return price, nil
}
