package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type DaoSignature struct {
	ID             int64  `json:"id"`
	NetworkId      int64  `json:"network_id"`
	OfflineDealId  int64  `json:"offline_deal_id"`
	Status         string `json:"status"`
	TxHash         string `json:"tx_hash"`
	SignerWalletId int64  `json:"signer_wallet_id"`
	DaoRecipientId int64  `json:"dao_recipient_id"`
	CreateAt       int64  `json:"create_at"`
	UpdateAt       int64  `json:"update_at"`
}

func GetDaoSignaturesByOfflineDealId(offlineDealId int64) ([]*DaoSignature, error) {
	var daoSignatures []*DaoSignature
	err := database.GetDB().Where("offline_deal_id=?", offlineDealId).Order("create_at desc").Find(&daoSignatures).Error

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
