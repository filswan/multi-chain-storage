package service

import (
	"context"
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"multi-chain-storage/database"
	"multi-chain-storage/models"
	"multi-chain-storage/on-chain/client"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/filswan/go-swan-lib/utils"
)

type Deal2Sign struct {
	DealId int64    `json:"deal_id"`
	WCids  []string `json:"w_cid"`
}

func GetDeals2Sign(signerWalletAddress string) ([]*Deal2Sign, error) {
	signerWallet, err := models.GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	offlineDeals, err := models.GetOfflineDeals2BeSigned(signerWallet.ID)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	deals2Sign := []*Deal2Sign{}

	for _, offlineDeal := range offlineDeals {
		deal2BeSigned := &Deal2Sign{
			DealId: *offlineDeal.DealId,
			WCids:  []string{},
		}

		sourceFileUploads, err := models.GetSourceFileUploadOutsByCarFileId(offlineDeal.CarFileId)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		for _, sourceFileUpload := range sourceFileUploads {
			deal2BeSigned.WCids = append(deal2BeSigned.WCids, sourceFileUpload.Uuid+sourceFileUpload.PayloadCid)
		}

		deals2Sign = append(deals2Sign, deal2BeSigned)
	}

	return deals2Sign, nil
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

	walletSigner, err := models.GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if walletSigner.IsDao == nil || !*walletSigner.IsDao {
		err = models.SetWalletAsDao(walletSigner.ID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	walletContract, err := models.GetWalletByAddress(daoContractAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	walletRecipient, err := models.GetWalletByAddress(recipientWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	transactionReceipt, err := ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	offlineDeal, err := models.GetOfflineDealByDealId(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if offlineDeal == nil {
		err := fmt.Errorf("offline deal with deal id: %d not exists", dealId)
		logs.GetLogger().Error(err)
		return nil
	}

	daoSignature, err := models.GetDaoSignaturesByOfflineDealIdTxHash(offlineDeal.Id, txHash)
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

func RegisterDao(daoWalletAddress string) error {
	daoWallet, err := models.GetWalletByAddress(daoWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if daoWallet.IsDao == nil || !*daoWallet.IsDao {
		err = models.SetWalletAsDao(daoWallet.ID)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	return nil
}
