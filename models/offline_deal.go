package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"
)

type OfflineDeal struct {
	Id             int64   `json:"id"`
	CarFileId      int64   `json:"car_file_id"`
	DealCid        string  `json:"deal_cid"`
	MinerId        int64   `json:"miner_id"`
	Verified       bool    `json:"verified"`
	StartEpoch     int     `json:"start_epoch"`
	SenderWalletId int64   `json:"sender_wallet_id"`
	DealId         int64   `json:"deal_id"`
	Status         string  `json:"status"`
	Note           *string `json:"note"`
	OnChainStatus  *string `json:"on_chain_status"`
	CreateAt       int64   `json:"create_at"`
	UpdateAt       int64   `json:"update_at"`
	UnlockAt       int64   `json:"unlock_at"`
}

type OfflineDealOut struct {
	OfflineDeal
	MinerFid string `json:"miner_fid"`
}

func GetOfflineDeals2BeScanned() ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	err := database.GetDB().Where("on_chain_status is null ||(on_chain_status!=? and on_chain_status!=?)", constants.DEAL_STATUS_ACTIVE, constants.DEAL_STATUS_ERROR).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDeals2BeUnlocked() ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select a.* from offline_deal a where a.deal_id>0 and a.status=?"
	err := database.GetDB().Raw(sql, constants.OFFLINE_DEAL_STATUS_CREATED).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDealsNotUnlockedByDealFileId(dealFileId int64) ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select a.* from offline_deal a where a.deal_file_id=? and a.unlock_status in (?,?)"
	err := database.GetDB().Raw(sql, dealFileId, constants.OFFLINE_DEAL_STATUS_CREATED, constants.OFFLINE_DEAL_STATUS_UNLOCK_FAILED).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDealsByCarFileId(carFileId int64) ([]*OfflineDealOut, error) {
	var offlineDeals []*OfflineDealOut
	sql := "select a.*,b.fid miner_fid from offline_deal a,miner b where a.car_file_id=? and a.miner_id=b.id"
	err := database.GetDB().Raw(sql, carFileId).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func UpdateOfflineDealUnlockStatus(id int64, unlockStatus string, messages ...string) error {
	sql := "update offline_deal set status=?,note=?,update_at=?,unlock_at=? where id=?"

	note := ""
	for _, message := range messages {
		note = note + "," + message
	}

	note = strings.Trim(note, ",")

	curUtcMilliSec := libutils.GetCurrentUtcSecond()

	params := []interface{}{}
	params = append(params, unlockStatus)
	params = append(params, note)
	params = append(params, curUtcMilliSec)
	params = append(params, curUtcMilliSec)
	params = append(params, id)

	err := database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
