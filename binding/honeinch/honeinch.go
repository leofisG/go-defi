// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package honeinch

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

// HoneinchABI is the input ABI used to generate the binding from.
const HoneinchABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ETH_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ONEINCH_PROXY\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"REFERRER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_SPENDER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"guaranteedAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"callAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"callDataConcat\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"starts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"gasLimitsAndValues\",\"type\":\"uint256[]\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"returnAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Honeinch is an auto generated Go binding around an Ethereum contract.
type Honeinch struct {
	HoneinchCaller     // Read-only binding to the contract
	HoneinchTransactor // Write-only binding to the contract
	HoneinchFilterer   // Log filterer for contract events
}

// HoneinchCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoneinchCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneinchTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoneinchTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneinchFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoneinchFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoneinchSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoneinchSession struct {
	Contract     *Honeinch         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoneinchCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoneinchCallerSession struct {
	Contract *HoneinchCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HoneinchTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoneinchTransactorSession struct {
	Contract     *HoneinchTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HoneinchRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoneinchRaw struct {
	Contract *Honeinch // Generic contract binding to access the raw methods on
}

// HoneinchCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoneinchCallerRaw struct {
	Contract *HoneinchCaller // Generic read-only contract binding to access the raw methods on
}

// HoneinchTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoneinchTransactorRaw struct {
	Contract *HoneinchTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoneinch creates a new instance of Honeinch, bound to a specific deployed contract.
func NewHoneinch(address common.Address, backend bind.ContractBackend) (*Honeinch, error) {
	contract, err := bindHoneinch(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Honeinch{HoneinchCaller: HoneinchCaller{contract: contract}, HoneinchTransactor: HoneinchTransactor{contract: contract}, HoneinchFilterer: HoneinchFilterer{contract: contract}}, nil
}

// NewHoneinchCaller creates a new read-only instance of Honeinch, bound to a specific deployed contract.
func NewHoneinchCaller(address common.Address, caller bind.ContractCaller) (*HoneinchCaller, error) {
	contract, err := bindHoneinch(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoneinchCaller{contract: contract}, nil
}

// NewHoneinchTransactor creates a new write-only instance of Honeinch, bound to a specific deployed contract.
func NewHoneinchTransactor(address common.Address, transactor bind.ContractTransactor) (*HoneinchTransactor, error) {
	contract, err := bindHoneinch(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoneinchTransactor{contract: contract}, nil
}

// NewHoneinchFilterer creates a new log filterer instance of Honeinch, bound to a specific deployed contract.
func NewHoneinchFilterer(address common.Address, filterer bind.ContractFilterer) (*HoneinchFilterer, error) {
	contract, err := bindHoneinch(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoneinchFilterer{contract: contract}, nil
}

// bindHoneinch binds a generic wrapper to an already deployed contract.
func bindHoneinch(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoneinchABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Honeinch *HoneinchRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Honeinch.Contract.HoneinchCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Honeinch *HoneinchRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honeinch.Contract.HoneinchTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Honeinch *HoneinchRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Honeinch.Contract.HoneinchTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Honeinch *HoneinchCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Honeinch.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Honeinch *HoneinchTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honeinch.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Honeinch *HoneinchTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Honeinch.Contract.contract.Transact(opts, method, params...)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_Honeinch *HoneinchCaller) ETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Honeinch.contract.Call(opts, &out, "ETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_Honeinch *HoneinchSession) ETHADDRESS() (common.Address, error) {
	return _Honeinch.Contract.ETHADDRESS(&_Honeinch.CallOpts)
}

// ETHADDRESS is a free data retrieval call binding the contract method 0xa734f06e.
//
// Solidity: function ETH_ADDRESS() view returns(address)
func (_Honeinch *HoneinchCallerSession) ETHADDRESS() (common.Address, error) {
	return _Honeinch.Contract.ETHADDRESS(&_Honeinch.CallOpts)
}

// ONEINCHPROXY is a free data retrieval call binding the contract method 0x4add32fb.
//
// Solidity: function ONEINCH_PROXY() view returns(address)
func (_Honeinch *HoneinchCaller) ONEINCHPROXY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Honeinch.contract.Call(opts, &out, "ONEINCH_PROXY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ONEINCHPROXY is a free data retrieval call binding the contract method 0x4add32fb.
//
// Solidity: function ONEINCH_PROXY() view returns(address)
func (_Honeinch *HoneinchSession) ONEINCHPROXY() (common.Address, error) {
	return _Honeinch.Contract.ONEINCHPROXY(&_Honeinch.CallOpts)
}

// ONEINCHPROXY is a free data retrieval call binding the contract method 0x4add32fb.
//
// Solidity: function ONEINCH_PROXY() view returns(address)
func (_Honeinch *HoneinchCallerSession) ONEINCHPROXY() (common.Address, error) {
	return _Honeinch.Contract.ONEINCHPROXY(&_Honeinch.CallOpts)
}

// REFERRER is a free data retrieval call binding the contract method 0xc0560374.
//
// Solidity: function REFERRER() view returns(address)
func (_Honeinch *HoneinchCaller) REFERRER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Honeinch.contract.Call(opts, &out, "REFERRER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REFERRER is a free data retrieval call binding the contract method 0xc0560374.
//
// Solidity: function REFERRER() view returns(address)
func (_Honeinch *HoneinchSession) REFERRER() (common.Address, error) {
	return _Honeinch.Contract.REFERRER(&_Honeinch.CallOpts)
}

// REFERRER is a free data retrieval call binding the contract method 0xc0560374.
//
// Solidity: function REFERRER() view returns(address)
func (_Honeinch *HoneinchCallerSession) REFERRER() (common.Address, error) {
	return _Honeinch.Contract.REFERRER(&_Honeinch.CallOpts)
}

// TOKENSPENDER is a free data retrieval call binding the contract method 0x1e485f11.
//
// Solidity: function TOKEN_SPENDER() view returns(address)
func (_Honeinch *HoneinchCaller) TOKENSPENDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Honeinch.contract.Call(opts, &out, "TOKEN_SPENDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENSPENDER is a free data retrieval call binding the contract method 0x1e485f11.
//
// Solidity: function TOKEN_SPENDER() view returns(address)
func (_Honeinch *HoneinchSession) TOKENSPENDER() (common.Address, error) {
	return _Honeinch.Contract.TOKENSPENDER(&_Honeinch.CallOpts)
}

// TOKENSPENDER is a free data retrieval call binding the contract method 0x1e485f11.
//
// Solidity: function TOKEN_SPENDER() view returns(address)
func (_Honeinch *HoneinchCallerSession) TOKENSPENDER() (common.Address, error) {
	return _Honeinch.Contract.TOKENSPENDER(&_Honeinch.CallOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Honeinch *HoneinchTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Honeinch.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Honeinch *HoneinchSession) PostProcess() (*types.Transaction, error) {
	return _Honeinch.Contract.PostProcess(&_Honeinch.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Honeinch *HoneinchTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Honeinch.Contract.PostProcess(&_Honeinch.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0xf88309d7.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromTokenAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer, address[] callAddresses, bytes callDataConcat, uint256[] starts, uint256[] gasLimitsAndValues) payable returns(uint256 returnAmount)
func (_Honeinch *HoneinchTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, fromTokenAmount *big.Int, minReturnAmount *big.Int, guaranteedAmount *big.Int, referrer common.Address, callAddresses []common.Address, callDataConcat []byte, starts []*big.Int, gasLimitsAndValues []*big.Int) (*types.Transaction, error) {
	return _Honeinch.contract.Transact(opts, "swap", fromToken, toToken, fromTokenAmount, minReturnAmount, guaranteedAmount, referrer, callAddresses, callDataConcat, starts, gasLimitsAndValues)
}

// Swap is a paid mutator transaction binding the contract method 0xf88309d7.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromTokenAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer, address[] callAddresses, bytes callDataConcat, uint256[] starts, uint256[] gasLimitsAndValues) payable returns(uint256 returnAmount)
func (_Honeinch *HoneinchSession) Swap(fromToken common.Address, toToken common.Address, fromTokenAmount *big.Int, minReturnAmount *big.Int, guaranteedAmount *big.Int, referrer common.Address, callAddresses []common.Address, callDataConcat []byte, starts []*big.Int, gasLimitsAndValues []*big.Int) (*types.Transaction, error) {
	return _Honeinch.Contract.Swap(&_Honeinch.TransactOpts, fromToken, toToken, fromTokenAmount, minReturnAmount, guaranteedAmount, referrer, callAddresses, callDataConcat, starts, gasLimitsAndValues)
}

// Swap is a paid mutator transaction binding the contract method 0xf88309d7.
//
// Solidity: function swap(address fromToken, address toToken, uint256 fromTokenAmount, uint256 minReturnAmount, uint256 guaranteedAmount, address referrer, address[] callAddresses, bytes callDataConcat, uint256[] starts, uint256[] gasLimitsAndValues) payable returns(uint256 returnAmount)
func (_Honeinch *HoneinchTransactorSession) Swap(fromToken common.Address, toToken common.Address, fromTokenAmount *big.Int, minReturnAmount *big.Int, guaranteedAmount *big.Int, referrer common.Address, callAddresses []common.Address, callDataConcat []byte, starts []*big.Int, gasLimitsAndValues []*big.Int) (*types.Transaction, error) {
	return _Honeinch.Contract.Swap(&_Honeinch.TransactOpts, fromToken, toToken, fromTokenAmount, minReturnAmount, guaranteedAmount, referrer, callAddresses, callDataConcat, starts, gasLimitsAndValues)
}
