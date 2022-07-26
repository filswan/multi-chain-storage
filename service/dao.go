package service

import (
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"

	"github.com/filswan/go-swan-lib/logs"
)

func GetDeals2PreSign(signerWalletAddress string) ([]*models.Deal2PreSign, error) {
	signerWallet, err := models.GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	deals2PreSign, err := models.GetDeals2PreSign(signerWallet.ID)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals2PreSign, nil
}

func GetDeals2Sign(signerWalletAddress string) ([]*models.Deal2Sign, error) {
	signerWallet, err := models.GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	deals2Sign, err := models.GetDeals2Sign(signerWallet.ID)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	for _, deal2Sign := range deals2Sign {
		daoSignatures, err := models.GetDaoSignaturesByOfflineDealIdWalletIdSigner(deal2Sign.OfflineDealId, signerWallet.ID)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		for i := 0; i < deal2Sign.BatchCount; i++ {
			signed := false
			for _, daodaoSignature := range daoSignatures {
				if daodaoSignature.BatchNo == i {
					signed = true
				}
			}

			if !signed {
				sourceFileUploads, err := models.GetSourceFileUploadsByCarFileId(deal2Sign.CarFileId, &i)
				if err != nil {
					logs.GetLogger().Error(err)
					return nil, err
				}
				wCids := []string{}
				for _, sourceFileUpload := range sourceFileUploads {
					wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
					wCids = append(wCids, wCid)
				}

				deal2SignBatchInfo := &models.Deal2SignBatchInfo{
					BatchNo: i,
					WCid:    wCids,
				}

				deal2Sign.BatchInfo = append(deal2Sign.BatchInfo, deal2SignBatchInfo)
			}
		}
	}

	return deals2Sign, nil
}

func GetDeals2SignHash(signerWalletAddress string) ([]*models.Deal2Sign, error) {
	signerWallet, err := models.GetWalletByAddress(signerWalletAddress, constants.WALLET_TYPE_META_MASK)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	deals2Sign, err := models.GetDeals2SignHash(signerWallet.ID)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	for _, deal2Sign := range deals2Sign {
		sourceFileUploads, err := models.GetSourceFileUploadsByCarFileId(deal2Sign.CarFileId, nil)
		if err != nil {
			logs.GetLogger().Error(err)
			return nil, err
		}

		wCids := []string{}
		for _, sourceFileUpload := range sourceFileUploads {
			wCid := sourceFileUpload.Uuid + sourceFileUpload.PayloadCid
			wCids = append(wCids, wCid)
		}

		deal2SignBatchInfo := &models.Deal2SignBatchInfo{
			BatchNo: 0,
			WCid:    wCids,
		}

		deal2Sign.BatchInfo = append(deal2Sign.BatchInfo, deal2SignBatchInfo)
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
