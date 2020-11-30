// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hctoken

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

// HctokenABI is the input ABI used to generate the binding from.
const HctokenABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Hctoken is an auto generated Go binding around an Ethereum contract.
type Hctoken struct {
	HctokenCaller     // Read-only binding to the contract
	HctokenTransactor // Write-only binding to the contract
	HctokenFilterer   // Log filterer for contract events
}

// HctokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type HctokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HctokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HctokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HctokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HctokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HctokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HctokenSession struct {
	Contract     *Hctoken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HctokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HctokenCallerSession struct {
	Contract *HctokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HctokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HctokenTransactorSession struct {
	Contract     *HctokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HctokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type HctokenRaw struct {
	Contract *Hctoken // Generic contract binding to access the raw methods on
}

// HctokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HctokenCallerRaw struct {
	Contract *HctokenCaller // Generic read-only contract binding to access the raw methods on
}

// HctokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HctokenTransactorRaw struct {
	Contract *HctokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHctoken creates a new instance of Hctoken, bound to a specific deployed contract.
func NewHctoken(address common.Address, backend bind.ContractBackend) (*Hctoken, error) {
	contract, err := bindHctoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hctoken{HctokenCaller: HctokenCaller{contract: contract}, HctokenTransactor: HctokenTransactor{contract: contract}, HctokenFilterer: HctokenFilterer{contract: contract}}, nil
}

// NewHctokenCaller creates a new read-only instance of Hctoken, bound to a specific deployed contract.
func NewHctokenCaller(address common.Address, caller bind.ContractCaller) (*HctokenCaller, error) {
	contract, err := bindHctoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HctokenCaller{contract: contract}, nil
}

// NewHctokenTransactor creates a new write-only instance of Hctoken, bound to a specific deployed contract.
func NewHctokenTransactor(address common.Address, transactor bind.ContractTransactor) (*HctokenTransactor, error) {
	contract, err := bindHctoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HctokenTransactor{contract: contract}, nil
}

// NewHctokenFilterer creates a new log filterer instance of Hctoken, bound to a specific deployed contract.
func NewHctokenFilterer(address common.Address, filterer bind.ContractFilterer) (*HctokenFilterer, error) {
	contract, err := bindHctoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HctokenFilterer{contract: contract}, nil
}

// bindHctoken binds a generic wrapper to an already deployed contract.
func bindHctoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HctokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hctoken *HctokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hctoken.Contract.HctokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hctoken *HctokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hctoken.Contract.HctokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hctoken *HctokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hctoken.Contract.HctokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hctoken *HctokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hctoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hctoken *HctokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hctoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hctoken *HctokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hctoken.Contract.contract.Transact(opts, method, params...)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 mintAmount) payable returns()
func (_Hctoken *HctokenTransactor) Mint(opts *bind.TransactOpts, cToken common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.contract.Transact(opts, "mint", cToken, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 mintAmount) payable returns()
func (_Hctoken *HctokenSession) Mint(cToken common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.Mint(&_Hctoken.TransactOpts, cToken, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address cToken, uint256 mintAmount) payable returns()
func (_Hctoken *HctokenTransactorSession) Mint(cToken common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.Mint(&_Hctoken.TransactOpts, cToken, mintAmount)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hctoken *HctokenTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hctoken.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hctoken *HctokenSession) PostProcess() (*types.Transaction, error) {
	return _Hctoken.Contract.PostProcess(&_Hctoken.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hctoken *HctokenTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Hctoken.Contract.PostProcess(&_Hctoken.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address cToken, uint256 redeemTokens) payable returns()
func (_Hctoken *HctokenTransactor) Redeem(opts *bind.TransactOpts, cToken common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Hctoken.contract.Transact(opts, "redeem", cToken, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address cToken, uint256 redeemTokens) payable returns()
func (_Hctoken *HctokenSession) Redeem(cToken common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.Redeem(&_Hctoken.TransactOpts, cToken, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address cToken, uint256 redeemTokens) payable returns()
func (_Hctoken *HctokenTransactorSession) Redeem(cToken common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.Redeem(&_Hctoken.TransactOpts, cToken, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address cToken, uint256 redeemAmount) payable returns()
func (_Hctoken *HctokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, cToken common.Address, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.contract.Transact(opts, "redeemUnderlying", cToken, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address cToken, uint256 redeemAmount) payable returns()
func (_Hctoken *HctokenSession) RedeemUnderlying(cToken common.Address, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.RedeemUnderlying(&_Hctoken.TransactOpts, cToken, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x96294178.
//
// Solidity: function redeemUnderlying(address cToken, uint256 redeemAmount) payable returns()
func (_Hctoken *HctokenTransactorSession) RedeemUnderlying(cToken common.Address, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.RedeemUnderlying(&_Hctoken.TransactOpts, cToken, redeemAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x59086a5e.
//
// Solidity: function repayBorrowBehalf(address cToken, address borrower, uint256 repayAmount) payable returns()
func (_Hctoken *HctokenTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.contract.Transact(opts, "repayBorrowBehalf", cToken, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x59086a5e.
//
// Solidity: function repayBorrowBehalf(address cToken, address borrower, uint256 repayAmount) payable returns()
func (_Hctoken *HctokenSession) RepayBorrowBehalf(cToken common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.RepayBorrowBehalf(&_Hctoken.TransactOpts, cToken, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x59086a5e.
//
// Solidity: function repayBorrowBehalf(address cToken, address borrower, uint256 repayAmount) payable returns()
func (_Hctoken *HctokenTransactorSession) RepayBorrowBehalf(cToken common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Hctoken.Contract.RepayBorrowBehalf(&_Hctoken.TransactOpts, cToken, borrower, repayAmount)
}
