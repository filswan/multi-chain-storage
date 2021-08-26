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

// StateSenderABI is the input ABI used to generate the binding from.
const StateSenderABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"StateSynced\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"syncState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StateSender is an auto generated Go binding around an Ethereum contract.
type StateSender struct {
	StateSenderCaller     // Read-only binding to the contract
	StateSenderTransactor // Write-only binding to the contract
	StateSenderFilterer   // Log filterer for contract events
}

// StateSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateSenderSession struct {
	Contract     *StateSender      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateSenderCallerSession struct {
	Contract *StateSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StateSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateSenderTransactorSession struct {
	Contract     *StateSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StateSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateSenderRaw struct {
	Contract *StateSender // Generic contract binding to access the raw methods on
}

// StateSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateSenderCallerRaw struct {
	Contract *StateSenderCaller // Generic read-only contract binding to access the raw methods on
}

// StateSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateSenderTransactorRaw struct {
	Contract *StateSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateSender creates a new instance of StateSender, bound to a specific deployed contract.
func NewStateSender(address common.Address, backend bind.ContractBackend) (*StateSender, error) {
	contract, err := bindStateSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateSender{StateSenderCaller: StateSenderCaller{contract: contract}, StateSenderTransactor: StateSenderTransactor{contract: contract}, StateSenderFilterer: StateSenderFilterer{contract: contract}}, nil
}

// NewStateSenderCaller creates a new read-only instance of StateSender, bound to a specific deployed contract.
func NewStateSenderCaller(address common.Address, caller bind.ContractCaller) (*StateSenderCaller, error) {
	contract, err := bindStateSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateSenderCaller{contract: contract}, nil
}

// NewStateSenderTransactor creates a new write-only instance of StateSender, bound to a specific deployed contract.
func NewStateSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*StateSenderTransactor, error) {
	contract, err := bindStateSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateSenderTransactor{contract: contract}, nil
}

// NewStateSenderFilterer creates a new log filterer instance of StateSender, bound to a specific deployed contract.
func NewStateSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*StateSenderFilterer, error) {
	contract, err := bindStateSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateSenderFilterer{contract: contract}, nil
}

// bindStateSender binds a generic wrapper to an already deployed contract.
func bindStateSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateSender *StateSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateSender.Contract.StateSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateSender *StateSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateSender.Contract.StateSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateSender *StateSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateSender.Contract.StateSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateSender *StateSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateSender *StateSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateSender *StateSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateSender.Contract.contract.Transact(opts, method, params...)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_StateSender *StateSenderTransactor) SyncState(opts *bind.TransactOpts, receiver common.Address, data []byte) (*types.Transaction, error) {
	return _StateSender.contract.Transact(opts, "syncState", receiver, data)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_StateSender *StateSenderSession) SyncState(receiver common.Address, data []byte) (*types.Transaction, error) {
	return _StateSender.Contract.SyncState(&_StateSender.TransactOpts, receiver, data)
}

// SyncState is a paid mutator transaction binding the contract method 0x16f19831.
//
// Solidity: function syncState(address receiver, bytes data) returns()
func (_StateSender *StateSenderTransactorSession) SyncState(receiver common.Address, data []byte) (*types.Transaction, error) {
	return _StateSender.Contract.SyncState(&_StateSender.TransactOpts, receiver, data)
}

// StateSenderStateSyncedIterator is returned from FilterStateSynced and is used to iterate over the raw logs and unpacked data for StateSynced events raised by the StateSender contract.
type StateSenderStateSyncedIterator struct {
	Event *StateSenderStateSynced // Event containing the contract specifics and raw log

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
func (it *StateSenderStateSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StateSenderStateSynced)
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
		it.Event = new(StateSenderStateSynced)
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
func (it *StateSenderStateSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StateSenderStateSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StateSenderStateSynced represents a StateSynced event raised by the StateSender contract.
type StateSenderStateSynced struct {
	Id              *big.Int
	ContractAddress common.Address
	Data            []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStateSynced is a free log retrieval operation binding the contract event 0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392.
//
// Solidity: event StateSynced(uint256 indexed id, address indexed contractAddress, bytes data)
func (_StateSender *StateSenderFilterer) FilterStateSynced(opts *bind.FilterOpts, id []*big.Int, contractAddress []common.Address) (*StateSenderStateSyncedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _StateSender.contract.FilterLogs(opts, "StateSynced", idRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &StateSenderStateSyncedIterator{contract: _StateSender.contract, event: "StateSynced", logs: logs, sub: sub}, nil
}

// WatchStateSynced is a free log subscription operation binding the contract event 0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392.
//
// Solidity: event StateSynced(uint256 indexed id, address indexed contractAddress, bytes data)
func (_StateSender *StateSenderFilterer) WatchStateSynced(opts *bind.WatchOpts, sink chan<- *StateSenderStateSynced, id []*big.Int, contractAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _StateSender.contract.WatchLogs(opts, "StateSynced", idRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StateSenderStateSynced)
				if err := _StateSender.contract.UnpackLog(event, "StateSynced", log); err != nil {
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

// ParseStateSynced is a log parse operation binding the contract event 0x103fed9db65eac19c4d870f49ab7520fe03b99f1838e5996caf47e9e43308392.
//
// Solidity: event StateSynced(uint256 indexed id, address indexed contractAddress, bytes data)
func (_StateSender *StateSenderFilterer) ParseStateSynced(log types.Log) (*StateSenderStateSynced, error) {
	event := new(StateSenderStateSynced)
	if err := _StateSender.contract.UnpackLog(event, "StateSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
