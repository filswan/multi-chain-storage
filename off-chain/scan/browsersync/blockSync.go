package browsersync

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	common2 "payment-bridge/off-chain/common/utils"
	"payment-bridge/off-chain/database"
	"payment-bridge/off-chain/logs"
	models2 "payment-bridge/off-chain/models"
	"payment-bridge/off-chain/scan/eth"
	"strconv"
	"time"
)

func GetRewardPerBlock() *big.Int {
	rewardBig, _ := new(big.Int).SetString("35000000000000000000", 10) // the unit is wei
	return rewardBig
}

func BlockBrowserSyncAndEventLogsSync() {

	for {
		var lastCunrrentNumber int64
		blockScanRecord := models2.BlockScanRecord{}

		blockScanRecordList, err := blockScanRecord.FindLastCurrentBlockNumber()
		if err != nil {
			logs.GetLogger().Error(err)
			lastCunrrentNumber = 1
		}

		if len(blockScanRecordList) >= 1 {
			lastCunrrentNumber = blockScanRecordList[0].LastCurrentBlockNumber
		}

		blockNoCurrent, err := eth.WebConn.GetBlockNumber()
		if err != nil {
			eth.ClientInit()
			logs.GetLogger().Error(err)
			continue
		}

		err = ScanEventFromChainAndSaveEventLogData(lastCunrrentNumber, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		//Change to record the current block number in talbe block_scan_records
		/*UpdateFromLastToCurrent(blockNoDB, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
		}*/

		err = updateOrSaveBlockScanRecord(blockScanRecordList, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
			continue
		}

		time.Sleep(time.Second * 5)
	}
}

func updateOrSaveBlockScanRecord(blockScanRecordList []*models2.BlockScanRecord, lastCurrentBlockNumber int64) error {
	currentTime := strconv.FormatInt(common2.GetEpochInMillis(), 10)
	if len(blockScanRecordList) >= 1 {
		_, err := models2.UpdateBlockScanRecord("", map[string]interface{}{"update_at": currentTime, "last_current_block_number": lastCurrentBlockNumber})
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	} else {
		newRecord := new(models2.BlockScanRecord)
		newRecord.LastCurrentBlockNumber = lastCurrentBlockNumber
		newRecord.UpdateAt = string(currentTime)
		err := database.SaveOne(newRecord)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}
	return nil
}

func UpdateFromLastToCurrent(lastBlockNo, currentBlockNo int64) {

	for lastBlockNo < currentBlockNo {
		nextBlock, err := eth.WebConn.GetBlockByNumber(big.NewInt(lastBlockNo + 1))
		if err != nil {
			logs.GetLogger().Error(err)
			break
		}

		err = SaveBlock(nextBlock)
		if err != nil {
			logs.GetLogger().Error(err)
		}

		uncles := nextBlock.Uncles()
		for _, uncle := range uncles {
			logs.GetLogger().Debug("uncle at ", nextBlock.Number())

			reward := calculateUncleReward(nextBlock, uncle)
			//fmt.Println("Reward =" + reward)

			uncleBlock := models2.UncleBlock{

				Number:               uncle.Number.Int64(),
				Author:               "",
				Difficulty:           uncle.Difficulty.String(),
				ExtraData:            "0x" + hex.EncodeToString(uncle.Extra),
				GasLimit:             strconv.FormatUint(uncle.GasLimit, 10),
				GasUsed:              strconv.FormatUint(uncle.GasUsed, 10),
				Hash:                 uncle.Hash().String(),
				LogsBloom:            "0x" + hex.EncodeToString(uncle.Bloom.Bytes()),
				Miner:                uncle.Coinbase.String(),
				MixHash:              uncle.MixDigest.String(),
				Nonce:                strconv.FormatUint(uncle.Nonce.Uint64(), 10),
				ParentHash:           uncle.ParentHash.String(),
				Size:                 uncle.Size().String(),
				Timestamp:            string(uncle.Time),
				CanonicalBlockNumber: nextBlock.Number().Int64(),
				UncleBlockReward:     reward,
			}
			err = SaveUncleBlock(&uncleBlock)
			if err != nil {
				logs.GetLogger().Error(err)
			}

		}

		lastBlockNo++

	}

}

func SaveBlock(block *types.Block) error {
	browserBlock := &models2.Block{}
	browserBlock.Number = block.Number().Int64()
	browserBlock.Hash = block.Hash().String()
	browserBlock.ParentHash = block.ParentHash().String()
	browserBlock.Author = ""
	browserBlock.Miner = block.Coinbase().String()
	browserBlock.Size = block.Size().String()
	browserBlock.GasLimit = strconv.FormatUint(block.GasLimit(), 10)
	browserBlock.GasUsed = strconv.FormatUint(block.GasUsed(), 10)
	browserBlock.Nonce = strconv.FormatUint(block.Nonce(), 10)
	browserBlock.Timestamp = string(block.Time())
	browserBlock.Difficulty = block.Difficulty().String()

	_findblock, _ := browserBlock.FindOneBlock(browserBlock)
	if _findblock.ID == 0 {
		database.GetDB().Save(browserBlock)
	}

	return nil
}

func SaveUncleBlock(UncleBlock *models2.UncleBlock) error {

	_findblock, _ := UncleBlock.FindOneUncleBlock(UncleBlock)
	if _findblock.ID == 0 {
		err := database.SaveOne(UncleBlock)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func SaveTransaction(block *types.Block) error {

	chainID, err := eth.WebConn.ConnWeb.NetworkID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
	}

	for i, tx := range block.Transactions() {
		browserTx := &models2.Transaction{}

		txMsg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil)
		if err != nil {
			logs.GetLogger().Error(err)
		}

		browserTx.Hash = tx.Hash().String()
		browserTx.Nonce = strconv.FormatUint(txMsg.Nonce(), 10)
		browserTx.BlockNumber = block.Number().Int64()
		browserTx.TransactionIndex = strconv.Itoa(int(i))
		browserTx.TFrom = txMsg.From().String()

		browserTx.Timestamp = string(block.Time())

		if txMsg.To() != nil { //if this is a contract creation
			browserTx.TTo = txMsg.To().String()
		} else {
			browserTx.TTo = ""
			receipt, err := eth.WebConn.ConnWeb.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				logs.GetLogger().Error(err)
			}
			browserTx.ContractAddress = receipt.ContractAddress.String()
		}

		browserTx.Input = tx.Data()
		browserTx.Value = txMsg.Value().String()
		browserTx.GasPrice = txMsg.GasPrice().String()
		browserTx.Gas = strconv.FormatUint(txMsg.Gas(), 10)
		browserTx.BlockHash = block.Hash().String()
		_browserTx, _ := browserTx.FindOneTransaction(browserTx)
		if _browserTx.ID == 0 {
			database.GetDB().Save(browserTx)
		}

	}

	return nil
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func GetUnclesHashes(canonicalBlock *types.Block) []common.Hash {
	unclesLength := len(canonicalBlock.Uncles())
	var UnclesHashes []common.Hash

	for i := 0; i < unclesLength; i++ {
		fmt.Println(canonicalBlock.Number())
		singleUncle := canonicalBlock.Uncles()[i]
		singleUncleHash := singleUncle.Hash()
		//fmt.Println("\nGet one uncle block \nuncle number=",singleUncle.Number)
		//fmt.Println("hash =",singleUncleHash.String(),"\n")
		UnclesHashes = append(UnclesHashes, singleUncleHash)
	}

	return UnclesHashes
}

func calculateUncleReward(block *types.Block, singleuncle *types.Header) string {

	canonicalRewardBigInt := GetRewardPerBlock()
	blockNumberBigInt := block.Number()
	uncleBlockNumberBigInt := singleuncle.Number
	deltaBigInt := new(big.Int).Sub(blockNumberBigInt, uncleBlockNumberBigInt)
	//fmt.Println("deltaBigInt  =\t",deltaBigInt)
	deltaBigIntMulCReward := new(big.Int).Mul(canonicalRewardBigInt, deltaBigInt)
	//fmt.Println("deltaBigIntMulCReward  =\t",deltaBigIntMulCReward)
	factor1 := new(big.Int).Div(deltaBigIntMulCReward, big.NewInt(8))
	//fmt.Println("factor1  =\t\t",factor1)
	finalReward := new(big.Int).Sub(canonicalRewardBigInt, factor1)
	//fmt.Println("finalReward  =\t",finalReward)
	finalRewardString := finalReward.String()
	//fmt.Println("finalRewardString  =\t",finalRewardString)

	//fmt.Println("Canonical block number =",block.Number()," uncle block number=", singleuncle.Number, " uncleRewardString =",finalRewardString)
	//fmt.Println("\n\n\n\n")

	return finalRewardString
}
