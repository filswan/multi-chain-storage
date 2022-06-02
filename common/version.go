package common

type HostInfo struct {
	Version         string `json:"version"`
	OperatingSystem string `json:"operating_system"`
	Architecture    string `json:"architecture"`
	CpuNumber       int    `json:"cpu_number"`
	CurrentUnixNano int64  `json:"current_unix_nano"`
}
