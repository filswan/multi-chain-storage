package service

import (
	"multi-chain-storage/common"
	"multi-chain-storage/common/constants"
	"runtime"
	"time"
)

func GetHostInfo() *common.HostInfo {
	hostInfo := common.HostInfo{
		Version:         constants.VERSION,
		OperatingSystem: runtime.GOOS,
		Architecture:    runtime.GOARCH,
		CpuNumber:       runtime.NumCPU(),
		CurrentUnixNano: time.Now().UnixNano(),
	}

	return &hostInfo
}
