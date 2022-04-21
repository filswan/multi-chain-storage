package service

import (
	"encoding/json"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/httpClient"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"net/http"

	"github.com/filswan/go-swan-lib/logs"
)

func WriteLockPayment(sourceFileUploadId int64, txHash string) error {
	err := models.CreateTransaction(sourceFileUploadId, txHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func GetTransactions(walletAddress, txHash, fileName, orderBy string, isAscend bool, limit, offset int) ([]*models.Billing, *int, error) {
	wallet, err := models.GetWalletByAddress(walletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	billings, totalRecordCount, err := models.GetTransactions(wallet.ID, txHash, fileName, orderBy, isAscend, limit, offset)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return billings, totalRecordCount, nil
}

func getBillHistoriesByWalletAddress(walletAddress string, fileName string) ([]*BillingResult, error) {
	sql := "select a.tx_hash,a.locked_fee,b.cn_name coin_type,h.file_name,d.payload_cid,h.wallet_address address_from,c.network_name network,a.lock_payment_time,a.deadline"
	sql = sql + " from event_lock_payment a, coin b, network c, source_file d, source_file_upload_history h"
	sql = sql + " where a.coin_id=b.id and a.network_id=c.id and a.source_file_id=d.id and d.id=h.source_file_id and a.address_from=h.wallet_address and h.wallet_address=?"

	if fileName != "" {
		sql = sql + " and h.file_name like lower('%" + fileName + "%') "
	}

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

type BillingResult struct {
	TxHash              string `json:"tx_hash"`
	LockedFee           string `json:"locked_fee"`
	Deadline            string `json:"deadline"`
	PayloadCid          string `json:"payload_cid"`
	AddressFrom         string `json:"address_from"`
	Network             string `json:"network"`
	CoinType            string `json:"coin_type"`
	LockPaymentTime     string `json:"lock_payment_time"`
	CreateAt            string `json:"create_at"`
	UnlockToUserAddress string `json:"unlock_to_user_address"`
	UnlockToUserAmount  string `json:"unlock_to_user_amount"`
	UnlockTime          string `json:"unlock_time"`
	FileName            string `json:"file_name"`
	SourceFileId        int64  `json:"source_file_id"`
}

type BillingRequest struct {
	TxHash        string `json:"tx_hash"`
	WalletAddress string `json:"wallet_address"`
	PageNumber    string `json:"page_number"`
	PageSize      string `json:"page_size"`
}

type PriceResult struct {
	Filecoin struct {
		Usd float64 `json:"usd"`
	} `json:"filecoin"`
}
