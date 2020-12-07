// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hyearn

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

// HyearnABI is the input ABI used to generate the binding from.
const HyearnABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Hyearn is an auto generated Go binding around an Ethereum contract.
type Hyearn struct {
	HyearnCaller     // Read-only binding to the contract
	HyearnTransactor // Write-only binding to the contract
	HyearnFilterer   // Log filterer for contract events
}

// HyearnCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyearnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyearnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyearnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyearnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyearnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyearnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyearnSession struct {
	Contract     *Hyearn           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyearnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyearnCallerSession struct {
	Contract *HyearnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HyearnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyearnTransactorSession struct {
	Contract     *HyearnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyearnRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyearnRaw struct {
	Contract *Hyearn // Generic contract binding to access the raw methods on
}

// HyearnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyearnCallerRaw struct {
	Contract *HyearnCaller // Generic read-only contract binding to access the raw methods on
}

// HyearnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyearnTransactorRaw struct {
	Contract *HyearnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyearn creates a new instance of Hyearn, bound to a specific deployed contract.
func NewHyearn(address common.Address, backend bind.ContractBackend) (*Hyearn, error) {
	contract, err := bindHyearn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hyearn{HyearnCaller: HyearnCaller{contract: contract}, HyearnTransactor: HyearnTransactor{contract: contract}, HyearnFilterer: HyearnFilterer{contract: contract}}, nil
}

// NewHyearnCaller creates a new read-only instance of Hyearn, bound to a specific deployed contract.
func NewHyearnCaller(address common.Address, caller bind.ContractCaller) (*HyearnCaller, error) {
	contract, err := bindHyearn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyearnCaller{contract: contract}, nil
}

// NewHyearnTransactor creates a new write-only instance of Hyearn, bound to a specific deployed contract.
func NewHyearnTransactor(address common.Address, transactor bind.ContractTransactor) (*HyearnTransactor, error) {
	contract, err := bindHyearn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyearnTransactor{contract: contract}, nil
}

// NewHyearnFilterer creates a new log filterer instance of Hyearn, bound to a specific deployed contract.
func NewHyearnFilterer(address common.Address, filterer bind.ContractFilterer) (*HyearnFilterer, error) {
	contract, err := bindHyearn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyearnFilterer{contract: contract}, nil
}

// bindHyearn binds a generic wrapper to an already deployed contract.
func bindHyearn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HyearnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyearn *HyearnRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyearn.Contract.HyearnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyearn *HyearnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyearn.Contract.HyearnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyearn *HyearnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyearn.Contract.HyearnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyearn *HyearnCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyearn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyearn *HyearnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyearn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyearn *HyearnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyearn.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address vault, uint256 _amount) payable returns()
func (_Hyearn *HyearnTransactor) Deposit(opts *bind.TransactOpts, vault common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hyearn.contract.Transact(opts, "deposit", vault, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address vault, uint256 _amount) payable returns()
func (_Hyearn *HyearnSession) Deposit(vault common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hyearn.Contract.Deposit(&_Hyearn.TransactOpts, vault, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address vault, uint256 _amount) payable returns()
func (_Hyearn *HyearnTransactorSession) Deposit(vault common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Hyearn.Contract.Deposit(&_Hyearn.TransactOpts, vault, _amount)
}

// DepositETH is a paid mutator transaction binding the contract method 0x56150edf.
//
// Solidity: function depositETH(uint256 value, address vault) payable returns()
func (_Hyearn *HyearnTransactor) DepositETH(opts *bind.TransactOpts, value *big.Int, vault common.Address) (*types.Transaction, error) {
	return _Hyearn.contract.Transact(opts, "depositETH", value, vault)
}

// DepositETH is a paid mutator transaction binding the contract method 0x56150edf.
//
// Solidity: function depositETH(uint256 value, address vault) payable returns()
func (_Hyearn *HyearnSession) DepositETH(value *big.Int, vault common.Address) (*types.Transaction, error) {
	return _Hyearn.Contract.DepositETH(&_Hyearn.TransactOpts, value, vault)
}

// DepositETH is a paid mutator transaction binding the contract method 0x56150edf.
//
// Solidity: function depositETH(uint256 value, address vault) payable returns()
func (_Hyearn *HyearnTransactorSession) DepositETH(value *big.Int, vault common.Address) (*types.Transaction, error) {
	return _Hyearn.Contract.DepositETH(&_Hyearn.TransactOpts, value, vault)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hyearn *HyearnTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyearn.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hyearn *HyearnSession) PostProcess() (*types.Transaction, error) {
	return _Hyearn.Contract.PostProcess(&_Hyearn.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hyearn *HyearnTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Hyearn.Contract.PostProcess(&_Hyearn.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address vault, uint256 _shares) payable returns()
func (_Hyearn *HyearnTransactor) Withdraw(opts *bind.TransactOpts, vault common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Hyearn.contract.Transact(opts, "withdraw", vault, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address vault, uint256 _shares) payable returns()
func (_Hyearn *HyearnSession) Withdraw(vault common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Hyearn.Contract.Withdraw(&_Hyearn.TransactOpts, vault, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address vault, uint256 _shares) payable returns()
func (_Hyearn *HyearnTransactorSession) Withdraw(vault common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Hyearn.Contract.Withdraw(&_Hyearn.TransactOpts, vault, _shares)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address vault, uint256 _shares) payable returns()
func (_Hyearn *HyearnTransactor) WithdrawETH(opts *bind.TransactOpts, vault common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Hyearn.contract.Transact(opts, "withdrawETH", vault, _shares)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address vault, uint256 _shares) payable returns()
func (_Hyearn *HyearnSession) WithdrawETH(vault common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Hyearn.Contract.WithdrawETH(&_Hyearn.TransactOpts, vault, _shares)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address vault, uint256 _shares) payable returns()
func (_Hyearn *HyearnTransactorSession) WithdrawETH(vault common.Address, _shares *big.Int) (*types.Transaction, error) {
	return _Hyearn.Contract.WithdrawETH(&_Hyearn.TransactOpts, vault, _shares)
}
