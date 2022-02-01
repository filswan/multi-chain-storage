package client

import (
	"errors"
	"math/big"
	"payment-bridge/config"
	"payment-bridge/on-chain/goBind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filswan/go-swan-lib/logs"
)

func GetWfilPriceFromSushiPrice(client *ethclient.Client, wfilPrice string) (*big.Int, error) {
	//routerAddress sushiswap mumbai address
	routerAddress := config.GetConfig().Polygon.RouterAddressOfSushiswapOnPolygon

	//pairAddress sushiswap mumbai address
	pairAddress := config.GetConfig().Polygon.PairAddressBetweenWfilUsdcOfSushiswapOnPolygon

	contractRouter, _ := goBind.NewRouter(common.HexToAddress(routerAddress), client)
	contractPool, _ := goBind.NewPair(common.HexToAddress(pairAddress), client)

	reserves, _ := contractPool.GetReserves(nil)

	//amt,_:=  new(big.Int).SetString("1000000000000000000", 10)
	amt, flag := new(big.Int).SetString(wfilPrice, 10)
	if flag == false {
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
