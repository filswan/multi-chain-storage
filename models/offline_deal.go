package models

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"

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
	Status         string  `json:"status"`
	DealId         *int64  `json:"deal_id"`
	OnChainStatus  *string `json:"on_chain_status"`
	UnlockTxHash   *string `json:"unlock_tx_hash"`
	UnlockAt       *int64  `json:"unlock_at"`
	CreateAt       int64   `json:"create_at"`
	UpdateAt       int64   `json:"update_at"`
}

type OfflineDealOut struct {
	OfflineDeal
	MinerFid string `json:"miner_fid"`
}

func GetOfflineDeals2BeScanned() ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	err := database.GetDB().Where("status not in (?,?)", constants.OFFLINE_DEAL_STATUS_SUCCESS, constants.OFFLINE_DEAL_STATUS_FAILED).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDeals2BeSigned(signerWalletId int64) ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select * from offline_deal a\n" +
		"left outer join dao_signature b on a.id=b.offline_deal_id and b.status=? and b.wallet_id_signer=?\n" +
		"where a.status=? and b.id is null \n"
	err := database.GetDB().Raw(sql, constants.DAO_SIGNATURE_STATUS_SUCCESS, signerWalletId, constants.OFFLINE_DEAL_STATUS_ACTIVE).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDeals2BeUnlocked() ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	err := database.GetDB().Where("deal_id>0 and status=?", constants.OFFLINE_DEAL_STATUS_ACTIVE).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDealsByCarFileId(carFileId int64) ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	err := database.GetDB().Where("car_file_id=?", carFileId).Find(&offlineDeals).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDealOutsByCarFileId(carFileId int64) ([]*OfflineDealOut, error) {
	var offlineDeals []*OfflineDealOut
	sql := "select a.*,b.fid miner_fid from offline_deal a,miner b where a.car_file_id=? and a.miner_id=b.id"
	err := database.GetDB().Raw(sql, carFileId).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDealByDealId(dealId int64) (*OfflineDeal, error) {
	if dealId <= 0 {
		err := fmt.Errorf("deal id must be greater than 0")
		logs.GetLogger().Error(err)
		return nil, err
	}

	var offlineDeals []*OfflineDeal
	err := database.GetDB().Where("deal_id=?", dealId).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(offlineDeals) > 0 {
		return offlineDeals[0], nil
	}

	return nil, nil
}

func UpdateOfflineDealUnlockInfo(id int64, txHashUnlock string) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["status"] = constants.OFFLINE_DEAL_STATUS_SUCCESS
	fields2BeUpdated["unlock_tx_hash"] = txHashUnlock
	fields2BeUpdated["unlock_at"] = currentUtcSecond
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(OfflineDeal{}).Where("id=?", id).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
