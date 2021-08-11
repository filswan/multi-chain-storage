package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"payment-bridge/off-chain/logs"
)

/**
 * created on 10/10/18.
 * author: nebula-ai-chengkun
 * Copyright defined in blockchainwebbrowser/LICENSE.txt
 */

func (conn *ConnSetup) GetBlockNumber() (*big.Int, error) {
	//return conn.ConnWeb.Eth.GetBlockNumber()
	block, err := conn.ConnWeb.HeaderByNumber(context.Background(), nil)
	if block != nil {
		return block.Number, err
	}
	return nil, err
}

func (conn *ConnSetup) GetBlockByNumber(blockNo *big.Int) (*types.Block, error) {
	block, err := conn.ConnWeb.BlockByNumber(context.Background(), blockNo)
	logs.GetLogger().Debug(blockNo)
	return block, err
}

func (conn *ConnSetup) GetUncleBlockByHash(blockHash common.Hash) (*types.Block, error) {
	block, err := conn.ConnWeb.BlockByHash(context.Background(), blockHash)
	logs.GetLogger().Debug(blockHash.String())

	return block, err
}

func (conn *ConnSetup) GetTransactionByBlockHashAndIndex(blockHash common.Hash, index uint) (*types.Transaction, error) {
	tx, err := conn.ConnWeb.TransactionInBlock(context.Background(), blockHash, index)
	return tx, err
}

func (conn *ConnSetup) GetTransactionCountByBlockHash(blockHash common.Hash) (uint, error) {
	txCount, err := conn.ConnWeb.TransactionCount(context.Background(), blockHash)
	return txCount, err
}
