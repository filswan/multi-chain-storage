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

func ChangeNbaiToBnb(data []byte) {
	pk := os.Getenv("privateKey")
	fromAddress := common.HexToAddress(config.GetConfig().BscMainnetNode.BscAdminWallet)
	client := bscclient.WebConn.ConnWeb
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		logs.GetLogger().Error(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
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
	}

	childChainTX.TxHash = tx.Hash().Hex()
	childChainTX.CreateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
	childChainTX.UpdateAt = strconv.FormatInt(utils.GetEpochInMillis(), 10)
	childChainTX.Status = constants.TRANSACTION_STATUS_SUCCESS

	database.SaveOne(childChainTX)

}
