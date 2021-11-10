package config

import (
	"github.com/BurntSushi/toml"
	"github.com/shopspring/decimal"
	"log"
	"strings"
)

type Configuration struct {
	Port                        string
	Database                    database
	Dev                         bool
	SwanApi                     swanApi      `toml:"swan_api"`
	IpfsServer                  ipfsServer   `toml:"ipfs_server"`
	Lotus                       lotus        `toml:"lotus"`
	SwanTask                    swanTask     `toml:"swan_task"`
	ScheduleRule                ScheduleRule `toml:"schedule_rule"`
	AdminWalletOnPolygon        string       //pay for gas
	SwanPaymentAddressOnPolygon string
	FileCoinWallet              string
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
	DirDeal         string          `toml:"dir_deal"`
	Description     string          `toml:"description"`
	CuratedDataset  string          `toml:"curated_dataset"`
	Tags            string          `toml:"tags"`
	MinPrice        decimal.Decimal `toml:"min_price"`
	MaxPrice        decimal.Decimal `toml:"max_price"`
	ExpireDays      int             `toml:"expire_days"`
	VerifiedDeal    bool            `toml:"verified_deal"`
	FastRetrieval   bool            `toml:"fast_retrieval"`
	StartEpochHours int             `toml:"start_epoch_hours"`
	MinerId         string          `toml:"miner_id"`
}

type swanApi struct {
	ApiUrl                     string `toml:"api_url"`
	ApiKey                     string `toml:"api_key"`
	AccessToken                string `toml:"access_token"`
	GetShouldSendTaskUrlSuffix string `toml:"get_should_send_task_url_suffix"`
}

type ipfsServer struct {
	DownloadUrlPrefix string `toml:"download_url_prefix"`
	UploadUrl         string `toml:"upload_url"`
}

type ScheduleRule struct {
	UnlockPaymentRule string `toml:"unlock_payment_rule"`
	SendDealRule      string `toml:"send_deal_rule"`
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
