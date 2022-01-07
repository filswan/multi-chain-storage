package bsc

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/BurntSushi/toml"
)

type ConfigurationForBsc struct {
	BscMainnetNode BscMainnetNode
}

type BscMainnetNode struct {
	RpcUrl                    string
	ChainID                   int64
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
	CycleTimeInterval         time.Duration
}

var bscConfig *ConfigurationForBsc

func initConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcp/config_bsc.toml")
	if metaData, err := toml.DecodeFile(configFile, &bscConfig); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"BscMainnetNode", "rpcUrl"},
		{"BscMainnetNode", "paymentContractAddress"},
		{"BscMainnetNode", "contractFunctionSignature"},
		{"BscMainnetNode", "scanStep"},
		{"BscMainnetNode", "cycleTimeInterval"},
		{"BscMainnetNode", "chainID"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}

func GetConfig() ConfigurationForBsc {
	if bscConfig == nil {
		initConfig()
	}
	return *bscConfig
}
