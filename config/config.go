package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"strings"
	"time"
)

type Configuration struct {
	Port     string
	Database database
	Dev      bool
}

type database struct {
	DbUsername   string
	DbPwd        string
	DbHost       string
	DbPort       string
	DbSchemaName string
	DbArgs       string
}

type NbaiMainnetNode struct {
	RpcUrl                    string
	PaymentContractAddress    string
	ContractFunctionSignature string
	ScanStep                  int64
	StartFromBlockNo          int64
	CycleTimeInterval         time.Duration
	ChainID                   int64
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

func GetConfig() Configuration {
	if config == nil {
		InitConfig("")
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
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			log.Fatal("required fields ", v)
		}
	}

	return true
}
