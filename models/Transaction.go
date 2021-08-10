package models

import (
	"errors"
	"payment-bridge/database"
	"strconv"
	"strings"
)

/**
 * created on 10/10/18.
 * author: nebula-ai-chengkun
 * Copyright defined in blockchainwebbrowser/LICENSE.txt
 */

type Transaction struct {
	ID               int64  `json:"id"`
	Hash             string `json:"hash"`
	BlockHash        string `json:"block_hash"`
	Creates          string `json:"creates"`
	TFrom            string `json:"t_from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gas_price"`
	Input            []byte `json:"input"`
	Nonce            string `json:"nonce"`
	PublicKey        string `json:"public_key"`
	R                string `json:"r"`
	Raw              string `json:"raw"`
	S                string `json:"s"`
	TTo              string `json:"t_to"`
	TransactionIndex string `json:"transaction_index"`
	V                int64  `json:"v"`
	Value            string `json:"value"`
	BlockNumber      int64  `json:"block_number"`
	Timestamp        string `json:"timestamp"`
	ContractAddress  string `json:"contract_address"`

	/*
		type:	(within an address slice)
			1 - is a 'from' address (make tx)
			2 - is a 'to' address	(get tx)
			3 - is both
	*/
	Type int `json:"type" gorm:"-"`
	//BlockNumberRaw string `json:"block_number_raw"`
	//ChainId string `json:"chain_id"`
	//GasPriceRaw string `json:"gas_price_raw"`
	//GasRaw string `json:"gas_raw"`
	//NonceRaw string `json:"nonce_raw"`
	//TransactionIndexRaw string `json:"transaction_index_raw"`
	//ValueRaw string `json:"value_raw"`
}

func (self *Transaction) FindOneTransaction(condition interface{}) (*Transaction, error) {
	db := database.GetDB()
	tx := db.Begin()
	tx.Where(condition).First(&self)
	err := tx.Commit().Error
	return self, err
}

func FindMultipleTransactions(transactions *[]Transaction, condition interface{}) (*[]Transaction, error) {
	db := database.GetDB()
	//tx := db.Begin()
	err := db.Where(condition).Find(&transactions).Error
	//err := tx.Commit().Error
	return transactions, err
}

func (self *Transaction) FindLastTxId() (*Transaction, error) {
	db := database.GetDB()
	err := db.Select("id").Last(&self).Error
	return self, err
}

func FindTransactionsById(transactions *[]Transaction, startId int64, endId int64) (*[]Transaction, error) {
	db := database.GetDB()
	err := db.Where("transaction.id between ? and ?", startId, endId).Find(&transactions).Error
	return transactions, err
}

func FindWalletTransactionByWalletAddr(walletAddr, limit, offset string) ([]Transaction, int, error, int) { //the last int returned represents error situation
	db := database.GetDB()
	var models []Transaction
	var hasOffset bool = len(offset) > 0
	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		return models, 0, err, 1
	}
	offset_int, err := strconv.Atoi(offset)
	if err != nil && hasOffset {
		return models, 0, err, 2
	}
	count := 0
	db.Order("id desc").Where("t_to = ? or t_from = ?", walletAddr, walletAddr).Find(&models).Count(&count)
	if offset_int >= count && hasOffset {
		return models, 0, errors.New("offset exceeds count"), 3
	}
	db.Order("id desc").Where("t_to = ? or t_from = ?", walletAddr, walletAddr).Offset(offset_int).Limit(limit_int).Find(&models)
	return models, count, err, 4
}

func FindManyTransactions(limit, offset string) ([]*Transaction, error) {
	dbHandler := database.GetDB()
	var models []*Transaction

	offset_int, err := strconv.Atoi(offset)
	if err != nil {
		offset_int = 0
	}

	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		limit_int = 10
	}
	tx := dbHandler.Begin()
	dbHandler.Order("id desc").Offset(offset_int).Limit(limit_int).Find(&models)
	for i := range models {
		tx.Model(&models[i])
	}
	err = tx.Commit().Error
	return models, err
}

func FindTransactionsByMultipleWallet(addresses string, limit string, offset string) ([]*Transaction, int, error, int) {

	//Make a filter to ensure the legitablity of addresses (i.e.: addresses can easily be divided into strings)
	db := database.GetDB()
	var models []*Transaction

	addressSlice := strings.Split(addresses, ",")
	for _, element := range addressSlice {
		if !IsValidAddress(element) {
			return models, 0, errors.New("At least one of the addresses input is invalid"), 4
		}
		//&element = "'"+element+"'"
	}

	var hasOffset bool = len(offset) > 0
	limit_int, err := strconv.Atoi(limit)
	if err != nil {
		return models, 0, err, 1
	}
	offset_int, err := strconv.Atoi(offset)
	if err != nil && hasOffset {
		return models, 0, err, 2
	}

	//var replacer = strings.NewReplacer(",", "','")
	//addresses = replacer.Replace(addresses) //this step is to give the correct syntax to SQL query

	count := 0
	db.Order("id desc").Where("t_to in (?) or t_from in (?)", addressSlice, addressSlice).Find(&models).Count(&count)
	if offset_int >= count && hasOffset {
		return models, 0, errors.New("Offset exceeds count"), 3
	}
	db.Order("id desc").Where("t_to in (?) or t_from in (?)", addressSlice, addressSlice).Offset(offset_int).Limit(limit_int).Find(&models)
	for _, tx := range models {
		isFrom := false
		isTo := false
		for _, address := range addressSlice {
			if strings.EqualFold(tx.TFrom, address) {
				isFrom = true
			}
			if strings.EqualFold(tx.TTo, address) {
				isTo = true
			}
		}
		if isTo && isFrom {
			tx.Type = 3
		} else {
			if isFrom {
				tx.Type = 1
			}
			if isTo {
				tx.Type = 2
			}
		}
	}
	return models, count, err, 5

}
