// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hbalancer_exchange

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

// IExchangeProxySwap is an auto generated low-level Go binding around an user-defined struct.
type IExchangeProxySwap struct {
	Pool              common.Address
	TokenIn           common.Address
	TokenOut          common.Address
	SwapAmount        *big.Int
	LimitReturnAmount *big.Int
	MaxPrice          *big.Int
}

// HbalancerExchangeABI is the input ABI used to generate the binding from.
const HbalancerExchangeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXCHANGE_PROXY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"}],\"name\":\"batchSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeProxy.Swap[]\",\"name\":\"swaps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"}],\"name\":\"batchSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeProxy.Swap[][]\",\"name\":\"swapSequences\",\"type\":\"tuple[][]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"}],\"name\":\"multihopBatchSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"swapAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIExchangeProxy.Swap[][]\",\"name\":\"swapSequences\",\"type\":\"tuple[][]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"}],\"name\":\"multihopBatchSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minTotalAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"smartSwapExactIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nPools\",\"type\":\"uint256\"}],\"name\":\"smartSwapExactOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmountIn\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// HbalancerExchange is an auto generated Go binding around an Ethereum contract.
type HbalancerExchange struct {
	HbalancerExchangeCaller     // Read-only binding to the contract
	HbalancerExchangeTransactor // Write-only binding to the contract
	HbalancerExchangeFilterer   // Log filterer for contract events
}

// HbalancerExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type HbalancerExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HbalancerExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HbalancerExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HbalancerExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HbalancerExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HbalancerExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HbalancerExchangeSession struct {
	Contract     *HbalancerExchange // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HbalancerExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HbalancerExchangeCallerSession struct {
	Contract *HbalancerExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// HbalancerExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HbalancerExchangeTransactorSession struct {
	Contract     *HbalancerExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// HbalancerExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type HbalancerExchangeRaw struct {
	Contract *HbalancerExchange // Generic contract binding to access the raw methods on
}

// HbalancerExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HbalancerExchangeCallerRaw struct {
	Contract *HbalancerExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// HbalancerExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HbalancerExchangeTransactorRaw struct {
	Contract *HbalancerExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHbalancerExchange creates a new instance of HbalancerExchange, bound to a specific deployed contract.
func NewHbalancerExchange(address common.Address, backend bind.ContractBackend) (*HbalancerExchange, error) {
	contract, err := bindHbalancerExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HbalancerExchange{HbalancerExchangeCaller: HbalancerExchangeCaller{contract: contract}, HbalancerExchangeTransactor: HbalancerExchangeTransactor{contract: contract}, HbalancerExchangeFilterer: HbalancerExchangeFilterer{contract: contract}}, nil
}

// NewHbalancerExchangeCaller creates a new read-only instance of HbalancerExchange, bound to a specific deployed contract.
func NewHbalancerExchangeCaller(address common.Address, caller bind.ContractCaller) (*HbalancerExchangeCaller, error) {
	contract, err := bindHbalancerExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HbalancerExchangeCaller{contract: contract}, nil
}

// NewHbalancerExchangeTransactor creates a new write-only instance of HbalancerExchange, bound to a specific deployed contract.
func NewHbalancerExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*HbalancerExchangeTransactor, error) {
	contract, err := bindHbalancerExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HbalancerExchangeTransactor{contract: contract}, nil
}

// NewHbalancerExchangeFilterer creates a new log filterer instance of HbalancerExchange, bound to a specific deployed contract.
func NewHbalancerExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*HbalancerExchangeFilterer, error) {
	contract, err := bindHbalancerExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HbalancerExchangeFilterer{contract: contract}, nil
}

// bindHbalancerExchange binds a generic wrapper to an already deployed contract.
func bindHbalancerExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HbalancerExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HbalancerExchange *HbalancerExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HbalancerExchange.Contract.HbalancerExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HbalancerExchange *HbalancerExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.HbalancerExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HbalancerExchange *HbalancerExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.HbalancerExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HbalancerExchange *HbalancerExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HbalancerExchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HbalancerExchange *HbalancerExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HbalancerExchange *HbalancerExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.contract.Transact(opts, method, params...)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_HbalancerExchange *HbalancerExchangeCaller) ETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HbalancerExchange.contract.Call(opts, &out, "ETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_HbalancerExchange *HbalancerExchangeSession) ETHADDRESS() (common.Address, error) {
	return _HbalancerExchange.Contract.ETHADDRESS(&_HbalancerExchange.CallOpts)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_HbalancerExchange *HbalancerExchangeCallerSession) ETHADDRESS() (common.Address, error) {
	return _HbalancerExchange.Contract.ETHADDRESS(&_HbalancerExchange.CallOpts)
}

// EXCHANGEPROXY is a free data retrieval call binding the contract method 0x6fca4f8f.
//
// Solidity: function EXCHANGE_PROXY() view returns(address)
func (_HbalancerExchange *HbalancerExchangeCaller) EXCHANGEPROXY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HbalancerExchange.contract.Call(opts, &out, "EXCHANGE_PROXY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXCHANGEPROXY is a free data retrieval call binding the contract method 0x6fca4f8f.
//
// Solidity: function EXCHANGE_PROXY() view returns(address)
func (_HbalancerExchange *HbalancerExchangeSession) EXCHANGEPROXY() (common.Address, error) {
	return _HbalancerExchange.Contract.EXCHANGEPROXY(&_HbalancerExchange.CallOpts)
}

// EXCHANGEPROXY is a free data retrieval call binding the contract method 0x6fca4f8f.
//
// Solidity: function EXCHANGE_PROXY() view returns(address)
func (_HbalancerExchange *HbalancerExchangeCallerSession) EXCHANGEPROXY() (common.Address, error) {
	return _HbalancerExchange.Contract.EXCHANGEPROXY(&_HbalancerExchange.CallOpts)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeTransactor) BatchSwapExactIn(opts *bind.TransactOpts, swaps []IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "batchSwapExactIn", swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeSession) BatchSwapExactIn(swaps []IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.BatchSwapExactIn(&_HbalancerExchange.TransactOpts, swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactIn is a paid mutator transaction binding the contract method 0x8743ad58.
//
// Solidity: function batchSwapExactIn((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeTransactorSession) BatchSwapExactIn(swaps []IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.BatchSwapExactIn(&_HbalancerExchange.TransactOpts, swaps, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeTransactor) BatchSwapExactOut(opts *bind.TransactOpts, swaps []IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "batchSwapExactOut", swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeSession) BatchSwapExactOut(swaps []IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.BatchSwapExactOut(&_HbalancerExchange.TransactOpts, swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// BatchSwapExactOut is a paid mutator transaction binding the contract method 0x2db58134.
//
// Solidity: function batchSwapExactOut((address,address,address,uint256,uint256,uint256)[] swaps, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeTransactorSession) BatchSwapExactOut(swaps []IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.BatchSwapExactOut(&_HbalancerExchange.TransactOpts, swaps, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeTransactor) MultihopBatchSwapExactIn(opts *bind.TransactOpts, swapSequences [][]IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "multihopBatchSwapExactIn", swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeSession) MultihopBatchSwapExactIn(swapSequences [][]IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.MultihopBatchSwapExactIn(&_HbalancerExchange.TransactOpts, swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactIn is a paid mutator transaction binding the contract method 0xe2b39746.
//
// Solidity: function multihopBatchSwapExactIn((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeTransactorSession) MultihopBatchSwapExactIn(swapSequences [][]IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.MultihopBatchSwapExactIn(&_HbalancerExchange.TransactOpts, swapSequences, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeTransactor) MultihopBatchSwapExactOut(opts *bind.TransactOpts, swapSequences [][]IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "multihopBatchSwapExactOut", swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeSession) MultihopBatchSwapExactOut(swapSequences [][]IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.MultihopBatchSwapExactOut(&_HbalancerExchange.TransactOpts, swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// MultihopBatchSwapExactOut is a paid mutator transaction binding the contract method 0x86b2ecc4.
//
// Solidity: function multihopBatchSwapExactOut((address,address,address,uint256,uint256,uint256)[][] swapSequences, address tokenIn, address tokenOut, uint256 maxTotalAmountIn) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeTransactorSession) MultihopBatchSwapExactOut(swapSequences [][]IExchangeProxySwap, tokenIn common.Address, tokenOut common.Address, maxTotalAmountIn *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.MultihopBatchSwapExactOut(&_HbalancerExchange.TransactOpts, swapSequences, tokenIn, tokenOut, maxTotalAmountIn)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_HbalancerExchange *HbalancerExchangeTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_HbalancerExchange *HbalancerExchangeSession) PostProcess() (*types.Transaction, error) {
	return _HbalancerExchange.Contract.PostProcess(&_HbalancerExchange.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_HbalancerExchange *HbalancerExchangeTransactorSession) PostProcess() (*types.Transaction, error) {
	return _HbalancerExchange.Contract.PostProcess(&_HbalancerExchange.TransactOpts)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeTransactor) SmartSwapExactIn(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "smartSwapExactIn", tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeSession) SmartSwapExactIn(tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.SmartSwapExactIn(&_HbalancerExchange.TransactOpts, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactIn is a paid mutator transaction binding the contract method 0x21b0eb85.
//
// Solidity: function smartSwapExactIn(address tokenIn, address tokenOut, uint256 totalAmountIn, uint256 minTotalAmountOut, uint256 nPools) payable returns(uint256 totalAmountOut)
func (_HbalancerExchange *HbalancerExchangeTransactorSession) SmartSwapExactIn(tokenIn common.Address, tokenOut common.Address, totalAmountIn *big.Int, minTotalAmountOut *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.SmartSwapExactIn(&_HbalancerExchange.TransactOpts, tokenIn, tokenOut, totalAmountIn, minTotalAmountOut, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeTransactor) SmartSwapExactOut(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.contract.Transact(opts, "smartSwapExactOut", tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeSession) SmartSwapExactOut(tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.SmartSwapExactOut(&_HbalancerExchange.TransactOpts, tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}

// SmartSwapExactOut is a paid mutator transaction binding the contract method 0xb40f39ee.
//
// Solidity: function smartSwapExactOut(address tokenIn, address tokenOut, uint256 totalAmountOut, uint256 maxTotalAmountIn, uint256 nPools) payable returns(uint256 totalAmountIn)
func (_HbalancerExchange *HbalancerExchangeTransactorSession) SmartSwapExactOut(tokenIn common.Address, tokenOut common.Address, totalAmountOut *big.Int, maxTotalAmountIn *big.Int, nPools *big.Int) (*types.Transaction, error) {
	return _HbalancerExchange.Contract.SmartSwapExactOut(&_HbalancerExchange.TransactOpts, tokenIn, tokenOut, totalAmountOut, maxTotalAmountIn, nPools)
}
