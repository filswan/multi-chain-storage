package client

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"multi-chain-storage/config"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-swan-lib/logs"

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
