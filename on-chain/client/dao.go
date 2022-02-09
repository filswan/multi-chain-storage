package client

import (
	"context"
	"payment-bridge/config"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
)

func GetThreshHold() (uint8, error) {
	daoContractAddress := common.HexToAddress(config.GetConfig().Polygon.DaoContractAddress)
	client, _, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}

	callOpts := new(bind.CallOpts)
	callOpts.From = daoContractAddress
	callOpts.Context = context.Background()

	daoOracleContractInstance, err := goBind.NewFilswanOracle(daoContractAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}

	threshHold, err := daoOracleContractInstance.GetThreshold(callOpts)
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	logs.GetLogger().Info("dao threshHold is : ", threshHold)
	return threshHold, nil
}
