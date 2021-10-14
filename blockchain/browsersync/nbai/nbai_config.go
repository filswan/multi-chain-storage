package nbai

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

type ConfigurationForNbai struct {
	NbaiMainnetNode NbaiMainnetNode
}

type NbaiMainnetNode struct {
	RpcUrl                    string
	ChainID                   int64
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
	CycleTimeInterval         time.Duration
}

var nbaiConfig *ConfigurationForNbai

func initCofig() {
	configFile := "./config/polygon/config_polygon.toml"
	if metaData, err := toml.DecodeFile(configFile, &nbaiConfig); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"NbaiMainnetNode", "rpcUrl"},
		{"NbaiMainnetNode", "paymentContractAddress"},
		{"NbaiMainnetNode", "contractFunctionSignature"},
		{"NbaiMainnetNode", "scanStep"},
		{"NbaiMainnetNode", "cycleTimeInterval"},
		{"NbaiMainnetNode", "chainID"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}

func GetConfig() ConfigurationForNbai {
	if nbaiConfig == nil {
		initCofig()
	}
	return *nbaiConfig
}
