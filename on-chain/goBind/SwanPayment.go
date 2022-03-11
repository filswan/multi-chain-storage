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

// IPaymentMinimalTxInfo is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimalTxInfo struct {
	Id         string
	Token      common.Address
	MinPayment *big.Int
	LockedFee  *big.Int
	Owner      common.Address
	Recipient  common.Address
	Deadline   *big.Int
	IsExisted  bool
	Size       *big.Int
	CopyLimit  uint8
}

// IPaymentMinimallockPaymentParam is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimallockPaymentParam struct {
	Id         string
	MinPayment *big.Int
	Amount     *big.Int
	LockTime   *big.Int
	Recipient  common.Address
	Size       *big.Int
	CopyLimit  uint8
}

// IPaymentMinimalunlockPaymentParam is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimalunlockPaymentParam struct {
	Id        string
	OrderId   string
	DealId    string
	Amount    *big.Int
	Recipient common.Address
}

// FilswanOracleABI is the input ABI used to generate the binding from.
const FilswanOracleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ExpirePayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"LockPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UnlockPayment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cId\",\"type\":\"string\"}],\"name\":\"getLockedPaymentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isExisted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"copyLimit\",\"type\":\"uint8\"}],\"internalType\":\"structIPaymentMinimal.TxInfo\",\"name\":\"tx\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20_TOKEN\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chainlinkOracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"copyLimit\",\"type\":\"uint8\"}],\"internalType\":\"structIPaymentMinimal.lockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"lockTokenPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlinkOracle\",\"type\":\"address\"}],\"name\":\"setChainlinkOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unlockCarPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIPaymentMinimal.unlockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"unlockTokenPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_FilswanOracle *FilswanOracleCaller) NATIVETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "NATIVE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_FilswanOracle *FilswanOracleSession) NATIVETOKEN() (common.Address, error) {
	return _FilswanOracle.Contract.NATIVETOKEN(&_FilswanOracle.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_FilswanOracle *FilswanOracleCallerSession) NATIVETOKEN() (common.Address, error) {
	return _FilswanOracle.Contract.NATIVETOKEN(&_FilswanOracle.CallOpts)
}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string cId) view returns((string,address,uint256,uint256,address,address,uint256,bool,uint256,uint8) tx)
func (_FilswanOracle *FilswanOracleCaller) GetLockedPaymentInfo(opts *bind.CallOpts, cId string) (IPaymentMinimalTxInfo, error) {
	var out []interface{}
	err := _FilswanOracle.contract.Call(opts, &out, "getLockedPaymentInfo", cId)

	if err != nil {
		return *new(IPaymentMinimalTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPaymentMinimalTxInfo)).(*IPaymentMinimalTxInfo)

	return out0, err

}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string cId) view returns((string,address,uint256,uint256,address,address,uint256,bool,uint256,uint8) tx)
func (_FilswanOracle *FilswanOracleSession) GetLockedPaymentInfo(cId string) (IPaymentMinimalTxInfo, error) {
	return _FilswanOracle.Contract.GetLockedPaymentInfo(&_FilswanOracle.CallOpts, cId)
}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string cId) view returns((string,address,uint256,uint256,address,address,uint256,bool,uint256,uint8) tx)
func (_FilswanOracle *FilswanOracleCallerSession) GetLockedPaymentInfo(cId string) (IPaymentMinimalTxInfo, error) {
	return _FilswanOracle.Contract.GetLockedPaymentInfo(&_FilswanOracle.CallOpts, cId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed, address chainlinkOracle) returns()
func (_FilswanOracle *FilswanOracleTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, ERC20_TOKEN common.Address, oracle common.Address, priceFeed common.Address, chainlinkOracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "initialize", owner, ERC20_TOKEN, oracle, priceFeed, chainlinkOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed, address chainlinkOracle) returns()
func (_FilswanOracle *FilswanOracleSession) Initialize(owner common.Address, ERC20_TOKEN common.Address, oracle common.Address, priceFeed common.Address, chainlinkOracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Initialize(&_FilswanOracle.TransactOpts, owner, ERC20_TOKEN, oracle, priceFeed, chainlinkOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed, address chainlinkOracle) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) Initialize(owner common.Address, ERC20_TOKEN common.Address, oracle common.Address, priceFeed common.Address, chainlinkOracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Initialize(&_FilswanOracle.TransactOpts, owner, ERC20_TOKEN, oracle, priceFeed, chainlinkOracle)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) LockTokenPayment(opts *bind.TransactOpts, param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "lockTokenPayment", param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_FilswanOracle *FilswanOracleSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _FilswanOracle.Contract.LockTokenPayment(&_FilswanOracle.TransactOpts, param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xf4d98717.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256,uint8) param) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _FilswanOracle.Contract.LockTokenPayment(&_FilswanOracle.TransactOpts, param)
}

// Refund is a paid mutator transaction binding the contract method 0x7d29985b.
//
// Solidity: function refund(string[] cidList) returns()
func (_FilswanOracle *FilswanOracleTransactor) Refund(opts *bind.TransactOpts, cidList []string) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "refund", cidList)
}

// Refund is a paid mutator transaction binding the contract method 0x7d29985b.
//
// Solidity: function refund(string[] cidList) returns()
func (_FilswanOracle *FilswanOracleSession) Refund(cidList []string) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Refund(&_FilswanOracle.TransactOpts, cidList)
}

// Refund is a paid mutator transaction binding the contract method 0x7d29985b.
//
// Solidity: function refund(string[] cidList) returns()
func (_FilswanOracle *FilswanOracleTransactorSession) Refund(cidList []string) (*types.Transaction, error) {
	return _FilswanOracle.Contract.Refund(&_FilswanOracle.TransactOpts, cidList)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x7a9b0412.
//
// Solidity: function setChainlinkOracle(address _chainlinkOracle) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) SetChainlinkOracle(opts *bind.TransactOpts, _chainlinkOracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "setChainlinkOracle", _chainlinkOracle)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x7a9b0412.
//
// Solidity: function setChainlinkOracle(address _chainlinkOracle) returns(bool)
func (_FilswanOracle *FilswanOracleSession) SetChainlinkOracle(_chainlinkOracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetChainlinkOracle(&_FilswanOracle.TransactOpts, _chainlinkOracle)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x7a9b0412.
//
// Solidity: function setChainlinkOracle(address _chainlinkOracle) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) SetChainlinkOracle(_chainlinkOracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetChainlinkOracle(&_FilswanOracle.TransactOpts, _chainlinkOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_FilswanOracle *FilswanOracleSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetOracle(&_FilswanOracle.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetOracle(&_FilswanOracle.TransactOpts, oracle)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) SetPriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "setPriceFeed", priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns(bool)
func (_FilswanOracle *FilswanOracleSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetPriceFeed(&_FilswanOracle.TransactOpts, priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.SetPriceFeed(&_FilswanOracle.TransactOpts, priceFeed)
}

// UnlockCarPayment is a paid mutator transaction binding the contract method 0x38b3c23b.
//
// Solidity: function unlockCarPayment(string dealId, address recipient) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) UnlockCarPayment(opts *bind.TransactOpts, dealId string, recipient common.Address) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "unlockCarPayment", dealId, recipient)
}

// UnlockCarPayment is a paid mutator transaction binding the contract method 0x38b3c23b.
//
// Solidity: function unlockCarPayment(string dealId, address recipient) returns(bool)
func (_FilswanOracle *FilswanOracleSession) UnlockCarPayment(dealId string, recipient common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UnlockCarPayment(&_FilswanOracle.TransactOpts, dealId, recipient)
}

// UnlockCarPayment is a paid mutator transaction binding the contract method 0x38b3c23b.
//
// Solidity: function unlockCarPayment(string dealId, address recipient) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) UnlockCarPayment(dealId string, recipient common.Address) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UnlockCarPayment(&_FilswanOracle.TransactOpts, dealId, recipient)
}

// UnlockTokenPayment is a paid mutator transaction binding the contract method 0x5c95e7e1.
//
// Solidity: function unlockTokenPayment((string,string,string,uint256,address) param) returns(bool)
func (_FilswanOracle *FilswanOracleTransactor) UnlockTokenPayment(opts *bind.TransactOpts, param IPaymentMinimalunlockPaymentParam) (*types.Transaction, error) {
	return _FilswanOracle.contract.Transact(opts, "unlockTokenPayment", param)
}

// UnlockTokenPayment is a paid mutator transaction binding the contract method 0x5c95e7e1.
//
// Solidity: function unlockTokenPayment((string,string,string,uint256,address) param) returns(bool)
func (_FilswanOracle *FilswanOracleSession) UnlockTokenPayment(param IPaymentMinimalunlockPaymentParam) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UnlockTokenPayment(&_FilswanOracle.TransactOpts, param)
}

// UnlockTokenPayment is a paid mutator transaction binding the contract method 0x5c95e7e1.
//
// Solidity: function unlockTokenPayment((string,string,string,uint256,address) param) returns(bool)
func (_FilswanOracle *FilswanOracleTransactorSession) UnlockTokenPayment(param IPaymentMinimalunlockPaymentParam) (*types.Transaction, error) {
	return _FilswanOracle.Contract.UnlockTokenPayment(&_FilswanOracle.TransactOpts, param)
}

// FilswanOracleExpirePaymentIterator is returned from FilterExpirePayment and is used to iterate over the raw logs and unpacked data for ExpirePayment events raised by the FilswanOracle contract.
type FilswanOracleExpirePaymentIterator struct {
	Event *FilswanOracleExpirePayment // Event containing the contract specifics and raw log

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
func (it *FilswanOracleExpirePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleExpirePayment)
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
		it.Event = new(FilswanOracleExpirePayment)
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
func (it *FilswanOracleExpirePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleExpirePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleExpirePayment represents a ExpirePayment event raised by the FilswanOracle contract.
type FilswanOracleExpirePayment struct {
	Id     string
	Token  common.Address
	Amount *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExpirePayment is a free log retrieval operation binding the contract event 0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a.
//
// Solidity: event ExpirePayment(string id, address token, uint256 amount, address owner)
func (_FilswanOracle *FilswanOracleFilterer) FilterExpirePayment(opts *bind.FilterOpts) (*FilswanOracleExpirePaymentIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "ExpirePayment")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleExpirePaymentIterator{contract: _FilswanOracle.contract, event: "ExpirePayment", logs: logs, sub: sub}, nil
}

// WatchExpirePayment is a free log subscription operation binding the contract event 0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a.
//
// Solidity: event ExpirePayment(string id, address token, uint256 amount, address owner)
func (_FilswanOracle *FilswanOracleFilterer) WatchExpirePayment(opts *bind.WatchOpts, sink chan<- *FilswanOracleExpirePayment) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "ExpirePayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleExpirePayment)
				if err := _FilswanOracle.contract.UnpackLog(event, "ExpirePayment", log); err != nil {
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

// ParseExpirePayment is a log parse operation binding the contract event 0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a.
//
// Solidity: event ExpirePayment(string id, address token, uint256 amount, address owner)
func (_FilswanOracle *FilswanOracleFilterer) ParseExpirePayment(log types.Log) (*FilswanOracleExpirePayment, error) {
	event := new(FilswanOracleExpirePayment)
	if err := _FilswanOracle.contract.UnpackLog(event, "ExpirePayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleLockPaymentIterator is returned from FilterLockPayment and is used to iterate over the raw logs and unpacked data for LockPayment events raised by the FilswanOracle contract.
type FilswanOracleLockPaymentIterator struct {
	Event *FilswanOracleLockPayment // Event containing the contract specifics and raw log

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
func (it *FilswanOracleLockPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleLockPayment)
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
		it.Event = new(FilswanOracleLockPayment)
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
func (it *FilswanOracleLockPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleLockPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleLockPayment represents a LockPayment event raised by the FilswanOracle contract.
type FilswanOracleLockPayment struct {
	Id         string
	Token      common.Address
	LockedFee  *big.Int
	MinPayment *big.Int
	Recipient  common.Address
	Deadline   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockPayment is a free log retrieval operation binding the contract event 0x87b076911cb43e46012dbf762607ad5d0c5d6eb5e1d8fec72be99002451c18ba.
//
// Solidity: event LockPayment(string id, address token, uint256 lockedFee, uint256 minPayment, address recipient, uint256 deadline)
func (_FilswanOracle *FilswanOracleFilterer) FilterLockPayment(opts *bind.FilterOpts) (*FilswanOracleLockPaymentIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "LockPayment")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleLockPaymentIterator{contract: _FilswanOracle.contract, event: "LockPayment", logs: logs, sub: sub}, nil
}

// WatchLockPayment is a free log subscription operation binding the contract event 0x87b076911cb43e46012dbf762607ad5d0c5d6eb5e1d8fec72be99002451c18ba.
//
// Solidity: event LockPayment(string id, address token, uint256 lockedFee, uint256 minPayment, address recipient, uint256 deadline)
func (_FilswanOracle *FilswanOracleFilterer) WatchLockPayment(opts *bind.WatchOpts, sink chan<- *FilswanOracleLockPayment) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "LockPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleLockPayment)
				if err := _FilswanOracle.contract.UnpackLog(event, "LockPayment", log); err != nil {
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

// ParseLockPayment is a log parse operation binding the contract event 0x87b076911cb43e46012dbf762607ad5d0c5d6eb5e1d8fec72be99002451c18ba.
//
// Solidity: event LockPayment(string id, address token, uint256 lockedFee, uint256 minPayment, address recipient, uint256 deadline)
func (_FilswanOracle *FilswanOracleFilterer) ParseLockPayment(log types.Log) (*FilswanOracleLockPayment, error) {
	event := new(FilswanOracleLockPayment)
	if err := _FilswanOracle.contract.UnpackLog(event, "LockPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FilswanOracleUnlockPaymentIterator is returned from FilterUnlockPayment and is used to iterate over the raw logs and unpacked data for UnlockPayment events raised by the FilswanOracle contract.
type FilswanOracleUnlockPaymentIterator struct {
	Event *FilswanOracleUnlockPayment // Event containing the contract specifics and raw log

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
func (it *FilswanOracleUnlockPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FilswanOracleUnlockPayment)
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
		it.Event = new(FilswanOracleUnlockPayment)
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
func (it *FilswanOracleUnlockPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FilswanOracleUnlockPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FilswanOracleUnlockPayment represents a UnlockPayment event raised by the FilswanOracle contract.
type FilswanOracleUnlockPayment struct {
	Id        string
	Token     common.Address
	Cost      *big.Int
	RestToken *big.Int
	Recipient common.Address
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockPayment is a free log retrieval operation binding the contract event 0xd9aab088cd193ab56c9a682b9f2e42fb39206b42c3a6ba78130dc97a2a8d03ee.
//
// Solidity: event UnlockPayment(string id, address token, uint256 cost, uint256 restToken, address recipient, address owner)
func (_FilswanOracle *FilswanOracleFilterer) FilterUnlockPayment(opts *bind.FilterOpts) (*FilswanOracleUnlockPaymentIterator, error) {

	logs, sub, err := _FilswanOracle.contract.FilterLogs(opts, "UnlockPayment")
	if err != nil {
		return nil, err
	}
	return &FilswanOracleUnlockPaymentIterator{contract: _FilswanOracle.contract, event: "UnlockPayment", logs: logs, sub: sub}, nil
}

// WatchUnlockPayment is a free log subscription operation binding the contract event 0xd9aab088cd193ab56c9a682b9f2e42fb39206b42c3a6ba78130dc97a2a8d03ee.
//
// Solidity: event UnlockPayment(string id, address token, uint256 cost, uint256 restToken, address recipient, address owner)
func (_FilswanOracle *FilswanOracleFilterer) WatchUnlockPayment(opts *bind.WatchOpts, sink chan<- *FilswanOracleUnlockPayment) (event.Subscription, error) {

	logs, sub, err := _FilswanOracle.contract.WatchLogs(opts, "UnlockPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FilswanOracleUnlockPayment)
				if err := _FilswanOracle.contract.UnpackLog(event, "UnlockPayment", log); err != nil {
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

// ParseUnlockPayment is a log parse operation binding the contract event 0xd9aab088cd193ab56c9a682b9f2e42fb39206b42c3a6ba78130dc97a2a8d03ee.
//
// Solidity: event UnlockPayment(string id, address token, uint256 cost, uint256 restToken, address recipient, address owner)
func (_FilswanOracle *FilswanOracleFilterer) ParseUnlockPayment(log types.Log) (*FilswanOracleUnlockPayment, error) {
	event := new(FilswanOracleUnlockPayment)
	if err := _FilswanOracle.contract.UnpackLog(event, "UnlockPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
