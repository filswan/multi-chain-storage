package nbai

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"os"
	"payment-bridge/blockchain/initclient/bscclient"
	"payment-bridge/common/constants"
	"payment-bridge/common/utils"
	"payment-bridge/config"
	"payment-bridge/database"
	"payment-bridge/logs"
	"payment-bridge/models"
	"payment-bridge/on-chain/goBind"
	"strconv"
	"strings"
)

func ChangeNbaiToBnb(data []byte, txHashInNbai string, blockNo uint64, childChainTractionID int64) error {
	pk := os.Getenv("privateKey")
	fromAddress := common.HexToAddress(config.GetConfig().BscMainnetNode.BscAdminWallet)
	client := bscclient.WebConn.ConnWeb
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if strings.HasPrefix(strings.ToLower(pk), "0x") {
		pk = pk[2:]
	}

	privateKey, _ := crypto.HexToECDSA(pk)
	callOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(config.GetConfig().BscMainnetNode.ChainID))

	//callOpts := new(bind.TransactOpts)
	callOpts.Nonce = big.NewInt(int64(nonce))
	callOpts.GasPrice = gasPrice
	callOpts.GasLimit = config.GetConfig().BscMainnetNode.GasLimit
	callOpts.Context = context.Background()

	childManagerAddress := common.HexToAddress(config.GetConfig().BscMainnetNode.ChildChainManageContractAddress) //to config：想要调用的合约地址
	childInstance, _ := goBind.NewChildChainManagerContract(childManagerAddress, client)

	childChainTX := new(models.ChildChainTransaction)
	tx, err := childInstance.OnStateReceive(callOpts, big.NewInt(int64(1)), data)
	if err != nil {
		logs.GetLogger().Error(err)
		childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
		childChainTX.TxHashInBsc = ""
		childChainTX.Status = constants.TRANSACTION_STATUS_FAIL
	} else {
		childChainTX.TxHashInBsc = tx.Hash().Hex()
		childChainTX.Status = constants.TRANSACTION_STATUS_SUCCESS
	}

	if childChainTractionID > 0 {
		childChainTX.ID = childChainTractionID
	}

	childChainTX.TxHashInNbai = txHashInNbai
	childChainTX.BlockNo = blockNo
	currenTime := utils.GetEpochInMillis()
	childChainTX.CreateAt = strconv.FormatInt(currenTime, 10)
	childChainTX.UpdateAt = strconv.FormatInt(currenTime, 10)
	err = database.SaveOne(childChainTX)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}
