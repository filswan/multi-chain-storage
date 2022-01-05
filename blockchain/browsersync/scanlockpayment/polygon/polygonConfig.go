package polygon

import (
	"log"
	"os"
	"path/filepath"
	"payment-bridge/logs"
	"time"

	"github.com/BurntSushi/toml"
)

type ConfigurationForPolygon struct {
	PolygonMainnetNode PolygonMainnetNode `toml:"polygon_mainnet_node"`
}

type PolygonMainnetNode struct {
	RpcUrl                                         string        `toml:"rpc_url"`
	PaymentContractAddress                         string        `toml:"payment_contract_address"`
	ContractLockFunctionSignature                  string        `toml:"contract_lock_function_signature"`
	ContractUnlockFunctionSignature                string        `toml:"contract_unlock_function_signature"`
	PaymentUnlockContractAddress                   string        `toml:"payment_unlock_contract_address"`
	ScanStep                                       int64         `toml:"scan_step"`
	StartFromBlockNo                               int64         `toml:"start_from_blockNo"`
	DaoSwanOracleAddress                           string        `toml:"dao_swan_oracle_address"`
	DaoEventFunctionSignature                      string        `toml:"dao_event_function_signature"`
	CycleTimeInterval                              time.Duration `toml:"cycle_time_interval"`
	GasLimit                                       int64         `toml:"gas_limit"`
	RouterAddressOfSushiswapOnPolygon              string        `toml:"router_address_of_sushiswap_on_polygon"`
	PairAddressBetweenWfilUsdcOfSushiswapOnPolygon string        `toml:"pair_address_between_wfil_usdc_of_sushiswap_on_polygon"`
}

var polygonConfig *ConfigurationForPolygon

func initCofig() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		logs.GetLogger().Fatal("Cannot get home directory.")
	}

	configFile := filepath.Join(homedir, ".swan/mcp/config_polygon.toml")
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
		{"polygon_mainnet_node", "rpc_url"},
		{"polygon_mainnet_node", "payment_contract_address"},
		{"polygon_mainnet_node", "contract_lock_function_signature"},
		{"polygon_mainnet_node", "scan_step"},
		{"polygon_mainnet_node", "cycle_time_interval"},
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
