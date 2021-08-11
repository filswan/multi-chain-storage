package eth

import (
	"fmt"
	"github.com/nebulaai/nbai-node/common"
	"payment-bridge/logs"
	"testing"
)

func TestUncleBlock(t *testing.T) {

	//var testCanonicalBlockHash common.Hash
	//testCanonicalBlockHash = "0xa8313f5580e8f89edf97c77294ab35589cc6dc60ac2f5c40efaeb09e6154a4b5")
	//nextBlock := webService.GetBlockByNumber(599033)
	//nextBlock := &models.Block{}
	//nextBlock.Number = 599033
	////var nextBlock models.Block{
	////	nextBlock.Number :=
	////}
	//
	//nextBlock.FindOneBlock(nextBlock)
	//
	//if browsersync.HasUncleBlock(nextBlock) {
	// for id = 940652, number = 599032, sha3uncles = 0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347

	ClientInit()

	//testsha3uncles := "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347"
	testsha3uncles := "0x03446ff62b20b217f68a613fe434946e4bcc0ada36e9190de72158fc98e07b51"
	testhash := common.HexToHash(testsha3uncles)
	fmt.Println("testhash=", testhash)

	uncleBlock, err := WebConn.GetUncleBlockByHash(testhash)
	if err != nil {
		logs.GetLogger().Error(err)
	}
	fmt.Println(uncleBlock.Hash())
	//err = browsersync.SaveUncleBlock(uncleBlock)
}

//}
