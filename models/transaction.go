package models

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"
	"sort"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
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

type Billing struct {
	PayId        int64  `json:"pay_id"`
	PayTxHash    string `json:"pay_tx_hash"`
	PayAmount    string `json:"pay_amount"`
	UnlockAmount string `json:"unlock_amount"`
	FileName     string `json:"file_name"`
	PayloadCid   string `json:"payload_cid"`
	PayAt        int64  `json:"pay_at"`
	UnlockAt     int64  `json:"unlock_at"`
	Deadline     int64  `json:"deadline"`
	NetworkName  string `json:"network_name"`
}

type BillingByPayAt []*Billing

func (a BillingByPayAt) Len() int           { return len(a) }
func (a BillingByPayAt) Less(i, j int) bool { return a[i].PayAt < a[j].PayAt }
func (a BillingByPayAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetTransactions(walletId int64, txHash, fileName, orderBy string, isAscend bool, limit, offset int) ([]*Billing, *int, error) {
	sql := "select\n" +
		"a.id pay_id,a.tx_hash pay_tx_hash,a.amount pay_amount,e.amount unlock_amount,b.file_name,d.payload_cid,\n" +
		"a.create_at pay_at,e.create_at unlock_at,a.deadline,f.name network_name\n" +
		"from transaction_pay a\n" +
		"left join source_file_upload b on a.source_file_upload_id=b.id\n" +
		"left outer join car_file_source c on c.source_file_upload_id=a.source_file_upload_id\n" +
		"left outer join car_file d on c.car_file_id=d.id\n" +
		"left outer join transaction_unlock e on e.source_file_upload_id=a.source_file_upload_id\n" +
		"left join network f on a.network_id=f.id\n" +
		"where a.wallet_id_from=?"

	params := []interface{}{}
	params = append(params, walletId)

	if !libutils.IsStrEmpty(&txHash) {
		sql = sql + " and a.tx_hash =?"
		params = append(params, txHash)
	}

	if !libutils.IsStrEmpty(&fileName) {
		sql = sql + " and b.file_name like '%" + fileName + "%') "
	}

	var billings []*Billing
	err := database.GetDB().Raw(sql, params...).Scan(&billings).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err

	}

	sort.Sort(BillingByPayAt(billings))

	totalRecordCount := len(billings)
	start := (offset - 1) * limit
	end := start + limit
	result := billings[start:end]

	return result, &totalRecordCount, nil
}
