package models

import (
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/on-chain/client"

	"multi-chain-storage/database"
	"sort"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/shopspring/decimal"

	"github.com/filswan/go-swan-lib/logs"
)

type Transaction struct {
	ID                 int64   `json:"id"`
	SourceFileUploadId int64   `json:"source_file_upload_id"`
	NetworkId          int64   `json:"network_id"`
	TokenId            int64   `json:"token_id"`
	WalletIdPay        int64   `json:"wallet_id_pay"`
	WalletIdRecipient  int64   `json:"wallet_id_recipient"`
	WalletIdContract   int64   `json:"wallet_id_contract"`
	PayTxHash          string  `json:"pay_tx_hash"`
	PayAmount          string  `json:"pay_amount"`
	PayAt              int64   `json:"pay_at"`
	Deadline           int64   `json:"deadline"`
	UnlockAmount       *string `json:"unlock_amount"`
	LastUnlockAt       *int64  `json:"last_unlock_at"`
	RefundTxHash       *string `json:"refund_tx_hash"`
	RefundAmount       *string `json:"refund_amount"`
	RefundAt           *int64  `json:"refund_at"`
	CreateAt           int64   `json:"create_at"`
	UpdateAt           int64   `json:"update_at"`
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

func CreateTransaction4PayByWCid(wCid, txHash string, lockTime int64) error {
	sourceFile, sourceFileUpload, err := GetSourceFileUploadByWCid(wCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if sourceFile == nil || sourceFileUpload == nil {
		logs.GetLogger().Info("wCid:", wCid, " not exists")
		return nil
	}

	lockedPayment, err := client.GetLockedPaymentInfo(wCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if lockedPayment == nil {
		msg := fmt.Sprintf("payment not exists for wCid:%s ,tx hash:%s", wCid, txHash)
		logs.GetLogger().Info(msg)
		return nil
	}

	transaction, err := GetTransactionBySourceFileUploadId(sourceFileUpload.Id)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletPay, err := GetWalletByAddress(lockedPayment.AddressFrom, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletRecipient, err := GetWalletByAddress(lockedPayment.AddressTo, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletContract, err := GetWalletByAddress(config.GetConfig().Polygon.PaymentContractAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	token, err := GetTokenByAddress(lockedPayment.TokenAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if token == nil {
		err := fmt.Errorf("token address:%s not exists", lockedPayment.TokenAddress)
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := libutils.GetCurrentUtcSecond()

	if transaction == nil {
		transaction = &Transaction{}
		transaction.CreateAt = currentUtcSecond
	}
	transaction.SourceFileUploadId = sourceFileUpload.Id
	transaction.NetworkId = token.NetworkId
	transaction.TokenId = token.ID
	transaction.WalletIdPay = walletPay.ID
	transaction.WalletIdRecipient = walletRecipient.ID
	transaction.WalletIdContract = walletContract.ID
	transaction.PayTxHash = txHash
	transaction.PayAmount = lockedPayment.LockedFee.String()
	transaction.PayAt = lockedPayment.Deadline - lockTime
	transaction.Deadline = lockedPayment.Deadline
	transaction.UpdateAt = currentUtcSecond

	err = database.SaveOne(&transaction)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["status"] = constants.SOURCE_FILE_UPLOAD_STATUS_PAID
	fields2BeUpdated["update_at"] = currentUtcSecond

	err = database.GetDB().Model(SourceFileUpload{}).Where("id=?", sourceFileUpload.Id).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func UpdateTransactionUnlockInfo(sourceFileUploadId int64, lockedAmount decimal.Decimal, lastUnlockAt int64) error {
	transaction, err := GetTransactionBySourceFileUploadId(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if transaction == nil {
		err := fmt.Errorf("transaction not exists for source file upload:%d", sourceFileUploadId)
		logs.GetLogger().Error(err)
		return nil
	}

	payAmount, err := decimal.NewFromString(transaction.PayAmount)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockAmount := payAmount.Sub(lockedAmount)

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["unlock_amount"] = unlockAmount.String()
	fields2BeUpdated["last_unlock_at"] = lastUnlockAt
	fields2BeUpdated["update_at"] = currentUtcSecond

	err = database.GetDB().Model(Transaction{}).Where("id=?", transaction.ID).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func UpdateTransactionRefundInfo(wCid, refundTxHash string, refundAt int64) error {
	lockedPayment, err := client.GetLockedPaymentInfo(wCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if lockedPayment != nil {
		msg := fmt.Sprintf("payment still exists for w_cid:%s ,tx hash:%s", wCid, refundTxHash)
		logs.GetLogger().Info(msg)
		return nil
	}

	sourceFile, sourceFileUpload, err := GetSourceFileUploadByWCid(wCid)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if sourceFile == nil || sourceFileUpload == nil {
		logs.GetLogger().Info("wCid:", wCid, " not exists")
		return nil
	}

	transaction, err := GetTransactionBySourceFileUploadId(sourceFileUpload.Id)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if transaction == nil {
		msg := fmt.Sprintf("transaction not exists for source file upload:%d", sourceFileUpload.Id)
		logs.GetLogger().Info(msg)
		return nil
	}

	payAmount, err := decimal.NewFromString(transaction.PayAmount)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	unlockAmount := decimal.NewFromInt(0)
	if transaction.UnlockAmount != nil {
		unlockAmount, err = decimal.NewFromString(*transaction.UnlockAmount)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	refundAmount := payAmount.Sub(unlockAmount)

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["refund_tx_hash"] = refundTxHash
	fields2BeUpdated["refund_amount"] = refundAmount
	fields2BeUpdated["refund_at"] = refundAt
	fields2BeUpdated["update_at"] = currentUtcSecond

	err = database.GetDB().Model(Transaction{}).Where("id=?", transaction.ID).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	offlineDealsNotSuccess, err := GetOfflineDealsNotSuccessBySourceFileUploadId(sourceFileUpload.Id)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	sourceFileStatus := constants.SOURCE_FILE_UPLOAD_STATUS_REFUNDED
	if len(offlineDealsNotSuccess) == 0 {
		sourceFileStatus = constants.SOURCE_FILE_UPLOAD_STATUS_SUCCESS
	}

	err = UpdateSourceFileUploadStatus(sourceFileUpload.Id, sourceFileStatus)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	carFile, err := GetCarFileBySourceFileUploadId(sourceFileUpload.Id)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if carFile != nil {
		sourceFileUploads, err := GetSourceFileUploadsNotCompletedByCarFileId(carFile.ID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		if len(sourceFileUploads) == 0 {
			err = UpdateCarFileStatus(carFile.ID, constants.CAR_FILE_STATUS_COMPLETED)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
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
		"a.id pay_id,a.pay_tx_hash,a.pay_amount,a.unlock_amount,b.file_name,d.payload_cid,\n" +
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
		sql = sql + " and a.pay_tx_hash =?"
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
