package service

import (
	"context"
	"fmt"
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
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

func WriteRefundAfterExpired(txHash string) (*models.EventExpirePayment, error) {
	ethClient, rpcClient, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var rpcTransaction *models.RpcTransaction
	err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	transactionReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	event := new(models.EventExpirePayment)
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
		return nil, err
	}
	event.BlockNo = strconv.FormatUint(blockNumberInt64, 10)
	wfilCoinId, err := models.GetTokenByName(constants.TOKEN_USDC_NAME)
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

	eventList, err := models.FindEventExpirePayments(&models.EventExpirePayment{TxHash: txHash, BlockNo: strconv.
		FormatUint(blockNumberInt64, 10)}, "id desc", "10", "0")
	if err != nil {
		logs.GetLogger().Error(err)
	}
	if len(eventList) <= 0 {
		err = database.SaveOneWithTransaction(event)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}
	return event, nil
}
