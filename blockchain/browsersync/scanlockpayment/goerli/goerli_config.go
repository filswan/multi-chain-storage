package goerli

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/BurntSushi/toml"
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

func initConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcp/config_goerli.toml")
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
		initConfig()
	}
	return *goerliConfig
}
