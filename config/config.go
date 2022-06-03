package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/filswan/go-swan-lib/logs"

	"github.com/BurntSushi/toml"
	"github.com/shopspring/decimal"
)

type Configuration struct {
	Port                  int          `toml:"port"`
	Release               bool         `toml:"release"`
	SwanPlatformFilWallet string       `toml:"swan_platform_fil_wallet"`
	FlinkUrl              string       `toml:"flink_url"`
	FilecoinNetwork       string       `toml:"filecoin_network"`
	Polygon               polygon      `toml:"polygon"`
	Database              database     `toml:"database"`
	SwanApi               swanApi      `toml:"swan_api"`
	Lotus                 lotus        `toml:"lotus"`
	IpfsServer            ipfsServer   `toml:"ipfs_server"`
	SwanTask              swanTask     `toml:"swan_task"`
	ScheduleRule          ScheduleRule `toml:"schedule_rule"`
}

type polygon struct {
	PolygonRpcUrl           string        `toml:"polygon_rpc_url"`
	PaymentContractAddress  string        `toml:"payment_contract_address"`
	PaymentRecipientAddress string        `toml:"payment_recipient_address"`
	DaoContractAddress      string        `toml:"dao_contract_address"`
	MintContractAddress     string        `toml:"mint_contract_address"`
	SushiDexAddress         string        `toml:"sushi_dex_address"`
	UsdcWFilPoolContract    string        `toml:"usdc_wFil_pool_contract"`
	GasLimit                uint64        `toml:"gas_limit"`
	LockTime                int           `toml:"lock_time"`
	PayMultiplyFactor       float32       `toml:"pay_multiply_factor"`
	UnlockIntervalMinute    time.Duration `toml:"unlock_interval_minute"`
	DaoUnlockIntervalBlock  int64         `toml:"dao_unlock_interval_block"`
	ScanPolygonBlockStep    uint64        `toml:"scan_polygon_block_step"`
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
	DirDeal         string          `toml:"dir_deal"`
	Description     string          `toml:"description"`
	CuratedDataset  string          `toml:"curated_dataset"`
	MaxPrice        decimal.Decimal `toml:"max_price"`
	ExpireDays      int             `toml:"expire_days"`
	VerifiedDeal    bool            `toml:"verified_deal"`
	FastRetrieval   bool            `toml:"fast_retrieval"`
	StartEpochHours int             `toml:"start_epoch_hours"`
	MinFileSize     int64           `toml:"min_file_size"`
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
	UnlockIntervalSecond         time.Duration `toml:"unlock_interval_second"`
	CreateTaskIntervalSecond     time.Duration `toml:"create_task_interval_second"`
	SendDealIntervalSecond       time.Duration `toml:"send_deal_interval_second"`
	ScanDealStatusIntervalSecond time.Duration `toml:"scan_deal_status_interval_second"`
	ScanPolygonIntervalSecond    time.Duration `toml:"scan_polygon_interval_second"`
}

var config *Configuration

func InitConfig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcs/config.toml")

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
		{"swan_platform_fil_wallet"},
		{"flink_url"},
		{"filecoin_network"},

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

		{"schedule_rule", "unlock_interval_second"},
		{"schedule_rule", "create_task_interval_second"},
		{"schedule_rule", "send_deal_interval_second"},
		{"schedule_rule", "scan_deal_status_interval_second"},
		{"schedule_rule", "scan_polygon_interval_second"},

		{"polygon", "polygon_rpc_url"},
		{"polygon", "payment_contract_address"},
		{"polygon", "payment_recipient_address"},
		{"polygon", "dao_contract_address"},
		{"polygon", "mint_contract_address"},
		{"polygon", "sushi_dex_address"},
		{"polygon", "usdc_wFil_pool_contract"},
		{"polygon", "gas_limit"},
		{"polygon", "lock_time"},
		{"polygon", "pay_multiply_factor"},
		{"polygon", "unlock_interval_minute"},
		{"polygon", "dao_unlock_interval_block"},
		{"polygon", "scan_polygon_block_step"},
	}

	for _, v := range requiredFields {
		if !metaData.IsDefined(v...) {
			logs.GetLogger().Fatal("required fields ", v)
		}
	}

	return true
}
