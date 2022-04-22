package service

import (
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
