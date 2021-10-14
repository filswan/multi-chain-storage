package models

import "fmt"

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
	fmt.Println(Registry)
	for _, v := range Registry {
		go v.ScanEventFromChainAndSaveDataToDb()
	}
}
