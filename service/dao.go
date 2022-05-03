package service

import (
	"context"
	"errors"
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/filswan/go-swan-lib/utils"
)

type DealForDaoSignResult struct {
	DealFileId            int64    `json:"deal_file_id"`
	PayloadCid            string   `json:"payload_cid"`
	DealCid               string   `json:"deal_cid"`
	DealId                int64    `json:"deal_id"`
	PieceCid              string   `json:"piece_cid"`
	MinerFid              string   `json:"miner_fid"`
	Duration              int      `json:"duration"`
	Cost                  string   `json:"cost"`
	CreateAt              string   `json:"create_at"`
	Verified              bool     `json:"verified"`
	ClientWalletAddress   string   `json:"client_wallet_address"`
	SourceFilePayloadCids []string `json:"payload_cids_source"`
}

func GetShoulBeSignDealListFromDB() ([]*DealForDaoSignResult, error) {
	finalSql := "select a.id as deal_file_id, b.deal_id,a.deal_cid,a.piece_cid,a.payload_cid,a.cost,a.verified,a.miner_fid,duration,a.client_wallet_address,a.create_at from deal_file a left join offline_deal b on a.id = b.deal_file_id left join event_lock_payment c on a.payload_cid=c.payload_cid " +
		" where b.deal_id not in  ( " +
		" select  deal_id from dao_fetched_deal ) " +
		" and b.deal_id > 0 and IFNULL(c.deadline,0) < " + strconv.FormatInt(time.Now().Unix(), 10) +
		" order by a.create_at desc"
	var dealForDaoSignResultList []*DealForDaoSignResult
	err := database.GetDB().Raw(finalSql).Scan(&dealForDaoSignResultList).Limit(0).Offset(constants.DEFAULT_SELECT_LIMIT).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}
	sourceSQL := "select a.id, a.payload_cid, b.deal_file_id from source_file a, source_file_deal_file_map b where a.id = b.source_file_id"
	var sourceFileExt []*models.SourceFileExt
	err = database.GetDB().Raw(sourceSQL).Scan(&sourceFileExt).Limit(0).Offset(constants.DEFAULT_SELECT_LIMIT).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}
	for _, deal := range dealForDaoSignResultList {
		var cids []string
		for _, file := range sourceFileExt {
			if deal.DealFileId == file.DealFileId && file.PayloadCid != "" {
				cids = append(cids, file.PayloadCid)
			}
		}
		if len(cids) > 0 {
			deal.SourceFilePayloadCids = cids
		}
	}
	return dealForDaoSignResultList, nil
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

type RpcTransaction struct {
	tx *types.Transaction
	txExtraInfo
}

type txExtraInfo struct {
	BlockNumber *string         `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash    `json:"blockHash,omitempty"`
	From        *common.Address `json:"from,omitempty"`
}

func SaveDaoEventFromTxHash(txHash string, recipent string, dealId int64, verification bool) error {
	ethClient, _, err := client.GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	if txHash == "" || !strings.HasPrefix(txHash, "0x") {
		err := fmt.Errorf("invalid tx hash:%s", txHash)
		logs.GetLogger().Error(err)
		return err
	}

	transaction, _, err := ethClient.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	addrInfo, err := client.GetFromAndToAddressByTxHash(ethClient, transaction.ChainId(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("addrInfo.AddrFrom:", addrInfo.AddrFrom)
	logs.GetLogger().Info("addrInfo.AddrTo:", addrInfo.AddrTo)

	network, err := models.GetNetworkByName(constants.NETWORK_NAME_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if network == nil {
		err := fmt.Errorf("network:%s not exists", constants.NETWORK_NAME_POLYGON)
		logs.GetLogger().Error(err)
		return err
	}

	walletRecipient, err := models.GetWalletByAddress(recipent, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletDao, err := models.GetWalletByAddress(config.GetConfig().Polygon.DaoContractAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	transReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	daoSignature, err := models.GetDaoSignaturesByDealIdTxHash(dealId, txHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := utils.GetCurrentUtcSecond()
	if daoSignature == nil {
		daoSignature = &models.DaoSignature{
			CreateAt: currentUtcSecond,
		}
	}

	daoSignature.NetworkId = network.ID
	daoSignature.DealId = dealId
	if transReceipt.Status == 1 {
		daoSignature.Status = constants.DAO_SIGNATURE_STATUS_SUCCESS
	} else {
		daoSignature.Status = constants.DAO_SIGNATURE_STATUS_FAIL
	}
	daoSignature.TxHash = txHash
	daoSignature.RecipientId = walletRecipient.ID
	daoSignature.DaoAddressId = walletDao.ID
	daoSignature.UpdateAt = currentUtcSecond

	err = database.SaveOne(daoSignature)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
