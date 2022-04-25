package service

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"

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

type SourceFileUploadInfo struct {
	WCid              string   `json:"w_cid"`
	LockedFee         string   `json:"locked_fee"`
	TxHash            string   `json:"tx_hash"`
	TokenAddress      string   `json:"token_address"`
	WCidInSameCarFile []string `json:"w_cid_in_same_car_file"`
}

func GetSourceFileUploadInfo(sourceFileUploadId int64) (*SourceFileUploadInfo, error) {
	sourceFileUpload, err := models.GetSourceFileUploadById(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if sourceFileUpload == nil {
		err := fmt.Errorf("source file upload:%d not exists", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFile, err := models.GetSourceFileById(sourceFileUpload.SourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	transactionPay, err := models.GetTransactionBySourceFileUploadIdType(sourceFileUploadId, constants.TRANSACTION_TYPE_PAY)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	token, err := models.GetCoinById(transactionPay.CoinId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFileUploadInfo := &SourceFileUploadInfo{
		WCid:              sourceFileUpload.Uuid + sourceFile.PayloadCid,
		LockedFee:         transactionPay.Amount,
		TxHash:            transactionPay.TxHash,
		TokenAddress:      token.Address,
		WCidInSameCarFile: []string{},
	}

	carFileSource, err := models.GetCarFileSourceBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if carFileSource == nil {
		return sourceFileUploadInfo, nil
	}

	sourceFileUploads, err := models.GetSourceFileUploadsByCarFileId(carFileSource.CarFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	for _, sourceFileUpload1 := range sourceFileUploads {
		sourceFileUploadInfo.WCidInSameCarFile = append(sourceFileUploadInfo.WCidInSameCarFile, sourceFileUpload1.Uuid+sourceFileUpload1.PayloadCid)
	}

	return sourceFileUploadInfo, nil
}
