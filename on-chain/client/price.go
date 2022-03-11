package client

import (
	"errors"
	"math/big"
	"payment-bridge/config"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filswan/go-swan-lib/logs"
)

func GetWfilPriceFromSushiPrice(wfilPrice string) (*big.Int, error) {
	ethClient, _, err := GetEthClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	// sushiswap mumbai address
	sushiDexAddress := common.HexToAddress(config.GetConfig().Polygon.SushiDexAddress)

	//pairAddress sushiswap mumbai address
	usdcWFilPoolContract := common.HexToAddress(config.GetConfig().Polygon.UsdcWFilPoolContract)

	contractRouter, _ := goBind.NewRouter(sushiDexAddress, ethClient)
	contractPair, _ := goBind.NewPair(usdcWFilPoolContract, ethClient)

	reserves, _ := contractPair.GetReserves(nil)

	//amt,_:=  new(big.Int).SetString("1000000000000000000", 10)
	amt, flag := new(big.Int).SetString(wfilPrice, 10)
	if !flag {
		err := errors.New("calculating filecoin to usdc pring occurred error")
		logs.GetLogger().Error(err)
		return nil, err
	}
	dyByContract, err := contractRouter.GetAmountOut(nil, amt, reserves.Reserve0, reserves.Reserve1)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return dyByContract, nil
}

func GetFileCoinLastestPrice() (*float64, error) {
	price, err := GetWfilPriceFromSushiPrice("1")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	priceFloat, _ := new(big.Float).SetInt(price).Float64()

	return &priceFloat, nil
}
