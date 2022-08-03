package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"multi-chain-storage/config"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

func DownloadFile(sourceUrl string, destFilepath string) error {
	// Create the file
	out, err := os.Create(destFilepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(sourceUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

type SystemParam struct {
	PaymentContractAddress  string  `json:"payment_contract_address"`
	PaymentRecipientAddress string  `json:"payment_recipient_address"`
	MintContractAddress     string  `json:"mint_contract_address"`
	UsdcAddress             string  `json:"usdc_address"`
	GasLimit                uint64  `json:"gas_limit"`
	LockTime                int     `json:"lock_time"`
	PayMultiplyFactor       float32 `json:"pay_multiply_factor"`
	DaoThreshold            int     `json:"dao_threshold"`
	FilecoinPrice           int64   `json:"filecoin_price"`
}

type SystemParamResponse struct {
	Status  string      `json:"status"`
	Data    SystemParam `json:"data"`
	Message string      `json:"message"`
}

func GetSystemParam() (*SystemParam, error) {
	web3ApiUrl := libutils.UrlJoin(config.GetConfig().Web3ApiUrl, "api/v1/common/system/params")
	params := url.Values{}
	response, err := web.HttpGetNoToken(web3ApiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var systemParamResponse SystemParamResponse
	err = json.Unmarshal(response, &systemParamResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &systemParamResponse.Data, nil
}
