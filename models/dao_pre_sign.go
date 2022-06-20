package models

import (
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
