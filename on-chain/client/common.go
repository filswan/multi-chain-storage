package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"payment-bridge/config"
	"payment-bridge/on-chain/goBind"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-swan-lib/logs"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetPrivateKeyPublicKey(privateKeyEnvName string) (*ecdsa.PrivateKey, *common.Address, error) {
	privateKeyOnPolygon := os.Getenv(privateKeyEnvName)
	if len(privateKeyOnPolygon) <= 0 {
		err := fmt.Errorf("env variable %s is not defined", privateKeyEnvName)
		logs.GetLogger().Fatal(err)
	}

	if strings.HasPrefix(strings.ToLower(privateKeyOnPolygon), "0x") {
		privateKeyOnPolygon = privateKeyOnPolygon[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyOnPolygon)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err := fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	publicKeyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	logs.GetLogger().Info(publicKeyAddress.Hex())

	return privateKey, &publicKeyAddress, nil
}

func GetEthClient() (*ethclient.Client, *rpc.Client, error) {
	polygonRpcUrl := config.GetConfig().Polygon.PolygonRpcUrl
	rpcClient, err := rpc.DialContext(context.Background(), polygonRpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	ethClient, err := ethclient.Dial(polygonRpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, rpcClient, err
	}

	return ethClient, rpcClient, nil
}

func GetSwanPaymentTransactor(ethClient *ethclient.Client) (*goBind.SwanPaymentTransactor, error) {
	contractAddress := common.HexToAddress(config.GetConfig().Polygon.PaymentContractAddress)
	swanPaymentTransactor, err := goBind.NewSwanPaymentTransactor(contractAddress, ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return swanPaymentTransactor, nil
}

func GetFilswanOracleSession(ethClient *ethclient.Client) (*goBind.FilswanOracleSession, error) {
	daoContractAddress := common.HexToAddress(config.GetConfig().Polygon.DaoContractAddress)
	filswanOracle, err := goBind.NewFilswanOracle(daoContractAddress, ethClient)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	filswanOracleSession := &goBind.FilswanOracleSession{
		Contract: filswanOracle,
	}

	return filswanOracleSession, nil
}

func GetTransactOpts(ethClient *ethclient.Client, privateKey *ecdsa.PrivateKey, publicKeyAddress common.Address) (*bind.TransactOpts, error) {
	nonce, err := ethClient.PendingNonceAt(context.Background(), publicKeyAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	gasPrice, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = config.GetConfig().Polygon.GasLimit
	transactOpts.Context = context.Background()

	return transactOpts, nil
}

func GetContractAbi() (*abi.ABI, error) {
	paymentAbiString := goBind.SwanPaymentABI

	contractAbi, err := abi.JSON(strings.NewReader(paymentAbiString))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &contractAbi, nil
}

func CheckTx(client *ethclient.Client, tx *types.Transaction) (*types.Receipt, error) {
retry:
	rp, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		if err == ethereum.NotFound {
			logs.GetLogger().Error("tx %v not found, check it later", tx.Hash().String())
			time.Sleep(1 * time.Second)
			goto retry
		} else {
			logs.GetLogger().Error("TransactionReceipt fail: %s", err)
			return nil, err
		}
	}
	return rp, nil
}

func GetFromAndToAddressByTxHash(client *ethclient.Client, chainID *big.Int, txHash common.Hash) (*addressInfo, error) {
	addrInfo := new(addressInfo)
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	addrInfo.AddrTo = tx.To().Hex()
	txMsg, err := tx.AsMessage(types.LatestSignerForChainID(chainID), nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	addrInfo.AddrFrom = txMsg.From().Hex()
	return addrInfo, nil
}

type addressInfo struct {
	AddrFrom string
	AddrTo   string
}
