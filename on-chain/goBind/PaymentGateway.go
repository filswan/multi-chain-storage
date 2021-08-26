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
	MinPayment *big.Int
	LockedFee  *big.Int
	Owner      common.Address
	Recipient  common.Address
	Deadline   *big.Int
	IsExisted  bool
}

// IPaymentMinimallockPaymentParam is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimallockPaymentParam struct {
	Id         string
	MinPayment *big.Int
	LockTime   *big.Int
	Recipient  common.Address
}

// PaymentGatewayABI is the input ABI used to generate the binding from.
const PaymentGatewayABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"getLockedPaymentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isExisted\",\"type\":\"bool\"}],\"internalType\":\"structIPaymentMinimal.TxInfo\",\"name\":\"tx\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIPaymentMinimal.lockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"lockPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"unlockPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PaymentGateway is an auto generated Go binding around an Ethereum contract.
type PaymentGateway struct {
	PaymentGatewayCaller     // Read-only binding to the contract
	PaymentGatewayTransactor // Write-only binding to the contract
	PaymentGatewayFilterer   // Log filterer for contract events
}

// PaymentGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentGatewaySession struct {
	Contract     *PaymentGateway   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentGatewayCallerSession struct {
	Contract *PaymentGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PaymentGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentGatewayTransactorSession struct {
	Contract     *PaymentGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PaymentGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentGatewayRaw struct {
	Contract *PaymentGateway // Generic contract binding to access the raw methods on
}

// PaymentGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentGatewayCallerRaw struct {
	Contract *PaymentGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentGatewayTransactorRaw struct {
	Contract *PaymentGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentGateway creates a new instance of PaymentGateway, bound to a specific deployed contract.
func NewPaymentGateway(address common.Address, backend bind.ContractBackend) (*PaymentGateway, error) {
	contract, err := bindPaymentGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentGateway{PaymentGatewayCaller: PaymentGatewayCaller{contract: contract}, PaymentGatewayTransactor: PaymentGatewayTransactor{contract: contract}, PaymentGatewayFilterer: PaymentGatewayFilterer{contract: contract}}, nil
}

// NewPaymentGatewayCaller creates a new read-only instance of PaymentGateway, bound to a specific deployed contract.
func NewPaymentGatewayCaller(address common.Address, caller bind.ContractCaller) (*PaymentGatewayCaller, error) {
	contract, err := bindPaymentGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentGatewayCaller{contract: contract}, nil
}

// NewPaymentGatewayTransactor creates a new write-only instance of PaymentGateway, bound to a specific deployed contract.
func NewPaymentGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentGatewayTransactor, error) {
	contract, err := bindPaymentGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentGatewayTransactor{contract: contract}, nil
}

// NewPaymentGatewayFilterer creates a new log filterer instance of PaymentGateway, bound to a specific deployed contract.
func NewPaymentGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentGatewayFilterer, error) {
	contract, err := bindPaymentGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentGatewayFilterer{contract: contract}, nil
}

// bindPaymentGateway binds a generic wrapper to an already deployed contract.
func bindPaymentGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentGatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentGateway *PaymentGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentGateway.Contract.PaymentGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentGateway *PaymentGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentGateway.Contract.PaymentGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentGateway *PaymentGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentGateway.Contract.PaymentGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentGateway *PaymentGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentGateway *PaymentGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentGateway *PaymentGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentGateway.Contract.contract.Transact(opts, method, params...)
}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string txId) view returns((string,uint256,uint256,address,address,uint256,bool) tx)
func (_PaymentGateway *PaymentGatewayCaller) GetLockedPaymentInfo(opts *bind.CallOpts, txId string) (IPaymentMinimalTxInfo, error) {
	var out []interface{}
	err := _PaymentGateway.contract.Call(opts, &out, "getLockedPaymentInfo", txId)

	if err != nil {
		return *new(IPaymentMinimalTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPaymentMinimalTxInfo)).(*IPaymentMinimalTxInfo)

	return out0, err

}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string txId) view returns((string,uint256,uint256,address,address,uint256,bool) tx)
func (_PaymentGateway *PaymentGatewaySession) GetLockedPaymentInfo(txId string) (IPaymentMinimalTxInfo, error) {
	return _PaymentGateway.Contract.GetLockedPaymentInfo(&_PaymentGateway.CallOpts, txId)
}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string txId) view returns((string,uint256,uint256,address,address,uint256,bool) tx)
func (_PaymentGateway *PaymentGatewayCallerSession) GetLockedPaymentInfo(txId string) (IPaymentMinimalTxInfo, error) {
	return _PaymentGateway.Contract.GetLockedPaymentInfo(&_PaymentGateway.CallOpts, txId)
}

// LockPayment is a paid mutator transaction binding the contract method 0x0fe4b536.
//
// Solidity: function lockPayment((string,uint256,uint256,address) param) payable returns(bool)
func (_PaymentGateway *PaymentGatewayTransactor) LockPayment(opts *bind.TransactOpts, param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _PaymentGateway.contract.Transact(opts, "lockPayment", param)
}

// LockPayment is a paid mutator transaction binding the contract method 0x0fe4b536.
//
// Solidity: function lockPayment((string,uint256,uint256,address) param) payable returns(bool)
func (_PaymentGateway *PaymentGatewaySession) LockPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _PaymentGateway.Contract.LockPayment(&_PaymentGateway.TransactOpts, param)
}

// LockPayment is a paid mutator transaction binding the contract method 0x0fe4b536.
//
// Solidity: function lockPayment((string,uint256,uint256,address) param) payable returns(bool)
func (_PaymentGateway *PaymentGatewayTransactorSession) LockPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _PaymentGateway.Contract.LockPayment(&_PaymentGateway.TransactOpts, param)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_PaymentGateway *PaymentGatewayTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _PaymentGateway.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_PaymentGateway *PaymentGatewaySession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _PaymentGateway.Contract.SetOracle(&_PaymentGateway.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_PaymentGateway *PaymentGatewayTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _PaymentGateway.Contract.SetOracle(&_PaymentGateway.TransactOpts, oracle)
}

// UnlockPayment is a paid mutator transaction binding the contract method 0xe01d7646.
//
// Solidity: function unlockPayment(string txId) returns(bool)
func (_PaymentGateway *PaymentGatewayTransactor) UnlockPayment(opts *bind.TransactOpts, txId string) (*types.Transaction, error) {
	return _PaymentGateway.contract.Transact(opts, "unlockPayment", txId)
}

// UnlockPayment is a paid mutator transaction binding the contract method 0xe01d7646.
//
// Solidity: function unlockPayment(string txId) returns(bool)
func (_PaymentGateway *PaymentGatewaySession) UnlockPayment(txId string) (*types.Transaction, error) {
	return _PaymentGateway.Contract.UnlockPayment(&_PaymentGateway.TransactOpts, txId)
}

// UnlockPayment is a paid mutator transaction binding the contract method 0xe01d7646.
//
// Solidity: function unlockPayment(string txId) returns(bool)
func (_PaymentGateway *PaymentGatewayTransactorSession) UnlockPayment(txId string) (*types.Transaction, error) {
	return _PaymentGateway.Contract.UnlockPayment(&_PaymentGateway.TransactOpts, txId)
}
