package utils

import (
	"encoding/json"
	"fmt"
	"multi-chain-storage/common/constants"
	"multi-chain-storage/config"
	"net/url"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type SystemParam struct {
	ChainName               string  `json:"chain_name"`
	PaymentContractAddress  string  `json:"payment_contract_address"`
	PaymentRecipientAddress string  `json:"payment_recipient_address"`
	DaoContractAddress      string  `json:"dao_contract_address"`
	MintContractAddress     string  `json:"mint_contract_address"`
	DexAddress              string  `json:"dex_address"`
	UsdcWFilPoolContract    string  `json:"usdc_wFil_pool_contract"`
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

func GetSystemParam(web3ApiUrlPrefix string) (*SystemParam, error) {
	if web3ApiUrlPrefix == "" {
		switch config.GetConfig().PaymentChainName {
		case constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI:
			web3ApiUrlPrefix = config.GetConfig().Web3ApiUrlPolygonMumbai
		case constants.PAYMENT_CHAIN_NAME_BSC_TESTNET:
			web3ApiUrlPrefix = config.GetConfig().Web3ApiUrlBscTestnet
		case constants.PAYMENT_CHAIN_NAME_POLYGON_MAINNET:
			web3ApiUrlPrefix = config.GetConfig().Web3ApiUrlPolygonMainnet
		default:
			err := fmt.Errorf("chain:%s not supported now", config.GetConfig().PaymentChainName)
			logs.GetLogger().Error(err)
			return nil, err
		}
	}

	web3ApiUrl := libutils.UrlJoin(web3ApiUrlPrefix, "api/v1/common/system/params")
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
