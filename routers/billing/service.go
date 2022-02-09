package billing

import (
	"encoding/json"
	"net/http"
	"payment-bridge/common/httpClient"
	"payment-bridge/database"
	"payment-bridge/models"

	"github.com/filswan/go-swan-lib/logs"
)

func getBillHistoryList(walletAddress, limit, offset string, txHash string) ([]*BillingResult, error) {
	sql := "select a.tx_hash,a.locked_fee,b.cn_name coin_type,h.file_name,d.payload_cid,h.wallet_address address_from,c.network_name network,a.lock_payment_time,a.deadline,h.source_file_id"
	sql = sql + " from event_lock_payment a, coin b, network c, source_file d, source_file_upload_history h"
	sql = sql + " where a.coin_id=b.id and a.network_id=c.id and a.source_file_id=d.id and d.id=h.source_file_id and a.address_from=h.wallet_address and h.wallet_address=?"

	if txHash != "" {
		sql = sql + " and a.tx_hash ='" + txHash + "'"
	}

	var billingResults []*BillingResult
	err := database.GetDB().Raw(sql, walletAddress).Limit(limit).Offset(offset).Order("lock_payment_time desc").Scan(&billingResults).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}

	for _, billingRec := range billingResults {
		unlockInfo, err := models.GetEventUnlockPaymentsByPayloadCidUserWallet(billingRec.SourceFileId, billingRec.AddressFrom)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err

		}

		if unlockInfo != nil {
			billingRec.UnlockToUserAddress = unlockInfo.UnlockToUserAddress
			billingRec.UnlockToUserAmount = unlockInfo.UnlockToUserAmount
		}
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
