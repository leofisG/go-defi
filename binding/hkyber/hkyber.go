// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hkyber

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

// HkyberABI is the input ABI used to generate the binding from.
const HkyberABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRate\",\"type\":\"uint256\"}],\"name\":\"swapEtherToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenQty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"srcToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"srcQty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minRate\",\"type\":\"uint256\"}],\"name\":\"swapTokenToToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Hkyber is an auto generated Go binding around an Ethereum contract.
type Hkyber struct {
	HkyberCaller     // Read-only binding to the contract
	HkyberTransactor // Write-only binding to the contract
	HkyberFilterer   // Log filterer for contract events
}

// HkyberCaller is an auto generated read-only Go binding around an Ethereum contract.
type HkyberCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HkyberTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HkyberTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HkyberFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HkyberFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HkyberSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HkyberSession struct {
	Contract     *Hkyber           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HkyberCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HkyberCallerSession struct {
	Contract *HkyberCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HkyberTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HkyberTransactorSession struct {
	Contract     *HkyberTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HkyberRaw is an auto generated low-level Go binding around an Ethereum contract.
type HkyberRaw struct {
	Contract *Hkyber // Generic contract binding to access the raw methods on
}

// HkyberCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HkyberCallerRaw struct {
	Contract *HkyberCaller // Generic read-only contract binding to access the raw methods on
}

// HkyberTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HkyberTransactorRaw struct {
	Contract *HkyberTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHkyber creates a new instance of Hkyber, bound to a specific deployed contract.
func NewHkyber(address common.Address, backend bind.ContractBackend) (*Hkyber, error) {
	contract, err := bindHkyber(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hkyber{HkyberCaller: HkyberCaller{contract: contract}, HkyberTransactor: HkyberTransactor{contract: contract}, HkyberFilterer: HkyberFilterer{contract: contract}}, nil
}

// NewHkyberCaller creates a new read-only instance of Hkyber, bound to a specific deployed contract.
func NewHkyberCaller(address common.Address, caller bind.ContractCaller) (*HkyberCaller, error) {
	contract, err := bindHkyber(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HkyberCaller{contract: contract}, nil
}

// NewHkyberTransactor creates a new write-only instance of Hkyber, bound to a specific deployed contract.
func NewHkyberTransactor(address common.Address, transactor bind.ContractTransactor) (*HkyberTransactor, error) {
	contract, err := bindHkyber(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HkyberTransactor{contract: contract}, nil
}

// NewHkyberFilterer creates a new log filterer instance of Hkyber, bound to a specific deployed contract.
func NewHkyberFilterer(address common.Address, filterer bind.ContractFilterer) (*HkyberFilterer, error) {
	contract, err := bindHkyber(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HkyberFilterer{contract: contract}, nil
}

// bindHkyber binds a generic wrapper to an already deployed contract.
func bindHkyber(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HkyberABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hkyber *HkyberRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hkyber.Contract.HkyberCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hkyber *HkyberRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hkyber.Contract.HkyberTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hkyber *HkyberRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hkyber.Contract.HkyberTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hkyber *HkyberCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hkyber.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hkyber *HkyberTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hkyber.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hkyber *HkyberTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hkyber.Contract.contract.Transact(opts, method, params...)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Hkyber *HkyberCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Hkyber.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Hkyber *HkyberSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Hkyber.Contract.Tokens(&_Hkyber.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Hkyber *HkyberCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Hkyber.Contract.Tokens(&_Hkyber.CallOpts, arg0)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x43ac1dc6.
//
// Solidity: function swapEtherToToken(uint256 value, address token, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberTransactor) SwapEtherToToken(opts *bind.TransactOpts, value *big.Int, token common.Address, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.contract.Transact(opts, "swapEtherToToken", value, token, minRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x43ac1dc6.
//
// Solidity: function swapEtherToToken(uint256 value, address token, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberSession) SwapEtherToToken(value *big.Int, token common.Address, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.Contract.SwapEtherToToken(&_Hkyber.TransactOpts, value, token, minRate)
}

// SwapEtherToToken is a paid mutator transaction binding the contract method 0x43ac1dc6.
//
// Solidity: function swapEtherToToken(uint256 value, address token, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberTransactorSession) SwapEtherToToken(value *big.Int, token common.Address, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.Contract.SwapEtherToToken(&_Hkyber.TransactOpts, value, token, minRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 tokenQty, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberTransactor) SwapTokenToEther(opts *bind.TransactOpts, token common.Address, tokenQty *big.Int, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.contract.Transact(opts, "swapTokenToEther", token, tokenQty, minRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 tokenQty, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberSession) SwapTokenToEther(token common.Address, tokenQty *big.Int, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.Contract.SwapTokenToEther(&_Hkyber.TransactOpts, token, tokenQty, minRate)
}

// SwapTokenToEther is a paid mutator transaction binding the contract method 0x3bba21dc.
//
// Solidity: function swapTokenToEther(address token, uint256 tokenQty, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberTransactorSession) SwapTokenToEther(token common.Address, tokenQty *big.Int, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.Contract.SwapTokenToEther(&_Hkyber.TransactOpts, token, tokenQty, minRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address srcToken, uint256 srcQty, address destToken, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberTransactor) SwapTokenToToken(opts *bind.TransactOpts, srcToken common.Address, srcQty *big.Int, destToken common.Address, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.contract.Transact(opts, "swapTokenToToken", srcToken, srcQty, destToken, minRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address srcToken, uint256 srcQty, address destToken, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberSession) SwapTokenToToken(srcToken common.Address, srcQty *big.Int, destToken common.Address, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.Contract.SwapTokenToToken(&_Hkyber.TransactOpts, srcToken, srcQty, destToken, minRate)
}

// SwapTokenToToken is a paid mutator transaction binding the contract method 0x7409e2eb.
//
// Solidity: function swapTokenToToken(address srcToken, uint256 srcQty, address destToken, uint256 minRate) payable returns(uint256 destAmount)
func (_Hkyber *HkyberTransactorSession) SwapTokenToToken(srcToken common.Address, srcQty *big.Int, destToken common.Address, minRate *big.Int) (*types.Transaction, error) {
	return _Hkyber.Contract.SwapTokenToToken(&_Hkyber.TransactOpts, srcToken, srcQty, destToken, minRate)
}
