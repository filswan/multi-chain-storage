package models

import "github.com/filswan/go-swan-lib/logs"

var Registry = make(map[string]*RegInfo)

type RegInfo struct {
	Network                           string
	CoinName                          string
	ScanEventFromChainAndSaveDataToDb func() `json:"-"`
}

func Register(info *RegInfo) {
	Registry[info.Network] = info
}

func RunAllTheScan() {
	logs.GetLogger().Info(Registry)
	for _, v := range Registry {
		go v.ScanEventFromChainAndSaveDataToDb()
	}
}
