package goBind

import (
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

	infuraKey := os.LookupEnv("INFURA")
	client, err := ethclient.Dial("https://goerli.infura.io/v3/" + infuraKey)
	if err != nil {
		log.Fatal(err)
	}

	gatewayAddress := "0xad8cE271beE7b917F2a1870C8b64EDfF6aAF3342"

	paymentGateway, _ := PaymentGateway.NewPaymentGateway(common.HexToAddress(gatewayAddress), client)

	cid := "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72"

	info, _ := paymentGateway.GetLockedPaymentInfo(nil, cid)

	result, _ := paymentGateway.unlockPayment(nil, cid)
	// info.LockedFee
}
