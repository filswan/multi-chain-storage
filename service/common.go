package service

import (
	"fmt"
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
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

func GetSystemParams() (map[string]interface{}, error) {
	systemParams := map[string]interface{}{}
	systemParams[constants.SYSTEM_PARAM_PAYMENT_CONTRACT_ADDRESS] = config.GetConfig().Polygon.PaymentContractAddress
	systemParams[constants.SYSTEM_PARAM_PAYMENT_RECIPIENT_ADDRESS] = config.GetConfig().Polygon.PaymentRecipientAddress
	systemParams[constants.SYSTEM_PARAM_MINT_CONTRACT_ADDRESS] = config.GetConfig().Polygon.MintContractAddress
	systemParams[constants.SYSTEM_PARAM_GAS_LIMIT] = config.GetConfig().Polygon.GasLimit
	systemParams[constants.SYSTEM_PARAM_LOCK_TIME] = config.GetConfig().Polygon.LockTime
	systemParams[constants.SYSTEM_PARAM_PAY_MULTIPLY_FACTOR] = config.GetConfig().Polygon.PayMultiplyFactor

	token, err := models.GetTokenByName(constants.TOKEN_USDC_NAME)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if token == nil {
		err := fmt.Errorf("token:%s not found", constants.TOKEN_USDC_NAME)
		logs.GetLogger().Error(err)
		return nil, err
	}

	systemParams[constants.TOKEN_USDC_ADRESS] = token.Address

	return systemParams, nil
}
