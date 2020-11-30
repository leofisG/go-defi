// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package furucombo

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

// FurucomboABI is the input ABI used to generate the binding from.
const FurucomboABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"batchExec\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tos\",\"type\":\"address[]\"},{\"internalType\":\"bytes[]\",\"name\":\"datas\",\"type\":\"bytes[]\"}],\"name\":\"execs\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Furucombo is an auto generated Go binding around an Ethereum contract.
type Furucombo struct {
	FurucomboCaller     // Read-only binding to the contract
	FurucomboTransactor // Write-only binding to the contract
	FurucomboFilterer   // Log filterer for contract events
}

// FurucomboCaller is an auto generated read-only Go binding around an Ethereum contract.
type FurucomboCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FurucomboTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FurucomboTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FurucomboFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FurucomboFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FurucomboSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FurucomboSession struct {
	Contract     *Furucombo        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FurucomboCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FurucomboCallerSession struct {
	Contract *FurucomboCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FurucomboTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FurucomboTransactorSession struct {
	Contract     *FurucomboTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FurucomboRaw is an auto generated low-level Go binding around an Ethereum contract.
type FurucomboRaw struct {
	Contract *Furucombo // Generic contract binding to access the raw methods on
}

// FurucomboCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FurucomboCallerRaw struct {
	Contract *FurucomboCaller // Generic read-only contract binding to access the raw methods on
}

// FurucomboTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FurucomboTransactorRaw struct {
	Contract *FurucomboTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFurucombo creates a new instance of Furucombo, bound to a specific deployed contract.
func NewFurucombo(address common.Address, backend bind.ContractBackend) (*Furucombo, error) {
	contract, err := bindFurucombo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Furucombo{FurucomboCaller: FurucomboCaller{contract: contract}, FurucomboTransactor: FurucomboTransactor{contract: contract}, FurucomboFilterer: FurucomboFilterer{contract: contract}}, nil
}

// NewFurucomboCaller creates a new read-only instance of Furucombo, bound to a specific deployed contract.
func NewFurucomboCaller(address common.Address, caller bind.ContractCaller) (*FurucomboCaller, error) {
	contract, err := bindFurucombo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FurucomboCaller{contract: contract}, nil
}

// NewFurucomboTransactor creates a new write-only instance of Furucombo, bound to a specific deployed contract.
func NewFurucomboTransactor(address common.Address, transactor bind.ContractTransactor) (*FurucomboTransactor, error) {
	contract, err := bindFurucombo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FurucomboTransactor{contract: contract}, nil
}

// NewFurucomboFilterer creates a new log filterer instance of Furucombo, bound to a specific deployed contract.
func NewFurucomboFilterer(address common.Address, filterer bind.ContractFilterer) (*FurucomboFilterer, error) {
	contract, err := bindFurucombo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FurucomboFilterer{contract: contract}, nil
}

// bindFurucombo binds a generic wrapper to an already deployed contract.
func bindFurucombo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FurucomboABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Furucombo *FurucomboRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Furucombo.Contract.FurucomboCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Furucombo *FurucomboRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Furucombo.Contract.FurucomboTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Furucombo *FurucomboRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Furucombo.Contract.FurucomboTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Furucombo *FurucomboCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Furucombo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Furucombo *FurucomboTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Furucombo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Furucombo *FurucomboTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Furucombo.Contract.contract.Transact(opts, method, params...)
}

// BatchExec is a paid mutator transaction binding the contract method 0x74a28f79.
//
// Solidity: function batchExec(address[] tos, bytes[] datas) payable returns()
func (_Furucombo *FurucomboTransactor) BatchExec(opts *bind.TransactOpts, tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Furucombo.contract.Transact(opts, "batchExec", tos, datas)
}

// BatchExec is a paid mutator transaction binding the contract method 0x74a28f79.
//
// Solidity: function batchExec(address[] tos, bytes[] datas) payable returns()
func (_Furucombo *FurucomboSession) BatchExec(tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Furucombo.Contract.BatchExec(&_Furucombo.TransactOpts, tos, datas)
}

// BatchExec is a paid mutator transaction binding the contract method 0x74a28f79.
//
// Solidity: function batchExec(address[] tos, bytes[] datas) payable returns()
func (_Furucombo *FurucomboTransactorSession) BatchExec(tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Furucombo.Contract.BatchExec(&_Furucombo.TransactOpts, tos, datas)
}

// Execs is a paid mutator transaction binding the contract method 0x94da7864.
//
// Solidity: function execs(address[] tos, bytes[] datas) payable returns()
func (_Furucombo *FurucomboTransactor) Execs(opts *bind.TransactOpts, tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Furucombo.contract.Transact(opts, "execs", tos, datas)
}

// Execs is a paid mutator transaction binding the contract method 0x94da7864.
//
// Solidity: function execs(address[] tos, bytes[] datas) payable returns()
func (_Furucombo *FurucomboSession) Execs(tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Furucombo.Contract.Execs(&_Furucombo.TransactOpts, tos, datas)
}

// Execs is a paid mutator transaction binding the contract method 0x94da7864.
//
// Solidity: function execs(address[] tos, bytes[] datas) payable returns()
func (_Furucombo *FurucomboTransactorSession) Execs(tos []common.Address, datas [][]byte) (*types.Transaction, error) {
	return _Furucombo.Contract.Execs(&_Furucombo.TransactOpts, tos, datas)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Furucombo *FurucomboTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Furucombo.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Furucombo *FurucomboSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Furucombo.Contract.Fallback(&_Furucombo.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Furucombo *FurucomboTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Furucombo.Contract.Fallback(&_Furucombo.TransactOpts, calldata)
}
