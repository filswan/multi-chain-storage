package nbai

import (
	"log"
	"os"
	"path/filepath"
	"payment-bridge/logs"
	"time"

	"github.com/BurntSushi/toml"
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
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcp/config_nbai.toml")
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
