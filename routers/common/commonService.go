package common

import (
	"payment-bridge/common"
	"payment-bridge/models"
	"runtime"

	"github.com/filswan/go-swan-lib/logs"
)

func getSwanMinerHostInfo() *common.HostInfo {
	info := new(common.HostInfo)
	info.SwanMinerVersion = common.GetVersion()
	info.OperatingSystem = runtime.GOOS
	info.Architecture = runtime.GOARCH
	info.CPUnNumber = runtime.NumCPU()
	return info
}

func getSystemConfigParams(limit string) ([]*models.SystemConfigParam, error) {
	sysconfigs, err := models.FindSystemConfigParam("", "id desc", limit, "0")
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	return sysconfigs, nil
}
