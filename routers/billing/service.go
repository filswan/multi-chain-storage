package billing

import (
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"payment-bridge/common/httpClient"
	"payment-bridge/common/utils"
	"payment-bridge/config"
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

func getBillHistoryList(walletAddress, limit, offset string) ([]*BillingResult, error) {
	sql := "select a.tx_hash,a.locked_fee,b.cn_name coin_type,h.file_name,d.payload_cid,h.wallet_address address_from,c.network_name network,a.lock_payment_time"
	sql = sql + " from event_lock_payment a, coin b, network c, source_file d, source_file_upload_history h"
	sql = sql + " where a.coin_id=b.id and a.network_id=c.id and a.payload_cid=d.payload_cid and a.address_from=h.wallet_address and h.wallet_address=?"

	var billingResultList []*BillingResult
	err := database.GetDB().Raw(sql, walletAddress).Limit(limit).Offset(offset).Order("lock_payment_time desc").Scan(&billingResultList).Error
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err

	}
	return billingResultList, nil
}
func getBillHistoriesByWalletAddress(walletAddress string) ([]*BillingResult, error) {
	sql := "select a.tx_hash,a.locked_fee,b.cn_name coin_type,h.file_name,d.payload_cid,h.wallet_address address_from,c.network_name network,a.lock_payment_time"
	sql = sql + " from event_lock_payment a, coin b, network c, source_file d, source_file_upload_history h"
	sql = sql + " where a.coin_id=b.id and a.network_id=c.id and a.payload_cid=d.payload_cid and a.address_from=h.wallet_address and h.wallet_address=?"

	var billingResultList []*BillingResult
	err := database.GetDB().Raw(sql, walletAddress).Order("lock_payment_time desc").Scan(&billingResultList).Error
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
	paymentGatewayAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)
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
	callOpts.GasLimit = uint64(config.GetConfig().Polygon.GasLimit)
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
