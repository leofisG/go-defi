// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package haave

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

// HaaveABI is the input ABI used to generate the binding from.
const HaaveABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_aToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Haave is an auto generated Go binding around an Ethereum contract.
type Haave struct {
	HaaveCaller     // Read-only binding to the contract
	HaaveTransactor // Write-only binding to the contract
	HaaveFilterer   // Log filterer for contract events
}

// HaaveCaller is an auto generated read-only Go binding around an Ethereum contract.
type HaaveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HaaveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HaaveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HaaveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HaaveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HaaveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HaaveSession struct {
	Contract     *Haave            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HaaveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HaaveCallerSession struct {
	Contract *HaaveCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HaaveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HaaveTransactorSession struct {
	Contract     *HaaveTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HaaveRaw is an auto generated low-level Go binding around an Ethereum contract.
type HaaveRaw struct {
	Contract *Haave // Generic contract binding to access the raw methods on
}

// HaaveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HaaveCallerRaw struct {
	Contract *HaaveCaller // Generic read-only contract binding to access the raw methods on
}

// HaaveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HaaveTransactorRaw struct {
	Contract *HaaveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHaave creates a new instance of Haave, bound to a specific deployed contract.
func NewHaave(address common.Address, backend bind.ContractBackend) (*Haave, error) {
	contract, err := bindHaave(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Haave{HaaveCaller: HaaveCaller{contract: contract}, HaaveTransactor: HaaveTransactor{contract: contract}, HaaveFilterer: HaaveFilterer{contract: contract}}, nil
}

// NewHaaveCaller creates a new read-only instance of Haave, bound to a specific deployed contract.
func NewHaaveCaller(address common.Address, caller bind.ContractCaller) (*HaaveCaller, error) {
	contract, err := bindHaave(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HaaveCaller{contract: contract}, nil
}

// NewHaaveTransactor creates a new write-only instance of Haave, bound to a specific deployed contract.
func NewHaaveTransactor(address common.Address, transactor bind.ContractTransactor) (*HaaveTransactor, error) {
	contract, err := bindHaave(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HaaveTransactor{contract: contract}, nil
}

// NewHaaveFilterer creates a new log filterer instance of Haave, bound to a specific deployed contract.
func NewHaaveFilterer(address common.Address, filterer bind.ContractFilterer) (*HaaveFilterer, error) {
	contract, err := bindHaave(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HaaveFilterer{contract: contract}, nil
}

// bindHaave binds a generic wrapper to an already deployed contract.
func bindHaave(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HaaveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Haave *HaaveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Haave.Contract.HaaveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Haave *HaaveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Haave.Contract.HaaveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Haave *HaaveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Haave.Contract.HaaveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Haave *HaaveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Haave.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Haave *HaaveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Haave.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Haave *HaaveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Haave.Contract.contract.Transact(opts, method, params...)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Haave *HaaveCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Haave.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Haave *HaaveSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Haave.Contract.Tokens(&_Haave.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_Haave *HaaveCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _Haave.Contract.Tokens(&_Haave.CallOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _reserve, uint256 _amount) payable returns()
func (_Haave *HaaveTransactor) Deposit(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Haave.contract.Transact(opts, "deposit", _reserve, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _reserve, uint256 _amount) payable returns()
func (_Haave *HaaveSession) Deposit(_reserve common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Haave.Contract.Deposit(&_Haave.TransactOpts, _reserve, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address _reserve, uint256 _amount) payable returns()
func (_Haave *HaaveTransactorSession) Deposit(_reserve common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Haave.Contract.Deposit(&_Haave.TransactOpts, _reserve, _amount)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) payable returns()
func (_Haave *HaaveTransactor) ExecuteOperation(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _Haave.contract.Transact(opts, "executeOperation", _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) payable returns()
func (_Haave *HaaveSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _Haave.Contract.ExecuteOperation(&_Haave.TransactOpts, _reserve, _amount, _fee, _params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0xee872558.
//
// Solidity: function executeOperation(address _reserve, uint256 _amount, uint256 _fee, bytes _params) payable returns()
func (_Haave *HaaveTransactorSession) ExecuteOperation(_reserve common.Address, _amount *big.Int, _fee *big.Int, _params []byte) (*types.Transaction, error) {
	return _Haave.Contract.ExecuteOperation(&_Haave.TransactOpts, _reserve, _amount, _fee, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address _reserve, uint256 _amount, bytes _params) payable returns()
func (_Haave *HaaveTransactor) FlashLoan(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Haave.contract.Transact(opts, "flashLoan", _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address _reserve, uint256 _amount, bytes _params) payable returns()
func (_Haave *HaaveSession) FlashLoan(_reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Haave.Contract.FlashLoan(&_Haave.TransactOpts, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xe0232b42.
//
// Solidity: function flashLoan(address _reserve, uint256 _amount, bytes _params) payable returns()
func (_Haave *HaaveTransactorSession) FlashLoan(_reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Haave.Contract.FlashLoan(&_Haave.TransactOpts, _reserve, _amount, _params)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address _aToken, uint256 _amount) payable returns()
func (_Haave *HaaveTransactor) Redeem(opts *bind.TransactOpts, _aToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Haave.contract.Transact(opts, "redeem", _aToken, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address _aToken, uint256 _amount) payable returns()
func (_Haave *HaaveSession) Redeem(_aToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Haave.Contract.Redeem(&_Haave.TransactOpts, _aToken, _amount)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address _aToken, uint256 _amount) payable returns()
func (_Haave *HaaveTransactorSession) Redeem(_aToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Haave.Contract.Redeem(&_Haave.TransactOpts, _aToken, _amount)
}
