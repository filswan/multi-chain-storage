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
	Note           *string `json:"note"`
	CreateAt       int64   `json:"create_at"`
	UpdateAt       int64   `json:"update_at"`
}

type OfflineDealOut struct {
	OfflineDeal
	MinerFid string `json:"miner_fid"`
}

func GetOfflineDeals2BeScanned() ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	err := database.GetDB().Where("status not in (?,?,?)", constants.OFFLINE_DEAL_STATUS_ACTIVE, constants.OFFLINE_DEAL_STATUS_SUCCESS, constants.OFFLINE_DEAL_STATUS_FAILED).Find(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

type Deal2PreSign struct {
	DealId              int64 `json:"deal_id"`
	SourceFileUploadCnt int   `json:"source_file_upload_cnt"`
	BatchCount          int   `json:"batch_count"`
}

func GetDeals2PreSign(signerWalletId int64) ([]*Deal2PreSign, error) {
	var deal2PreSign []*Deal2PreSign
	sql := "select a.deal_id,ceil(count(*)/?) batch_count,count(*) source_file_upload_cnt\n" +
		"from (select a.* from offline_deal a\n" +
		"left outer join dao_pre_sign b on a.id=b.offline_deal_id and b.status=? and b.wallet_id_signer=?\n" +
		"where a.status=? and b.id is null \n" +
		"and exists (select 1 from dao_signature c where a.id=c.offline_deal_id and c.status='Success' and signed_by_hash=true)) a\n" +
		"left join car_file_source b on a.car_file_id=b.car_file_id\n" +
		"group by a.deal_id"
	err := database.GetDB().Raw(sql, constants.MAX_WCID_COUNT_IN_TRANSACTION, constants.DAO_PRE_SIGN_STATUS_SUCCESS, signerWalletId, constants.OFFLINE_DEAL_STATUS_ACTIVE).Scan(&deal2PreSign).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deal2PreSign, nil
}

type Deal2SignBatchInfo struct {
	BatchNo int      `json:"batch_no"`
	WCid    []string `json:"w_cid"`
}

type Deal2Sign struct {
	OfflineDealId int64                 `json:"offline_deal_id"`
	CarFileId     int64                 `json:"car_file_id"`
	DealId        int64                 `json:"deal_id"`
	BatchCount    int                   `json:"batch_count"`
	BatchSizeMax  int                   `json:"batch_size_max"`
	BatchInfo     []*Deal2SignBatchInfo `json:"batch_info"`
}

func GetDeals2Sign(signerWalletId int64) ([]*Deal2Sign, error) {
	var deals2Sign []*Deal2Sign
	sql := "select a.offline_deal_id,b.deal_id,b.car_file_id,a.batch_count,a.batch_size_max from dao_pre_sign a\n" +
		"left join offline_deal b on a.offline_deal_id=b.id\n" +
		"where a.source_file_upload_cnt_sign<a.source_file_upload_cnt_total and a.status=? and a.wallet_id_signer=?\n"
	err := database.GetDB().Raw(sql, constants.DAO_PRE_SIGN_STATUS_SUCCESS, signerWalletId).Scan(&deals2Sign).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals2Sign, nil
}

func GetDeals2SignHash(signerWalletId int64) ([]*Deal2Sign, error) {
	var deals2Sign []*Deal2Sign
	sql := "select a.id offline_deal_id,a.deal_id,a.car_file_id from offline_deal a\n" +
		"left outer join dao_signature b on b.offline_deal_id=a.id and b.status=? and b.wallet_id_signer=?\n" +
		"where a.status=? and b.id is null\n"
	err := database.GetDB().Raw(sql, constants.DAO_SIGNATURE_STATUS_SUCCESS, signerWalletId, constants.OFFLINE_DEAL_STATUS_ACTIVE).Scan(&deals2Sign).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals2Sign, nil
}

type Deal2Unlock struct {
	DealId int64 `json:"deal_id"`
}

func GetDeals2Unlock() ([]*Deal2Unlock, error) {
	var deals2Unlock []*Deal2Unlock
	sql := "select a.deal_id from offline_deal a\n" +
		"left join dao_pre_sign b on a.id=b.offline_deal_id\n" +
		"where b.source_file_upload_cnt_sign=b.source_file_upload_cnt_total and b.status=? and a.status=?\n" +
		"group by a.deal_id having count(*)>=?"
	err := database.GetDB().Raw(sql, constants.DAO_PRE_SIGN_STATUS_SUCCESS, constants.OFFLINE_DEAL_STATUS_ACTIVE, constants.DAO_SIGNATURE_THRESHOLD).Scan(&deals2Unlock).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals2Unlock, nil
}

func GetOfflineDealsNotSuccessBySourceFileUploadId(sourceFileUploadId int64) ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select c.* from source_file_upload a, car_file_source b, offline_deal c\n" +
		"where a.id=b.source_file_upload_id and b.car_file_id=c.car_file_id\n" +
		"  and c.status!=? and a.id=?"
	err := database.GetDB().Raw(sql, constants.OFFLINE_DEAL_STATUS_SUCCESS, sourceFileUploadId).Scan(&offlineDeals).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return offlineDeals, nil
}

func GetOfflineDealsBySourceFileUploadId(sourceFileUploadId int64) ([]*OfflineDeal, error) {
	var offlineDeals []*OfflineDeal
	sql := "select c.* from source_file_upload a, car_file_source b, offline_deal c\n" +
		"where a.id=b.source_file_upload_id and b.car_file_id=c.car_file_id\n" +
		"  and a.id=?"
	err := database.GetDB().Raw(sql, sourceFileUploadId).Scan(&offlineDeals).Error

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

func GetOfflineDealOutsBySourceFileUploadId(sourceFileUploadId int64) ([]*OfflineDealOut, error) {
	var offlineDeals []*OfflineDealOut
	sql := "select b.*,c.fid miner_fid from car_file_source a,offline_deal b,miner c\n" +
		"where a.source_file_upload_id=? and a.car_file_id=b.car_file_id and b.miner_id=c.id\n"
	err := database.GetDB().Raw(sql, sourceFileUploadId).Scan(&offlineDeals).Error

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

func UpdateOfflineDealUnlockInfo(id int64, unlockTxHash string, unlockAt int64) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["status"] = constants.OFFLINE_DEAL_STATUS_SUCCESS
	fields2BeUpdated["unlock_tx_hash"] = unlockTxHash
	fields2BeUpdated["unlock_at"] = unlockAt
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(OfflineDeal{}).Where("id=?", id).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
