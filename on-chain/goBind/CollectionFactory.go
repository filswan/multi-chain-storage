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

// CollectionFactoryMetaData contains all meta data concerning the CollectionFactory contract.
var CollectionFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"defaultCollection\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectionOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"}],\"name\":\"CreateCollection\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"}],\"name\":\"createCollection\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultCollectionAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCollections\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"contractURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintToNewCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"}]",
}

// CollectionFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CollectionFactoryMetaData.ABI instead.
var CollectionFactoryABI = CollectionFactoryMetaData.ABI

// CollectionFactory is an auto generated Go binding around an Ethereum contract.
type CollectionFactory struct {
	CollectionFactoryCaller     // Read-only binding to the contract
	CollectionFactoryTransactor // Write-only binding to the contract
	CollectionFactoryFilterer   // Log filterer for contract events
}

// CollectionFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CollectionFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CollectionFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CollectionFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CollectionFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CollectionFactorySession struct {
	Contract     *CollectionFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CollectionFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CollectionFactoryCallerSession struct {
	Contract *CollectionFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CollectionFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CollectionFactoryTransactorSession struct {
	Contract     *CollectionFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CollectionFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CollectionFactoryRaw struct {
	Contract *CollectionFactory // Generic contract binding to access the raw methods on
}

// CollectionFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CollectionFactoryCallerRaw struct {
	Contract *CollectionFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CollectionFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CollectionFactoryTransactorRaw struct {
	Contract *CollectionFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCollectionFactory creates a new instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactory(address common.Address, backend bind.ContractBackend) (*CollectionFactory, error) {
	contract, err := bindCollectionFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CollectionFactory{CollectionFactoryCaller: CollectionFactoryCaller{contract: contract}, CollectionFactoryTransactor: CollectionFactoryTransactor{contract: contract}, CollectionFactoryFilterer: CollectionFactoryFilterer{contract: contract}}, nil
}

// NewCollectionFactoryCaller creates a new read-only instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactoryCaller(address common.Address, caller bind.ContractCaller) (*CollectionFactoryCaller, error) {
	contract, err := bindCollectionFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryCaller{contract: contract}, nil
}

// NewCollectionFactoryTransactor creates a new write-only instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CollectionFactoryTransactor, error) {
	contract, err := bindCollectionFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryTransactor{contract: contract}, nil
}

// NewCollectionFactoryFilterer creates a new log filterer instance of CollectionFactory, bound to a specific deployed contract.
func NewCollectionFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CollectionFactoryFilterer, error) {
	contract, err := bindCollectionFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryFilterer{contract: contract}, nil
}

// bindCollectionFactory binds a generic wrapper to an already deployed contract.
func bindCollectionFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CollectionFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionFactory *CollectionFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionFactory.Contract.CollectionFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionFactory *CollectionFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CollectionFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionFactory *CollectionFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CollectionFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CollectionFactory *CollectionFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CollectionFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CollectionFactory *CollectionFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CollectionFactory *CollectionFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CollectionFactory.Contract.contract.Transact(opts, method, params...)
}

// DefaultCollectionAddress is a free data retrieval call binding the contract method 0x72ba30e0.
//
// Solidity: function defaultCollectionAddress() view returns(address)
func (_CollectionFactory *CollectionFactoryCaller) DefaultCollectionAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionFactory.contract.Call(opts, &out, "defaultCollectionAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DefaultCollectionAddress is a free data retrieval call binding the contract method 0x72ba30e0.
//
// Solidity: function defaultCollectionAddress() view returns(address)
func (_CollectionFactory *CollectionFactorySession) DefaultCollectionAddress() (common.Address, error) {
	return _CollectionFactory.Contract.DefaultCollectionAddress(&_CollectionFactory.CallOpts)
}

// DefaultCollectionAddress is a free data retrieval call binding the contract method 0x72ba30e0.
//
// Solidity: function defaultCollectionAddress() view returns(address)
func (_CollectionFactory *CollectionFactoryCallerSession) DefaultCollectionAddress() (common.Address, error) {
	return _CollectionFactory.Contract.DefaultCollectionAddress(&_CollectionFactory.CallOpts)
}

// GetCollections is a free data retrieval call binding the contract method 0x127f1498.
//
// Solidity: function getCollections(address user) view returns(address[])
func (_CollectionFactory *CollectionFactoryCaller) GetCollections(opts *bind.CallOpts, user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _CollectionFactory.contract.Call(opts, &out, "getCollections", user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCollections is a free data retrieval call binding the contract method 0x127f1498.
//
// Solidity: function getCollections(address user) view returns(address[])
func (_CollectionFactory *CollectionFactorySession) GetCollections(user common.Address) ([]common.Address, error) {
	return _CollectionFactory.Contract.GetCollections(&_CollectionFactory.CallOpts, user)
}

// GetCollections is a free data retrieval call binding the contract method 0x127f1498.
//
// Solidity: function getCollections(address user) view returns(address[])
func (_CollectionFactory *CollectionFactoryCallerSession) GetCollections(user common.Address) ([]common.Address, error) {
	return _CollectionFactory.Contract.GetCollections(&_CollectionFactory.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollectionFactory *CollectionFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CollectionFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollectionFactory *CollectionFactorySession) Owner() (common.Address, error) {
	return _CollectionFactory.Contract.Owner(&_CollectionFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CollectionFactory *CollectionFactoryCallerSession) Owner() (common.Address, error) {
	return _CollectionFactory.Contract.Owner(&_CollectionFactory.CallOpts)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x059dfe13.
//
// Solidity: function createCollection(string contractURI) returns(address)
func (_CollectionFactory *CollectionFactoryTransactor) CreateCollection(opts *bind.TransactOpts, contractURI string) (*types.Transaction, error) {
	return _CollectionFactory.contract.Transact(opts, "createCollection", contractURI)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x059dfe13.
//
// Solidity: function createCollection(string contractURI) returns(address)
func (_CollectionFactory *CollectionFactorySession) CreateCollection(contractURI string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CreateCollection(&_CollectionFactory.TransactOpts, contractURI)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x059dfe13.
//
// Solidity: function createCollection(string contractURI) returns(address)
func (_CollectionFactory *CollectionFactoryTransactorSession) CreateCollection(contractURI string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.CreateCollection(&_CollectionFactory.TransactOpts, contractURI)
}

// Mint is a paid mutator transaction binding the contract method 0xb85cbc79.
//
// Solidity: function mint(address collection, address recipient, uint256 amount, string uri) returns()
func (_CollectionFactory *CollectionFactoryTransactor) Mint(opts *bind.TransactOpts, collection common.Address, recipient common.Address, amount *big.Int, uri string) (*types.Transaction, error) {
	return _CollectionFactory.contract.Transact(opts, "mint", collection, recipient, amount, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xb85cbc79.
//
// Solidity: function mint(address collection, address recipient, uint256 amount, string uri) returns()
func (_CollectionFactory *CollectionFactorySession) Mint(collection common.Address, recipient common.Address, amount *big.Int, uri string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.Mint(&_CollectionFactory.TransactOpts, collection, recipient, amount, uri)
}

// Mint is a paid mutator transaction binding the contract method 0xb85cbc79.
//
// Solidity: function mint(address collection, address recipient, uint256 amount, string uri) returns()
func (_CollectionFactory *CollectionFactoryTransactorSession) Mint(collection common.Address, recipient common.Address, amount *big.Int, uri string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.Mint(&_CollectionFactory.TransactOpts, collection, recipient, amount, uri)
}

// MintToNewCollection is a paid mutator transaction binding the contract method 0x9c2fcd47.
//
// Solidity: function mintToNewCollection(string contractURI, address recipient, uint256 amount, string uri) returns()
func (_CollectionFactory *CollectionFactoryTransactor) MintToNewCollection(opts *bind.TransactOpts, contractURI string, recipient common.Address, amount *big.Int, uri string) (*types.Transaction, error) {
	return _CollectionFactory.contract.Transact(opts, "mintToNewCollection", contractURI, recipient, amount, uri)
}

// MintToNewCollection is a paid mutator transaction binding the contract method 0x9c2fcd47.
//
// Solidity: function mintToNewCollection(string contractURI, address recipient, uint256 amount, string uri) returns()
func (_CollectionFactory *CollectionFactorySession) MintToNewCollection(contractURI string, recipient common.Address, amount *big.Int, uri string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.MintToNewCollection(&_CollectionFactory.TransactOpts, contractURI, recipient, amount, uri)
}

// MintToNewCollection is a paid mutator transaction binding the contract method 0x9c2fcd47.
//
// Solidity: function mintToNewCollection(string contractURI, address recipient, uint256 amount, string uri) returns()
func (_CollectionFactory *CollectionFactoryTransactorSession) MintToNewCollection(contractURI string, recipient common.Address, amount *big.Int, uri string) (*types.Transaction, error) {
	return _CollectionFactory.Contract.MintToNewCollection(&_CollectionFactory.TransactOpts, contractURI, recipient, amount, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollectionFactory *CollectionFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CollectionFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollectionFactory *CollectionFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _CollectionFactory.Contract.RenounceOwnership(&_CollectionFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CollectionFactory *CollectionFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CollectionFactory.Contract.RenounceOwnership(&_CollectionFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollectionFactory *CollectionFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CollectionFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollectionFactory *CollectionFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollectionFactory.Contract.TransferOwnership(&_CollectionFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CollectionFactory *CollectionFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CollectionFactory.Contract.TransferOwnership(&_CollectionFactory.TransactOpts, newOwner)
}

// CollectionFactoryCreateCollectionIterator is returned from FilterCreateCollection and is used to iterate over the raw logs and unpacked data for CreateCollection events raised by the CollectionFactory contract.
type CollectionFactoryCreateCollectionIterator struct {
	Event *CollectionFactoryCreateCollection // Event containing the contract specifics and raw log

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
func (it *CollectionFactoryCreateCollectionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionFactoryCreateCollection)
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
		it.Event = new(CollectionFactoryCreateCollection)
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
func (it *CollectionFactoryCreateCollectionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionFactoryCreateCollectionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionFactoryCreateCollection represents a CreateCollection event raised by the CollectionFactory contract.
type CollectionFactoryCreateCollection struct {
	CollectionOwner   common.Address
	CollectionAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCreateCollection is a free log retrieval operation binding the contract event 0xb4c4734e81b1e2f6da6e94416410e635e762ea72591c3fe936f08dfdbe57479c.
//
// Solidity: event CreateCollection(address collectionOwner, address collectionAddress)
func (_CollectionFactory *CollectionFactoryFilterer) FilterCreateCollection(opts *bind.FilterOpts) (*CollectionFactoryCreateCollectionIterator, error) {

	logs, sub, err := _CollectionFactory.contract.FilterLogs(opts, "CreateCollection")
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryCreateCollectionIterator{contract: _CollectionFactory.contract, event: "CreateCollection", logs: logs, sub: sub}, nil
}

// WatchCreateCollection is a free log subscription operation binding the contract event 0xb4c4734e81b1e2f6da6e94416410e635e762ea72591c3fe936f08dfdbe57479c.
//
// Solidity: event CreateCollection(address collectionOwner, address collectionAddress)
func (_CollectionFactory *CollectionFactoryFilterer) WatchCreateCollection(opts *bind.WatchOpts, sink chan<- *CollectionFactoryCreateCollection) (event.Subscription, error) {

	logs, sub, err := _CollectionFactory.contract.WatchLogs(opts, "CreateCollection")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionFactoryCreateCollection)
				if err := _CollectionFactory.contract.UnpackLog(event, "CreateCollection", log); err != nil {
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

// ParseCreateCollection is a log parse operation binding the contract event 0xb4c4734e81b1e2f6da6e94416410e635e762ea72591c3fe936f08dfdbe57479c.
//
// Solidity: event CreateCollection(address collectionOwner, address collectionAddress)
func (_CollectionFactory *CollectionFactoryFilterer) ParseCreateCollection(log types.Log) (*CollectionFactoryCreateCollection, error) {
	event := new(CollectionFactoryCreateCollection)
	if err := _CollectionFactory.contract.UnpackLog(event, "CreateCollection", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CollectionFactory contract.
type CollectionFactoryOwnershipTransferredIterator struct {
	Event *CollectionFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CollectionFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionFactoryOwnershipTransferred)
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
		it.Event = new(CollectionFactoryOwnershipTransferred)
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
func (it *CollectionFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the CollectionFactory contract.
type CollectionFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollectionFactory *CollectionFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CollectionFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollectionFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryOwnershipTransferredIterator{contract: _CollectionFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CollectionFactory *CollectionFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CollectionFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CollectionFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionFactoryOwnershipTransferred)
				if err := _CollectionFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_CollectionFactory *CollectionFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*CollectionFactoryOwnershipTransferred, error) {
	event := new(CollectionFactoryOwnershipTransferred)
	if err := _CollectionFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionFactoryTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the CollectionFactory contract.
type CollectionFactoryTransferBatchIterator struct {
	Event *CollectionFactoryTransferBatch // Event containing the contract specifics and raw log

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
func (it *CollectionFactoryTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionFactoryTransferBatch)
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
		it.Event = new(CollectionFactoryTransferBatch)
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
func (it *CollectionFactoryTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionFactoryTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionFactoryTransferBatch represents a TransferBatch event raised by the CollectionFactory contract.
type CollectionFactoryTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_CollectionFactory *CollectionFactoryFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*CollectionFactoryTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollectionFactory.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryTransferBatchIterator{contract: _CollectionFactory.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_CollectionFactory *CollectionFactoryFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *CollectionFactoryTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollectionFactory.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionFactoryTransferBatch)
				if err := _CollectionFactory.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_CollectionFactory *CollectionFactoryFilterer) ParseTransferBatch(log types.Log) (*CollectionFactoryTransferBatch, error) {
	event := new(CollectionFactoryTransferBatch)
	if err := _CollectionFactory.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CollectionFactoryTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the CollectionFactory contract.
type CollectionFactoryTransferSingleIterator struct {
	Event *CollectionFactoryTransferSingle // Event containing the contract specifics and raw log

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
func (it *CollectionFactoryTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CollectionFactoryTransferSingle)
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
		it.Event = new(CollectionFactoryTransferSingle)
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
func (it *CollectionFactoryTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CollectionFactoryTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CollectionFactoryTransferSingle represents a TransferSingle event raised by the CollectionFactory contract.
type CollectionFactoryTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_CollectionFactory *CollectionFactoryFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*CollectionFactoryTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollectionFactory.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CollectionFactoryTransferSingleIterator{contract: _CollectionFactory.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_CollectionFactory *CollectionFactoryFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *CollectionFactoryTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CollectionFactory.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CollectionFactoryTransferSingle)
				if err := _CollectionFactory.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_CollectionFactory *CollectionFactoryFilterer) ParseTransferSingle(log types.Log) (*CollectionFactoryTransferSingle, error) {
	event := new(CollectionFactoryTransferSingle)
	if err := _CollectionFactory.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
