// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yweth

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

// YwethABI is the input ABI used to generate the binding from.
const YwethABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"earn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPricePerFullShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"harvest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"max\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"min\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_controller\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_min\",\"type\":\"uint256\"}],\"name\":\"setMin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAllETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shares\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Yweth is an auto generated Go binding around an Ethereum contract.
type Yweth struct {
	YwethCaller     // Read-only binding to the contract
	YwethTransactor // Write-only binding to the contract
	YwethFilterer   // Log filterer for contract events
}

// YwethCaller is an auto generated read-only Go binding around an Ethereum contract.
type YwethCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YwethTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YwethTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YwethFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YwethFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YwethSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YwethSession struct {
	Contract     *Yweth            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YwethCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YwethCallerSession struct {
	Contract *YwethCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// YwethTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YwethTransactorSession struct {
	Contract     *YwethTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YwethRaw is an auto generated low-level Go binding around an Ethereum contract.
type YwethRaw struct {
	Contract *Yweth // Generic contract binding to access the raw methods on
}

// YwethCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YwethCallerRaw struct {
	Contract *YwethCaller // Generic read-only contract binding to access the raw methods on
}

// YwethTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YwethTransactorRaw struct {
	Contract *YwethTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYweth creates a new instance of Yweth, bound to a specific deployed contract.
func NewYweth(address common.Address, backend bind.ContractBackend) (*Yweth, error) {
	contract, err := bindYweth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yweth{YwethCaller: YwethCaller{contract: contract}, YwethTransactor: YwethTransactor{contract: contract}, YwethFilterer: YwethFilterer{contract: contract}}, nil
}

// NewYwethCaller creates a new read-only instance of Yweth, bound to a specific deployed contract.
func NewYwethCaller(address common.Address, caller bind.ContractCaller) (*YwethCaller, error) {
	contract, err := bindYweth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YwethCaller{contract: contract}, nil
}

// NewYwethTransactor creates a new write-only instance of Yweth, bound to a specific deployed contract.
func NewYwethTransactor(address common.Address, transactor bind.ContractTransactor) (*YwethTransactor, error) {
	contract, err := bindYweth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YwethTransactor{contract: contract}, nil
}

// NewYwethFilterer creates a new log filterer instance of Yweth, bound to a specific deployed contract.
func NewYwethFilterer(address common.Address, filterer bind.ContractFilterer) (*YwethFilterer, error) {
	contract, err := bindYweth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YwethFilterer{contract: contract}, nil
}

// bindYweth binds a generic wrapper to an already deployed contract.
func bindYweth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YwethABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yweth *YwethRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yweth.Contract.YwethCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yweth *YwethRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.Contract.YwethTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yweth *YwethRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yweth.Contract.YwethTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yweth *YwethCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yweth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yweth *YwethTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yweth *YwethTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yweth.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yweth *YwethCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yweth *YwethSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Yweth.Contract.Allowance(&_Yweth.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Yweth *YwethCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Yweth.Contract.Allowance(&_Yweth.CallOpts, owner, spender)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Yweth *YwethCaller) Available(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "available")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Yweth *YwethSession) Available() (*big.Int, error) {
	return _Yweth.Contract.Available(&_Yweth.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x48a0d754.
//
// Solidity: function available() view returns(uint256)
func (_Yweth *YwethCallerSession) Available() (*big.Int, error) {
	return _Yweth.Contract.Available(&_Yweth.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Yweth *YwethCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Yweth *YwethSession) Balance() (*big.Int, error) {
	return _Yweth.Contract.Balance(&_Yweth.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Yweth *YwethCallerSession) Balance() (*big.Int, error) {
	return _Yweth.Contract.Balance(&_Yweth.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Yweth *YwethCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Yweth *YwethSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Yweth.Contract.BalanceOf(&_Yweth.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Yweth *YwethCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Yweth.Contract.BalanceOf(&_Yweth.CallOpts, account)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Yweth *YwethCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Yweth *YwethSession) Controller() (common.Address, error) {
	return _Yweth.Contract.Controller(&_Yweth.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Yweth *YwethCallerSession) Controller() (common.Address, error) {
	return _Yweth.Contract.Controller(&_Yweth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yweth *YwethCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yweth *YwethSession) Decimals() (uint8, error) {
	return _Yweth.Contract.Decimals(&_Yweth.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Yweth *YwethCallerSession) Decimals() (uint8, error) {
	return _Yweth.Contract.Decimals(&_Yweth.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Yweth *YwethCaller) GetPricePerFullShare(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "getPricePerFullShare")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Yweth *YwethSession) GetPricePerFullShare() (*big.Int, error) {
	return _Yweth.Contract.GetPricePerFullShare(&_Yweth.CallOpts)
}

// GetPricePerFullShare is a free data retrieval call binding the contract method 0x77c7b8fc.
//
// Solidity: function getPricePerFullShare() view returns(uint256)
func (_Yweth *YwethCallerSession) GetPricePerFullShare() (*big.Int, error) {
	return _Yweth.Contract.GetPricePerFullShare(&_Yweth.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yweth *YwethCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yweth *YwethSession) Governance() (common.Address, error) {
	return _Yweth.Contract.Governance(&_Yweth.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yweth *YwethCallerSession) Governance() (common.Address, error) {
	return _Yweth.Contract.Governance(&_Yweth.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_Yweth *YwethCaller) Max(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "max")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_Yweth *YwethSession) Max() (*big.Int, error) {
	return _Yweth.Contract.Max(&_Yweth.CallOpts)
}

// Max is a free data retrieval call binding the contract method 0x6ac5db19.
//
// Solidity: function max() view returns(uint256)
func (_Yweth *YwethCallerSession) Max() (*big.Int, error) {
	return _Yweth.Contract.Max(&_Yweth.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_Yweth *YwethCaller) Min(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "min")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_Yweth *YwethSession) Min() (*big.Int, error) {
	return _Yweth.Contract.Min(&_Yweth.CallOpts)
}

// Min is a free data retrieval call binding the contract method 0xf8897945.
//
// Solidity: function min() view returns(uint256)
func (_Yweth *YwethCallerSession) Min() (*big.Int, error) {
	return _Yweth.Contract.Min(&_Yweth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yweth *YwethCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yweth *YwethSession) Name() (string, error) {
	return _Yweth.Contract.Name(&_Yweth.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Yweth *YwethCallerSession) Name() (string, error) {
	return _Yweth.Contract.Name(&_Yweth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yweth *YwethCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yweth *YwethSession) Symbol() (string, error) {
	return _Yweth.Contract.Symbol(&_Yweth.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Yweth *YwethCallerSession) Symbol() (string, error) {
	return _Yweth.Contract.Symbol(&_Yweth.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yweth *YwethCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yweth *YwethSession) Token() (common.Address, error) {
	return _Yweth.Contract.Token(&_Yweth.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Yweth *YwethCallerSession) Token() (common.Address, error) {
	return _Yweth.Contract.Token(&_Yweth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yweth *YwethCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yweth.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yweth *YwethSession) TotalSupply() (*big.Int, error) {
	return _Yweth.Contract.TotalSupply(&_Yweth.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Yweth *YwethCallerSession) TotalSupply() (*big.Int, error) {
	return _Yweth.Contract.TotalSupply(&_Yweth.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yweth *YwethTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yweth *YwethSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Approve(&_Yweth.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Yweth *YwethTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Approve(&_Yweth.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Yweth *YwethTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Yweth *YwethSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.DecreaseAllowance(&_Yweth.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Yweth *YwethTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.DecreaseAllowance(&_Yweth.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Yweth *YwethTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Yweth *YwethSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Deposit(&_Yweth.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Yweth *YwethTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Deposit(&_Yweth.TransactOpts, _amount)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Yweth *YwethTransactor) DepositAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "depositAll")
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Yweth *YwethSession) DepositAll() (*types.Transaction, error) {
	return _Yweth.Contract.DepositAll(&_Yweth.TransactOpts)
}

// DepositAll is a paid mutator transaction binding the contract method 0xde5f6268.
//
// Solidity: function depositAll() returns()
func (_Yweth *YwethTransactorSession) DepositAll() (*types.Transaction, error) {
	return _Yweth.Contract.DepositAll(&_Yweth.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Yweth *YwethTransactor) DepositETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "depositETH")
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Yweth *YwethSession) DepositETH() (*types.Transaction, error) {
	return _Yweth.Contract.DepositETH(&_Yweth.TransactOpts)
}

// DepositETH is a paid mutator transaction binding the contract method 0xf6326fb3.
//
// Solidity: function depositETH() payable returns()
func (_Yweth *YwethTransactorSession) DepositETH() (*types.Transaction, error) {
	return _Yweth.Contract.DepositETH(&_Yweth.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Yweth *YwethTransactor) Earn(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "earn")
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Yweth *YwethSession) Earn() (*types.Transaction, error) {
	return _Yweth.Contract.Earn(&_Yweth.TransactOpts)
}

// Earn is a paid mutator transaction binding the contract method 0xd389800f.
//
// Solidity: function earn() returns()
func (_Yweth *YwethTransactorSession) Earn() (*types.Transaction, error) {
	return _Yweth.Contract.Earn(&_Yweth.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_Yweth *YwethTransactor) Harvest(opts *bind.TransactOpts, reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "harvest", reserve, amount)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_Yweth *YwethSession) Harvest(reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Harvest(&_Yweth.TransactOpts, reserve, amount)
}

// Harvest is a paid mutator transaction binding the contract method 0x018ee9b7.
//
// Solidity: function harvest(address reserve, uint256 amount) returns()
func (_Yweth *YwethTransactorSession) Harvest(reserve common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Harvest(&_Yweth.TransactOpts, reserve, amount)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Yweth *YwethTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Yweth *YwethSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.IncreaseAllowance(&_Yweth.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Yweth *YwethTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.IncreaseAllowance(&_Yweth.TransactOpts, spender, addedValue)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Yweth *YwethTransactor) SetController(opts *bind.TransactOpts, _controller common.Address) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "setController", _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Yweth *YwethSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Yweth.Contract.SetController(&_Yweth.TransactOpts, _controller)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address _controller) returns()
func (_Yweth *YwethTransactorSession) SetController(_controller common.Address) (*types.Transaction, error) {
	return _Yweth.Contract.SetController(&_Yweth.TransactOpts, _controller)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Yweth *YwethTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Yweth *YwethSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Yweth.Contract.SetGovernance(&_Yweth.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Yweth *YwethTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Yweth.Contract.SetGovernance(&_Yweth.TransactOpts, _governance)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_Yweth *YwethTransactor) SetMin(opts *bind.TransactOpts, _min *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "setMin", _min)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_Yweth *YwethSession) SetMin(_min *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.SetMin(&_Yweth.TransactOpts, _min)
}

// SetMin is a paid mutator transaction binding the contract method 0x45dc3dd8.
//
// Solidity: function setMin(uint256 _min) returns()
func (_Yweth *YwethTransactorSession) SetMin(_min *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.SetMin(&_Yweth.TransactOpts, _min)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Yweth *YwethTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Yweth *YwethSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Transfer(&_Yweth.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Yweth *YwethTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Transfer(&_Yweth.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Yweth *YwethTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Yweth *YwethSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.TransferFrom(&_Yweth.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Yweth *YwethTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.TransferFrom(&_Yweth.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Yweth *YwethTransactor) Withdraw(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "withdraw", _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Yweth *YwethSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Withdraw(&_Yweth.TransactOpts, _shares)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _shares) returns()
func (_Yweth *YwethTransactorSession) Withdraw(_shares *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.Withdraw(&_Yweth.TransactOpts, _shares)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Yweth *YwethTransactor) WithdrawAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "withdrawAll")
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Yweth *YwethSession) WithdrawAll() (*types.Transaction, error) {
	return _Yweth.Contract.WithdrawAll(&_Yweth.TransactOpts)
}

// WithdrawAll is a paid mutator transaction binding the contract method 0x853828b6.
//
// Solidity: function withdrawAll() returns()
func (_Yweth *YwethTransactorSession) WithdrawAll() (*types.Transaction, error) {
	return _Yweth.Contract.WithdrawAll(&_Yweth.TransactOpts)
}

// WithdrawAllETH is a paid mutator transaction binding the contract method 0x90386bbf.
//
// Solidity: function withdrawAllETH() returns()
func (_Yweth *YwethTransactor) WithdrawAllETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "withdrawAllETH")
}

// WithdrawAllETH is a paid mutator transaction binding the contract method 0x90386bbf.
//
// Solidity: function withdrawAllETH() returns()
func (_Yweth *YwethSession) WithdrawAllETH() (*types.Transaction, error) {
	return _Yweth.Contract.WithdrawAllETH(&_Yweth.TransactOpts)
}

// WithdrawAllETH is a paid mutator transaction binding the contract method 0x90386bbf.
//
// Solidity: function withdrawAllETH() returns()
func (_Yweth *YwethTransactorSession) WithdrawAllETH() (*types.Transaction, error) {
	return _Yweth.Contract.WithdrawAllETH(&_Yweth.TransactOpts)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 _shares) returns()
func (_Yweth *YwethTransactor) WithdrawETH(opts *bind.TransactOpts, _shares *big.Int) (*types.Transaction, error) {
	return _Yweth.contract.Transact(opts, "withdrawETH", _shares)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 _shares) returns()
func (_Yweth *YwethSession) WithdrawETH(_shares *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.WithdrawETH(&_Yweth.TransactOpts, _shares)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xf14210a6.
//
// Solidity: function withdrawETH(uint256 _shares) returns()
func (_Yweth *YwethTransactorSession) WithdrawETH(_shares *big.Int) (*types.Transaction, error) {
	return _Yweth.Contract.WithdrawETH(&_Yweth.TransactOpts, _shares)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Yweth *YwethTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Yweth.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Yweth *YwethSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Yweth.Contract.Fallback(&_Yweth.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Yweth *YwethTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Yweth.Contract.Fallback(&_Yweth.TransactOpts, calldata)
}

// YwethApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Yweth contract.
type YwethApprovalIterator struct {
	Event *YwethApproval // Event containing the contract specifics and raw log

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
func (it *YwethApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YwethApproval)
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
		it.Event = new(YwethApproval)
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
func (it *YwethApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YwethApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YwethApproval represents a Approval event raised by the Yweth contract.
type YwethApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yweth *YwethFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YwethApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yweth.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YwethApprovalIterator{contract: _Yweth.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yweth *YwethFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YwethApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Yweth.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YwethApproval)
				if err := _Yweth.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Yweth *YwethFilterer) ParseApproval(log types.Log) (*YwethApproval, error) {
	event := new(YwethApproval)
	if err := _Yweth.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// YwethTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Yweth contract.
type YwethTransferIterator struct {
	Event *YwethTransfer // Event containing the contract specifics and raw log

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
func (it *YwethTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YwethTransfer)
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
		it.Event = new(YwethTransfer)
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
func (it *YwethTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YwethTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YwethTransfer represents a Transfer event raised by the Yweth contract.
type YwethTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Yweth *YwethFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YwethTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yweth.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YwethTransferIterator{contract: _Yweth.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Yweth *YwethFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YwethTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Yweth.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YwethTransfer)
				if err := _Yweth.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Yweth *YwethFilterer) ParseTransfer(log types.Log) (*YwethTransfer, error) {
	event := new(YwethTransfer)
	if err := _Yweth.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
