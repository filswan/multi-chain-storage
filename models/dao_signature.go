package models

import (
	"context"
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/filswan/go-swan-lib/utils"

	"github.com/ethereum/go-ethereum/common"
)

type DaoSignature struct {
	ID                int64  `json:"id"`
	OfflineDealId     int64  `json:"offline_deal_id"`
	NetworkId         int64  `json:"network_id"`
	WalletIdSigner    int64  `json:"wallet_id_signer"`
	WalletIdRecipient int64  `json:"wallet_id_recipient"`
	WalletIdContract  int64  `json:"wallet_id_contract"`
	TxHash            string `json:"tx_hash"`
	Status            string `json:"status"`
	CreateAt          int64  `json:"create_at"`
	UpdateAt          int64  `json:"update_at"`
}

type DaoSignatureOut struct {
	WalletSigner string  `json:"wallet_signer"`
	TxHash       *string `json:"tx_hash"`
	Status       *string `json:"status"`
	CreateAt     *int64  `json:"create_at"`
}

func GetDaoSignaturesByDealId(dealId int64) ([]*DaoSignatureOut, error) {
	sql := "select a.address wallet_signer,b.tx_hash,b.status,b.create_at from wallet a\n" +
		"left outer join (\n" +
		"  select b.wallet_id_signer,b.tx_hash,b.status,b.create_at from offline_deal a,dao_signature b\n" +
		"  where a.deal_id=? and a.id=b.offline_deal_id\n" +
		") b on a.id=b.wallet_id_signer\n" +
		"where a.is_dao=true order by a.id"

	var daoSignatures []*DaoSignatureOut
	err := database.GetDB().Raw(sql, dealId).Scan(&daoSignatures).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return daoSignatures, nil
}

func GetDaoSignaturesByOfflineDealIdTxHash(offlineDealId int64, txHash string) (*DaoSignature, error) {
	var eventDaoSignatures []*DaoSignature
	err := database.GetDB().Where("offline_deal_id=? and tx_hash=?", offlineDealId, txHash).Find(&eventDaoSignatures).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(eventDaoSignatures) >= 1 {
		return eventDaoSignatures[0], nil
	}

	return nil, nil
}

func WriteDaoSignature(txHash string, recipientWalletAddress string, dealId int64) error {
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
		err := fmt.Errorf("failed to get transaction by tx hash: %s, %s", txHash, err.Error())
		logs.GetLogger().Error(err)
		return err
	}
	addrInfo, err := client.GetFromAndToAddressByTxHash(ethClient, transaction.ChainId(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("addrInfo.AddrFrom:", addrInfo.AddrFrom) //this is signer wallet address
	logs.GetLogger().Info("addrInfo.AddrTo:", addrInfo.AddrTo)     //this is dao contract address
	signerWalletAddress := addrInfo.AddrFrom
	daoContractAddress := addrInfo.AddrTo
	daoContractAddressConfig := config.GetConfig().Polygon.DaoContractAddress
	if !strings.EqualFold(daoContractAddress, daoContractAddressConfig) {
		err := fmt.Errorf("dao contract address:%s not match config:%s for tx hash:%s", daoContractAddress, daoContractAddressConfig, txHash)
		logs.GetLogger().Error(err)
		//return err
	}

	network, err := GetNetworkByName(constants.NETWORK_NAME_POLYGON)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if network == nil {
		err := fmt.Errorf("network:%s not exists", constants.NETWORK_NAME_POLYGON)
		logs.GetLogger().Error(err)
		return err
	}

	walletSigner, err := GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if walletSigner.IsDao == nil || !*walletSigner.IsDao {
		err = SetWalletAsDao(walletSigner.ID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	walletContract, err := GetWalletByAddress(daoContractAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletRecipient, err := GetWalletByAddress(recipientWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	transactionReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	offlineDeal, err := GetOfflineDealByDealId(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if offlineDeal == nil {
		err := fmt.Errorf("offline deal with deal id: %d not exists", dealId)
		logs.GetLogger().Error(err)
		return err
	}

	daoSignature, err := GetDaoSignaturesByOfflineDealIdTxHash(offlineDeal.Id, txHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := utils.GetCurrentUtcSecond()
	if daoSignature == nil {
		daoSignature = &DaoSignature{
			CreateAt: currentUtcSecond,
		}
	}

	daoSignature.NetworkId = network.ID
	daoSignature.OfflineDealId = offlineDeal.Id
	if transactionReceipt.Status == 1 {
		daoSignature.Status = constants.DAO_SIGNATURE_STATUS_SUCCESS
	} else {
		daoSignature.Status = constants.DAO_SIGNATURE_STATUS_FAILED
	}
	daoSignature.TxHash = txHash
	daoSignature.WalletIdSigner = walletSigner.ID
	daoSignature.WalletIdRecipient = walletRecipient.ID
	daoSignature.WalletIdContract = walletContract.ID
	daoSignature.UpdateAt = currentUtcSecond

	err = database.SaveOne(daoSignature)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
