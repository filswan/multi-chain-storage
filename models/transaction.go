package models

import (
	"context"
	"math/big"
	"multi-chain-storage/common/constants"
	"strconv"

	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"
	"sort"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	ID                 int64  `json:"id"`
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	Type               int    `json:"type"`
	NetworkId          int64  `json:"network_id"`
	TokenId            int64  `json:"token_id"`
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

func CreateTransaction4Pay(sourceFileUploadId int64, txHash string) error {
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
		Type:               constants.TRANSACTION_TYPE_PAY,
		NetworkId:          network.ID,
		TokenId:            coin.ID,
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

func CreateTransaction4RefundAfterExpired(txHash string) error {
	ethClient, rpcClient, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	var rpcTransaction *RpcTransaction
	err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	transaction, _, _ := ethClient.TransactionByHash(context.Background(), common.HexToHash(txHash))
	transactionReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("rpcTransaction.From:", rpcTransaction.From)
	logs.GetLogger().Info("rpcTransaction.BlockNumber:", rpcTransaction.BlockNumber)
	logs.GetLogger().Info("transaction.To():", transaction.To())
	logs.GetLogger().Info("transaction.Value():", transaction.Value())
	logs.GetLogger().Info("transactionReceipt.ContractAddress:", transactionReceipt.ContractAddress)

	event := new(EventExpirePayment)
	event.TxHash = txHash

	block, err := ethClient.BlockByHash(context.Background(), *rpcTransaction.BlockHash)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		event.BlockTime = strconv.FormatUint(block.Time(), 10)
	}
	blockNumberStr := strings.Replace(*rpcTransaction.BlockNumber, "0x", "", -1)
	blockNumberInt64, err := strconv.ParseUint(blockNumberStr, 16, 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	event.BlockNo = strconv.FormatUint(blockNumberInt64, 10)
	wfilCoinId, err := GetTokenByName(constants.TOKEN_USDC_NAME)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		event.CoinId = wfilCoinId.ID
		event.NetworkId = wfilCoinId.NetworkId
	}

	contrackABI, err := client.GetContractAbi()

	if err != nil {
		logs.GetLogger().Error(err)
	}

	for _, v := range transactionReceipt.Logs {
		if v.Topics[0].Hex() == "0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a" {
			dataList, err := contrackABI.Unpack("ExpirePayment", v.Data)
			if err != nil {
				logs.GetLogger().Error(err)
			}
			event.PayloadCid = dataList[0].(string)
			event.TokenAddress = dataList[1].(common.Address).Hex()
			event.ExpireUserAmount = dataList[2].(*big.Int).String()
			event.UserAddress = dataList[3].(common.Address).Hex()
		}
	}
	event.CreateAt = strconv.FormatInt(libutils.GetCurrentUtcSecond(), 10)
	event.ContractAddress = transactionReceipt.ContractAddress.Hex()

	return nil
}
