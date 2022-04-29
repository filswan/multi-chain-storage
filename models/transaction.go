package models

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"

	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"
	"sort"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/shopspring/decimal"

	"github.com/filswan/go-swan-lib/logs"
)

type Transaction struct {
	ID                       int64  `json:"id"`
	SourceFileUploadId       int64  `json:"source_file_upload_id"`
	Status                   string `json:"status"`
	NetworkId                int64  `json:"network_id"`
	TokenId                  int64  `json:"token_id"`
	WalletIdPay              int64  `json:"wallet_id_pay"`
	WalletIdRecipient        int64  `json:"wallet_id_recipient"`
	WalletIdContract         int64  `json:"wallet_id_contract"`
	TxHashPay                string `json:"tx_hash_pay"`
	TxHashRefundAfterExpired string `json:"tx_hash_refund_after_expired"`
	TxHashRefundAfterUnlock  string `json:"tx_hash_refund_after_unlock"`
	AmountLock               string `json:"amount_lock"`
	AmountUnlock             string `json:"amount_unlock"`
	AmountRefundAfterExpired string `json:"amount_refund_after_expired"`
	AmountRefundAfterUnlock  string `json:"amount_refund_after_unlock"`
	Deadline                 int64  `json:"deadline"`
	PayAt                    int64  `json:"pay_at"`
	LastUnlockAt             int64  `json:"last_unlock_at"`
	RefundAfterExpiredAt     int64  `json:"refund_after_expired_at"`
	RefundAfterUnlockAt      int64  `json:"refund_after_unlock_at"`
	CreateAt                 int64  `json:"create_at"`
	UpdateAt                 int64  `json:"update_at"`
}

func GetTransactionBySourceFileUploadId(sourceFileUploadId int64) (*Transaction, error) {
	var transactions []*Transaction
	err := database.GetDB().Where("source_file_upload_id=?", sourceFileUploadId).Find(&transactions).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(transactions) > 0 {
		return transactions[0], nil
	}

	return nil, nil
}

func CreateTransaction4Pay(sourceFileUploadId int64, txHash string) error {
	transactionOld, err := GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if transactionOld != nil {
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

	walletIdPay, err := GetWalletByAddress(lockPayment.AddressFrom, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	walletIdRecipient, err := GetWalletByAddress(lockPayment.AddressTo, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	walletIdContract, err := GetWalletByAddress(config.GetConfig().Polygon.PaymentContractAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	coin, err := GetTokenByName(constants.TOKEN_USDC_NAME)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	network, err := GetNetworkByName(constants.NETWORK_NAME_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	transaction := Transaction{
		SourceFileUploadId: sourceFileUploadId,
		Status:             constants.TRANSACTION_STATUS_PAID,
		NetworkId:          network.ID,
		TokenId:            coin.ID,
		WalletIdPay:        walletIdPay.ID,
		WalletIdRecipient:  walletIdRecipient.ID,
		WalletIdContract:   walletIdContract.ID,
		TxHashPay:          txHash,
		AmountLock:         lockPayment.LockedFee.String(),
		Deadline:           lockPayment.Deadline,
		PayAt:              currentUtcSecond,
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

func UpdateTransactionUnlockInfo(sourceFileUploadId int64, unlockAmount decimal.Decimal) error {
	transaction, err := GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if transaction != nil {
		err := fmt.Errorf("transaction not exists for source file upload:%d", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil
	}

	unlockAmountBefore, err := decimal.NewFromString(transaction.AmountUnlock)
	if err != nil {
		logs.GetLogger().Error(err)
		unlockAmountBefore = decimal.NewFromInt(0)
	}

	unlockAmountNew := unlockAmountBefore.Add(unlockAmount)

	sql := "update transaction set status=?,amount_unlock=?,last_unlock_at=?,update_at=? where id=?"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	params := []interface{}{}
	params = append(params, constants.TRANSACTION_STATUS_UNLOCKING)
	params = append(params, unlockAmountNew.String())
	params = append(params, currentUtcSecond)
	params = append(params, currentUtcSecond)
	params = append(params, transaction.ID)

	err = database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func UpdateTransactionRefundAfterExpired(sourceFileUploadId int64, refundTxHash string, refundAmount decimal.Decimal) error {
	transaction, err := GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if transaction != nil {
		err := fmt.Errorf("transaction not exists for source file upload:%d", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil
	}

	sql := "update transaction set status=?,tx_hash_refund_after_expired=?,amount_refund_after_expired=?,refund_after_expired_at=?,update_at=? where id=?"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	params := []interface{}{}
	params = append(params, constants.TRANSACTION_STATUS_REFUNDED_AFTER_EXPIRED)
	params = append(params, refundTxHash)
	params = append(params, refundAmount.String())
	params = append(params, currentUtcSecond)
	params = append(params, currentUtcSecond)
	params = append(params, transaction.ID)

	err = database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func UpdateTransactionRefundAfterUnlock(sourceFileUploadId int64, refundTxHash string, refundAmount decimal.Decimal) error {
	transaction, err := GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if transaction != nil {
		err := fmt.Errorf("transaction not exists for source file upload:%d", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil
	}

	sql := "update transaction set status=?,tx_hash_refund_after_unlock=?,amount_refund_after_unlock=?,refund_after_unlock_at=?,update_at=? where id=?"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	params := []interface{}{}
	params = append(params, constants.TRANSACTION_STATUS_REFUNDED_AFTER_UNLOCK)
	params = append(params, refundTxHash)
	params = append(params, refundAmount.String())
	params = append(params, currentUtcSecond)
	params = append(params, currentUtcSecond)
	params = append(params, transaction.ID)

	err = database.GetDB().Exec(sql, params...).Error
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
	TokenName    string `json:"token_name"`
}

type BillingByPayAt []*Billing

func (a BillingByPayAt) Len() int           { return len(a) }
func (a BillingByPayAt) Less(i, j int) bool { return a[i].PayAt < a[j].PayAt }
func (a BillingByPayAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BillingByPayAmount []*Billing

func (a BillingByPayAmount) Len() int           { return len(a) }
func (a BillingByPayAmount) Less(i, j int) bool { return a[i].PayAmount < a[j].PayAmount }
func (a BillingByPayAmount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BillingByUnlockAmount []*Billing

func (a BillingByUnlockAmount) Len() int           { return len(a) }
func (a BillingByUnlockAmount) Less(i, j int) bool { return a[i].UnlockAmount < a[j].UnlockAmount }
func (a BillingByUnlockAmount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BillingByFileName []*Billing

func (a BillingByFileName) Len() int           { return len(a) }
func (a BillingByFileName) Less(i, j int) bool { return a[i].FileName < a[j].FileName }
func (a BillingByFileName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BillingByUnlockAt []*Billing

func (a BillingByUnlockAt) Len() int           { return len(a) }
func (a BillingByUnlockAt) Less(i, j int) bool { return a[i].UnlockAt < a[j].UnlockAt }
func (a BillingByUnlockAt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type BillingByDeadline []*Billing

func (a BillingByDeadline) Len() int           { return len(a) }
func (a BillingByDeadline) Less(i, j int) bool { return a[i].Deadline < a[j].Deadline }
func (a BillingByDeadline) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetTransactions(walletId int64, txHash, fileName, orderBy string, isAscend bool, limit, offset int) ([]*Billing, *int, error) {
	sql := "select\n" +
		"a.id pay_id,a.tx_hash pay_tx_hash,a.amount pay_amount,e.amount unlock_amount,b.file_name,d.payload_cid,\n" +
		"a.create_at pay_at,e.create_at unlock_at,a.deadline,f.name network_name,g.name token_name\n" +
		"from transaction_pay a\n" +
		"left join source_file_upload b on a.source_file_upload_id=b.id\n" +
		"left outer join car_file_source c on c.source_file_upload_id=a.source_file_upload_id\n" +
		"left outer join car_file d on c.car_file_id=d.id\n" +
		"left outer join transaction_unlock e on e.source_file_upload_id=a.source_file_upload_id\n" +
		"left join network f on a.network_id=f.id\n" +
		"left join token g on a.token_id=g.id\n" +
		"where a.wallet_id_from=?"

	params := []interface{}{}
	params = append(params, walletId)

	if !libutils.IsStrEmpty(&txHash) {
		sql = sql + " and a.tx_hash =?"
		params = append(params, txHash)
	}

	if !libutils.IsStrEmpty(&fileName) {
		sql = sql + " and b.file_name like '%" + fileName + "%' "
	}

	var billings []*Billing
	err := database.GetDB().Raw(sql, params...).Scan(&billings).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err

	}

	switch strings.Trim(orderBy, " ") {
	case "pay_amount":
		if isAscend {
			sort.Sort(BillingByPayAmount(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByPayAmount(billings)))
		}
	case "unlock_amount":
		if isAscend {
			sort.Sort(BillingByUnlockAmount(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByUnlockAmount(billings)))
		}
	case "file_name":
		if isAscend {
			sort.Sort(BillingByFileName(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByFileName(billings)))
		}
	case "pay_at":
		if isAscend {
			sort.Sort(BillingByPayAt(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByPayAt(billings)))
		}
	case "unlock_at":
		if isAscend {
			sort.Sort(BillingByUnlockAt(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByUnlockAt(billings)))
		}
	case "deadline":
		if isAscend {
			sort.Sort(BillingByDeadline(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByDeadline(billings)))
		}
	default:
		if isAscend {
			sort.Sort(BillingByPayAt(billings))
		} else {
			sort.Sort(sort.Reverse(BillingByPayAt(billings)))
		}
	}

	totalRecordCount := len(billings)
	start := (offset - 1) * limit
	end := start + limit
	if start >= totalRecordCount {
		return nil, &totalRecordCount, nil
	}

	if end >= totalRecordCount {
		end = totalRecordCount
	}

	result := billings[start:end]
	return result, &totalRecordCount, nil
}
