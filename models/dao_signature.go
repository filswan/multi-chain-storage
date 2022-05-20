package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
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
		"where a.is_dao=true"

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
