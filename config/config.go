package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)

type Configuration struct {
	Port               string
	Database           database
	GoerliMainnetNode  GoerliMainnetNode
	PolygonMainnetNode PolygonMainnetNode
	Dev                bool
}

type database struct {
	DbUsername   string
	DbPwd        string
	DbHost       string
	DbPort       string
	DbSchemaName string
	DbArgs       string
}

type GoerliMainnetNode struct {
	RpcUrl                    string
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
}

type PolygonMainnetNode struct {
	RpcUrl                    string
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
}

var config *Configuration

func InitConfig(configFile string) {
	if strings.Trim(configFile, " ") == "" {
		configFile = "./config/config.toml"
	}
	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			log.Fatal("required fields not given")
		}
	}
}

func (c *Configuration) GetGoerliMainnetNode() string {
	return c.GoerliMainnetNode.RpcUrl
}

func GetConfig() Configuration {
	if config == nil {
		InitConfig("")
	}
	return *config
}

func GetConfigFromMainParams(configFile string) Configuration {
	if config == nil {
		InitConfig(configFile)
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},

		{"DataBase", "dbHost"},
		{"DataBase", "dbPort"},
		{"DataBase", "dbSchemaName"},
		{"DataBase", "dbUsername"},
		{"DataBase", "dbPwd"},

		{"GoerliMainnetNode", "rpcUrl"},
		{"GoerliMainnetNode", "paymentContractAddress"},
		{"GoerliMainnetNode", "contractFunctionSignature"},
		{"GoerliMainnetNode", "scanStep"},

		{"PolygonMainnetNode", "rpcUrl"},
		{"PolygonMainnetNode", "paymentContractAddress"},
		{"PolygonMainnetNode", "contractFunctionSignature"},
		{"PolygonMainnetNode", "scanStep"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
