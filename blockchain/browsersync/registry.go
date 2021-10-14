// Filesystem registry and backend options

package browsersync

import (
	"payment-bridge/blockchain/browsersync/bsc"
	"payment-bridge/blockchain/browsersync/goerli"
	"payment-bridge/blockchain/browsersync/nbai"
	"payment-bridge/blockchain/browsersync/polygon"
)

func Init() {
	bsc.Init()
	goerli.Init()
	nbai.Init()
	polygon.Init()

}
