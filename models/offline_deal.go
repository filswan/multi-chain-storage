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
	err := database.GetDB().Where("status not in (?,?)", constants.OFFLINE_DEAL_STATUS_ACTIVE, constants.OFFLINE_DEAL_STATUS_FAILED).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

type Deal2PreSign struct {
	DealId              int64 `json:"deal_id"`
	SourceFileUploadCnt int   `json:"source_file_upload_cnt"`
}

func GetDeals2PreSign(signerWalletId int64) ([]*Deal2PreSign, error) {
	var deal2PreSign []*Deal2PreSign
	sql := "select a.deal_id,count(*) source_file_upload_cnt from (select a.* from offline_deal a\n" +
		"left outer join dao_pre_sign b on a.id=b.offline_deal_id and b.status=? and b.wallet_id_signer=?\n" +
		"where a.status=? and b.id is null ) a\n" +
		"left join car_file_source on a.car_file_id=b.car_file_id\n" +
		"group by a.deal_id"
	err := database.GetDB().Raw(sql, constants.DAO_PRE_SIGN_STATUS_SUCCESS, signerWalletId, constants.OFFLINE_DEAL_STATUS_ACTIVE).Scan(&deal2PreSign).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deal2PreSign, nil
}

type Deal2Sign struct {
	OfflineDealId int64 `json:"offline_deal_id"`
	DealId        int64 `json:"deal_id"`
	BatchCount    int   `json:"batch_count"`
	BatchNo       []int `json:"batch_no"`
}

func GetDeals2Sign(signerWalletId int64) ([]*Deal2Sign, error) {
	var deals2Sign []*Deal2Sign
	sql := "select a.offline_deal_id,b.deal_id,a.batch_count from dao_pre_sign a\n" +
		"left join offline_deal b on a.offline_deal_id=b.id\n" +
		"where a.source_file_upload_cnt_sign<a.source_file_upload_cnt_total and a.status=? and a.wallet_id_signer=?\n"
	err := database.GetDB().Raw(sql, constants.DAO_PRE_SIGN_STATUS_SUCCESS, signerWalletId).Scan(&deals2Sign).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals2Sign, nil
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
