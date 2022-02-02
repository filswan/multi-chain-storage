package models

import (
	"fmt"
	"payment-bridge/database"

	"github.com/filswan/go-swan-lib/logs"
)

type Network struct {
	ID          int64  `json:"id"`
	NetworkName string `json:"network_name"`
	Uuid        string `json:"uuid"`
	RpcUrl      string `json:"rpc_url"`
	NativeCoin  string `json:"native_coin"`
	Description string `json:"description"`
}

func GetNetworkByName(name string) (*Network, error) {
	var networks []*Network
	err := database.GetDB().Where("network_name=?", name).Find(&networks).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if len(networks) > 0 {
		return networks[0], nil
	}

	err = fmt.Errorf("no network for name:%s", name)

	return nil, err
}
