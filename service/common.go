package service

import (
	"multi-chain-storage/common"
	"multi-chain-storage/models"
	"runtime"

	"github.com/filswan/go-swan-lib/logs"
)

func GetHostInfo() *common.HostInfo {
	hostInfo := common.HostInfo{
		Version:         common.GetVersion(),
		OperatingSystem: runtime.GOOS,
		Architecture:    runtime.GOARCH,
		CpuNumber:       runtime.NumCPU(),
	}

	return &hostInfo
}

func GetSystemConfigParams(limit string) ([]*models.SystemConfigParam, error) {
	sysconfigs, err := models.GetSystemConfigParams()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return sysconfigs, nil
}
