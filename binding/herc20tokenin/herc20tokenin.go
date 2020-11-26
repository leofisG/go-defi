// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package herc20tokenin

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

// Herc20tokeninABI is the input ABI used to generate the binding from.
const Herc20tokeninABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"inject\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Herc20tokenin is an auto generated Go binding around an Ethereum contract.
type Herc20tokenin struct {
	Herc20tokeninCaller     // Read-only binding to the contract
	Herc20tokeninTransactor // Write-only binding to the contract
	Herc20tokeninFilterer   // Log filterer for contract events
}

// Herc20tokeninCaller is an auto generated read-only Go binding around an Ethereum contract.
type Herc20tokeninCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Herc20tokeninTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Herc20tokeninTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Herc20tokeninFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Herc20tokeninFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Herc20tokeninSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Herc20tokeninSession struct {
	Contract     *Herc20tokenin    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Herc20tokeninCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Herc20tokeninCallerSession struct {
	Contract *Herc20tokeninCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// Herc20tokeninTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Herc20tokeninTransactorSession struct {
	Contract     *Herc20tokeninTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Herc20tokeninRaw is an auto generated low-level Go binding around an Ethereum contract.
type Herc20tokeninRaw struct {
	Contract *Herc20tokenin // Generic contract binding to access the raw methods on
}

// Herc20tokeninCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Herc20tokeninCallerRaw struct {
	Contract *Herc20tokeninCaller // Generic read-only contract binding to access the raw methods on
}

// Herc20tokeninTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Herc20tokeninTransactorRaw struct {
	Contract *Herc20tokeninTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHerc20tokenin creates a new instance of Herc20tokenin, bound to a specific deployed contract.
func NewHerc20tokenin(address common.Address, backend bind.ContractBackend) (*Herc20tokenin, error) {
	contract, err := bindHerc20tokenin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Herc20tokenin{Herc20tokeninCaller: Herc20tokeninCaller{contract: contract}, Herc20tokeninTransactor: Herc20tokeninTransactor{contract: contract}, Herc20tokeninFilterer: Herc20tokeninFilterer{contract: contract}}, nil
}

// NewHerc20tokeninCaller creates a new read-only instance of Herc20tokenin, bound to a specific deployed contract.
func NewHerc20tokeninCaller(address common.Address, caller bind.ContractCaller) (*Herc20tokeninCaller, error) {
	contract, err := bindHerc20tokenin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Herc20tokeninCaller{contract: contract}, nil
}

// NewHerc20tokeninTransactor creates a new write-only instance of Herc20tokenin, bound to a specific deployed contract.
func NewHerc20tokeninTransactor(address common.Address, transactor bind.ContractTransactor) (*Herc20tokeninTransactor, error) {
	contract, err := bindHerc20tokenin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Herc20tokeninTransactor{contract: contract}, nil
}

// NewHerc20tokeninFilterer creates a new log filterer instance of Herc20tokenin, bound to a specific deployed contract.
func NewHerc20tokeninFilterer(address common.Address, filterer bind.ContractFilterer) (*Herc20tokeninFilterer, error) {
	contract, err := bindHerc20tokenin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Herc20tokeninFilterer{contract: contract}, nil
}

// bindHerc20tokenin binds a generic wrapper to an already deployed contract.
func bindHerc20tokenin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Herc20tokeninABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Herc20tokenin *Herc20tokeninRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Herc20tokenin.Contract.Herc20tokeninCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Herc20tokenin *Herc20tokeninRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Herc20tokenin.Contract.Herc20tokeninTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Herc20tokenin *Herc20tokeninRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Herc20tokenin.Contract.Herc20tokeninTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Herc20tokenin *Herc20tokeninCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Herc20tokenin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Herc20tokenin *Herc20tokeninTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Herc20tokenin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Herc20tokenin *Herc20tokeninTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Herc20tokenin.Contract.contract.Transact(opts, method, params...)
}

// Inject is a paid mutator transaction binding the contract method 0xd0797f84.
//
// Solidity: function inject(address[] tokens, uint256[] amounts) payable returns()
func (_Herc20tokenin *Herc20tokeninTransactor) Inject(opts *bind.TransactOpts, tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Herc20tokenin.contract.Transact(opts, "inject", tokens, amounts)
}

// Inject is a paid mutator transaction binding the contract method 0xd0797f84.
//
// Solidity: function inject(address[] tokens, uint256[] amounts) payable returns()
func (_Herc20tokenin *Herc20tokeninSession) Inject(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Herc20tokenin.Contract.Inject(&_Herc20tokenin.TransactOpts, tokens, amounts)
}

// Inject is a paid mutator transaction binding the contract method 0xd0797f84.
//
// Solidity: function inject(address[] tokens, uint256[] amounts) payable returns()
func (_Herc20tokenin *Herc20tokeninTransactorSession) Inject(tokens []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Herc20tokenin.Contract.Inject(&_Herc20tokenin.TransactOpts, tokens, amounts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Herc20tokenin *Herc20tokeninTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Herc20tokenin.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Herc20tokenin *Herc20tokeninSession) PostProcess() (*types.Transaction, error) {
	return _Herc20tokenin.Contract.PostProcess(&_Herc20tokenin.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Herc20tokenin *Herc20tokeninTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Herc20tokenin.Contract.PostProcess(&_Herc20tokenin.TransactOpts)
}
