// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hcether

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

// HcetherABI is the input ABI used to generate the binding from.
const HcetherABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"CETHER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Hcether is an auto generated Go binding around an Ethereum contract.
type Hcether struct {
	HcetherCaller     // Read-only binding to the contract
	HcetherTransactor // Write-only binding to the contract
	HcetherFilterer   // Log filterer for contract events
}

// HcetherCaller is an auto generated read-only Go binding around an Ethereum contract.
type HcetherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HcetherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HcetherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HcetherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HcetherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HcetherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HcetherSession struct {
	Contract     *Hcether          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HcetherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HcetherCallerSession struct {
	Contract *HcetherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// HcetherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HcetherTransactorSession struct {
	Contract     *HcetherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// HcetherRaw is an auto generated low-level Go binding around an Ethereum contract.
type HcetherRaw struct {
	Contract *Hcether // Generic contract binding to access the raw methods on
}

// HcetherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HcetherCallerRaw struct {
	Contract *HcetherCaller // Generic read-only contract binding to access the raw methods on
}

// HcetherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HcetherTransactorRaw struct {
	Contract *HcetherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHcether creates a new instance of Hcether, bound to a specific deployed contract.
func NewHcether(address common.Address, backend bind.ContractBackend) (*Hcether, error) {
	contract, err := bindHcether(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hcether{HcetherCaller: HcetherCaller{contract: contract}, HcetherTransactor: HcetherTransactor{contract: contract}, HcetherFilterer: HcetherFilterer{contract: contract}}, nil
}

// NewHcetherCaller creates a new read-only instance of Hcether, bound to a specific deployed contract.
func NewHcetherCaller(address common.Address, caller bind.ContractCaller) (*HcetherCaller, error) {
	contract, err := bindHcether(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HcetherCaller{contract: contract}, nil
}

// NewHcetherTransactor creates a new write-only instance of Hcether, bound to a specific deployed contract.
func NewHcetherTransactor(address common.Address, transactor bind.ContractTransactor) (*HcetherTransactor, error) {
	contract, err := bindHcether(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HcetherTransactor{contract: contract}, nil
}

// NewHcetherFilterer creates a new log filterer instance of Hcether, bound to a specific deployed contract.
func NewHcetherFilterer(address common.Address, filterer bind.ContractFilterer) (*HcetherFilterer, error) {
	contract, err := bindHcether(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HcetherFilterer{contract: contract}, nil
}

// bindHcether binds a generic wrapper to an already deployed contract.
func bindHcether(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HcetherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hcether *HcetherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hcether.Contract.HcetherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hcether *HcetherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hcether.Contract.HcetherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hcether *HcetherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hcether.Contract.HcetherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hcether *HcetherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hcether.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hcether *HcetherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hcether.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hcether *HcetherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hcether.Contract.contract.Transact(opts, method, params...)
}

// CETHER is a free data retrieval call binding the contract method 0x092a5781.
//
// Solidity: function CETHER() view returns(address)
func (_Hcether *HcetherCaller) CETHER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hcether.contract.Call(opts, &out, "CETHER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CETHER is a free data retrieval call binding the contract method 0x092a5781.
//
// Solidity: function CETHER() view returns(address)
func (_Hcether *HcetherSession) CETHER() (common.Address, error) {
	return _Hcether.Contract.CETHER(&_Hcether.CallOpts)
}

// CETHER is a free data retrieval call binding the contract method 0x092a5781.
//
// Solidity: function CETHER() view returns(address)
func (_Hcether *HcetherCallerSession) CETHER() (common.Address, error) {
	return _Hcether.Contract.CETHER(&_Hcether.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) payable returns()
func (_Hcether *HcetherTransactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Hcether.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) payable returns()
func (_Hcether *HcetherSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _Hcether.Contract.Mint(&_Hcether.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) payable returns()
func (_Hcether *HcetherTransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _Hcether.Contract.Mint(&_Hcether.TransactOpts, value)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hcether *HcetherTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hcether.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hcether *HcetherSession) PostProcess() (*types.Transaction, error) {
	return _Hcether.Contract.PostProcess(&_Hcether.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hcether *HcetherTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Hcether.Contract.PostProcess(&_Hcether.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) payable returns()
func (_Hcether *HcetherTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Hcether.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) payable returns()
func (_Hcether *HcetherSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Hcether.Contract.Redeem(&_Hcether.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) payable returns()
func (_Hcether *HcetherTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Hcether.Contract.Redeem(&_Hcether.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) payable returns()
func (_Hcether *HcetherTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Hcether.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) payable returns()
func (_Hcether *HcetherSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Hcether.Contract.RedeemUnderlying(&_Hcether.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) payable returns()
func (_Hcether *HcetherTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Hcether.Contract.RedeemUnderlying(&_Hcether.TransactOpts, redeemAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xaa553b61.
//
// Solidity: function repayBorrowBehalf(uint256 amount, address borrower) payable returns()
func (_Hcether *HcetherTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, amount *big.Int, borrower common.Address) (*types.Transaction, error) {
	return _Hcether.contract.Transact(opts, "repayBorrowBehalf", amount, borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xaa553b61.
//
// Solidity: function repayBorrowBehalf(uint256 amount, address borrower) payable returns()
func (_Hcether *HcetherSession) RepayBorrowBehalf(amount *big.Int, borrower common.Address) (*types.Transaction, error) {
	return _Hcether.Contract.RepayBorrowBehalf(&_Hcether.TransactOpts, amount, borrower)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0xaa553b61.
//
// Solidity: function repayBorrowBehalf(uint256 amount, address borrower) payable returns()
func (_Hcether *HcetherTransactorSession) RepayBorrowBehalf(amount *big.Int, borrower common.Address) (*types.Transaction, error) {
	return _Hcether.Contract.RepayBorrowBehalf(&_Hcether.TransactOpts, amount, borrower)
}
