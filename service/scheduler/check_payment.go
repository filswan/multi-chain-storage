package scheduler

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
)

func CheckPayment() error {
	srcFileUploads, err := models.GetSourceFileUploadsByFileTypeStatus(constants.SOURCE_FILE_TYPE_NORMAL, constants.SOURCE_FILE_UPLOAD_STATUS_PENDING)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	for _, srcFileUpload := range srcFileUploads {
		err := models.CreateTransaction4Pay(srcFileUpload.Id, "")
		if err != nil {
			if strings.Contains(err.Error(), "payment not exists") {
				logs.GetLogger().Info(err)
			} else {
				logs.GetLogger().Error(err)
			}

			continue
		}
	}

	return nil
}
