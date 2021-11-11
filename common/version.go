package common

import "fmt"

const (
	MajorVersion = 2
	MinorVersion = 5
	FixVersion   = 0
	CommitHash   = ""
)

func GetVersion() string {
	if CommitHash != "" {
		return fmt.Sprintf("swan-miner-v%v.%v.%v-%s", MajorVersion, MinorVersion, FixVersion, CommitHash)
	} else {
		return fmt.Sprintf("swan-miner-v%v.%v.%v", MajorVersion, MinorVersion, FixVersion)
	}
}

type HostInfo struct {
	SwanMinerVersion string `json:"swan_miner_version""`
	OperatingSystem  string `json:"operating_system"`
	Architecture     string `json:"architecture"`
	CPUnNumber       int    `json:"cpu_number"`
}
