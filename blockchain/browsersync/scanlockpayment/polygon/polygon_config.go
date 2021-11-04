package polygon

import (
	"github.com/BurntSushi/toml"
	"log"
	"time"
)

type ConfigurationForPolygon struct {
	PolygonMainnetNode PolygonMainnetNode
}

type PolygonMainnetNode struct {
	RpcUrl                    string
	ChainID                   int64
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
	DaoSwanOracleAddress      string
	DaoEventFunctionSignature string
	CycleTimeInterval         time.Duration
	GasLimit                  uint64
}

var polygonConfig *ConfigurationForPolygon

func initCofig() {
	configFile := "./config/polygon/config_polygon.toml"
	if metaData, err := toml.DecodeFile(configFile, &polygonConfig); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"PolygonMainnetNode", "rpcUrl"},
		{"PolygonMainnetNode", "paymentContractAddress"},
		{"PolygonMainnetNode", "contractFunctionSignature"},
		{"PolygonMainnetNode", "scanStep"},
		{"PolygonMainnetNode", "cycleTimeInterval"},
		{"PolygonMainnetNode", "chainID"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}

func GetConfig() ConfigurationForPolygon {
	if polygonConfig == nil {
		initCofig()
	}
	return *polygonConfig
}
