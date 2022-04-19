package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"

	"github.com/filswan/go-swan-lib/logs"
)

type Transaction struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	Type               int    `json:"type"`
	TxHash             string `json:"tx_hash"`
	WalletIdFrom       int64  `json:"wallet_id_from"`
	WalletIdTo         int64  `json:"wallet_id_to"`
	CoinId             int64  `json:"coin_id"`
	Amount             string `json:"amount"`
	BlockNumber        int64  `json:"block_number"`
	TransactionAt      int64  `json:"transaction_at"`
	CreateAt           int64  `json:"create_at"`
}

func GetTransactionBySourceFileUploadIdType(sourceFileUploadId int64, transactionType int) ([]*Transaction, error) {
	var transactions []*Transaction
	err := database.GetDB().Where("source_file_upload_id=? and type=?", sourceFileUploadId, transactionType).Find(&transactions).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return transactions, nil
}

func CreateTransaction(sourceFileUploadId int64, txHash string) error {
	transactions, err := GetTransactionBySourceFileUploadIdType(sourceFileUploadId, constants.TRANSACTION_TYPE_PAY)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if len(transactions) > 0 {
		return nil
	}

	sourceFileUpload, err := GetSourceFileUploadById(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	sourceFile, err := GetSourceFileById(sourceFileUpload.SourceFileId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	wCid := sourceFileUpload.Uuid + sourceFile.PayloadCid

	lockPayment, err := client.GetLockedPaymentInfo(wCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletFrom, err := GetWalletByAddress(lockPayment.AddressFrom, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	walletTo, err := GetWalletByAddress(lockPayment.AddressTo, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	coin, err := GetCoinByName(constants.COIN_USDC_NAME)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := utils.GetCurrentUtcSecond()
	transaction := Transaction{
		SourceFileUploadId: sourceFileUploadId,
		Type:               constants.TRANSACTION_TYPE_PAY,
		TxHash:             txHash,
		WalletIdFrom:       walletFrom.ID,
		WalletIdTo:         walletTo.ID,
		CoinId:             coin.ID,
		Amount:             lockPayment.LockedFee.String(),
		BlockNumber:        lockPayment.BlockNumber,
		TransactionAt:      currentUtcSecond,
		CreateAt:           currentUtcSecond,
	}

	db := database.GetDBTransaction()
	err = database.SaveOneInTransaction(db, &transaction)
	if err != nil {
		db.Rollback()
		logs.GetLogger().Error(err)
		return err
	}

	sql := "update source_file_upload set status=?,update_at=? where id=?"

	params := []interface{}{}
	params = append(params, constants.SOURCE_FILE_UPLOAD_STATUS_PAID)
	params = append(params, currentUtcSecond)
	params = append(params, sourceFileUploadId)

	err = db.Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		db.Rollback()
		return err
	}

	err = db.Commit().Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
