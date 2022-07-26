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

// FilswanOracleTxOracleInfo is an auto generated low-level Go binding around an user-defined struct.
type FilswanOracleTxOracleInfo struct {
	Paid         *big.Int
	Terms        *big.Int
	Recipient    common.Address
	Status       bool
	Flag         bool
	CidList      []string
	Signer       common.Address
	Timestamp    *big.Int
	BlockNumber  *big.Int
	SignStatus   *big.Int
	Batch        uint8
	BatchCidList [256][]string
}

// FilswanOracleABI is the input ABI used to generate the binding from.
const FilswanOracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"batchCount\",\"type\":\"uint8\"}],\"name\":\"PreSign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"batchNo\",\"type\":\"uint8\"}],\"name\":\"Sign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SignCarTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"voteKey\",\"type\":\"bytes32\"}],\"name\":\"SignHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"SignTransaction\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAO_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"s1\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"s2\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"a1\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"sa\",\"type\":\"string[]\"}],\"name\":\"f\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"getCarPaymentVotes\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"}],\"name\":\"getCidList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"}],\"name\":\"getHashKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getOracleInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"terms\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signStatus\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"batch\",\"type\":\"uint8\"},{\"internalType\":\"string[][256]\",\"name\":\"batchCidList\",\"type\":\"string[][256]\"}],\"internalType\":\"structFilswanOracle.TxOracleInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"}],\"name\":\"getSignatureList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"terms\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signStatus\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"batch\",\"type\":\"uint8\"},{\"internalType\":\"string[][256]\",\"name\":\"batchCidList\",\"type\":\"string[][256]\"}],\"internalType\":\"structFilswanOracle.TxOracleInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"isCarPaymentAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"batchCount\",\"type\":\"uint8\"}],\"name\":\"preSign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"daoUsers\",\"type\":\"address[]\"}],\"name\":\"setDAOUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"filinkAddress\",\"type\":\"address\"}],\"name\":\"setFilinkOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"},{\"internalType\":\"uint8\",\"name\":\"batchNo\",\"type\":\"uint8\"}],\"name\":\"sign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"signCarTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"voteKey\",\"type\":\"bytes32\"}],\"name\":\"signHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"updateThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// DAOROLE is a free data retrieval call binding the contract method 0xe9c26518.
//
// Solidity: function DAO_ROLE() view returns(bytes32)
func (_FilswanOracle *FilswanOracleCaller) DAOROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "DAO_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAOROLE is a free data retrieval call binding the contract method 0xe9c26518.
//
// Solidity: function DAO_ROLE() view returns(bytes32)
func (_FilswanOracle *FilswanOracleSession) DAOROLE() ([32]byte, error) {
	return _FilswanOracle.Contract.DAOROLE(&_FilswanOracle.CallOpts)
}

// DAOROLE is a free data retrieval call binding the contract method 0xe9c26518.
//
// Solidity: function DAO_ROLE() view returns(bytes32)
func (_FilswanOracle *FilswanOracleCallerSession) DAOROLE() ([32]byte, error) {
	return _FilswanOracle.Contract.DAOROLE(&_FilswanOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FilswanOracle *FilswanOracleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FilswanOracle *FilswanOracleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FilswanOracle.Contract.DEFAULTADMINROLE(&_FilswanOracle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FilswanOracle *FilswanOracleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FilswanOracle.Contract.DEFAULTADMINROLE(&_FilswanOracle.CallOpts)
}

// GetCarPaymentVotes is a free data retrieval call binding the contract method 0x9c8ecdf3.
//
// Solidity: function getCarPaymentVotes(string dealId, string network, address recipient) view returns(uint8)
func (_FilswanOracle *FilswanOracleCaller) GetCarPaymentVotes(opts *bind.CallOpts, dealId string, network string, recipient common.Address) (uint8, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getCarPaymentVotes", dealId, network, recipient)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCarPaymentVotes is a free data retrieval call binding the contract method 0x9c8ecdf3.
//
// Solidity: function getCarPaymentVotes(string dealId, string network, address recipient) view returns(uint8)
func (_FilswanOracle *FilswanOracleSession) GetCarPaymentVotes(dealId string, network string, recipient common.Address) (uint8, error) {
	return _FilswanOracle.Contract.GetCarPaymentVotes(&_FilswanOracle.CallOpts, dealId, network, recipient)
}

// GetCarPaymentVotes is a free data retrieval call binding the contract method 0x9c8ecdf3.
//
// Solidity: function getCarPaymentVotes(string dealId, string network, address recipient) view returns(uint8)
func (_FilswanOracle *FilswanOracleCallerSession) GetCarPaymentVotes(dealId string, network string, recipient common.Address) (uint8, error) {
	return _FilswanOracle.Contract.GetCarPaymentVotes(&_FilswanOracle.CallOpts, dealId, network, recipient)
}

// GetCidList is a free data retrieval call binding the contract method 0x1ce1eaf9.
//
// Solidity: function getCidList(string dealId, string network) view returns(string[])
func (_FilswanOracle *FilswanOracleCaller) GetCidList(opts *bind.CallOpts, dealId string, network string) ([]string, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getCidList", dealId, network)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetCidList is a free data retrieval call binding the contract method 0x1ce1eaf9.
//
// Solidity: function getCidList(string dealId, string network) view returns(string[])
func (_FilswanOracle *FilswanOracleSession) GetCidList(dealId string, network string) ([]string, error) {
	return _FilswanOracle.Contract.GetCidList(&_FilswanOracle.CallOpts, dealId, network)
}

// GetCidList is a free data retrieval call binding the contract method 0x1ce1eaf9.
//
// Solidity: function getCidList(string dealId, string network) view returns(string[])
func (_FilswanOracle *FilswanOracleCallerSession) GetCidList(dealId string, network string) ([]string, error) {
	return _FilswanOracle.Contract.GetCidList(&_FilswanOracle.CallOpts, dealId, network)
}

// GetHashKey is a free data retrieval call binding the contract method 0xfcd797a1.
//
// Solidity: function getHashKey(string dealId, string network, address recipient, string[] cidList) view returns(bytes32)
func (_FilswanOracle *FilswanOracleCaller) GetHashKey(opts *bind.CallOpts, dealId string, network string, recipient common.Address, cidList []string) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getHashKey", dealId, network, recipient, cidList)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetHashKey is a free data retrieval call binding the contract method 0xfcd797a1.
//
// Solidity: function getHashKey(string dealId, string network, address recipient, string[] cidList) view returns(bytes32)
func (_FilswanOracle *FilswanOracleSession) GetHashKey(dealId string, network string, recipient common.Address, cidList []string) ([32]byte, error) {
	return _FilswanOracle.Contract.GetHashKey(&_FilswanOracle.CallOpts, dealId, network, recipient, cidList)
}

// GetHashKey is a free data retrieval call binding the contract method 0xfcd797a1.
//
// Solidity: function getHashKey(string dealId, string network, address recipient, string[] cidList) view returns(bytes32)
func (_FilswanOracle *FilswanOracleCallerSession) GetHashKey(dealId string, network string, recipient common.Address, cidList []string) ([32]byte, error) {
	return _FilswanOracle.Contract.GetHashKey(&_FilswanOracle.CallOpts, dealId, network, recipient, cidList)
}

// GetOracleInfo is a free data retrieval call binding the contract method 0x36c1106f.
//
// Solidity: function getOracleInfo(string dealId, string network, address sender) view returns((uint256,uint256,address,bool,bool,string[],address,uint256,uint256,uint256,uint8,string[][256]))
func (_FilswanOracle *FilswanOracleCaller) GetOracleInfo(opts *bind.CallOpts, dealId string, network string, sender common.Address) (FilswanOracleTxOracleInfo, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getOracleInfo", dealId, network, sender)

	if err != nil {
		return *new(FilswanOracleTxOracleInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(FilswanOracleTxOracleInfo)).(*FilswanOracleTxOracleInfo)

	return out0, err

}

// GetOracleInfo is a free data retrieval call binding the contract method 0x36c1106f.
//
// Solidity: function getOracleInfo(string dealId, string network, address sender) view returns((uint256,uint256,address,bool,bool,string[],address,uint256,uint256,uint256,uint8,string[][256]))
func (_FilswanOracle *FilswanOracleSession) GetOracleInfo(dealId string, network string, sender common.Address) (FilswanOracleTxOracleInfo, error) {
	return _FilswanOracle.Contract.GetOracleInfo(&_FilswanOracle.CallOpts, dealId, network, sender)
}

// GetOracleInfo is a free data retrieval call binding the contract method 0x36c1106f.
//
// Solidity: function getOracleInfo(string dealId, string network, address sender) view returns((uint256,uint256,address,bool,bool,string[],address,uint256,uint256,uint256,uint8,string[][256]))
func (_FilswanOracle *FilswanOracleCallerSession) GetOracleInfo(dealId string, network string, sender common.Address) (FilswanOracleTxOracleInfo, error) {
	return _FilswanOracle.Contract.GetOracleInfo(&_FilswanOracle.CallOpts, dealId, network, sender)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FilswanOracle *FilswanOracleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FilswanOracle *FilswanOracleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FilswanOracle.Contract.GetRoleAdmin(&_FilswanOracle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FilswanOracle *FilswanOracleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FilswanOracle.Contract.GetRoleAdmin(&_FilswanOracle.CallOpts, role)
}

// GetSignatureList is a free data retrieval call binding the contract method 0x0ba5a924.
//
// Solidity: function getSignatureList(string dealId, string network) view returns((uint256,uint256,address,bool,bool,string[],address,uint256,uint256,uint256,uint8,string[][256])[])
func (_FilswanOracle *FilswanOracleCaller) GetSignatureList(opts *bind.CallOpts, dealId string, network string) ([]FilswanOracleTxOracleInfo, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getSignatureList", dealId, network)

	if err != nil {
		return *new([]FilswanOracleTxOracleInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]FilswanOracleTxOracleInfo)).(*[]FilswanOracleTxOracleInfo)

	return out0, err

}

// GetSignatureList is a free data retrieval call binding the contract method 0x0ba5a924.
//
// Solidity: function getSignatureList(string dealId, string network) view returns((uint256,uint256,address,bool,bool,string[],address,uint256,uint256,uint256,uint8,string[][256])[])
func (_FilswanOracle *FilswanOracleSession) GetSignatureList(dealId string, network string) ([]FilswanOracleTxOracleInfo, error) {
	return _FilswanOracle.Contract.GetSignatureList(&_FilswanOracle.CallOpts, dealId, network)
}

// GetSignatureList is a free data retrieval call binding the contract method 0x0ba5a924.
//
// Solidity: function getSignatureList(string dealId, string network) view returns((uint256,uint256,address,bool,bool,string[],address,uint256,uint256,uint256,uint8,string[][256])[])
func (_FilswanOracle *FilswanOracleCallerSession) GetSignatureList(dealId string, network string) ([]FilswanOracleTxOracleInfo, error) {
	return _FilswanOracle.Contract.GetSignatureList(&_FilswanOracle.CallOpts, dealId, network)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_FilswanOracle *FilswanOracleCaller) GetThreshold(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getThreshold")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_FilswanOracle *FilswanOracleSession) GetThreshold() (uint8, error) {
	return _FilswanOracle.Contract.GetThreshold(&_FilswanOracle.CallOpts)
}

// GetThreshold is a free data retrieval call binding the contract method 0xe75235b8.
//
// Solidity: function getThreshold() view returns(uint8)
func (_FilswanOracle *FilswanOracleCallerSession) GetThreshold() (uint8, error) {
	return _FilswanOracle.Contract.GetThreshold(&_FilswanOracle.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FilswanOracle *FilswanOracleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FilswanOracle *FilswanOracleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FilswanOracle.Contract.HasRole(&_FilswanOracle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FilswanOracle *FilswanOracleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FilswanOracle.Contract.HasRole(&_FilswanOracle.CallOpts, role, account)
}

// IsCarPaymentAvailable is a free data retrieval call binding the contract method 0xe1296541.
//
// Solidity: function isCarPaymentAvailable(string dealId, string network, address recipient) view returns(bool)
func (_FilswanOracle *FilswanOracleCaller) IsCarPaymentAvailable(opts *bind.CallOpts, dealId string, network string, recipient common.Address) (bool, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "isCarPaymentAvailable", dealId, network, recipient)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCarPaymentAvailable is a free data retrieval call binding the contract method 0xe1296541.
//
// Solidity: function isCarPaymentAvailable(string dealId, string network, address recipient) view returns(bool)
func (_FilswanOracle *FilswanOracleSession) IsCarPaymentAvailable(dealId string, network string, recipient common.Address) (bool, error) {
	return _FilswanOracle.Contract.IsCarPaymentAvailable(&_FilswanOracle.CallOpts, dealId, network, recipient)
}

// IsCarPaymentAvailable is a free data retrieval call binding the contract method 0xe1296541.
//
// Solidity: function isCarPaymentAvailable(string dealId, string network, address recipient) view returns(bool)
func (_FilswanOracle *FilswanOracleCallerSession) IsCarPaymentAvailable(dealId string, network string, recipient common.Address) (bool, error) {
	return _FilswanOracle.Contract.IsCarPaymentAvailable(&_FilswanOracle.CallOpts, dealId, network, recipient)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilswanOracle *FilswanOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilswanOracle *FilswanOracleSession) Owner() (common.Address, error) {
	return _FilswanOracle.Contract.Owner(&_FilswanOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilswanOracle *FilswanOracleCallerSession) Owner() (common.Address, error) {
	return _FilswanOracle.Contract.Owner(&_FilswanOracle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FilswanOracle *FilswanOracleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FilswanOracle *FilswanOracleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FilswanOracle.Contract.SupportsInterface(&_FilswanOracle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FilswanOracle *FilswanOracleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FilswanOracle.Contract.SupportsInterface(&_FilswanOracle.CallOpts, interfaceId)
}

// F is a paid mutator transaction binding the contract method 0xa33f9c67.
//
// Solidity: function f(string s1, string s2, address a1, string[] sa) returns()
func (_FilswanOracle *FilswanOracleTransactor) F(opts *bind.TransactOpts, s1 string, s2 string, a1 common.Address, sa []string) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "f", s1, s2, a1, sa)
}

// F is a paid mutator transaction binding the contract method 0xa33f9c67.
//
// Solidity: function f(string s1, string s2, address a1, string[] sa) returns()
func (_FilswanOracle *FilswanOracleSession) F(s1 string, s2 string, a1 common.Address, sa []string) (*types.Transaction, error) {
	return _FilswanOracle.Contract.F(&_FilswanOracle.TransactOpts, s1, s2, a1, sa)
}

// F is a paid mutator transaction binding the contract method 0xa33f9c67.
//
// Solidity: function f(string s1, string s2, address a1, string[] sa) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) F(s1 string, s2 string, a1 common.Address, sa []string) (*types.Transaction, error) {
	return _FilswanOracle.Contract.F(&_FilswanOracle.TransactOpts, s1, s2, a1, sa)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.GrantRole(&_FilswanOracle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.GrantRole(&_FilswanOracle.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address admin, uint8 threshold) returns()
func (_FilswanOracle *FilswanOracleTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "initialize", admin, threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address admin, uint8 threshold) returns()
func (_FilswanOracle *FilswanOracleSession) Initialize(admin common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Initialize(&_FilswanOracle.TransactOpts, admin, threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address admin, uint8 threshold) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) Initialize(admin common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Initialize(&_FilswanOracle.TransactOpts, admin, threshold)
}

// PreSign is a paid mutator transaction binding the contract method 0xcc8a4c11.
//
// Solidity: function preSign(string dealId, string network, address recipient, uint8 batchCount) returns()
func (_FilswanOracle *FilswanOracleTransactor) PreSign(opts *bind.TransactOpts, dealId string, network string, recipient common.Address, batchCount uint8) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "preSign", dealId, network, recipient, batchCount)
}

// PreSign is a paid mutator transaction binding the contract method 0xcc8a4c11.
//
// Solidity: function preSign(string dealId, string network, address recipient, uint8 batchCount) returns()
func (_FilswanOracle *FilswanOracleSession) PreSign(dealId string, network string, recipient common.Address, batchCount uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.PreSign(&_FilswanOracle.TransactOpts, dealId, network, recipient, batchCount)
}

// PreSign is a paid mutator transaction binding the contract method 0xcc8a4c11.
//
// Solidity: function preSign(string dealId, string network, address recipient, uint8 batchCount) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) PreSign(dealId string, network string, recipient common.Address, batchCount uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.PreSign(&_FilswanOracle.TransactOpts, dealId, network, recipient, batchCount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FilswanOracle *FilswanOracleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FilswanOracle *FilswanOracleSession) RenounceOwnership() (*types.Transaction, error) {
	return _FilswanOracle.Contract.RenounceOwnership(&_FilswanOracle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FilswanOracle *FilswanOracleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FilswanOracle.Contract.RenounceOwnership(&_FilswanOracle.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.RenounceRole(&_FilswanOracle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.RenounceRole(&_FilswanOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.RevokeRole(&_FilswanOracle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.RevokeRole(&_FilswanOracle.TransactOpts, role, account)
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

// SetFilinkOracle is a paid mutator transaction binding the contract method 0x214fe62e.
//
// Solidity: function setFilinkOracle(address filinkAddress) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) SetFilinkOracle(opts *bind.TransactOpts, filinkAddress common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "setFilinkOracle", filinkAddress)
}

// SetFilinkOracle is a paid mutator transaction binding the contract method 0x214fe62e.
//
// Solidity: function setFilinkOracle(address filinkAddress) returns(bool)
func (_FilswanOracle *FilswanOracleSession) SetFilinkOracle(filinkAddress common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetFilinkOracle(&_FilswanOracle.TransactOpts, filinkAddress)
}

// SetFilinkOracle is a paid mutator transaction binding the contract method 0x214fe62e.
//
// Solidity: function setFilinkOracle(address filinkAddress) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) SetFilinkOracle(filinkAddress common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetFilinkOracle(&_FilswanOracle.TransactOpts, filinkAddress)
}

// Sign is a paid mutator transaction binding the contract method 0xe7c2999a.
//
// Solidity: function sign(string dealId, string network, string[] cidList, uint8 batchNo) returns()
func (_FilswanOracle *FilswanOracleTransactor) Sign(opts *bind.TransactOpts, dealId string, network string, cidList []string, batchNo uint8) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "sign", dealId, network, cidList, batchNo)
}

// Sign is a paid mutator transaction binding the contract method 0xe7c2999a.
//
// Solidity: function sign(string dealId, string network, string[] cidList, uint8 batchNo) returns()
func (_FilswanOracle *FilswanOracleSession) Sign(dealId string, network string, cidList []string, batchNo uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Sign(&_FilswanOracle.TransactOpts, dealId, network, cidList, batchNo)
}

// Sign is a paid mutator transaction binding the contract method 0xe7c2999a.
//
// Solidity: function sign(string dealId, string network, string[] cidList, uint8 batchNo) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) Sign(dealId string, network string, cidList []string, batchNo uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Sign(&_FilswanOracle.TransactOpts, dealId, network, cidList, batchNo)
}

// SignCarTransaction is a paid mutator transaction binding the contract method 0x4d043ef6.
//
// Solidity: function signCarTransaction(string[] cidList, string dealId, string network, address recipient) returns()
func (_FilswanOracle *FilswanOracleTransactor) SignCarTransaction(opts *bind.TransactOpts, cidList []string, dealId string, network string, recipient common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "signCarTransaction", cidList, dealId, network, recipient)
}

// SignCarTransaction is a paid mutator transaction binding the contract method 0x4d043ef6.
//
// Solidity: function signCarTransaction(string[] cidList, string dealId, string network, address recipient) returns()
func (_FilswanOracle *FilswanOracleSession) SignCarTransaction(cidList []string, dealId string, network string, recipient common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SignCarTransaction(&_FilswanOracle.TransactOpts, cidList, dealId, network, recipient)
}

// SignCarTransaction is a paid mutator transaction binding the contract method 0x4d043ef6.
//
// Solidity: function signCarTransaction(string[] cidList, string dealId, string network, address recipient) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) SignCarTransaction(cidList []string, dealId string, network string, recipient common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SignCarTransaction(&_FilswanOracle.TransactOpts, cidList, dealId, network, recipient)
}

// SignHash is a paid mutator transaction binding the contract method 0x4a239fdb.
//
// Solidity: function signHash(string dealId, string network, address recipient, bytes32 voteKey) returns()
func (_FilswanOracle *FilswanOracleTransactor) SignHash(opts *bind.TransactOpts, dealId string, network string, recipient common.Address, voteKey [32]byte) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "signHash", dealId, network, recipient, voteKey)
}

// SignHash is a paid mutator transaction binding the contract method 0x4a239fdb.
//
// Solidity: function signHash(string dealId, string network, address recipient, bytes32 voteKey) returns()
func (_FilswanOracle *FilswanOracleSession) SignHash(dealId string, network string, recipient common.Address, voteKey [32]byte) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SignHash(&_FilswanOracle.TransactOpts, dealId, network, recipient, voteKey)
}

// SignHash is a paid mutator transaction binding the contract method 0x4a239fdb.
//
// Solidity: function signHash(string dealId, string network, address recipient, bytes32 voteKey) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) SignHash(dealId string, network string, recipient common.Address, voteKey [32]byte) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SignHash(&_FilswanOracle.TransactOpts, dealId, network, recipient, voteKey)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FilswanOracle *FilswanOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FilswanOracle *FilswanOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.TransferOwnership(&_FilswanOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.TransferOwnership(&_FilswanOracle.TransactOpts, newOwner)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0xf3df5b69.
//
// Solidity: function updateThreshold(uint8 threshold) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) UpdateThreshold(opts *bind.TransactOpts, threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "updateThreshold", threshold)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0xf3df5b69.
//
// Solidity: function updateThreshold(uint8 threshold) returns(bool)
func (_FilswanOracle *FilswanOracleSession) UpdateThreshold(threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UpdateThreshold(&_FilswanOracle.TransactOpts, threshold)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0xf3df5b69.
//
// Solidity: function updateThreshold(uint8 threshold) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) UpdateThreshold(threshold uint8) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UpdateThreshold(&_FilswanOracle.TransactOpts, threshold)
}

// FilswanOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FilswanOracle contract.
type FilswanOracleOwnershipTransferredIterator struct {
	Event *FilswanOracleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleOwnershipTransferred represents a OwnershipTransferred event raised by the FilswanOracle contract.
type FilswanOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FilswanOracle *FilswanOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FilswanOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleOwnershipTransferredIterator{contract: _FilswanOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FilswanOracle *FilswanOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FilswanOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleOwnershipTransferred)
				if err := _FilswanOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FilswanOracle *FilswanOracleFilterer) ParseOwnershipTransferred(log types.Log) (*FilswanOracleOwnershipTransferred, error) {
	event := new(FilswanOracleOwnershipTransferred)
	if err := _FilswanOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOraclePreSignIterator is returned from FilterPreSign and is used to iterate over the raw logs and unpacked data for PreSign events raised by the FilswanOracle contract.
type FilswanOraclePreSignIterator struct {
	Event *FilswanOraclePreSign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOraclePreSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOraclePreSign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOraclePreSign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOraclePreSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOraclePreSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOraclePreSign represents a PreSign event raised by the FilswanOracle contract.
type FilswanOraclePreSign struct {
	DealId     string
	Network    string
	Recipient  common.Address
	BatchCount uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPreSign is a free log retrieval operation binding the contract event 0xc69d6239860503167422092fbe6d5ad3e6843763d8fc9c6b876b4117cd366960.
//
// Solidity: event PreSign(string dealId, string network, address recipient, uint8 batchCount)
func (_FilswanOracle *FilswanOracleFilterer) FilterPreSign(opts *bind.FilterOpts) (*FilswanOraclePreSignIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "PreSign")
	if err != nil {
		return nil, err
	}
	return &FilswanOraclePreSignIterator{contract: _FilswanOracle.contract, event: "PreSign", logs: logs, sub: sub}, nil
}

// WatchPreSign is a free log subscription operation binding the contract event 0xc69d6239860503167422092fbe6d5ad3e6843763d8fc9c6b876b4117cd366960.
//
// Solidity: event PreSign(string dealId, string network, address recipient, uint8 batchCount)
func (_FilswanOracle *FilswanOracleFilterer) WatchPreSign(opts *bind.WatchOpts, sink chan<- *FilswanOraclePreSign) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "PreSign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOraclePreSign)
				if err := _FilswanOracle.contract.UnpackLog(event, "PreSign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePreSign is a log parse operation binding the contract event 0xc69d6239860503167422092fbe6d5ad3e6843763d8fc9c6b876b4117cd366960.
//
// Solidity: event PreSign(string dealId, string network, address recipient, uint8 batchCount)
func (_FilswanOracle *FilswanOracleFilterer) ParsePreSign(log types.Log) (*FilswanOraclePreSign, error) {
	event := new(FilswanOraclePreSign)
	if err := _FilswanOracle.contract.UnpackLog(event, "PreSign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FilswanOracle contract.
type FilswanOracleRoleAdminChangedIterator struct {
	Event *FilswanOracleRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleRoleAdminChanged represents a RoleAdminChanged event raised by the FilswanOracle contract.
type FilswanOracleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FilswanOracle *FilswanOracleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FilswanOracleRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleRoleAdminChangedIterator{contract: _FilswanOracle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FilswanOracle *FilswanOracleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FilswanOracleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleRoleAdminChanged)
				if err := _FilswanOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FilswanOracle *FilswanOracleFilterer) ParseRoleAdminChanged(log types.Log) (*FilswanOracleRoleAdminChanged, error) {
	event := new(FilswanOracleRoleAdminChanged)
	if err := _FilswanOracle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FilswanOracle contract.
type FilswanOracleRoleGrantedIterator struct {
	Event *FilswanOracleRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleRoleGranted represents a RoleGranted event raised by the FilswanOracle contract.
type FilswanOracleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOracle *FilswanOracleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FilswanOracleRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleRoleGrantedIterator{contract: _FilswanOracle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOracle *FilswanOracleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FilswanOracleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleRoleGranted)
				if err := _FilswanOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOracle *FilswanOracleFilterer) ParseRoleGranted(log types.Log) (*FilswanOracleRoleGranted, error) {
	event := new(FilswanOracleRoleGranted)
	if err := _FilswanOracle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FilswanOracle contract.
type FilswanOracleRoleRevokedIterator struct {
	Event *FilswanOracleRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleRoleRevoked represents a RoleRevoked event raised by the FilswanOracle contract.
type FilswanOracleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOracle *FilswanOracleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FilswanOracleRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOracleRoleRevokedIterator{contract: _FilswanOracle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOracle *FilswanOracleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FilswanOracleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleRoleRevoked)
				if err := _FilswanOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOracle *FilswanOracleFilterer) ParseRoleRevoked(log types.Log) (*FilswanOracleRoleRevoked, error) {
	event := new(FilswanOracleRoleRevoked)
	if err := _FilswanOracle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleSignIterator is returned from FilterSign and is used to iterate over the raw logs and unpacked data for Sign events raised by the FilswanOracle contract.
type FilswanOracleSignIterator struct {
	Event *FilswanOracleSign // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleSignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleSign)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleSign)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleSignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleSignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleSign represents a Sign event raised by the FilswanOracle contract.
type FilswanOracleSign struct {
	DealId  string
	Network string
	CidList []string
	BatchNo uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSign is a free log retrieval operation binding the contract event 0xbe0b994713ae4ca03ec61948b3b53c07529d5ce656d04fb5559f70298351203b.
//
// Solidity: event Sign(string dealId, string network, string[] cidList, uint8 batchNo)
func (_FilswanOracle *FilswanOracleFilterer) FilterSign(opts *bind.FilterOpts) (*FilswanOracleSignIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleSignIterator{contract: _FilswanOracle.contract, event: "Sign", logs: logs, sub: sub}, nil
}

// WatchSign is a free log subscription operation binding the contract event 0xbe0b994713ae4ca03ec61948b3b53c07529d5ce656d04fb5559f70298351203b.
//
// Solidity: event Sign(string dealId, string network, string[] cidList, uint8 batchNo)
func (_FilswanOracle *FilswanOracleFilterer) WatchSign(opts *bind.WatchOpts, sink chan<- *FilswanOracleSign) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "Sign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleSign)
				if err := _FilswanOracle.contract.UnpackLog(event, "Sign", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSign is a log parse operation binding the contract event 0xbe0b994713ae4ca03ec61948b3b53c07529d5ce656d04fb5559f70298351203b.
//
// Solidity: event Sign(string dealId, string network, string[] cidList, uint8 batchNo)
func (_FilswanOracle *FilswanOracleFilterer) ParseSign(log types.Log) (*FilswanOracleSign, error) {
	event := new(FilswanOracleSign)
	if err := _FilswanOracle.contract.UnpackLog(event, "Sign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleSignCarTransactionIterator is returned from FilterSignCarTransaction and is used to iterate over the raw logs and unpacked data for SignCarTransaction events raised by the FilswanOracle contract.
type FilswanOracleSignCarTransactionIterator struct {
	Event *FilswanOracleSignCarTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleSignCarTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleSignCarTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleSignCarTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleSignCarTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleSignCarTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleSignCarTransaction represents a SignCarTransaction event raised by the FilswanOracle contract.
type FilswanOracleSignCarTransaction struct {
	CidList   []string
	DealId    string
	Network   string
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignCarTransaction is a free log retrieval operation binding the contract event 0x2a521717e606d77fdfbcf23c89749aae1af7f5943a56f0b0c34ad77478fcc4f1.
//
// Solidity: event SignCarTransaction(string[] cidList, string dealId, string network, address recipient)
func (_FilswanOracle *FilswanOracleFilterer) FilterSignCarTransaction(opts *bind.FilterOpts) (*FilswanOracleSignCarTransactionIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "SignCarTransaction")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleSignCarTransactionIterator{contract: _FilswanOracle.contract, event: "SignCarTransaction", logs: logs, sub: sub}, nil
}

// WatchSignCarTransaction is a free log subscription operation binding the contract event 0x2a521717e606d77fdfbcf23c89749aae1af7f5943a56f0b0c34ad77478fcc4f1.
//
// Solidity: event SignCarTransaction(string[] cidList, string dealId, string network, address recipient)
func (_FilswanOracle *FilswanOracleFilterer) WatchSignCarTransaction(opts *bind.WatchOpts, sink chan<- *FilswanOracleSignCarTransaction) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "SignCarTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleSignCarTransaction)
				if err := _FilswanOracle.contract.UnpackLog(event, "SignCarTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignCarTransaction is a log parse operation binding the contract event 0x2a521717e606d77fdfbcf23c89749aae1af7f5943a56f0b0c34ad77478fcc4f1.
//
// Solidity: event SignCarTransaction(string[] cidList, string dealId, string network, address recipient)
func (_FilswanOracle *FilswanOracleFilterer) ParseSignCarTransaction(log types.Log) (*FilswanOracleSignCarTransaction, error) {
	event := new(FilswanOracleSignCarTransaction)
	if err := _FilswanOracle.contract.UnpackLog(event, "SignCarTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleSignHashIterator is returned from FilterSignHash and is used to iterate over the raw logs and unpacked data for SignHash events raised by the FilswanOracle contract.
type FilswanOracleSignHashIterator struct {
	Event *FilswanOracleSignHash // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleSignHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleSignHash)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleSignHash)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleSignHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleSignHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleSignHash represents a SignHash event raised by the FilswanOracle contract.
type FilswanOracleSignHash struct {
	DealId    string
	Network   string
	Recipient common.Address
	VoteKey   [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignHash is a free log retrieval operation binding the contract event 0xd68d015de10c7d3d43d32911d7f15a2eca53f4a3d44390f6c38f97ae197ea98a.
//
// Solidity: event SignHash(string dealId, string network, address recipient, bytes32 voteKey)
func (_FilswanOracle *FilswanOracleFilterer) FilterSignHash(opts *bind.FilterOpts) (*FilswanOracleSignHashIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "SignHash")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleSignHashIterator{contract: _FilswanOracle.contract, event: "SignHash", logs: logs, sub: sub}, nil
}

// WatchSignHash is a free log subscription operation binding the contract event 0xd68d015de10c7d3d43d32911d7f15a2eca53f4a3d44390f6c38f97ae197ea98a.
//
// Solidity: event SignHash(string dealId, string network, address recipient, bytes32 voteKey)
func (_FilswanOracle *FilswanOracleFilterer) WatchSignHash(opts *bind.WatchOpts, sink chan<- *FilswanOracleSignHash) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "SignHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleSignHash)
				if err := _FilswanOracle.contract.UnpackLog(event, "SignHash", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignHash is a log parse operation binding the contract event 0xd68d015de10c7d3d43d32911d7f15a2eca53f4a3d44390f6c38f97ae197ea98a.
//
// Solidity: event SignHash(string dealId, string network, address recipient, bytes32 voteKey)
func (_FilswanOracle *FilswanOracleFilterer) ParseSignHash(log types.Log) (*FilswanOracleSignHash, error) {
	event := new(FilswanOracleSignHash)
	if err := _FilswanOracle.contract.UnpackLog(event, "SignHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleSignTransactionIterator is returned from FilterSignTransaction and is used to iterate over the raw logs and unpacked data for SignTransaction events raised by the FilswanOracle contract.
type FilswanOracleSignTransactionIterator struct {
	Event *FilswanOracleSignTransaction // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FilswanOracleSignTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleSignTransaction)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FilswanOracleSignTransaction)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FilswanOracleSignTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleSignTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleSignTransaction represents a SignTransaction event raised by the FilswanOracle contract.
type FilswanOracleSignTransaction struct {
	Cid       string
	DealId    string
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignTransaction is a free log retrieval operation binding the contract event 0xd7e913d5977f1930a4a767cfcc42dbf926c307ed710ee76eb27764cfd60e5d7f.
//
// Solidity: event SignTransaction(string cid, string dealId, address recipient)
func (_FilswanOracle *FilswanOracleFilterer) FilterSignTransaction(opts *bind.FilterOpts) (*FilswanOracleSignTransactionIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "SignTransaction")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleSignTransactionIterator{contract: _FilswanOracle.contract, event: "SignTransaction", logs: logs, sub: sub}, nil
}

// WatchSignTransaction is a free log subscription operation binding the contract event 0xd7e913d5977f1930a4a767cfcc42dbf926c307ed710ee76eb27764cfd60e5d7f.
//
// Solidity: event SignTransaction(string cid, string dealId, address recipient)
func (_FilswanOracle *FilswanOracleFilterer) WatchSignTransaction(opts *bind.WatchOpts, sink chan<- *FilswanOracleSignTransaction) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "SignTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleSignTransaction)
				if err := _FilswanOracle.contract.UnpackLog(event, "SignTransaction", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignTransaction is a log parse operation binding the contract event 0xd7e913d5977f1930a4a767cfcc42dbf926c307ed710ee76eb27764cfd60e5d7f.
//
// Solidity: event SignTransaction(string cid, string dealId, address recipient)
func (_FilswanOracle *FilswanOracleFilterer) ParseSignTransaction(log types.Log) (*FilswanOracleSignTransaction, error) {
	event := new(FilswanOracleSignTransaction)
	if err := _FilswanOracle.contract.UnpackLog(event, "SignTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
