package billing

import (
	"encoding/json"
	"net/http"
	"payment-bridge/common/httpClient"
	"payment-bridge/database"
	"payment-bridge/logs"
	"strings"
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

func getBillingCount(walletAddress string) (int64, error) {
	sql := "select  count(* ) as total_record " +
		"from event_polygon ep left join event_unlock_payment eup   on eup.payload_cid = ep.payload_cid where ep.address_from='" + walletAddress + "'"
	var recordCount RecordCount
	err := database.GetDB().Raw(sql).Scan(&recordCount).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return recordCount.TotalRecord, nil
}

func getBillHistoryList(walletAddress, limit, offset string) ([]*BillingResult, error) {
	finalSql := "select ep.address_from,ep.locked_fee,ep.deadline,ep.payload_cid,ep.lock_payment_time,ep.coin_type,eup.unlock_to_user_address,eup.unlock_to_user_amount,eup.unlock_time  " +
		"from event_polygon ep left join event_unlock_payment eup   on eup.payload_cid = ep.payload_cid where lower(ep.address_from)='" + strings.ToLower(walletAddress) + "'"
	//"limit " + limit + " offset " + offset

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
