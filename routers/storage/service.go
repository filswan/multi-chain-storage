package storage

import (
	"context"
	"errors"
	"math/big"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/common/utils"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
)

func GetSourceFileAndDealFileInfoByPayloadCid(payloadCid string) ([]*SourceFileAndDealFileInfo, error) {
	sql := "select h.wallet_address,s.ipfs_url,h.file_name,d.id,d.payload_cid,d.deal_cid,d.deal_id,d.lock_payment_status,s.create_at "
	sql = sql + "from source_file s,source_file_deal_file_map m,deal_file d, source_file_upload_history h "
	sql = sql + "where s.id = m.source_file_id and s.id=h.source_file_id and m.deal_file_id = d.id and d.payload_cid=?"
	var results []*SourceFileAndDealFileInfo
	err := database.GetDB().Raw(sql, payloadCid).Order("create_at desc").Limit(10).Offset(0).Order("create_at desc").Scan(&results).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func GetDealListThanGreaterDealID(dealId int64, offset, limit int) ([]*DaoDealResult, error) {
	whereCondition := "deal_id > " + strconv.FormatInt(dealId, 10)
	var results []*DaoDealResult
	err := database.GetDB().Table("deal_file").Where(whereCondition).Offset(offset).Limit(limit).Order("create_at").Scan(&results).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func GetDaoSignatureInfoByDealId(dealId int64) ([]*DaoSignResult, error) {
	whereCondition := "deal_id = " + strconv.FormatInt(dealId, 10)
	var results []*DaoSignResult
	err := database.GetDB().Table("event_dao_signature").Where(whereCondition).Offset(0).Limit(constants.DEFAULT_SELECT_LIMIT).Order("block_time desc").Scan(&results).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return results, nil
}

func SaveDaoEventFromTxHash(txHash string, payload_cid string, recipent string, deal_id int64, verification bool) error {
	ethClient, rpcClient, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if txHash != "" && strings.HasPrefix(txHash, "0x") {
		var rpcTransaction *models.RpcTransaction
		err = rpcClient.CallContext(context.Background(), &rpcTransaction, "eth_getTransactionByHash", common.HexToHash(txHash))
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		transaction, _, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(txHash))
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		transReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		eventDaoSignature, err := models.GetEventDaoSignatures(&models.EventDaoSignature{PayloadCid: payload_cid, TxHash: txHash})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}

		if eventDaoSignature == nil {
			eventDaoSignature = &models.EventDaoSignature{}
		}

		eventDaoSignature.TxHash = txHash
		eventDaoSignature.Recipient = recipent
		eventDaoSignature.PayloadCid = payload_cid
		wfilCoinId, err := models.GetCoinByName(constants.COIN_NAME_USDC)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			eventDaoSignature.CoinId = wfilCoinId.ID
			eventDaoSignature.NetworkId = wfilCoinId.NetworkId
		}
		eventDaoSignature.DealId = deal_id
		block, err := ethClient.BlockByHash(context.Background(), *rpcTransaction.BlockHash)
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			eventDaoSignature.BlockTime = strconv.FormatUint(block.Time(), 10)
			eventDaoSignature.DaoPassTime = strconv.FormatUint(block.Time(), 10)
		}
		blockNumberStr := strings.Replace(*rpcTransaction.BlockNumber, "0x", "", -1)
		blockNumberInt64, err := strconv.ParseUint(blockNumberStr, 16, 64)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		eventDaoSignature.BlockNo = blockNumberInt64
		if transReceipt.Status == 1 {
			eventDaoSignature.Status = true
		} else {
			eventDaoSignature.Status = false
		}

		if verification {
			eventDaoSignature.SignatureUnlockStatus = constants.SIGNATURE_SUCCESS_VALUE
		} else {
			eventDaoSignature.SignatureUnlockStatus = constants.SIGNATURE_FAILED_VALUE
		}
		addrInfo, err := client.GetFromAndToAddressByTxHash(ethClient, transaction.ChainId(), common.HexToHash(txHash))
		if err != nil {
			logs.GetLogger().Error(err)
		} else {
			eventDaoSignature.DaoAddress = addrInfo.AddrFrom
		}
		err = database.SaveOneWithTransaction(eventDaoSignature)
		if err != nil {
			logs.GetLogger().Error(err)
		}
	}
	return nil
}

func SaveExpirePaymentEvent(txHash string) (*models.EventExpirePayment, error) {

	ethClient, rpcClient, err := client.GetEthClient()

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if txHash != "" && strings.HasPrefix(txHash, "0x") {
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
		wfilCoinId, err := models.GetCoinByName(constants.COIN_NAME_USDC)
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
		event.CreateAt = strconv.FormatInt(utils.GetCurrentUtcMilliSecond(), 10)
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
	return nil, nil
}

func VerifyDaoSigOnContract(tx_hash string) (bool, error) {
	client, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return false, err
	}
	if tx_hash != "" && strings.HasPrefix(tx_hash, "0x") {
		transaction, err := client.TransactionReceipt(context.Background(), common.HexToHash(tx_hash))
		if err != nil {
			logs.GetLogger().Error(err)
			return false, err
		}
		if transaction.Status == 1 {
			return true, nil
		} else {
			return false, nil
		}
	} else {
		err := errors.New("invalid transaction hash:" + tx_hash)
		logs.GetLogger().Error(err)
		return false, err
	}
}
