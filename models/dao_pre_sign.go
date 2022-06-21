package models

import (
	"context"
	"fmt"
	"math"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/on-chain/client"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type DaoPreSign struct {
	Id                       int64  `json:"id"`
	OfflineDealId            int64  `json:"offline_deal_id"`
	BatchCount               int    `json:"batch_count"`
	BatchSizeMax             int    `json:"batch_size_max"`
	SourceFileUploadCntTotal int    `json:"source_file_upload_cnt_total"`
	SourceFileUploadCntSign  int    `json:"source_file_upload_cnt_sign"`
	NetworkId                int64  `json:"network_id"`
	WalletIdSigner           int64  `json:"wallet_id_signer"`
	WalletIdRecipient        int64  `json:"wallet_id_recipient"`
	WalletIdContract         int64  `json:"wallet_id_contract"`
	TxHash                   string `json:"tx_hash"`
	Status                   string `json:"status"`
	CreateAt                 int64  `json:"create_at"`
	UpdateAt                 int64  `json:"update_at"`
}

func GetDaoPreSignByOfflineDealId(offlineDealId int64) (*DaoPreSign, error) {
	var daoPreSigns []*DaoPreSign
	err := database.GetDB().Where("offline_deal_id=?", offlineDealId).Find(&daoPreSigns).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(daoPreSigns) >= 1 {
		return daoPreSigns[0], nil
	}

	return nil, nil
}

func GetDaoPreSignSourceFileUploadCntSign(offlineDealId int64) (*int, error) {
	var daoPreSigns []*DaoPreSign
	sql := "select count(*) source_file_upload_cnt_sign from (\n" +
		"select b.source_file_upload_id,count(*) from dao_signature a, dao_signature_source_file_upload b\n" +
		"where a.id=b.dao_signature_id and a.offline_deal_id=?\n" +
		"group by b.source_file_upload_id) a"
	err := database.GetDB().Raw(sql, offlineDealId).Scan(&daoPreSigns).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var sourceFileUploadCntSign int = 0
	if len(daoPreSigns) > 0 {
		sourceFileUploadCntSign = daoPreSigns[0].SourceFileUploadCntSign
	}

	return &sourceFileUploadCntSign, nil
}

func GetDaoPreSignSourceFileUploadCntTotal(offlineDealId int64) (*int, error) {
	var daoPreSigns []*DaoPreSign
	sql := "select count(*) source_file_upload_cnt_total from offline_deal a, car_file_source b where a.car_file_id=b.car_file_id and a.id=?"
	err := database.GetDB().Raw(sql, offlineDealId).Scan(&daoPreSigns).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var sourceFileUploadCntTotal int = 0
	if len(daoPreSigns) > 0 {
		sourceFileUploadCntTotal = daoPreSigns[0].SourceFileUploadCntTotal
	}

	return &sourceFileUploadCntTotal, nil
}

func UpdateDaoPreSignSourceFileUploadCntSign(offlineDealId int64) error {
	sourceFileUploadCntSign, err := GetDaoPreSignSourceFileUploadCntSign(offlineDealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["source_file_upload_cnt_sign"] = *sourceFileUploadCntSign
	fields2BeUpdated["update_at"] = libutils.GetCurrentUtcSecond

	err = database.GetDB().Model(DaoPreSign{}).Where("offline_deal_id=?", offlineDealId).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func WriteDaoPreSign(txHash string, recipientWalletAddress string, dealId int64, batchCount int) error {
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

	sourceFileUploadCntTotal, err := GetDaoPreSignSourceFileUploadCntTotal(offlineDeal.Id)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	transactionReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	batchSizeMax := float64(*sourceFileUploadCntTotal) / float64(batchCount)

	daoPreSign, err := GetDaoPreSignByOfflineDealId(offlineDeal.Id)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	currentUtcSecond := libutils.GetCurrentUtcSecond()
	if daoPreSign == nil {
		daoPreSign = &DaoPreSign{
			CreateAt:                currentUtcSecond,
			SourceFileUploadCntSign: 0,
		}
	}

	daoPreSign.OfflineDealId = offlineDeal.Id
	daoPreSign.BatchCount = batchCount
	daoPreSign.BatchSizeMax = int(math.Ceil(batchSizeMax))
	daoPreSign.SourceFileUploadCntTotal = *sourceFileUploadCntTotal
	daoPreSign.NetworkId = network.ID
	daoPreSign.WalletIdSigner = walletSigner.ID
	daoPreSign.WalletIdRecipient = walletRecipient.ID
	daoPreSign.WalletIdContract = walletContract.ID
	daoPreSign.TxHash = txHash

	if transactionReceipt.Status == 1 {
		daoPreSign.Status = constants.DAO_PRE_SIGN_STATUS_SUCCESS
	} else {
		daoPreSign.Status = constants.DAO_PRE_SIGN_STATUS_FAILED
	}

	daoPreSign.UpdateAt = currentUtcSecond

	err = database.SaveOne(daoPreSign)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
