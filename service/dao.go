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
	SignerWalletAddress string   `json:"signer_wallet_address"`
	DealId              int64    `json:"deal_id"`
	WCids               []string `json:"w_cid"`
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

		sourceFileUploads, err := models.GetSourceFileUploadsByCarFileId(offlineDeal.CarFileId)
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

func WriteDaoSignature(txHash string, signerWalletAddress string, dealId int64) error {
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

	walletSigner, err := models.GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
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

	offlineDeal, err := models.GetOfflineDealByDealId(dealId)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if offlineDeal == nil {
		err := fmt.Errorf("offline deal with deal id: %d not exists", dealId)
		logs.GetLogger().Error(err)
		return err
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
	if transReceipt.Status == 1 {
		daoSignature.Status = constants.DAO_SIGNATURE_STATUS_SUCCESS
	} else {
		daoSignature.Status = constants.DAO_SIGNATURE_STATUS_FAIL
	}
	daoSignature.TxHash = txHash
	daoSignature.SignerWalletId = walletSigner.ID
	daoSignature.DaoRecipientId = walletDao.ID
	daoSignature.UpdateAt = currentUtcSecond

	err = database.SaveOne(daoSignature)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
