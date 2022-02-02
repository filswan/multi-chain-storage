package common

import (
	"payment-bridge/common"
	"payment-bridge/models"
	"runtime"

	"github.com/filswan/go-swan-lib/logs"
)

func getHostInfo() *common.HostInfo {
	hostInfo := common.HostInfo{
		Version:         common.GetVersion(),
		OperatingSystem: runtime.GOOS,
		Architecture:    runtime.GOARCH,
		CpuNumber:       runtime.NumCPU(),
	}

	return &hostInfo
}

func getSystemConfigParams(limit string) ([]*models.SystemConfigParam, error) {
	sysconfigs, err := models.GetSystemConfigParams(limit, "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return sysconfigs, nil
}
