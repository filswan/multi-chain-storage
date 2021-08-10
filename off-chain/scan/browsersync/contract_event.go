package browsersync

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"payment-bridge/off-chain/config"
	"payment-bridge/off-chain/logs"
	"payment-bridge/off-chain/models"
)

/**
 * created on 04/03/19.
 * author: nebula-ai-chengkun
 * Copyright defined in navigator-api/LICENSE.txt
 */

type ContractEventResponse struct {
	Result models.ContractEventArray `json:"result"`
}

func GetContractEvent(ip string) (models.ContractEventArray, error) {
	config := config.GetConfig()
	evenPort := config.MainnetNode.ContractEventPort
	res, err := http.Get("http://" + ip + ":" + evenPort + "/address")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		logs.GetLogger().Error(err)
		return models.ContractEventArray{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.GetLogger().Error(err)
		return models.ContractEventArray{}, err
	}

	var contractEventResponse ContractEventResponse
	err = json.Unmarshal(resBody, &contractEventResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return models.ContractEventArray{}, err
	}

	contractEvent := contractEventResponse.Result

	return contractEvent, err
}
