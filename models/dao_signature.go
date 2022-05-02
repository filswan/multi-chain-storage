package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
)

type DaoSignature struct {
	ID         int64  `json:"id"`
	NetworkId  int64  `json:"network_id"`
	TxHash     string `json:"tx_hash"`
	Recipient  string `json:"recipient"`
	DealId     int64  `json:"deal_id"`
	Status     bool   `json:"status"`
	DaoAddress string `json:"dao_address"`
	CreateAt   int64  `json:"create_at"`
	UpdateAt   int64  `json:"update_at"`
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
