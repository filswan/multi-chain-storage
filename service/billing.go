package service

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/shopspring/decimal"
)

func WriteLockPayment(sourceFileUploadId int64, txHash string) error {
	err := models.CreateTransaction4Pay(sourceFileUploadId, txHash)
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

type SourceFileUploadInfoWcid struct {
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	WCID               string `json:"w_cid"`
}
type SourceFileUploadInfo struct {
	WCid              string                      `json:"w_cid"`
	LockedFee         string                      `json:"locked_fee"`
	TxHash            string                      `json:"tx_hash"`
	TokenAddress      string                      `json:"token_address"`
	WCidInSameCarFile []*SourceFileUploadInfoWcid `json:"w_cid_in_same_car_file"`
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

	transactionPay, err := models.GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	token, err := models.GetTokenById(transactionPay.TokenId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFileUploadInfo := &SourceFileUploadInfo{
		WCid:              sourceFileUpload.Uuid + sourceFile.PayloadCid,
		LockedFee:         transactionPay.AmountLock,
		TxHash:            transactionPay.TxHashPay,
		TokenAddress:      token.Address,
		WCidInSameCarFile: []*SourceFileUploadInfoWcid{},
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
		sourceFileUploadInfoWcid := &SourceFileUploadInfoWcid{
			SourceFileUploadId: sourceFileUpload1.Id,
			WCID:               sourceFileUpload1.Uuid + sourceFileUpload1.PayloadCid,
		}
		sourceFileUploadInfo.WCidInSameCarFile = append(sourceFileUploadInfo.WCidInSameCarFile, sourceFileUploadInfoWcid)
	}

	return sourceFileUploadInfo, nil
}

func WriteRefundAfterExpired(sourceFileUploadId int64, refundTxHash string, refundAmount decimal.Decimal) error {
	err := models.UpdateTransactionRefundAfterExpired(sourceFileUploadId, refundTxHash, refundAmount)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
