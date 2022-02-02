package models

import (
	"errors"
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

func FindNetworkIdByUUID(uuid string) (int64, error) {
	db := database.GetDB()
	var models []*Network
	err := db.Where("uuid='" + uuid + "'").Find(&models).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, nil
	}
	if len(models) > 0 {
		return models[0].ID, nil
	} else {
		err = errors.New("There is no network info by uuid :" + uuid)
		return 0, err
	}
}
