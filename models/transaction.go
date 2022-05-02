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
	NetworkId                int64  `json:"network_id"`
	TokenId                  int64  `json:"token_id"`
	WalletIdPay              int64  `json:"wallet_id_pay"`
	WalletIdRecipient        int64  `json:"wallet_id_recipient"`
	WalletIdContract         int64  `json:"wallet_id_contract"`
	PayTxHash                string `json:"pay_tx_hash"`
	PayAmount                string `json:"pay_amount"`
	PayAt                    int64  `json:"pay_at"`
	Deadline                 int64  `json:"deadline"`
	UnlockAmount             string `json:"unlock_amount"`
	LastUnlockAt             int64  `json:"last_unlock_at"`
	RefundAfterUnlockTxHash  string `json:"refund_after_unlock_tx_hash"`
	RefundAfterUnlockAmount  string `json:"refund_after_unlock_amount"`
	RefundAfterUnlockAt      int64  `json:"refund_after_unlock_at"`
	RefundAfterExpiredTxHash string `json:"refund_after_expired_tx_hash"`
	RefundAfterExpiredAmount string `json:"refund_after_expired_amount"`
	RefundAfterExpiredAt     int64  `json:"refund_after_expired_at"`
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
		NetworkId:          network.ID,
		TokenId:            coin.ID,
		WalletIdPay:        walletIdPay.ID,
		WalletIdRecipient:  walletIdRecipient.ID,
		WalletIdContract:   walletIdContract.ID,
		PayTxHash:          txHash,
		PayAmount:          lockPayment.LockedFee.String(),
		PayAt:              currentUtcSecond,
		Deadline:           lockPayment.Deadline,
		CreateAt:           currentUtcSecond,
		UpdateAt:           currentUtcSecond,
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

	unlockAmountBefore, err := decimal.NewFromString(transaction.UnlockAmount)
	if err != nil {
		logs.GetLogger().Error(err)
		unlockAmountBefore = decimal.NewFromInt(0)
	}

	unlockAmountNew := unlockAmountBefore.Add(unlockAmount)

	sql := "update transaction set unlock_amount=?,last_unlock_at=?,update_at=? where id=?"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	params := []interface{}{}
	params = append(params, unlockAmountNew.String())
	params = append(params, currentUtcSecond)
	params = append(params, currentUtcSecond)
	params = append(params, transaction.ID)

	err = database.GetDB().Exec(sql, params...).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	err = UpdateSourceFileUploadStatus(sourceFileUploadId, constants.SOURCE_FILE_UPLOAD_STATUS_UNLOCKING)
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

	sql := "update transaction set refund_after_unlock_tx_hash=?,refund_after_unlock_amount=?,refund_after_unlock_at=?,update_at=? where id=?"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	params := []interface{}{}
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

	err = UpdateSourceFileUploadStatus(sourceFileUploadId, constants.SOURCE_FILE_UPLOAD_STATUS_UNLOCKED)
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

	sql := "update transaction set refund_after_expired_tx_hash=?,refund_after_expired_amount=?,refund_after_expired_at=?,update_at=? where id=?"

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	params := []interface{}{}
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

	err = UpdateSourceFileUploadStatus(sourceFileUploadId, constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED)
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
		"a.id pay_id,a.tx_hash_pay,a.pay_amount,a.unlock_amount,b.file_name,d.payload_cid,\n" +
		"a.pay_at,a.last_unlock_at unlock_at,a.deadline,e.name network_name,f.name token_name\n" +
		"from transaction a\n" +
		"left join source_file_upload b on a.source_file_upload_id=b.id\n" +
		"left outer join car_file_source c on c.source_file_upload_id=a.source_file_upload_id\n" +
		"left outer join car_file d on c.car_file_id=d.id\n" +
		"left join network e on a.network_id=e.id\n" +
		"left join token f on a.token_id=f.id\n" +
		"where a.wallet_id_pay=?"

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
