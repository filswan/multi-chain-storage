// Filesystem registry and backend options

package browsersync

import (
	"payment-bridge/blockchain/browsersync/bsc"
	"payment-bridge/blockchain/browsersync/goerli"
)

func Init() {
	bsc.Init()
	goerli.Init()
}
