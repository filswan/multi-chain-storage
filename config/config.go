package config

import (
	"os"
	"path/filepath"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/BurntSushi/toml"
	"github.com/shopspring/decimal"
)

type Configuration struct {
	Port                            string       `toml:"port"`
	Release                         bool         `toml:"release"`
	AdminWalletOnPolygon            string       `toml:"admin_wallet_on_polygon"`
	FileCoinWallet                  string       `toml:"file_coin_wallet"`
	FilinkUrl                       string       `toml:"filink_url"`
	PolygonRpcUrl                   string       `toml:"polygon_rpc_url"`
	PaymentContractAddress          string       `toml:"payment_contract_address"`
	ContractLockFunctionSignature   string       `toml:"contract_lock_function_signature"`
	ContractUnlockFunctionSignature string       `toml:"contract_unlock_function_signature"`
	Database                        database     `toml:"database"`
	SwanApi                         swanApi      `toml:"swan_api"`
	Lotus                           lotus        `toml:"lotus"`
	IpfsServer                      ipfsServer   `toml:"ipfs_server"`
	SwanTask                        swanTask     `toml:"swan_task"`
	ScheduleRule                    ScheduleRule `toml:"schedule_rule"`
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
	DirDeal              string          `toml:"dir_deal"`
	Description          string          `toml:"description"`
	CuratedDataset       string          `toml:"curated_dataset"`
	Tags                 string          `toml:"tags"`
	MaxPrice             decimal.Decimal `toml:"max_price"`
	ExpireDays           int             `toml:"expire_days"`
	VerifiedDeal         bool            `toml:"verified_deal"`
	FastRetrieval        bool            `toml:"fast_retrieval"`
	StartEpochHours      int             `toml:"start_epoch_hours"`
	MaxAutoBidCopyNumber int             `toml:"max_auto_bid_copy_number"`
	MinFileSize          int64           `toml:"min_file_size"`
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
	UnlockPaymentRule  string `toml:"unlock_payment_rule"`
	CreateTaskRule     string `toml:"create_task_rule"`
	SendDealRule       string `toml:"send_deal_rule"`
	ScanDealStatusRule string `toml:"scan_deal_status_rule"`
}

var config *Configuration

func InitConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcp/config.toml")

	if metaData, err := toml.DecodeFile(configFile, &config); err != nil {
		logs.GetLogger().Fatal("error:", err)
	} else {
		if !requiredFieldsAreGiven(metaData) {
			logs.GetLogger().Fatal("required fields not given")
		}
	}
}

func GetConfig() Configuration {
	if config == nil {
		InitConfig()
	}
	return *config
}

func requiredFieldsAreGiven(metaData toml.MetaData) bool {
	requiredFields := [][]string{
		{"port"},
		{"release"},
		{"admin_wallet_on_polygon"},
		{"file_coin_wallet"},
		{"filink_url"},
		{"polygon_rpc_url"},

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
		{"swan_task", "tags"},
		{"swan_task", "max_price"},
		{"swan_task", "expire_days"},
		{"swan_task", "verified_deal"},
		{"swan_task", "fast_retrieval"},
		{"swan_task", "start_epoch_hours"},
		{"swan_task", "max_auto_bid_copy_number"},
		{"swan_task", "min_file_size"},

		{"schedule_rule", "unlock_payment_rule"},
		{"schedule_rule", "create_task_rule"},
		{"schedule_rule", "send_deal_rule"},
		{"schedule_rule", "scan_deal_status_rule"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			logs.GetLogger().Fatal("required fields ", v)
		}
	}

	return true
}
