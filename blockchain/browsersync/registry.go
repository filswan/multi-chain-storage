// Filesystem registry and backend options

package browsersync

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
	for _, v := range Registry {
		v.ScanEventFromChainAndSaveDataToDb()
	}
}
