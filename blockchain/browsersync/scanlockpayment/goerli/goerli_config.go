package goerli

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

type ConfigurationForGoerli struct {
	GoerliMainnetNode GoerliMainnetNode
}

type GoerliMainnetNode struct {
	RpcUrl                    string
	ChainID                   int64
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
	CycleTimeInterval         time.Duration
}

var goerliConfig *ConfigurationForGoerli

func initCofig() {
	configFile := "./config/goerli/config_goerli.toml"
	if metaData, err := toml.DecodeFile(configFile, &goerliConfig); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"GoerliMainnetNode", "rpcUrl"},
		{"GoerliMainnetNode", "paymentContractAddress"},
		{"GoerliMainnetNode", "contractFunctionSignature"},
		{"GoerliMainnetNode", "scanStep"},
		{"GoerliMainnetNode", "cycleTimeInterval"},
		{"GoerliMainnetNode", "chainID"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}

func GetConfig() ConfigurationForGoerli {
	if goerliConfig == nil {
		initCofig()
	}
	return *goerliConfig
}
