// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goBind

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FilswanOracleABI is the input ABI used to generate the binding from.
const FilswanOracleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"getPaymentInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"actualPaid\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"daoUsers\",\"type\":\"address[]\"}],\"name\":\"setDAOUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"actualPaid\",\"type\":\"uint256\"}],\"name\":\"updatePaymentInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// FilswanOracle is an auto generated Go binding around an Ethereum contract.
type FilswanOracle struct {
	FilswanOracleCaller     // Read-only binding to the contract
	FilswanOracleTransactor // Write-only binding to the contract
	FilswanOracleFilterer   // Log filterer for contract events
}

// FilswanOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type FilswanOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilswanOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FilswanOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilswanOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FilswanOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilswanOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FilswanOracleSession struct {
	Contract     *FilswanOracle    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FilswanOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FilswanOracleCallerSession struct {
	Contract *FilswanOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// FilswanOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FilswanOracleTransactorSession struct {
	Contract     *FilswanOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// FilswanOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type FilswanOracleRaw struct {
	Contract *FilswanOracle // Generic contract binding to access the raw methods on
}

// FilswanOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FilswanOracleCallerRaw struct {
	Contract *FilswanOracleCaller // Generic read-only contract binding to access the raw methods on
}

// FilswanOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FilswanOracleTransactorRaw struct {
	Contract *FilswanOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFilswanOracle creates a new instance of FilswanOracle, bound to a specific deployed contract.
func NewFilswanOracle(address common.Address, backend bind.ContractBackend) (*FilswanOracle, error) {
	contract, err := bindFilswanOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FilswanOracle{FilswanOracleCaller: FilswanOracleCaller{contract: contract}, FilswanOracleTransactor: FilswanOracleTransactor{contract: contract}, FilswanOracleFilterer: FilswanOracleFilterer{contract: contract}}, nil
}

// NewFilswanOracleCaller creates a new read-only instance of FilswanOracle, bound to a specific deployed contract.
func NewFilswanOracleCaller(address common.Address, caller bind.ContractCaller) (*FilswanOracleCaller, error) {
	contract, err := bindFilswanOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleCaller{contract: contract}, nil
}

// NewFilswanOracleTransactor creates a new write-only instance of FilswanOracle, bound to a specific deployed contract.
func NewFilswanOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*FilswanOracleTransactor, error) {
	contract, err := bindFilswanOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleTransactor{contract: contract}, nil
}

// NewFilswanOracleFilterer creates a new log filterer instance of FilswanOracle, bound to a specific deployed contract.
func NewFilswanOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*FilswanOracleFilterer, error) {
	contract, err := bindFilswanOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleFilterer{contract: contract}, nil
}

// bindFilswanOracle binds a generic wrapper to an already deployed contract.
func bindFilswanOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FilswanOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilswanOracle *FilswanOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilswanOracle.Contract.FilswanOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilswanOracle *FilswanOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilswanOracle.Contract.FilswanOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilswanOracle *FilswanOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilswanOracle.Contract.FilswanOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilswanOracle *FilswanOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilswanOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilswanOracle *FilswanOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilswanOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilswanOracle *FilswanOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilswanOracle.Contract.contract.Transact(opts, method, params...)
}

// GetPaymentInfo is a free data retrieval call binding the contract method 0x018fce0e.
//
// Solidity: function getPaymentInfo(string txId) view returns(uint256 actualPaid)
func (_FilswanOracle *FilswanOracleCaller) GetPaymentInfo(opts *bind.CallOpts, txId string) (*big.Int, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getPaymentInfo", txId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentInfo is a free data retrieval call binding the contract method 0x018fce0e.
//
// Solidity: function getPaymentInfo(string txId) view returns(uint256 actualPaid)
func (_FilswanOracle *FilswanOracleSession) GetPaymentInfo(txId string) (*big.Int, error) {
	return _FilswanOracle.Contract.GetPaymentInfo(&_FilswanOracle.CallOpts, txId)
}

// GetPaymentInfo is a free data retrieval call binding the contract method 0x018fce0e.
//
// Solidity: function getPaymentInfo(string txId) view returns(uint256 actualPaid)
func (_FilswanOracle *FilswanOracleCallerSession) GetPaymentInfo(txId string) (*big.Int, error) {
	return _FilswanOracle.Contract.GetPaymentInfo(&_FilswanOracle.CallOpts, txId)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 threshold) returns()
func (_FilswanOracle *FilswanOracleTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "initialize", owner, threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 threshold) returns()
func (_FilswanOracle *FilswanOracleSession) Initialize(owner common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Initialize(&_FilswanOracle.TransactOpts, owner, threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address owner, uint8 threshold) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) Initialize(owner common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Initialize(&_FilswanOracle.TransactOpts, owner, threshold)
}

// SetDAOUsers is a paid mutator transaction binding the contract method 0x6bcf35f9.
//
// Solidity: function setDAOUsers(address[] daoUsers) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) SetDAOUsers(opts *bind.TransactOpts, daoUsers []common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "setDAOUsers", daoUsers)
}

// SetDAOUsers is a paid mutator transaction binding the contract method 0x6bcf35f9.
//
// Solidity: function setDAOUsers(address[] daoUsers) returns(bool)
func (_FilswanOracle *FilswanOracleSession) SetDAOUsers(daoUsers []common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetDAOUsers(&_FilswanOracle.TransactOpts, daoUsers)
}

// SetDAOUsers is a paid mutator transaction binding the contract method 0x6bcf35f9.
//
// Solidity: function setDAOUsers(address[] daoUsers) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) SetDAOUsers(daoUsers []common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetDAOUsers(&_FilswanOracle.TransactOpts, daoUsers)
}

// UpdatePaymentInfo is a paid mutator transaction binding the contract method 0xa7dddedb.
//
// Solidity: function updatePaymentInfo(string txId, uint256 actualPaid) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) UpdatePaymentInfo(opts *bind.TransactOpts, txId string, actualPaid *big.Int) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "updatePaymentInfo", txId, actualPaid)
}

// UpdatePaymentInfo is a paid mutator transaction binding the contract method 0xa7dddedb.
//
// Solidity: function updatePaymentInfo(string txId, uint256 actualPaid) returns(bool)
func (_FilswanOracle *FilswanOracleSession) UpdatePaymentInfo(txId string, actualPaid *big.Int) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UpdatePaymentInfo(&_FilswanOracle.TransactOpts, txId, actualPaid)
}

// UpdatePaymentInfo is a paid mutator transaction binding the contract method 0xa7dddedb.
//
// Solidity: function updatePaymentInfo(string txId, uint256 actualPaid) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) UpdatePaymentInfo(txId string, actualPaid *big.Int) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UpdatePaymentInfo(&_FilswanOracle.TransactOpts, txId, actualPaid)
}
