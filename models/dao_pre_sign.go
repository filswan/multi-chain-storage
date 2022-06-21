package models

import (
	"math"
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type DaoPreSign struct {
	Id                       int64  `json:"id"`
	OfflineDealId            int64  `json:"offline_deal_id"`
	TxHash                   string `json:"tx_hash"`
	BatchNumber              int    `json:"batch_number"`
	BatchSizeMax             int    `json:"batch_size_max"`
	SourceFileUploadCntTotal int    `json:"source_file_upload_cnt_total"`
	SourceFileUploadCntSign  int    `json:"source_file_upload_cnt_sign"`
	CreateAt                 int64  `json:"create_at"`
	UpdateAt                 int64  `json:"update_at"`
}

func GetDaoPreSignSourceFileUploadCntSign(offlineDealId int64) (*int, error) {
	var daoPreSigns []*DaoPreSign
	sql := "select count(*) source_file_upload_cnt_sign from (\n" +
		"select b.source_file_upload_id,count(*) from dao_signature a, dao_signature_source_file_upload b\n" +
		"where a.id=b.dao_signature_id and a.offline_deal_id=?\n" +
		"group by b.source_file_upload_id) a"
	err := database.GetDB().Raw(sql, offlineDealId).Scan(&daoPreSigns).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var sourceFileUploadCntSign int = 0
	if len(daoPreSigns) > 0 {
		sourceFileUploadCntSign = daoPreSigns[0].SourceFileUploadCntSign
	}

	return &sourceFileUploadCntSign, nil
}

func GetDaoPreSignSourceFileUploadCntTotal(offlineDealId int64) (*int, error) {
	var daoPreSigns []*DaoPreSign
	sql := "select count(*) source_file_upload_cnt_total from offline_deal a, car_file_source b where a.car_file_id=b.car_file_id and a.id=?"
	err := database.GetDB().Raw(sql, offlineDealId).Scan(&daoPreSigns).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var sourceFileUploadCntTotal int = 0
	if len(daoPreSigns) > 0 {
		sourceFileUploadCntTotal = daoPreSigns[0].SourceFileUploadCntTotal
	}

	return &sourceFileUploadCntTotal, nil
}

func UpdateDaoPreSignSourceFileUploadCntSign(offlineDealId int64) error {
	sourceFileUploadCntSign, err := GetDaoPreSignSourceFileUploadCntSign(offlineDealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["source_file_upload_cnt_sign"] = *sourceFileUploadCntSign
	fields2BeUpdated["update_at"] = libutils.GetCurrentUtcSecond

	err = database.GetDB().Model(DaoPreSign{}).Where("offline_deal_id=?", offlineDealId).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func CreateDaoPreSign(offlineDealId int64, txHash string, batchNumber int) error {
	sourceFileUploadCntTotal, err := GetDaoPreSignSourceFileUploadCntTotal(offlineDealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	batchSizeMax := float64(*sourceFileUploadCntTotal) / float64(batchNumber)

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	daoPreSign := DaoPreSign{
		OfflineDealId:            offlineDealId,
		TxHash:                   txHash,
		BatchNumber:              batchNumber,
		BatchSizeMax:             int(math.Ceil(batchSizeMax)),
		SourceFileUploadCntTotal: *sourceFileUploadCntTotal,
		SourceFileUploadCntSign:  0,
		CreateAt:                 currentUtcSecond,
		UpdateAt:                 currentUtcSecond,
	}

	err = database.SaveOne(daoPreSign)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
