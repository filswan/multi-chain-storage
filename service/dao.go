package service

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"

	"github.com/filswan/go-swan-lib/logs"
)

type Deal2Sign struct {
	DealId int64    `json:"deal_id"`
	WCids  []string `json:"w_cid"`
}

func GetDeals2PreSign(signerWalletAddress string) ([]*Deal2Sign, error) {
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
