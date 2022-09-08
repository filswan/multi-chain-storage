package config

import (
	"multi-chain-storage/common/constants"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/BurntSushi/toml"
	"github.com/shopspring/decimal"
)

type Configuration struct {
	Port                    int          `toml:"port"`
	Release                 bool         `toml:"release"`
	FilecoinNetwork         string       `toml:"filecoin_network"`
	FilecoinWallet          string       `toml:"filecoin_wallet"`
	FlinkUrl                string       `toml:"flink_url"`
	Web3ApiUrlPolygonMumbai string       `toml:"web3_api_url_polygon_mumbai"`
	Web3ApiUrlBscTestnet    string       `toml:"web3_api_url_bsc_testnet"`
	Database                database     `toml:"database"`
	SwanApi                 swanApi      `toml:"swan_api"`
	Lotus                   lotus        `toml:"lotus"`
	IpfsServer              ipfsServer   `toml:"ipfs_server"`
	SwanTask                swanTask     `toml:"swan_task"`
	ScheduleRule            ScheduleRule `toml:"schedule_rule"`
	PaymentChainName        string
}

type database struct {
	DbHost       string `toml:"db_host"`
	DbPort       string `toml:"db_port"`
	DbSchemaName string `toml:"db_schema_name"`
	DbUsername   string `toml:"db_username"`
	DbPassword   string `toml:"db_password"`
	DbArgs       string `toml:"db_args"`
}

type lotus struct {
	ClientApiUrl      string `toml:"client_api_url"`
	ClientAccessToken string `toml:"client_access_token"`
}

type swanTask struct {
	DirDeal          string          `toml:"dir_deal"`
	Description      string          `toml:"description"`
	CuratedDataset   string          `toml:"curated_dataset"`
	MaxPrice         decimal.Decimal `toml:"max_price"`
	ExpireDays       int             `toml:"expire_days"`
	VerifiedDeal     bool            `toml:"verified_deal"`
	FastRetrieval    bool            `toml:"fast_retrieval"`
	StartEpochHours  int             `toml:"start_epoch_hours"`
	MinFileSize      int64           `toml:"min_file_size"`
	MaxFileNumPerCar int             `toml:"max_file_num_per_car"`
}

type swanApi struct {
	ApiUrl      string `toml:"api_url"`
	ApiKey      string `toml:"api_key"`
	AccessToken string `toml:"access_token"`
}

type ipfsServer struct {
	DownloadUrlPrefix string `toml:"download_url_prefix"`
	UploadUrlPrefix   string `toml:"upload_url_prefix"`
}

type ScheduleRule struct {
	CreateTaskIntervalSecond     time.Duration `toml:"create_task_interval_second"`
	SendDealIntervalSecond       time.Duration `toml:"send_deal_interval_second"`
	ScanDealStatusIntervalSecond time.Duration `toml:"scan_deal_status_interval_second"`
}

var config *Configuration

func InitConfig(paymentChainName string) {
	if !strings.EqualFold(paymentChainName, constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI) &&
		!strings.EqualFold(paymentChainName, constants.PAYMENT_CHAIN_NAME_POLYGON_MAINNET) &&
		!strings.EqualFold(paymentChainName, constants.PAYMENT_CHAIN_NAME_BSC_TESTNET) {
		logs.GetLogger().Fatal("invalid payment chain name:", paymentChainName)
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFilename := "config_" + paymentChainName + ".toml"
	configFile := filepath.Join(homedir, constants.CONFIG_PATH, configFilename)

	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		logs.GetLogger().Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			logs.GetLogger().Fatal("required fields not given")
		}
	}

	config.PaymentChainName = paymentChainName
}

func GetConfig() Configuration {
	if config == nil {
		logs.GetLogger().Fatal("config not initialized")
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},
		{"release"},
		{"filecoin_wallet"},
		{"flink_url"},
		{"filecoin_network"},
		{"web3_api_url_polygon_mumbai"},
		{"web3_api_url_bsc_testnet"},

		{"database", "db_host"},
		{"database", "db_port"},
		{"database", "db_schema_name"},
		{"database", "db_username"},
		{"database", "db_password"},
		{"database", "db_args"},

		{"swan_api", "api_url"},
		{"swan_api", "api_key"},
		{"swan_api", "access_token"},

		{"lotus", "client_api_url"},
		{"lotus", "client_access_token"},

		{"ipfs_server", "download_url_prefix"},
		{"ipfs_server", "upload_url_prefix"},

		{"swan_task", "dir_deal"},
		{"swan_task", "description"},
		{"swan_task", "curated_dataset"},
		{"swan_task", "max_price"},
		{"swan_task", "expire_days"},
		{"swan_task", "verified_deal"},
		{"swan_task", "fast_retrieval"},
		{"swan_task", "start_epoch_hours"},
		{"swan_task", "min_file_size"},
		{"swan_task", "max_file_num_per_car"},

		{"schedule_rule", "create_task_interval_second"},
		{"schedule_rule", "send_deal_interval_second"},
		{"schedule_rule", "scan_deal_status_interval_second"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			logs.GetLogger().Fatal("required fields ", v)
		}
	}

	return true
}
