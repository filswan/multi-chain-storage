package billing

import (
	"encoding/json"
	"net/http"
	"payment-bridge/common/httpClient"
	"payment-bridge/database"
	"strconv"

	"github.com/filswan/go-swan-lib/logs"
)

func getBillHistoryList(walletAddress, limit, offset string, txHash string, fileName string, orderByColumn int, ascdesc string) ([]*BillingResult, error) {
	sql := "select a.tx_hash,a.locked_fee,b.cn_name coin_type,h.file_name,d.payload_cid,h.wallet_address address_from,d.refund_amount unlock_to_user_amount," +
		"d.refund_at unlock_time,c.network_name network,a.lock_payment_time,a.deadline,h.source_file_id" +
		" from event_lock_payment a, coin b, network c, source_file d, source_file_upload_history h" +
		" where a.coin_id=b.id and a.network_id=c.id and a.source_file_id=d.id and d.id=h.source_file_id and a.address_from=h.wallet_address and h.wallet_address=?"

	if txHash != "" {
		sql = sql + " and a.tx_hash ='" + txHash + "' "
	}

	if fileName != "" {
		sql = sql + " and lower(h.file_name) =lower('" + fileName + "') "
	}

	orderClause := strconv.Itoa(orderByColumn) + " " + ascdesc

	var billingResults []*BillingResult
	err := database.GetDB().Raw(sql, walletAddress).Limit(limit).Offset(offset).Order(orderClause).Scan(&billingResults).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}

	return billingResults, nil
}

func getBillHistoriesByWalletAddress(walletAddress string) ([]*BillingResult, error) {
	sql := "select a.tx_hash,a.locked_fee,b.cn_name coin_type,h.file_name,d.payload_cid,h.wallet_address address_from,c.network_name network,a.lock_payment_time,a.deadline"
	sql = sql + " from event_lock_payment a, coin b, network c, source_file d, source_file_upload_history h"
	sql = sql + " where a.coin_id=b.id and a.network_id=c.id and a.source_file_id=d.id and d.id=h.source_file_id and a.address_from=h.wallet_address and h.wallet_address=?"

	var billingResultList []*BillingResult
	err := database.GetDB().Raw(sql, walletAddress).Order("lock_payment_time desc").Scan(&billingResultList).Error
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
