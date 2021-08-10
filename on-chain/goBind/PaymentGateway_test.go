package fileswan_payment

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
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/" + infuraKey)
	if err != nil {
		log.Fatal(err)
	}

	gatewayAddress := "0xad8cE271beE7b917F2a1870C8b64EDfF6aAF3342"

	paymentGateway, _ := PaymentGateway.NewPaymentGateway(common.HexToAddress(gatewayAddress), client)

	cid := "bafykbzaceafdasngafrordoboczbmp4enweo7omqelfgcjf3cty6tnlpjqw72"

	tests := []struct {
		input string
		sep   string
		diff  int
	}{
		{input: "10000", diff: 1},
		{input: "100000", diff: 2},
		{input: "1000000", diff: 3},
		{input: "10000000", diff: 5},
		{input: "100000000", diff: 8},
	}
	// var zero12, _ = new(big.Int).SetString("1000000000000", 10)
	// var zero0, _ = new(big.Int).SetString("1", 10)

	// multiplier := []*big.Int{zero12, zero0, zero0}
	info, _ := paymentGateway.GetLockedPaymentInfo(nil, cid)

}
