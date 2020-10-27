// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// UniswapABI is the input ABI used to generate the binding from.
const UniswapABI = "[{\"name\":\"NewExchange\",\"inputs\":[{\"type\":\"address\",\"name\":\"token\",\"indexed\":true},{\"type\":\"address\",\"name\":\"exchange\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"initializeFactory\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"template\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35725},{\"name\":\"createExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":187911},{\"name\":\"getExchange\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"token\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":715},{\"name\":\"getToken\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"exchange\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":745},{\"name\":\"getTokenWithId\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"token_id\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":736},{\"name\":\"exchangeTemplate\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":633},{\"name\":\"tokenCount\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":663}]"

// Uniswap is an auto generated Go binding around an Ethereum contract.
type Uniswap struct {
	UniswapCaller     // Read-only binding to the contract
	UniswapTransactor // Write-only binding to the contract
	UniswapFilterer   // Log filterer for contract events
}

// UniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapSession struct {
	Contract     *Uniswap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapCallerSession struct {
	Contract *UniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapTransactorSession struct {
	Contract     *UniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapRaw struct {
	Contract *Uniswap // Generic contract binding to access the raw methods on
}

// UniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapCallerRaw struct {
	Contract *UniswapCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapTransactorRaw struct {
	Contract *UniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswap creates a new instance of Uniswap, bound to a specific deployed contract.
func NewUniswap(address common.Address, backend bind.ContractBackend) (*Uniswap, error) {
	contract, err := bindUniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswap{UniswapCaller: UniswapCaller{contract: contract}, UniswapTransactor: UniswapTransactor{contract: contract}, UniswapFilterer: UniswapFilterer{contract: contract}}, nil
}

// NewUniswapCaller creates a new read-only instance of Uniswap, bound to a specific deployed contract.
func NewUniswapCaller(address common.Address, caller bind.ContractCaller) (*UniswapCaller, error) {
	contract, err := bindUniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapCaller{contract: contract}, nil
}

// NewUniswapTransactor creates a new write-only instance of Uniswap, bound to a specific deployed contract.
func NewUniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapTransactor, error) {
	contract, err := bindUniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapTransactor{contract: contract}, nil
}

// NewUniswapFilterer creates a new log filterer instance of Uniswap, bound to a specific deployed contract.
func NewUniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapFilterer, error) {
	contract, err := bindUniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapFilterer{contract: contract}, nil
}

// bindUniswap binds a generic wrapper to an already deployed contract.
func bindUniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap *UniswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswap.Contract.UniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap *UniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.Contract.UniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap *UniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap.Contract.UniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap *UniswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap *UniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap *UniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap.Contract.contract.Transact(opts, method, params...)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() returns(address out)
func (_Uniswap *UniswapCaller) ExchangeTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "exchangeTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() returns(address out)
func (_Uniswap *UniswapSession) ExchangeTemplate() (common.Address, error) {
	return _Uniswap.Contract.ExchangeTemplate(&_Uniswap.CallOpts)
}

// ExchangeTemplate is a free data retrieval call binding the contract method 0x1c2bbd18.
//
// Solidity: function exchangeTemplate() returns(address out)
func (_Uniswap *UniswapCallerSession) ExchangeTemplate() (common.Address, error) {
	return _Uniswap.Contract.ExchangeTemplate(&_Uniswap.CallOpts)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) returns(address out)
func (_Uniswap *UniswapCaller) GetExchange(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getExchange", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) returns(address out)
func (_Uniswap *UniswapSession) GetExchange(token common.Address) (common.Address, error) {
	return _Uniswap.Contract.GetExchange(&_Uniswap.CallOpts, token)
}

// GetExchange is a free data retrieval call binding the contract method 0x06f2bf62.
//
// Solidity: function getExchange(address token) returns(address out)
func (_Uniswap *UniswapCallerSession) GetExchange(token common.Address) (common.Address, error) {
	return _Uniswap.Contract.GetExchange(&_Uniswap.CallOpts, token)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) returns(address out)
func (_Uniswap *UniswapCaller) GetToken(opts *bind.CallOpts, exchange common.Address) (common.Address, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getToken", exchange)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) returns(address out)
func (_Uniswap *UniswapSession) GetToken(exchange common.Address) (common.Address, error) {
	return _Uniswap.Contract.GetToken(&_Uniswap.CallOpts, exchange)
}

// GetToken is a free data retrieval call binding the contract method 0x59770438.
//
// Solidity: function getToken(address exchange) returns(address out)
func (_Uniswap *UniswapCallerSession) GetToken(exchange common.Address) (common.Address, error) {
	return _Uniswap.Contract.GetToken(&_Uniswap.CallOpts, exchange)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) returns(address out)
func (_Uniswap *UniswapCaller) GetTokenWithId(opts *bind.CallOpts, token_id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getTokenWithId", token_id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) returns(address out)
func (_Uniswap *UniswapSession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _Uniswap.Contract.GetTokenWithId(&_Uniswap.CallOpts, token_id)
}

// GetTokenWithId is a free data retrieval call binding the contract method 0xaa65a6c0.
//
// Solidity: function getTokenWithId(uint256 token_id) returns(address out)
func (_Uniswap *UniswapCallerSession) GetTokenWithId(token_id *big.Int) (common.Address, error) {
	return _Uniswap.Contract.GetTokenWithId(&_Uniswap.CallOpts, token_id)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() returns(uint256 out)
func (_Uniswap *UniswapCaller) TokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "tokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() returns(uint256 out)
func (_Uniswap *UniswapSession) TokenCount() (*big.Int, error) {
	return _Uniswap.Contract.TokenCount(&_Uniswap.CallOpts)
}

// TokenCount is a free data retrieval call binding the contract method 0x9f181b5e.
//
// Solidity: function tokenCount() returns(uint256 out)
func (_Uniswap *UniswapCallerSession) TokenCount() (*big.Int, error) {
	return _Uniswap.Contract.TokenCount(&_Uniswap.CallOpts)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Uniswap *UniswapTransactor) CreateExchange(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "createExchange", token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Uniswap *UniswapSession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.CreateExchange(&_Uniswap.TransactOpts, token)
}

// CreateExchange is a paid mutator transaction binding the contract method 0x1648f38e.
//
// Solidity: function createExchange(address token) returns(address out)
func (_Uniswap *UniswapTransactorSession) CreateExchange(token common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.CreateExchange(&_Uniswap.TransactOpts, token)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Uniswap *UniswapTransactor) InitializeFactory(opts *bind.TransactOpts, template common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "initializeFactory", template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Uniswap *UniswapSession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.InitializeFactory(&_Uniswap.TransactOpts, template)
}

// InitializeFactory is a paid mutator transaction binding the contract method 0x538a3f0e.
//
// Solidity: function initializeFactory(address template) returns()
func (_Uniswap *UniswapTransactorSession) InitializeFactory(template common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.InitializeFactory(&_Uniswap.TransactOpts, template)
}

// UniswapNewExchangeIterator is returned from FilterNewExchange and is used to iterate over the raw logs and unpacked data for NewExchange events raised by the Uniswap contract.
type UniswapNewExchangeIterator struct {
	Event *UniswapNewExchange // Event containing the contract specifics and raw log

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
func (it *UniswapNewExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapNewExchange)
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
		it.Event = new(UniswapNewExchange)
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
func (it *UniswapNewExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapNewExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapNewExchange represents a NewExchange event raised by the Uniswap contract.
type UniswapNewExchange struct {
	Token    common.Address
	Exchange common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewExchange is a free log retrieval operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Uniswap *UniswapFilterer) FilterNewExchange(opts *bind.FilterOpts, token []common.Address, exchange []common.Address) (*UniswapNewExchangeIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return &UniswapNewExchangeIterator{contract: _Uniswap.contract, event: "NewExchange", logs: logs, sub: sub}, nil
}

// WatchNewExchange is a free log subscription operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Uniswap *UniswapFilterer) WatchNewExchange(opts *bind.WatchOpts, sink chan<- *UniswapNewExchange, token []common.Address, exchange []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var exchangeRule []interface{}
	for _, exchangeItem := range exchange {
		exchangeRule = append(exchangeRule, exchangeItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "NewExchange", tokenRule, exchangeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapNewExchange)
				if err := _Uniswap.contract.UnpackLog(event, "NewExchange", log); err != nil {
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

// ParseNewExchange is a log parse operation binding the contract event 0x9d42cb017eb05bd8944ab536a8b35bc68085931dd5f4356489801453923953f9.
//
// Solidity: event NewExchange(address indexed token, address indexed exchange)
func (_Uniswap *UniswapFilterer) ParseNewExchange(log types.Log) (*UniswapNewExchange, error) {
	event := new(UniswapNewExchange)
	if err := _Uniswap.contract.UnpackLog(event, "NewExchange", log); err != nil {
		return nil, err
	}
	return event, nil
}
