package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type DaoSignatureSourceFileUpload struct {
	Id                 *int64 `json:"id"`
	DaoSignatureId     int64  `json:"dao_signature_id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	CreateAt           int64  `json:"create_at"`
	UpdateAt           int64  `json:"update_at"`
}

func GetDaoSignatureSourceFileUpload(daoSignatureId int64, sourceFileUploadId int64) (*DaoSignatureSourceFileUpload, error) {
	var daoSignatureSourceFileUploads []*DaoSignatureSourceFileUpload
	err := database.GetDB().Where("dao_signature_id=? and source_file_upload_id=?", daoSignatureId, sourceFileUploadId).Find(&daoSignatureSourceFileUploads).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(daoSignatureSourceFileUploads) > 0 {
		return daoSignatureSourceFileUploads[0], nil
	}

	return nil, nil
}

func SaveDaoSignatureSourceFileUpload(daoSignatureId int64, sourceFileUploadId int64) error {
	daoSignatureSourceFileUpload, err := GetDaoSignatureSourceFileUpload(daoSignatureId, sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if daoSignatureSourceFileUpload != nil {
		return nil
	}

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	daoSignatureSourceFileUpload = &DaoSignatureSourceFileUpload{
		DaoSignatureId:     daoSignatureId,
		SourceFileUploadId: sourceFileUploadId,
		CreateAt:           currentUtcSecond,
		UpdateAt:           currentUtcSecond,
	}

	err = database.SaveOne(daoSignatureSourceFileUpload)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
