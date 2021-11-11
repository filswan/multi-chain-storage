package scanFactory

import (
	"payment-bridge/blockchain/browsersync/scanlockpayment/bsc"
	"payment-bridge/blockchain/browsersync/scanlockpayment/goerli"
	"payment-bridge/blockchain/browsersync/scanlockpayment/nbai"
	"payment-bridge/blockchain/browsersync/scanlockpayment/polygon"
	"payment-bridge/common/constants"
)

type IEventScanFactory struct {
}

type BlockChainNetwork interface {
	ScanEventFromChainAndSaveDataToDb()
}

type BscBlockChainNetwork struct {
}

func (b BscBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	bsc.ScanEventFromChainAndSaveDataToDbForBsc()
}

type GoerliBlockChainNetwork struct {
}

func (g GoerliBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	goerli.ScanEventFromChainAndSaveDataToDbForGoerli()
}

type NBAIBlockChainNetwork struct {
}

func (n NBAIBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	nbai.ScanEventFromChainAndSaveDataToDbForNBAI()
}

type PolygonBlockChainNetwork struct {
}

func (p PolygonBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	polygon.ScanEventFromChainAndSaveDataToDbForPolygon()
}

func (*IEventScanFactory) GenerateBlockChainNetwork(network string) BlockChainNetwork {
	switch network {
	case constants.NETWORK_TYPE_BSC:
		return BscBlockChainNetwork{}
	case constants.NETWORK_TYPE_GOERLI:
		return GoerliBlockChainNetwork{}
	case constants.NETWORK_TYPE_NBAI:
		return NBAIBlockChainNetwork{}
	case constants.NETWORK_TYPE_POLYGON:
		return PolygonBlockChainNetwork{}
	}
	return nil
}
