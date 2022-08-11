package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

const (
	LOTUS_JSON_RPC_ID      = 7878
	LOTUS_JSON_RPC_VERSION = "2.0"
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

	if !strings.EqualFold(systemParamResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", systemParamResponse.Status, systemParamResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &systemParamResponse.Data, nil
}

type DealState struct {
	Result *struct {
		State struct {
			SectorStartEpoch int
		}
	} `json:"result"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type LotusJsonRpcParams struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

func IsDealActive(dealId int64) (*bool, error) {
	lotusApiUrl := config.GetConfig().Lotus.ClientApiUrl

	var params []interface{}
	params = append(params, dealId)
	params = append(params, nil)

	jsonRpcParams := LotusJsonRpcParams{
		JsonRpc: LOTUS_JSON_RPC_VERSION,
		Method:  "Filecoin.StateMarketStorageDeal",
		Params:  params,
		Id:      LOTUS_JSON_RPC_ID,
	}

	response, err := web.HttpGetNoToken(lotusApiUrl, jsonRpcParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var dealState DealState
	err = json.Unmarshal(response, &dealState)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if dealState.Error != nil {
		message := fmt.Sprintf("no deal state, code:%d,message:%s", dealState.Error.Code, dealState.Error.Message)
		logs.GetLogger().Info(message)
		dealStateBool := false
		return &dealStateBool, nil
	}

	dealStateBool := dealState.Result.State.SectorStartEpoch > -1

	return &dealStateBool, nil
}

func GetMonthStart() int64 {
	currentTime := time.Now()
	monthStart := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.UTC().Location())
	monthStartUtc := monthStart.Unix()
	return monthStartUtc
}
