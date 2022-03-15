// Filesystem registry and backend options

package browsersync

import (
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/polygon"
)

func Init() {
	//bsc.Init()
	//goerli.Init()
	//nbai.Init()
	polygon.Init()

}
