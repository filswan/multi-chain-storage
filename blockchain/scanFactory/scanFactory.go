package scanFactory

import (
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/bsc"
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/goerli"
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/nbai"
	"multi-chain-storage/blockchain/browsersync/scanlockpayment/polygon"
	"multi-chain-storage/common/constants"
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
