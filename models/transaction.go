package models

import (
	"multi-chain-storage/common/constants"

	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"
	"sort"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"
)

type Transaction struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	Type               int    `json:"type"`
	NetworkId          int64  `json:"network_id"`
	CoinId             int64  `json:"coin_id"`
	TxHash             string `json:"tx_hash"`
	WalletIdFrom       int64  `json:"wallet_id_from"`
	WalletIdTo         int64  `json:"wallet_id_to"`
	Amount             string `json:"amount"`
	BlockNumber        int64  `json:"block_number"`
	Deadline           int64  `json:"deadline"`
	CreateAt           int64  `json:"create_at"`
}

func GetTransactionBySourceFileUploadIdType(sourceFileUploadId int64, transactionType int) (*Transaction, error) {
	var transactions []*Transaction
	err := database.GetDB().Where("source_file_upload_id=? and type=?", sourceFileUploadId, transactionType).Find(&transactions).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(transactions) > 0 {
		return transactions[0], nil
	}

	return nil, nil
}

func CreateTransaction(sourceFileUploadId int64, txHash string) error {
	transactionOld, err := GetTransactionBySourceFileUploadIdType(sourceFileUploadId, constants.TRANSACTION_TYPE_PAY)
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

	network, err := GetNetworkByName(constants.NETWORK_NAME_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	transaction := Transaction{
		SourceFileUploadId: sourceFileUploadId,
		Type:               constants.TRANSACTION_TYPE_PAY,
		NetworkId:          network.ID,
		CoinId:             coin.ID,
		TxHash:             txHash,
		WalletIdFrom:       walletFrom.ID,
		WalletIdTo:         walletTo.ID,
		Amount:             lockPayment.LockedFee.String(),
		BlockNumber:        lockPayment.BlockNumber,
		Deadline:           lockPayment.Deadline,
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
