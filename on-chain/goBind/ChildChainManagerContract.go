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

// ChildChainManagerContractABI is the input ABI used to generate the binding from.
const ChildChainManagerContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"}],\"name\":\"TokenMapped\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAPPER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAP_TOKEN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATE_SYNCER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"childToRootToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rootToChildToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"}],\"name\":\"mapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onStateReceive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rootToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"childToken\",\"type\":\"address\"}],\"name\":\"cleanMapToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ChildChainManagerContract is an auto generated Go binding around an Ethereum contract.
type ChildChainManagerContract struct {
	ChildChainManagerContractCaller     // Read-only binding to the contract
	ChildChainManagerContractTransactor // Write-only binding to the contract
	ChildChainManagerContractFilterer   // Log filterer for contract events
}

// ChildChainManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildChainManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildChainManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildChainManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildChainManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildChainManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildChainManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildChainManagerContractSession struct {
	Contract     *ChildChainManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChildChainManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildChainManagerContractCallerSession struct {
	Contract *ChildChainManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ChildChainManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildChainManagerContractTransactorSession struct {
	Contract     *ChildChainManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ChildChainManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildChainManagerContractRaw struct {
	Contract *ChildChainManagerContract // Generic contract binding to access the raw methods on
}

// ChildChainManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildChainManagerContractCallerRaw struct {
	Contract *ChildChainManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// ChildChainManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildChainManagerContractTransactorRaw struct {
	Contract *ChildChainManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildChainManagerContract creates a new instance of ChildChainManagerContract, bound to a specific deployed contract.
func NewChildChainManagerContract(address common.Address, backend bind.ContractBackend) (*ChildChainManagerContract, error) {
	contract, err := bindChildChainManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContract{ChildChainManagerContractCaller: ChildChainManagerContractCaller{contract: contract}, ChildChainManagerContractTransactor: ChildChainManagerContractTransactor{contract: contract}, ChildChainManagerContractFilterer: ChildChainManagerContractFilterer{contract: contract}}, nil
}

// NewChildChainManagerContractCaller creates a new read-only instance of ChildChainManagerContract, bound to a specific deployed contract.
func NewChildChainManagerContractCaller(address common.Address, caller bind.ContractCaller) (*ChildChainManagerContractCaller, error) {
	contract, err := bindChildChainManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractCaller{contract: contract}, nil
}

// NewChildChainManagerContractTransactor creates a new write-only instance of ChildChainManagerContract, bound to a specific deployed contract.
func NewChildChainManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildChainManagerContractTransactor, error) {
	contract, err := bindChildChainManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractTransactor{contract: contract}, nil
}

// NewChildChainManagerContractFilterer creates a new log filterer instance of ChildChainManagerContract, bound to a specific deployed contract.
func NewChildChainManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildChainManagerContractFilterer, error) {
	contract, err := bindChildChainManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractFilterer{contract: contract}, nil
}

// bindChildChainManagerContract binds a generic wrapper to an already deployed contract.
func bindChildChainManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildChainManagerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildChainManagerContract *ChildChainManagerContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildChainManagerContract.Contract.ChildChainManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildChainManagerContract *ChildChainManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.ChildChainManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildChainManagerContract *ChildChainManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.ChildChainManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildChainManagerContract *ChildChainManagerContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChildChainManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildChainManagerContract *ChildChainManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildChainManagerContract *ChildChainManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.DEFAULTADMINROLE(&_ChildChainManagerContract.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.DEFAULTADMINROLE(&_ChildChainManagerContract.CallOpts)
}

// DEPOSIT is a free data retrieval call binding the contract method 0xd81c8e52.
//
// Solidity: function DEPOSIT() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) DEPOSIT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "DEPOSIT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSIT is a free data retrieval call binding the contract method 0xd81c8e52.
//
// Solidity: function DEPOSIT() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractSession) DEPOSIT() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.DEPOSIT(&_ChildChainManagerContract.CallOpts)
}

// DEPOSIT is a free data retrieval call binding the contract method 0xd81c8e52.
//
// Solidity: function DEPOSIT() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) DEPOSIT() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.DEPOSIT(&_ChildChainManagerContract.CallOpts)
}

// MAPPERROLE is a free data retrieval call binding the contract method 0x568b80b5.
//
// Solidity: function MAPPER_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) MAPPERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "MAPPER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAPPERROLE is a free data retrieval call binding the contract method 0x568b80b5.
//
// Solidity: function MAPPER_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractSession) MAPPERROLE() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.MAPPERROLE(&_ChildChainManagerContract.CallOpts)
}

// MAPPERROLE is a free data retrieval call binding the contract method 0x568b80b5.
//
// Solidity: function MAPPER_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) MAPPERROLE() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.MAPPERROLE(&_ChildChainManagerContract.CallOpts)
}

// MAPTOKEN is a free data retrieval call binding the contract method 0x886a69ba.
//
// Solidity: function MAP_TOKEN() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) MAPTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "MAP_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MAPTOKEN is a free data retrieval call binding the contract method 0x886a69ba.
//
// Solidity: function MAP_TOKEN() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractSession) MAPTOKEN() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.MAPTOKEN(&_ChildChainManagerContract.CallOpts)
}

// MAPTOKEN is a free data retrieval call binding the contract method 0x886a69ba.
//
// Solidity: function MAP_TOKEN() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) MAPTOKEN() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.MAPTOKEN(&_ChildChainManagerContract.CallOpts)
}

// STATESYNCERROLE is a free data retrieval call binding the contract method 0x4dee4498.
//
// Solidity: function STATE_SYNCER_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) STATESYNCERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "STATE_SYNCER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STATESYNCERROLE is a free data retrieval call binding the contract method 0x4dee4498.
//
// Solidity: function STATE_SYNCER_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractSession) STATESYNCERROLE() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.STATESYNCERROLE(&_ChildChainManagerContract.CallOpts)
}

// STATESYNCERROLE is a free data retrieval call binding the contract method 0x4dee4498.
//
// Solidity: function STATE_SYNCER_ROLE() view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) STATESYNCERROLE() ([32]byte, error) {
	return _ChildChainManagerContract.Contract.STATESYNCERROLE(&_ChildChainManagerContract.CallOpts)
}

// ChildToRootToken is a free data retrieval call binding the contract method 0x6e86b770.
//
// Solidity: function childToRootToken(address ) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) ChildToRootToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "childToRootToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildToRootToken is a free data retrieval call binding the contract method 0x6e86b770.
//
// Solidity: function childToRootToken(address ) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractSession) ChildToRootToken(arg0 common.Address) (common.Address, error) {
	return _ChildChainManagerContract.Contract.ChildToRootToken(&_ChildChainManagerContract.CallOpts, arg0)
}

// ChildToRootToken is a free data retrieval call binding the contract method 0x6e86b770.
//
// Solidity: function childToRootToken(address ) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) ChildToRootToken(arg0 common.Address) (common.Address, error) {
	return _ChildChainManagerContract.Contract.ChildToRootToken(&_ChildChainManagerContract.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChildChainManagerContract.Contract.GetRoleAdmin(&_ChildChainManagerContract.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ChildChainManagerContract.Contract.GetRoleAdmin(&_ChildChainManagerContract.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChildChainManagerContract.Contract.GetRoleMember(&_ChildChainManagerContract.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ChildChainManagerContract.Contract.GetRoleMember(&_ChildChainManagerContract.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChildChainManagerContract *ChildChainManagerContractSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChildChainManagerContract.Contract.GetRoleMemberCount(&_ChildChainManagerContract.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ChildChainManagerContract.Contract.GetRoleMemberCount(&_ChildChainManagerContract.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChildChainManagerContract *ChildChainManagerContractSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChildChainManagerContract.Contract.HasRole(&_ChildChainManagerContract.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ChildChainManagerContract.Contract.HasRole(&_ChildChainManagerContract.CallOpts, role, account)
}

// RootToChildToken is a free data retrieval call binding the contract method 0xea60c7c4.
//
// Solidity: function rootToChildToken(address ) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractCaller) RootToChildToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ChildChainManagerContract.contract.Call(opts, &out, "rootToChildToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RootToChildToken is a free data retrieval call binding the contract method 0xea60c7c4.
//
// Solidity: function rootToChildToken(address ) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractSession) RootToChildToken(arg0 common.Address) (common.Address, error) {
	return _ChildChainManagerContract.Contract.RootToChildToken(&_ChildChainManagerContract.CallOpts, arg0)
}

// RootToChildToken is a free data retrieval call binding the contract method 0xea60c7c4.
//
// Solidity: function rootToChildToken(address ) view returns(address)
func (_ChildChainManagerContract *ChildChainManagerContractCallerSession) RootToChildToken(arg0 common.Address) (common.Address, error) {
	return _ChildChainManagerContract.Contract.RootToChildToken(&_ChildChainManagerContract.CallOpts, arg0)
}

// CleanMapToken is a paid mutator transaction binding the contract method 0x0c3894bb.
//
// Solidity: function cleanMapToken(address rootToken, address childToken) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) CleanMapToken(opts *bind.TransactOpts, rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "cleanMapToken", rootToken, childToken)
}

// CleanMapToken is a paid mutator transaction binding the contract method 0x0c3894bb.
//
// Solidity: function cleanMapToken(address rootToken, address childToken) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) CleanMapToken(rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.CleanMapToken(&_ChildChainManagerContract.TransactOpts, rootToken, childToken)
}

// CleanMapToken is a paid mutator transaction binding the contract method 0x0c3894bb.
//
// Solidity: function cleanMapToken(address rootToken, address childToken) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) CleanMapToken(rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.CleanMapToken(&_ChildChainManagerContract.TransactOpts, rootToken, childToken)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.GrantRole(&_ChildChainManagerContract.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.GrantRole(&_ChildChainManagerContract.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.Initialize(&_ChildChainManagerContract.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.Initialize(&_ChildChainManagerContract.TransactOpts, _owner)
}

// MapToken is a paid mutator transaction binding the contract method 0x47400269.
//
// Solidity: function mapToken(address rootToken, address childToken) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) MapToken(opts *bind.TransactOpts, rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "mapToken", rootToken, childToken)
}

// MapToken is a paid mutator transaction binding the contract method 0x47400269.
//
// Solidity: function mapToken(address rootToken, address childToken) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) MapToken(rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.MapToken(&_ChildChainManagerContract.TransactOpts, rootToken, childToken)
}

// MapToken is a paid mutator transaction binding the contract method 0x47400269.
//
// Solidity: function mapToken(address rootToken, address childToken) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) MapToken(rootToken common.Address, childToken common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.MapToken(&_ChildChainManagerContract.TransactOpts, rootToken, childToken)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0x26c53bea.
//
// Solidity: function onStateReceive(uint256 , bytes data) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) OnStateReceive(opts *bind.TransactOpts, arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "onStateReceive", arg0, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0x26c53bea.
//
// Solidity: function onStateReceive(uint256 , bytes data) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) OnStateReceive(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.OnStateReceive(&_ChildChainManagerContract.TransactOpts, arg0, data)
}

// OnStateReceive is a paid mutator transaction binding the contract method 0x26c53bea.
//
// Solidity: function onStateReceive(uint256 , bytes data) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) OnStateReceive(arg0 *big.Int, data []byte) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.OnStateReceive(&_ChildChainManagerContract.TransactOpts, arg0, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.RenounceRole(&_ChildChainManagerContract.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.RenounceRole(&_ChildChainManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.RevokeRole(&_ChildChainManagerContract.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ChildChainManagerContract *ChildChainManagerContractTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ChildChainManagerContract.Contract.RevokeRole(&_ChildChainManagerContract.TransactOpts, role, account)
}

// ChildChainManagerContractRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ChildChainManagerContract contract.
type ChildChainManagerContractRoleAdminChangedIterator struct {
	Event *ChildChainManagerContractRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ChildChainManagerContractRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainManagerContractRoleAdminChanged)
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
		it.Event = new(ChildChainManagerContractRoleAdminChanged)
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
func (it *ChildChainManagerContractRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainManagerContractRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainManagerContractRoleAdminChanged represents a RoleAdminChanged event raised by the ChildChainManagerContract contract.
type ChildChainManagerContractRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ChildChainManagerContractRoleAdminChangedIterator, error) {

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

	logs, sub, err := _ChildChainManagerContract.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractRoleAdminChangedIterator{contract: _ChildChainManagerContract.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ChildChainManagerContractRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _ChildChainManagerContract.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainManagerContractRoleAdminChanged)
				if err := _ChildChainManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) ParseRoleAdminChanged(log types.Log) (*ChildChainManagerContractRoleAdminChanged, error) {
	event := new(ChildChainManagerContractRoleAdminChanged)
	if err := _ChildChainManagerContract.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildChainManagerContractRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ChildChainManagerContract contract.
type ChildChainManagerContractRoleGrantedIterator struct {
	Event *ChildChainManagerContractRoleGranted // Event containing the contract specifics and raw log

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
func (it *ChildChainManagerContractRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainManagerContractRoleGranted)
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
		it.Event = new(ChildChainManagerContractRoleGranted)
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
func (it *ChildChainManagerContractRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainManagerContractRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainManagerContractRoleGranted represents a RoleGranted event raised by the ChildChainManagerContract contract.
type ChildChainManagerContractRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChildChainManagerContractRoleGrantedIterator, error) {

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

	logs, sub, err := _ChildChainManagerContract.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractRoleGrantedIterator{contract: _ChildChainManagerContract.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ChildChainManagerContractRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ChildChainManagerContract.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainManagerContractRoleGranted)
				if err := _ChildChainManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) ParseRoleGranted(log types.Log) (*ChildChainManagerContractRoleGranted, error) {
	event := new(ChildChainManagerContractRoleGranted)
	if err := _ChildChainManagerContract.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildChainManagerContractRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ChildChainManagerContract contract.
type ChildChainManagerContractRoleRevokedIterator struct {
	Event *ChildChainManagerContractRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ChildChainManagerContractRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainManagerContractRoleRevoked)
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
		it.Event = new(ChildChainManagerContractRoleRevoked)
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
func (it *ChildChainManagerContractRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainManagerContractRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainManagerContractRoleRevoked represents a RoleRevoked event raised by the ChildChainManagerContract contract.
type ChildChainManagerContractRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ChildChainManagerContractRoleRevokedIterator, error) {

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

	logs, sub, err := _ChildChainManagerContract.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractRoleRevokedIterator{contract: _ChildChainManagerContract.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ChildChainManagerContractRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ChildChainManagerContract.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainManagerContractRoleRevoked)
				if err := _ChildChainManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) ParseRoleRevoked(log types.Log) (*ChildChainManagerContractRoleRevoked, error) {
	event := new(ChildChainManagerContractRoleRevoked)
	if err := _ChildChainManagerContract.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildChainManagerContractTokenMappedIterator is returned from FilterTokenMapped and is used to iterate over the raw logs and unpacked data for TokenMapped events raised by the ChildChainManagerContract contract.
type ChildChainManagerContractTokenMappedIterator struct {
	Event *ChildChainManagerContractTokenMapped // Event containing the contract specifics and raw log

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
func (it *ChildChainManagerContractTokenMappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainManagerContractTokenMapped)
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
		it.Event = new(ChildChainManagerContractTokenMapped)
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
func (it *ChildChainManagerContractTokenMappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainManagerContractTokenMappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainManagerContractTokenMapped represents a TokenMapped event raised by the ChildChainManagerContract contract.
type ChildChainManagerContractTokenMapped struct {
	RootToken  common.Address
	ChildToken common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenMapped is a free log retrieval operation binding the contract event 0x85920d35e6c72f6b2affffa04298b0cecfeba86e4a9f407df661f1cb8ab5e617.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) FilterTokenMapped(opts *bind.FilterOpts, rootToken []common.Address, childToken []common.Address) (*ChildChainManagerContractTokenMappedIterator, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	logs, sub, err := _ChildChainManagerContract.contract.FilterLogs(opts, "TokenMapped", rootTokenRule, childTokenRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainManagerContractTokenMappedIterator{contract: _ChildChainManagerContract.contract, event: "TokenMapped", logs: logs, sub: sub}, nil
}

// WatchTokenMapped is a free log subscription operation binding the contract event 0x85920d35e6c72f6b2affffa04298b0cecfeba86e4a9f407df661f1cb8ab5e617.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) WatchTokenMapped(opts *bind.WatchOpts, sink chan<- *ChildChainManagerContractTokenMapped, rootToken []common.Address, childToken []common.Address) (event.Subscription, error) {

	var rootTokenRule []interface{}
	for _, rootTokenItem := range rootToken {
		rootTokenRule = append(rootTokenRule, rootTokenItem)
	}
	var childTokenRule []interface{}
	for _, childTokenItem := range childToken {
		childTokenRule = append(childTokenRule, childTokenItem)
	}

	logs, sub, err := _ChildChainManagerContract.contract.WatchLogs(opts, "TokenMapped", rootTokenRule, childTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainManagerContractTokenMapped)
				if err := _ChildChainManagerContract.contract.UnpackLog(event, "TokenMapped", log); err != nil {
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

// ParseTokenMapped is a log parse operation binding the contract event 0x85920d35e6c72f6b2affffa04298b0cecfeba86e4a9f407df661f1cb8ab5e617.
//
// Solidity: event TokenMapped(address indexed rootToken, address indexed childToken)
func (_ChildChainManagerContract *ChildChainManagerContractFilterer) ParseTokenMapped(log types.Log) (*ChildChainManagerContractTokenMapped, error) {
	event := new(ChildChainManagerContractTokenMapped)
	if err := _ChildChainManagerContract.contract.UnpackLog(event, "TokenMapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
