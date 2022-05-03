package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type DaoSignature struct {
	ID           int64  `json:"id"`
	NetworkId    int64  `json:"network_id"`
	DealId       int64  `json:"deal_id"`
	Status       string `json:"status"`
	TxHash       string `json:"tx_hash"`
	RecipientId  int64  `json:"recipient_id"`
	DaoAddressId int64  `json:"dao_address_id"`
	CreateAt     int64  `json:"create_at"`
	UpdateAt     int64  `json:"update_at"`
}

func GetDaoSignaturesByDealId(dealId int64) ([]*DaoSignature, error) {
	var daoSignatures []*DaoSignature
	err := database.GetDB().Where("deal_id=?", dealId).Order("create_at desc").Find(&daoSignatures).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return daoSignatures, nil
}

func GetDaoSignaturesByDealIdTxHash(dealId int64, txHash string) (*DaoSignature, error) {
	var eventDaoSignatures []*DaoSignature
	err := database.GetDB().Where("deal_id=? and tx_hash=?").Find(&eventDaoSignatures).Error

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(eventDaoSignatures) >= 1 {
		return eventDaoSignatures[0], nil
	}

	return nil, nil
}
