package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type Network struct {
	ID                         int64  `json:"id"`
	Name                       string `json:"name"`
	LastScanBlockNumberPayment *int64 `json:"last_scan_block_number_payment"`
	LastScanBlockNumberDao     *int64 `json:"last_scan_block_number_dao"`
	Description                string `json:"description"`
	CreateAt                   int64  `json:"create_at"`
	UpdateAt                   int64  `json:"update_at"`
}

func GetNetworkByName(name string) (*Network, error) {
	var networks []*Network
	err := database.GetDB().Where("name=?", name).Find(&networks).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(networks) > 0 {
		return networks[0], nil
	}

	network, err := SaveNetwork(name)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return network, nil
}

func SaveNetwork(name string) (*Network, error) {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	network := Network{
		Name:     name,
		CreateAt: currentUtcSecond,
		UpdateAt: currentUtcSecond,
	}

	networkResult := database.GetDB().Create(&network)
	err := networkResult.Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	networkCreated := networkResult.Value.(*Network)

	return networkCreated, nil
}

func UpdateNetworkLastScanBlockNumberDao(networkId int64, lastScanBlockNumberDao int64) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["last_scan_block_number_dao"] = lastScanBlockNumberDao
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(Network{}).Where("id=?", networkId).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func UpdateNetworkLastScanBlockNumberPayment(networkId int64, lastScanBlockNumberPayment int64) error {
	currentUtcSecond := libutils.GetCurrentUtcSecond()
	fields2BeUpdated := make(map[string]interface{})
	fields2BeUpdated["last_scan_block_number_payment"] = lastScanBlockNumberPayment
	fields2BeUpdated["update_at"] = currentUtcSecond

	err := database.GetDB().Model(Network{}).Where("id=?", networkId).Update(fields2BeUpdated).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
