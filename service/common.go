package service

import (
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/models"
	"runtime"

	"github.com/filswan/go-swan-lib/logs"
)

func GetHostInfo() *common.HostInfo {
	hostInfo := common.HostInfo{
		Version:         common.GetVersion(),
		OperatingSystem: runtime.GOOS,
		Architecture:    runtime.GOARCH,
		CpuNumber:       runtime.NumCPU(),
	}

	return &hostInfo
}

func GetSystemParams() (map[string]string, error) {
	systemConf, err := models.GetSystemParams()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	config := map[string]string{}

	for _, conf := range systemConf {
		config[conf.Name] = conf.Value
	}

	coin, err := models.GetCoinByName(constants.COIN_USDC_NAME)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if coin == nil {
		err := fmt.Errorf("coin:%s not found", constants.COIN_USDC_NAME)
		logs.GetLogger().Error(err)
		return nil, err
	}

	config[constants.COIN_USDC_ADRESS] = coin.Address

	return config, nil
}
