// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goBind

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FilswanOralcleMetaData contains all meta data concerning the FilswanOralcle contract.
var FilswanOralcleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"terms\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"SignTransaction\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DAO_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"isPaymentAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"daoUsers\",\"type\":\"address[]\"}],\"name\":\"setDAOUsers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"paid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"terms\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"signTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"threshold\",\"type\":\"uint8\"}],\"name\":\"updateThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FilswanOralcleABI is the input ABI used to generate the binding from.
// Deprecated: Use FilswanOralcleMetaData.ABI instead.
var FilswanOralcleABI = FilswanOralcleMetaData.ABI

// FilswanOralcle is an auto generated Go binding around an Ethereum contract.
type FilswanOralcle struct {
	FilswanOralcleCaller     // Read-only binding to the contract
	FilswanOralcleTransactor // Write-only binding to the contract
	FilswanOralcleFilterer   // Log filterer for contract events
}

// FilswanOralcleCaller is an auto generated read-only Go binding around an Ethereum contract.
type FilswanOralcleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilswanOralcleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FilswanOralcleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilswanOralcleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FilswanOralcleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FilswanOralcleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FilswanOralcleSession struct {
	Contract     *FilswanOralcle   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FilswanOralcleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FilswanOralcleCallerSession struct {
	Contract *FilswanOralcleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FilswanOralcleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FilswanOralcleTransactorSession struct {
	Contract     *FilswanOralcleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FilswanOralcleRaw is an auto generated low-level Go binding around an Ethereum contract.
type FilswanOralcleRaw struct {
	Contract *FilswanOralcle // Generic contract binding to access the raw methods on
}

// FilswanOralcleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FilswanOralcleCallerRaw struct {
	Contract *FilswanOralcleCaller // Generic read-only contract binding to access the raw methods on
}

// FilswanOralcleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FilswanOralcleTransactorRaw struct {
	Contract *FilswanOralcleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFilswanOralcle creates a new instance of FilswanOralcle, bound to a specific deployed contract.
func NewFilswanOralcle(address common.Address, backend bind.ContractBackend) (*FilswanOralcle, error) {
	contract, err := bindFilswanOralcle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcle{FilswanOralcleCaller: FilswanOralcleCaller{contract: contract}, FilswanOralcleTransactor: FilswanOralcleTransactor{contract: contract}, FilswanOralcleFilterer: FilswanOralcleFilterer{contract: contract}}, nil
}

// NewFilswanOralcleCaller creates a new read-only instance of FilswanOralcle, bound to a specific deployed contract.
func NewFilswanOralcleCaller(address common.Address, caller bind.ContractCaller) (*FilswanOralcleCaller, error) {
	contract, err := bindFilswanOralcle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleCaller{contract: contract}, nil
}

// NewFilswanOralcleTransactor creates a new write-only instance of FilswanOralcle, bound to a specific deployed contract.
func NewFilswanOralcleTransactor(address common.Address, transactor bind.ContractTransactor) (*FilswanOralcleTransactor, error) {
	contract, err := bindFilswanOralcle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleTransactor{contract: contract}, nil
}

// NewFilswanOralcleFilterer creates a new log filterer instance of FilswanOralcle, bound to a specific deployed contract.
func NewFilswanOralcleFilterer(address common.Address, filterer bind.ContractFilterer) (*FilswanOralcleFilterer, error) {
	contract, err := bindFilswanOralcle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleFilterer{contract: contract}, nil
}

// bindFilswanOralcle binds a generic wrapper to an already deployed contract.
func bindFilswanOralcle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FilswanOralcleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilswanOralcle *FilswanOralcleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilswanOralcle.Contract.FilswanOralcleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilswanOralcle *FilswanOralcleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.FilswanOralcleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilswanOralcle *FilswanOralcleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.FilswanOralcleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FilswanOralcle *FilswanOralcleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FilswanOralcle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FilswanOralcle *FilswanOralcleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FilswanOralcle *FilswanOralcleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.contract.Transact(opts, method, params...)
}

// DAOROLE is a free data retrieval call binding the contract method 0xe9c26518.
//
// Solidity: function DAO_ROLE() view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleCaller) DAOROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "DAO_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DAOROLE is a free data retrieval call binding the contract method 0xe9c26518.
//
// Solidity: function DAO_ROLE() view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleSession) DAOROLE() ([32]byte, error) {
	return _FilswanOralcle.Contract.DAOROLE(&_FilswanOralcle.CallOpts)
}

// DAOROLE is a free data retrieval call binding the contract method 0xe9c26518.
//
// Solidity: function DAO_ROLE() view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleCallerSession) DAOROLE() ([32]byte, error) {
	return _FilswanOralcle.Contract.DAOROLE(&_FilswanOralcle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FilswanOralcle.Contract.DEFAULTADMINROLE(&_FilswanOralcle.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FilswanOralcle.Contract.DEFAULTADMINROLE(&_FilswanOralcle.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FilswanOralcle.Contract.GetRoleAdmin(&_FilswanOralcle.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FilswanOralcle *FilswanOralcleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FilswanOralcle.Contract.GetRoleAdmin(&_FilswanOralcle.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FilswanOralcle *FilswanOralcleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FilswanOralcle *FilswanOralcleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FilswanOralcle.Contract.HasRole(&_FilswanOralcle.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FilswanOralcle *FilswanOralcleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FilswanOralcle.Contract.HasRole(&_FilswanOralcle.CallOpts, role, account)
}

// IsPaymentAvailable is a free data retrieval call binding the contract method 0x5a2b0ea6.
//
// Solidity: function isPaymentAvailable(string cid, string orderId, string dealId, uint256 paid, address recipient, bool status) view returns(bool)
func (_FilswanOralcle *FilswanOralcleCaller) IsPaymentAvailable(opts *bind.CallOpts, cid string, orderId string, dealId string, paid *big.Int, recipient common.Address, status bool) (bool, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "isPaymentAvailable", cid, orderId, dealId, paid, recipient, status)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaymentAvailable is a free data retrieval call binding the contract method 0x5a2b0ea6.
//
// Solidity: function isPaymentAvailable(string cid, string orderId, string dealId, uint256 paid, address recipient, bool status) view returns(bool)
func (_FilswanOralcle *FilswanOralcleSession) IsPaymentAvailable(cid string, orderId string, dealId string, paid *big.Int, recipient common.Address, status bool) (bool, error) {
	return _FilswanOralcle.Contract.IsPaymentAvailable(&_FilswanOralcle.CallOpts, cid, orderId, dealId, paid, recipient, status)
}

// IsPaymentAvailable is a free data retrieval call binding the contract method 0x5a2b0ea6.
//
// Solidity: function isPaymentAvailable(string cid, string orderId, string dealId, uint256 paid, address recipient, bool status) view returns(bool)
func (_FilswanOralcle *FilswanOralcleCallerSession) IsPaymentAvailable(cid string, orderId string, dealId string, paid *big.Int, recipient common.Address, status bool) (bool, error) {
	return _FilswanOralcle.Contract.IsPaymentAvailable(&_FilswanOralcle.CallOpts, cid, orderId, dealId, paid, recipient, status)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilswanOralcle *FilswanOralcleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilswanOralcle *FilswanOralcleSession) Owner() (common.Address, error) {
	return _FilswanOralcle.Contract.Owner(&_FilswanOralcle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FilswanOralcle *FilswanOralcleCallerSession) Owner() (common.Address, error) {
	return _FilswanOralcle.Contract.Owner(&_FilswanOralcle.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FilswanOralcle *FilswanOralcleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FilswanOralcle.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FilswanOralcle *FilswanOralcleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FilswanOralcle.Contract.SupportsInterface(&_FilswanOralcle.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FilswanOralcle *FilswanOralcleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FilswanOralcle.Contract.SupportsInterface(&_FilswanOralcle.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.GrantRole(&_FilswanOralcle.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.GrantRole(&_FilswanOralcle.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address admin, uint8 threshold) returns()
func (_FilswanOralcle *FilswanOralcleTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "initialize", admin, threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address admin, uint8 threshold) returns()
func (_FilswanOralcle *FilswanOralcleSession) Initialize(admin common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.Initialize(&_FilswanOralcle.TransactOpts, admin, threshold)
}

// Initialize is a paid mutator transaction binding the contract method 0x943b24b2.
//
// Solidity: function initialize(address admin, uint8 threshold) returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) Initialize(admin common.Address, threshold uint8) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.Initialize(&_FilswanOralcle.TransactOpts, admin, threshold)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FilswanOralcle *FilswanOralcleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FilswanOralcle *FilswanOralcleSession) RenounceOwnership() (*types.Transaction, error) {
	return _FilswanOralcle.Contract.RenounceOwnership(&_FilswanOralcle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FilswanOralcle.Contract.RenounceOwnership(&_FilswanOralcle.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.RenounceRole(&_FilswanOralcle.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.RenounceRole(&_FilswanOralcle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.RevokeRole(&_FilswanOralcle.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.RevokeRole(&_FilswanOralcle.TransactOpts, role, account)
}

// SetDAOUsers is a paid mutator transaction binding the contract method 0x6bcf35f9.
//
// Solidity: function setDAOUsers(address[] daoUsers) returns(bool)
func (_FilswanOralcle *FilswanOralcleTransactor) SetDAOUsers(opts *bind.TransactOpts, daoUsers []common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "setDAOUsers", daoUsers)
}

// SetDAOUsers is a paid mutator transaction binding the contract method 0x6bcf35f9.
//
// Solidity: function setDAOUsers(address[] daoUsers) returns(bool)
func (_FilswanOralcle *FilswanOralcleSession) SetDAOUsers(daoUsers []common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.SetDAOUsers(&_FilswanOralcle.TransactOpts, daoUsers)
}

// SetDAOUsers is a paid mutator transaction binding the contract method 0x6bcf35f9.
//
// Solidity: function setDAOUsers(address[] daoUsers) returns(bool)
func (_FilswanOralcle *FilswanOralcleTransactorSession) SetDAOUsers(daoUsers []common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.SetDAOUsers(&_FilswanOralcle.TransactOpts, daoUsers)
}

// SignTransaction is a paid mutator transaction binding the contract method 0x0c435a8b.
//
// Solidity: function signTransaction(string cid, string orderId, string dealId, uint256 paid, address recipient, uint256 terms, bool status) returns()
func (_FilswanOralcle *FilswanOralcleTransactor) SignTransaction(opts *bind.TransactOpts, cid string, orderId string, dealId string, paid *big.Int, recipient common.Address, terms *big.Int, status bool) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "signTransaction", cid, orderId, dealId, paid, recipient, terms, status)
}

// SignTransaction is a paid mutator transaction binding the contract method 0x0c435a8b.
//
// Solidity: function signTransaction(string cid, string orderId, string dealId, uint256 paid, address recipient, uint256 terms, bool status) returns()
func (_FilswanOralcle *FilswanOralcleSession) SignTransaction(cid string, orderId string, dealId string, paid *big.Int, recipient common.Address, terms *big.Int, status bool) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.SignTransaction(&_FilswanOralcle.TransactOpts, cid, orderId, dealId, paid, recipient, terms, status)
}

// SignTransaction is a paid mutator transaction binding the contract method 0x0c435a8b.
//
// Solidity: function signTransaction(string cid, string orderId, string dealId, uint256 paid, address recipient, uint256 terms, bool status) returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) SignTransaction(cid string, orderId string, dealId string, paid *big.Int, recipient common.Address, terms *big.Int, status bool) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.SignTransaction(&_FilswanOralcle.TransactOpts, cid, orderId, dealId, paid, recipient, terms, status)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FilswanOralcle *FilswanOralcleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FilswanOralcle *FilswanOralcleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.TransferOwnership(&_FilswanOralcle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FilswanOralcle *FilswanOralcleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.TransferOwnership(&_FilswanOralcle.TransactOpts, newOwner)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0xf3df5b69.
//
// Solidity: function updateThreshold(uint8 threshold) returns(bool)
func (_FilswanOralcle *FilswanOralcleTransactor) UpdateThreshold(opts *bind.TransactOpts, threshold uint8) (*types.Transaction, error) {
	return _FilswanOralcle.contract.Transact(opts, "updateThreshold", threshold)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0xf3df5b69.
//
// Solidity: function updateThreshold(uint8 threshold) returns(bool)
func (_FilswanOralcle *FilswanOralcleSession) UpdateThreshold(threshold uint8) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.UpdateThreshold(&_FilswanOralcle.TransactOpts, threshold)
}

// UpdateThreshold is a paid mutator transaction binding the contract method 0xf3df5b69.
//
// Solidity: function updateThreshold(uint8 threshold) returns(bool)
func (_FilswanOralcle *FilswanOralcleTransactorSession) UpdateThreshold(threshold uint8) (*types.Transaction, error) {
	return _FilswanOralcle.Contract.UpdateThreshold(&_FilswanOralcle.TransactOpts, threshold)
}

// FilswanOralcleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FilswanOralcle contract.
type FilswanOralcleOwnershipTransferredIterator struct {
	Event *FilswanOralcleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FilswanOralcleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOralcleOwnershipTransferred)
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
		it.Event = new(FilswanOralcleOwnershipTransferred)
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
func (it *FilswanOralcleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOralcleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOralcleOwnershipTransferred represents a OwnershipTransferred event raised by the FilswanOralcle contract.
type FilswanOralcleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FilswanOralcle *FilswanOralcleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FilswanOralcleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FilswanOralcle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleOwnershipTransferredIterator{contract: _FilswanOralcle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FilswanOralcle *FilswanOralcleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FilswanOralcleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FilswanOralcle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOralcleOwnershipTransferred)
				if err := _FilswanOralcle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FilswanOralcle *FilswanOralcleFilterer) ParseOwnershipTransferred(log types.Log) (*FilswanOralcleOwnershipTransferred, error) {
	event := new(FilswanOralcleOwnershipTransferred)
	if err := _FilswanOralcle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOralcleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FilswanOralcle contract.
type FilswanOralcleRoleAdminChangedIterator struct {
	Event *FilswanOralcleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FilswanOralcleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOralcleRoleAdminChanged)
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
		it.Event = new(FilswanOralcleRoleAdminChanged)
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
func (it *FilswanOralcleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOralcleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOralcleRoleAdminChanged represents a RoleAdminChanged event raised by the FilswanOralcle contract.
type FilswanOralcleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FilswanOralcle *FilswanOralcleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FilswanOralcleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _FilswanOralcle.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleRoleAdminChangedIterator{contract: _FilswanOralcle.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FilswanOralcle *FilswanOralcleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FilswanOralcleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _FilswanOralcle.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOralcleRoleAdminChanged)
				if err := _FilswanOralcle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_FilswanOralcle *FilswanOralcleFilterer) ParseRoleAdminChanged(log types.Log) (*FilswanOralcleRoleAdminChanged, error) {
	event := new(FilswanOralcleRoleAdminChanged)
	if err := _FilswanOralcle.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOralcleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FilswanOralcle contract.
type FilswanOralcleRoleGrantedIterator struct {
	Event *FilswanOralcleRoleGranted // Event containing the contract specifics and raw log

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
func (it *FilswanOralcleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOralcleRoleGranted)
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
		it.Event = new(FilswanOralcleRoleGranted)
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
func (it *FilswanOralcleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOralcleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOralcleRoleGranted represents a RoleGranted event raised by the FilswanOralcle contract.
type FilswanOralcleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOralcle *FilswanOralcleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FilswanOralcleRoleGrantedIterator, error) {

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

	logs, sub, err := _FilswanOralcle.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleRoleGrantedIterator{contract: _FilswanOralcle.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOralcle *FilswanOralcleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FilswanOralcleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FilswanOralcle.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOralcleRoleGranted)
				if err := _FilswanOralcle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_FilswanOralcle *FilswanOralcleFilterer) ParseRoleGranted(log types.Log) (*FilswanOralcleRoleGranted, error) {
	event := new(FilswanOralcleRoleGranted)
	if err := _FilswanOralcle.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOralcleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FilswanOralcle contract.
type FilswanOralcleRoleRevokedIterator struct {
	Event *FilswanOralcleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FilswanOralcleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOralcleRoleRevoked)
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
		it.Event = new(FilswanOralcleRoleRevoked)
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
func (it *FilswanOralcleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOralcleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOralcleRoleRevoked represents a RoleRevoked event raised by the FilswanOralcle contract.
type FilswanOralcleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOralcle *FilswanOralcleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FilswanOralcleRoleRevokedIterator, error) {

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

	logs, sub, err := _FilswanOralcle.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleRoleRevokedIterator{contract: _FilswanOralcle.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FilswanOralcle *FilswanOralcleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FilswanOralcleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _FilswanOralcle.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOralcleRoleRevoked)
				if err := _FilswanOralcle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_FilswanOralcle *FilswanOralcleFilterer) ParseRoleRevoked(log types.Log) (*FilswanOralcleRoleRevoked, error) {
	event := new(FilswanOralcleRoleRevoked)
	if err := _FilswanOralcle.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOralcleSignTransactionIterator is returned from FilterSignTransaction and is used to iterate over the raw logs and unpacked data for SignTransaction events raised by the FilswanOralcle contract.
type FilswanOralcleSignTransactionIterator struct {
	Event *FilswanOralcleSignTransaction // Event containing the contract specifics and raw log

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
func (it *FilswanOralcleSignTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOralcleSignTransaction)
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
		it.Event = new(FilswanOralcleSignTransaction)
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
func (it *FilswanOralcleSignTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOralcleSignTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOralcleSignTransaction represents a SignTransaction event raised by the FilswanOralcle contract.
type FilswanOralcleSignTransaction struct {
	Cid       string
	OrderId   string
	DealId    string
	Recipient common.Address
	Paid      *big.Int
	Terms     *big.Int
	Status    bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignTransaction is a free log retrieval operation binding the contract event 0x5bdc7ea86f89a3ea9f721b4361cc0dec10ec2981ff3a28092b27784d5812021c.
//
// Solidity: event SignTransaction(string cid, string orderId, string dealId, address recipient, uint256 paid, uint256 terms, bool status)
func (_FilswanOralcle *FilswanOralcleFilterer) FilterSignTransaction(opts *bind.FilterOpts) (*FilswanOralcleSignTransactionIterator, error) {

	logs, sub, err := _FilswanOralcle.contract.FilterLogs(opts, "SignTransaction")
	if err != nil {
		return nil, err
	}
	return &FilswanOralcleSignTransactionIterator{contract: _FilswanOralcle.contract, event: "SignTransaction", logs: logs, sub: sub}, nil
}

// WatchSignTransaction is a free log subscription operation binding the contract event 0x5bdc7ea86f89a3ea9f721b4361cc0dec10ec2981ff3a28092b27784d5812021c.
//
// Solidity: event SignTransaction(string cid, string orderId, string dealId, address recipient, uint256 paid, uint256 terms, bool status)
func (_FilswanOralcle *FilswanOralcleFilterer) WatchSignTransaction(opts *bind.WatchOpts, sink chan<- *FilswanOralcleSignTransaction) (event.Subscription, error) {

	logs, sub, err := _FilswanOralcle.contract.WatchLogs(opts, "SignTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOralcleSignTransaction)
				if err := _FilswanOralcle.contract.UnpackLog(event, "SignTransaction", log); err != nil {
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

// ParseSignTransaction is a log parse operation binding the contract event 0x5bdc7ea86f89a3ea9f721b4361cc0dec10ec2981ff3a28092b27784d5812021c.
//
// Solidity: event SignTransaction(string cid, string orderId, string dealId, address recipient, uint256 paid, uint256 terms, bool status)
func (_FilswanOralcle *FilswanOralcleFilterer) ParseSignTransaction(log types.Log) (*FilswanOralcleSignTransaction, error) {
	event := new(FilswanOralcleSignTransaction)
	if err := _FilswanOralcle.contract.UnpackLog(event, "SignTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
