package billing

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"net/http"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	common2 "payment-bridge/common"
	"payment-bridge/common/httpClient"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/on-chain/goBind"
	"strings"
)

func GetFileCoinLastestPriceService() (*PriceResult, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=filecoin&vs_currencies=usd"
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	var price *PriceResult
	err = json.Unmarshal(response, &price)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return price, nil
}

func getBillingCount(walletAddress string) (int64, error) {
	sql := "select  count(* ) as total_record " +
		"from event_polygon ep left join event_unlock_payment eup   on eup.payload_cid = ep.payload_cid where ep.address_from='" + walletAddress + "'"
	var recordCount common2.RecordCount
	err := database.GetDB().Raw(sql).Scan(&recordCount).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return recordCount.TotalRecord, nil
}

func getBillHistoryList(walletAddress, limit, offset string) ([]*BillingResult, error) {
	finalSql := "select ep.tx_hash,ep.address_from,ep.locked_fee,ep.deadline,ep.payload_cid,ep.lock_payment_time,ep.coin_type,eup.unlock_to_user_address,eup.unlock_to_user_amount,eup.unlock_time,'polygon' as network  " +
		"from event_polygon ep left join event_unlock_payment eup   on eup.payload_cid = ep.payload_cid where lower(ep.address_from)='" + strings.ToLower(walletAddress) + "'" +
		"order by lock_payment_time desc "

	var billingResultList []*BillingResult
	err := database.GetDB().Raw(finalSql).Scan(&billingResultList).Limit(limit).Offset(offset).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}
	return billingResultList, nil
}

func GetTaskDealsService(url string) (*PriceResult, error) {
	response, err := httpClient.SendRequestAndGetBytes(http.MethodGet, url, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	var price *PriceResult
	err = json.Unmarshal(response, &price)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return price, nil
}

func GetWfilPriceFromSushiPrice(client *ethclient.Client, wfilPrice string) (*big.Int, error) {
	//routerAddress sushiswap mumbai address
	routerAddress := polygon.GetConfig().PolygonMainnetNode.RouterAddressOfSushiswapOnPolygon

	//pairAddress sushiswap mumbai address
	pairAddress := polygon.GetConfig().PolygonMainnetNode.PairAddressBetweenWfilUsdcOfSushiswapOnPolygon

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
