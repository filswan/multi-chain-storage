package fileswan_payment

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog/log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Warn().Msg("Error loading .env file")
	}

}

func TestGateway(t *testing.T) {

	infuraKey, _ := os.LookupEnv("INFURA")
	client, err := ethclient.Dial("https://goerli.infura.io/v3/" + infuraKey)
	if err != nil {
		log.Err(err)
	}

	gatewayAddress := "0xad8cE271beE7b917F2a1870C8b64EDfF6aAF3342"

	paymentGateway, _ := NewPaymentGateway(common.HexToAddress(gatewayAddress), client)

	cid := "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72"

	info, _ := paymentGateway.GetLockedPaymentInfo(nil, cid)

	result, _ := paymentGateway.UnlockPayment(nil, cid)
	log.Info().Msg(info.Recipient.String())
	log.Info().Msg(result.Hash().String())
}

func TestPolygoncGateway(t *testing.T) {
	//https://polygon-mumbai.g.alchemy.com/v2/86HeefA3O9EF22t2NTLbmcpfN0hb9vlv
	infuraKey := "86HeefA3O9EF22t2NTLbmcpfN0hb9vlv"
	client, err := ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/" + infuraKey)
	if err != nil {

	}

	gatewayAddress := "0x5210ED929B5BEdBFFBA2F6b9A0b1B608eEAb3aa0"

	paymentGateway, _ := NewPaymentGateway(common.HexToAddress(gatewayAddress), client)

	cid := "bafk2bzaceb7cp727fxdrzudlgvsoivdwscydp35eb6wc3bzuflggfdhfa4rfe"

	info, _ := paymentGateway.GetLockedPaymentInfo(nil, cid)
	log.Info().Msg(info.Recipient.String() + " Locke fee:" + info.LockedFee.String())

	privateString := ""
	privateKey, _ := crypto.HexToECDSA(privateString)
	callOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))

	result, unlockErr := paymentGateway.UnlockPayment(callOpts, cid)
	if unlockErr != nil {
		log.Err(unlockErr)
	}
	log.Info().Msg(result.Hash().String())

}
