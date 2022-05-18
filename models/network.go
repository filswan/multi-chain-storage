package models

import (
	"multi-chain-storage/database"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type Network struct {
	ID                  int64  `json:"id"`
	Name                string `json:"name"`
	LastScanBlockNumber *int64 `json:"last_scan_block_number"`
	Description         string `json:"description"`
	CreateAt            int64  `json:"create_at"`
	UpdateAt            int64  `json:"update_at"`
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
