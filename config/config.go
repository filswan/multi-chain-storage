package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"strings"
)

type Configuration struct {
	Port                        string
	Database                    database
	Dev                         bool
	Temp                        temp       `toml:"temp"`
	SwanApi                     swanApi    `toml:"swan_api"`
	IpfsServer                  ipfsServer `toml:"ipfs_server"`
	Lotus                       lotus      `toml:"lotus"`
	SwanTask                    swanTask   `toml:"swan_task"`
	ScheduleRule                ScheduleRule
	AdminWalletOnPolygon        string //pay for gas
	SwanPaymentAddressOnPolygon string
}

type database struct {
	DbUsername   string
	DbPwd        string
	DbHost       string
	DbPort       string
	DbSchemaName string
	DbArgs       string
}

type lotus struct {
	ApiUrl      string `toml:"api_url"`
	AccessToken string `toml:"access_token"`
}

type swanTask struct {
	Description     string `toml:"description"`
	CuratedDataset  string `toml:"curated_dataset"`
	Tags            string `toml:"tags"`
	MinPrice        string `toml:"min_price"`
	MaxPrice        string `toml:"max_price"`
	ExpireDays      int    `toml:"expire_days"`
	VerifiedDeal    bool   `toml:"verified_deal"`
	FastRetrieval   bool   `toml:"fast_retrieval"`
	StartEpochHours int    `toml:"start_epoch_hours"`
	MinerId         string `toml:"miner_id"`
}

type temp struct {
	DirDeal string `toml:"dir_deal"`
}

type swanApi struct {
	ApiUrl string `toml:"api_url"`
}

type ipfsServer struct {
	DownloadUrlPrefix string `toml:"download_url_prefix"`
	UploadUrl         string `toml:"upload_url"`
}

type ScheduleRule struct {
	Nbai2BscMappingRedoRule string
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
