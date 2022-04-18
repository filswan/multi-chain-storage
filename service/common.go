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

func GetSystemConfigParams(limit string) (map[string]string, error) {
	wallets, err := models.GetWalletsByType(constants.WALLET_TYPE_PAY_CONTRACT)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	if len(wallets) == 0 {
		err := fmt.Errorf("payment contract wallet not found")
		logs.GetLogger().Error(err)
		return nil, err
	}

	config := map[string]string{}
	config[constants.CONFIG_SWAN_PAYMENT_CONTRACT_ADDRESS] = wallets[0].Address

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
