package billing

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	common2 "payment-bridge/common"
	"payment-bridge/common/httpClient"
	"payment-bridge/common/utils"
	"payment-bridge/database"
	"payment-bridge/on-chain/goBind"
	"strings"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
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

func getBillingCount(walletAddress, txHash string) (int64, error) {
	/*sql := "select  count(* ) as total_record " +
	"from event_lock_payment ep left join event_unlock_payment eup   on eup.payload_cid = ep.payload_cid where ep.address_from='" + walletAddress + "'"*/
	sql := " select count(* ) as total_record from ( " +
		"     select ep.* from event_lock_payment ep " +
		"     left join event_unlock_payment eup   on eup.payload_cid = ep.payload_cid " +
		"     where ep.address_from='" + walletAddress + "'"
	if txHash != "" {
		sql = sql + " and ep.tx_hash='" + txHash + "' "
	}
	endSql := " ) bh inner join  (select distinct substring_index(source_file_path, '/', -1) as file_name, payload_cid from deal_file) as df " +
		" on bh.payload_cid=df.payload_cid "
	finalSql := sql + endSql
	var recordCount common2.RecordCount
	err := database.GetDB().Raw(finalSql).Scan(&recordCount).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return 0, err
	}
	return recordCount.TotalRecord, nil
}

func getBillHistoryList(walletAddress, txHash, limit, offset string) ([]*BillingResult, error) {
	startSql := "select distinct t.*,co.short_name as coin_type,ne.network_name from ( " +
		"     select ep.tx_hash,ep.address_from,ep.locked_fee,ep.deadline,ep.payload_cid,ep.lock_payment_time,ep.coin_id,ep.network_id, IFNULL(eup.unlock_to_user_address, eupr.user_address) as unlock_to_user_address , IFNULL(eup.unlock_to_user_amount, eupr.expire_user_amount) as unlock_to_user_amount, IFNULL(eup.unlock_time,eupr.block_time) as unlock_time,'polygon' as network" +
		"     from event_lock_payment ep left join event_unlock_payment eup  on eup.payload_cid = ep.payload_cid left join event_expire_payment eupr on eupr.payload_cid = ep.payload_cid" +
		"     where lower(ep.address_from)=lower('" + walletAddress + "')"
	if txHash != "" {
		startSql = startSql + " and ep.tx_hash='" + txHash + "' "
	}

	endSql := " ) as t inner join coin co on t.coin_id=co.id " +
		" inner join network ne on t.network_id= ne.id" +
		" order by lock_payment_time desc"
	finalSql := " select bh.*,df.file_name from (" + startSql + endSql + " ) as bh inner join " +
		" ( select distinct substring_index(source_file_path, '/', -1) as file_name, payload_cid from deal_file ) as df " +
		" on bh.payload_cid=df.payload_cid"

	var billingResultList []*BillingResult
	err := database.GetDB().Raw(finalSql).Limit(limit).Offset(offset).Order("lock_payment_time desc").Scan(&billingResultList).Error
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

func LockPaymentService(client *ethclient.Client, userWalletAddress, privateKeyOfUser, payloadCid string, lockedFee *big.Int) error {
	paymentGatewayAddress := common.HexToAddress(polygon.GetConfig().PolygonMainnetNode.PaymentContractAddress)
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress(userWalletAddress))
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if strings.HasPrefix(strings.ToLower(privateKeyOfUser), "0x") {
		privateKeyOfUser = privateKeyOfUser[2:]
	}

	privateKey, _ := crypto.HexToECDSA(privateKeyOfUser)
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	callOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	//callOpts := new(bind.TransactOpts)
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = uint64(polygon.GetConfig().PolygonMainnetNode.GasLimit)
	callOpts.Context = context.Background()

	paymentGatewayInstance, err := goBind.NewSwanPayment(paymentGatewayAddress, client)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	paymentParam := new(goBind.IPaymentMinimallockPaymentParam)
	paymentParam.Id = payloadCid
	paymentParam.LockTime = big.NewInt(utils.GetCurrentUtcMilliSecond())
	paymentParam.Amount = lockedFee

	tx, err := paymentGatewayInstance.LockTokenPayment(callOpts, *paymentParam)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info("lock payment tx hash: ", tx.Hash())
	txRecept, err := utils.CheckTx(client, tx)
	if err != nil {
		logs.GetLogger().Error(err)
	} else {
		if txRecept.Status == uint64(1) {
			logs.GetLogger().Println("lock payment success! txHash=" + tx.Hash().Hex())
		} else {
			logs.GetLogger().Println("lock payment failed! txHash=" + tx.Hash().Hex())
		}
	}

	return err
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
