package browsersync

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/nebulaai/nbai-node/common"
	"github.com/nebulaai/nbai-node/core/types"
	"math/big"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/scan/eth"
	"sort"
	"strconv"
	"time"
)

//var RewardPerBlock  float64 = 35

func GetRewardPerBlock() *big.Int {
	rewardBig, _ := new(big.Int).SetString("35000000000000000000", 10) // the unit is wei
	return rewardBig
}

func UpdateContractAddress() {
	transactions := []models.Transaction{}
	_, err := models.FindMultipleTransactions(&transactions, models.Transaction{TTo: " "})
	if err != nil {
		logs.GetLogger().Error(err)
	}
	for _, element := range transactions {
		receipt, _ := eth.WebConn.ConnWeb.TransactionReceipt(context.Background(), common.HexToHash(element.Hash))
		content := fmt.Sprintf(`contract address is = "%#x" in block`, receipt.ContractAddress)
		fmt.Println(content)
		fmt.Println(element.BlockNumber)
	}
}

func ScanWallet() { //THIS STEP SHOULD ALSO BE USED FOR CONTRACT CODE DETECTION!
	var firstId int64
	var secondId int64
	var startId int64 //this is the query start id. The necessity of this is to distinguish first scan from subsequent scans. This MAY be rendered more concise, just saying.
	var endId int64   //this is the query end id. The necessity of this+ is to distinguish first scan from subsequent scans.
	var firstScan = true
	for {
		firstId = secondId + 1 //firstId is the end id in last iteration+1
		if firstScan {
			startId = 0
		} else {
			startId = firstId
		}
		transaction := models.Transaction{}
		_, err := transaction.FindLastTxId()
		if err != nil {
			logs.GetLogger().Error(err)
		}
		secondId = transaction.ID //this selects the latest transaction id at this moment. It's also where we are supposed to begin the query next time
		if firstScan {
			endId = 9223372036854775807
			firstScan = false
		} else {
			endId = secondId
		}
		transactions := []models.Transaction{}
		_, err = models.FindTransactionsById(&transactions, startId, endId)
		if err != nil {
			logs.GetLogger().Error(err)
		}
		for _, element := range transactions {
			wallet := models.Wallet{WalletAddress: element.TFrom}
			_wallet, _ := wallet.FindOneWallet(&wallet)
			if _wallet.ID == 0 {
				record := models.Wallet{WalletAddress: element.TFrom, FirstTxID: strconv.FormatInt(element.ID, 10)}
				err := database.GetDB().Save(&record).Error
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
			wallet2 := models.Wallet{WalletAddress: element.TTo}
			_wallet2, _ := wallet2.FindOneWallet(&wallet2)
			if _wallet2.ID == 0 {
				record := models.Wallet{WalletAddress: element.TTo, FirstTxID: strconv.FormatInt(element.ID, 10)}
				database.GetDB().Save(&record)
				err := database.GetDB().Save(&record).Error
				if err != nil {
					logs.GetLogger().Error(err)
				}
			}
		}
		UpdateAndSortWallet()
		time.Sleep(time.Hour * 24)
	}
}

func UpdateAndSortWallet() {
	var wallets []models.Wallet
	_, err := models.FindWalletAddresseses(&wallets)
	if err != nil {
		logs.GetLogger().Error(err)
	}
	totalBalance := new(big.Int).SetUint64(0)
	models.GetSecondSlice().WalletWithBalances = make([]*models.WalletWithBalance, 0, len(wallets))
	for _, wallet := range wallets {
		balance, err := getBalance(wallet.WalletAddress)
		if err != nil {
			logs.GetLogger().Error(err)
			return
		}
		totalBalance = new(big.Int).Add(totalBalance, balance)
		models.GetSecondSlice().StoreArray(models.WalletWithBalance{WalletAddress: wallet.WalletAddress, Balance: balance})
	}
	totalBalanceFloat := new(big.Float).Quo(new(big.Float).SetInt(totalBalance), new(big.Float).SetUint64(100))
	sort.Sort(models.ByBalance(models.GetSecondSlice().WalletWithBalances))
	for _, element := range models.GetSecondSlice().WalletWithBalances {
		element.Percentage = new(big.Float).Quo(new(big.Float).SetInt(element.Balance), totalBalanceFloat).String()
	}

	models.GetFirstSlice().WalletWithBalances = make([]*models.WalletWithBalance, 0, len(wallets))
	models.GetFirstSlice().WalletWithBalances = models.GetSecondSlice().WalletWithBalances
}

func getBalance(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := eth.WebConn.ConnWeb.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return balance, err
	}
	return balance, err
}

func BlockBrowserSync() {

	for {

		block := models.Block{}

		_blockDB, err := block.FindLatestBlock()

		blockNoDB := _blockDB.Number
		if err != nil {
			logs.GetLogger().Error(err)
		}

		blockNoCurrent, err := eth.WebConn.GetBlockNumber()
		if err != nil {
			eth.ClientInit()
			logs.GetLogger().Error(err)
			continue
		}

		UpdateFromLastToCurrent(blockNoDB, blockNoCurrent.Int64())
		if err != nil {
			logs.GetLogger().Error(err)
		}

		time.Sleep(time.Second * 5)
	}
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

		if nextBlock.Transactions().Len() > 0 {
			err = SaveTransaction(nextBlock)
			if err != nil {
				logs.GetLogger().Error(err)
			}

		}

		uncles := nextBlock.Uncles()
		for _, uncle := range uncles {
			logs.GetLogger().Debug("uncle at ", nextBlock.Number())

			reward := calculateUncleReward(nextBlock, uncle)
			//fmt.Println("Reward =" + reward)

			uncleBlock := models.UncleBlock{

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
				Timestamp:            uncle.Time.String(),
				CanonicalBlockNumber: nextBlock.Number().Int64(),
				UncleBlockReward:     reward,
			}
			err = SaveUncleBlock(&uncleBlock)
			if err != nil {
				logs.GetLogger().Error(err)
			}

		}
		// check and save uncle nextBlock if exists
		//unclesHashes := GetUnclesHashes(nextBlock)
		//unclesHashesLength := len(unclesHashes)
		//if unclesHashesLength >= 1 {
		//	for _, singleUncleHash := range unclesHashes {
		//		uncleBlock, err := eth.WebConn.GetUncleBlockByHash(singleUncleHash)
		//		correspondingCanonicalBlockNumber := nextBlock.Number()
		//		if err != nil {
		//			logs.GetLogger().Error(err,"You might need to run the rpc node in full syc mode.")
		//			panic(err)
		//		}
		//		err = SaveUncleBlock(uncleBlock,correspondingCanonicalBlockNumber)
		//		if err != nil {
		//			logs.GetLogger().Error(err)
		//		}
		//	}
		//}

		lastBlockNo++

	}

}

func SaveBlock(block *types.Block) error {
	browserBlock := &models.Block{}
	browserBlock.Number = block.Number().Int64()
	browserBlock.Hash = block.Hash().String()
	browserBlock.ParentHash = block.ParentHash().String()
	browserBlock.Author = ""
	browserBlock.Miner = block.Coinbase().String()
	browserBlock.Size = block.Size().String()
	browserBlock.GasLimit = strconv.FormatUint(block.GasLimit(), 10)
	browserBlock.GasUsed = strconv.FormatUint(block.GasUsed(), 10)
	browserBlock.Nonce = strconv.FormatUint(block.Nonce(), 10)
	browserBlock.Timestamp = block.Time().String()
	browserBlock.Difficulty = block.Difficulty().String()

	_findblock, _ := browserBlock.FindOneBlock(browserBlock)
	if _findblock.ID == 0 {
		database.GetDB().Save(browserBlock)
	}

	return nil
}

func SaveUncleBlock(UncleBlock *models.UncleBlock) error {

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
		browserTx := &models.Transaction{}

		txMsg, err := tx.AsMessage(types.NewEIP155Signer(chainID))
		if err != nil {
			logs.GetLogger().Error(err)
		}

		browserTx.Hash = tx.Hash().String()
		browserTx.Nonce = strconv.FormatUint(txMsg.Nonce(), 10)
		browserTx.BlockNumber = block.Number().Int64()
		browserTx.TransactionIndex = strconv.Itoa(int(i))
		browserTx.TFrom = txMsg.From().String()

		browserTx.Timestamp = block.Time().String()

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
