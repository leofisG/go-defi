// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswap

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

// UniswapABI is the input ABI used to generate the binding from.
const UniswapABI = "[{\"name\":\"TokenPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"EthPurchase\",\"inputs\":[{\"type\":\"address\",\"name\":\"buyer\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"tokens_sold\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_bought\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"AddLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"RemoveLiquidity\",\"inputs\":[{\"type\":\"address\",\"name\":\"provider\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"eth_amount\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"token_amount\",\"indexed\":true}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Transfer\",\"inputs\":[{\"type\":\"address\",\"name\":\"_from\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_to\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"Approval\",\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\",\"indexed\":true},{\"type\":\"address\",\"name\":\"_spender\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"_value\",\"indexed\":false}],\"anonymous\":false,\"type\":\"event\"},{\"name\":\"setup\",\"outputs\":[],\"inputs\":[{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":175875},{\"name\":\"addLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_liquidity\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":82616},{\"name\":\"removeLiquidity\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"},{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"amount\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":116814},{\"name\":\"__default__\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":true,\"type\":\"function\"},{\"name\":\"ethToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12757},{\"name\":\"ethToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"min_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":12965},{\"name\":\"ethToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50463},{\"name\":\"ethToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":true,\"type\":\"function\",\"gas\":50671},{\"name\":\"tokenToEthSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47503},{\"name\":\"tokenToEthTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_eth\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":47712},{\"name\":\"tokenToEthSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50175},{\"name\":\"tokenToEthTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":50384},{\"name\":\"tokenToTokenSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51007},{\"name\":\"tokenToTokenTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":51098},{\"name\":\"tokenToTokenSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":54928},{\"name\":\"tokenToTokenTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"token_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":55019},{\"name\":\"tokenToExchangeSwapInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49342},{\"name\":\"tokenToExchangeTransferInput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"},{\"type\":\"uint256\",\"name\":\"min_tokens_bought\"},{\"type\":\"uint256\",\"name\":\"min_eth_bought\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":49532},{\"name\":\"tokenToExchangeSwapOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53233},{\"name\":\"tokenToExchangeTransferOutput\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"},{\"type\":\"uint256\",\"name\":\"max_tokens_sold\"},{\"type\":\"uint256\",\"name\":\"max_eth_sold\"},{\"type\":\"uint256\",\"name\":\"deadline\"},{\"type\":\"address\",\"name\":\"recipient\"},{\"type\":\"address\",\"name\":\"exchange_addr\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":53423},{\"name\":\"getEthToTokenInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5542},{\"name\":\"getEthToTokenOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6872},{\"name\":\"getTokenToEthInputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"tokens_sold\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":5637},{\"name\":\"getTokenToEthOutputPrice\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"uint256\",\"name\":\"eth_bought\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":6897},{\"name\":\"tokenAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1413},{\"name\":\"factoryAddress\",\"outputs\":[{\"type\":\"address\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1443},{\"name\":\"balanceOf\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1645},{\"name\":\"transfer\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":75034},{\"name\":\"transferFrom\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_from\"},{\"type\":\"address\",\"name\":\"_to\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":110907},{\"name\":\"approve\",\"outputs\":[{\"type\":\"bool\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_spender\"},{\"type\":\"uint256\",\"name\":\"_value\"}],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":38769},{\"name\":\"allowance\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[{\"type\":\"address\",\"name\":\"_owner\"},{\"type\":\"address\",\"name\":\"_spender\"}],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1925},{\"name\":\"name\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1623},{\"name\":\"symbol\",\"outputs\":[{\"type\":\"bytes32\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1653},{\"name\":\"decimals\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1683},{\"name\":\"totalSupply\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":1713}]"

// Uniswap is an auto generated Go binding around an Ethereum contract.
type Uniswap struct {
	UniswapCaller     // Read-only binding to the contract
	UniswapTransactor // Write-only binding to the contract
	UniswapFilterer   // Log filterer for contract events
}

// UniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapSession struct {
	Contract     *Uniswap          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapCallerSession struct {
	Contract *UniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapTransactorSession struct {
	Contract     *UniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapRaw struct {
	Contract *Uniswap // Generic contract binding to access the raw methods on
}

// UniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapCallerRaw struct {
	Contract *UniswapCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapTransactorRaw struct {
	Contract *UniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswap creates a new instance of Uniswap, bound to a specific deployed contract.
func NewUniswap(address common.Address, backend bind.ContractBackend) (*Uniswap, error) {
	contract, err := bindUniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswap{UniswapCaller: UniswapCaller{contract: contract}, UniswapTransactor: UniswapTransactor{contract: contract}, UniswapFilterer: UniswapFilterer{contract: contract}}, nil
}

// NewUniswapCaller creates a new read-only instance of Uniswap, bound to a specific deployed contract.
func NewUniswapCaller(address common.Address, caller bind.ContractCaller) (*UniswapCaller, error) {
	contract, err := bindUniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapCaller{contract: contract}, nil
}

// NewUniswapTransactor creates a new write-only instance of Uniswap, bound to a specific deployed contract.
func NewUniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapTransactor, error) {
	contract, err := bindUniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapTransactor{contract: contract}, nil
}

// NewUniswapFilterer creates a new log filterer instance of Uniswap, bound to a specific deployed contract.
func NewUniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapFilterer, error) {
	contract, err := bindUniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapFilterer{contract: contract}, nil
}

// bindUniswap binds a generic wrapper to an already deployed contract.
func bindUniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UniswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap *UniswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswap.Contract.UniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap *UniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.Contract.UniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap *UniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap.Contract.UniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswap *UniswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswap *UniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswap *UniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswap.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_Uniswap *UniswapCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_Uniswap *UniswapSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Uniswap.Contract.Allowance(&_Uniswap.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) returns(uint256 out)
func (_Uniswap *UniswapCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Uniswap.Contract.Allowance(&_Uniswap.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 out)
func (_Uniswap *UniswapCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 out)
func (_Uniswap *UniswapSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Uniswap.Contract.BalanceOf(&_Uniswap.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) returns(uint256 out)
func (_Uniswap *UniswapCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Uniswap.Contract.BalanceOf(&_Uniswap.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_Uniswap *UniswapCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_Uniswap *UniswapSession) Decimals() (*big.Int, error) {
	return _Uniswap.Contract.Decimals(&_Uniswap.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() returns(uint256 out)
func (_Uniswap *UniswapCallerSession) Decimals() (*big.Int, error) {
	return _Uniswap.Contract.Decimals(&_Uniswap.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() returns(address out)
func (_Uniswap *UniswapCaller) FactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "factoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() returns(address out)
func (_Uniswap *UniswapSession) FactoryAddress() (common.Address, error) {
	return _Uniswap.Contract.FactoryAddress(&_Uniswap.CallOpts)
}

// FactoryAddress is a free data retrieval call binding the contract method 0x966dae0e.
//
// Solidity: function factoryAddress() returns(address out)
func (_Uniswap *UniswapCallerSession) FactoryAddress() (common.Address, error) {
	return _Uniswap.Contract.FactoryAddress(&_Uniswap.CallOpts)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) returns(uint256 out)
func (_Uniswap *UniswapCaller) GetEthToTokenInputPrice(opts *bind.CallOpts, eth_sold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getEthToTokenInputPrice", eth_sold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) returns(uint256 out)
func (_Uniswap *UniswapSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetEthToTokenInputPrice(&_Uniswap.CallOpts, eth_sold)
}

// GetEthToTokenInputPrice is a free data retrieval call binding the contract method 0xcd7724c3.
//
// Solidity: function getEthToTokenInputPrice(uint256 eth_sold) returns(uint256 out)
func (_Uniswap *UniswapCallerSession) GetEthToTokenInputPrice(eth_sold *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetEthToTokenInputPrice(&_Uniswap.CallOpts, eth_sold)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) returns(uint256 out)
func (_Uniswap *UniswapCaller) GetEthToTokenOutputPrice(opts *bind.CallOpts, tokens_bought *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getEthToTokenOutputPrice", tokens_bought)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) returns(uint256 out)
func (_Uniswap *UniswapSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetEthToTokenOutputPrice(&_Uniswap.CallOpts, tokens_bought)
}

// GetEthToTokenOutputPrice is a free data retrieval call binding the contract method 0x59e94862.
//
// Solidity: function getEthToTokenOutputPrice(uint256 tokens_bought) returns(uint256 out)
func (_Uniswap *UniswapCallerSession) GetEthToTokenOutputPrice(tokens_bought *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetEthToTokenOutputPrice(&_Uniswap.CallOpts, tokens_bought)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) returns(uint256 out)
func (_Uniswap *UniswapCaller) GetTokenToEthInputPrice(opts *bind.CallOpts, tokens_sold *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getTokenToEthInputPrice", tokens_sold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) returns(uint256 out)
func (_Uniswap *UniswapSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetTokenToEthInputPrice(&_Uniswap.CallOpts, tokens_sold)
}

// GetTokenToEthInputPrice is a free data retrieval call binding the contract method 0x95b68fe7.
//
// Solidity: function getTokenToEthInputPrice(uint256 tokens_sold) returns(uint256 out)
func (_Uniswap *UniswapCallerSession) GetTokenToEthInputPrice(tokens_sold *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetTokenToEthInputPrice(&_Uniswap.CallOpts, tokens_sold)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) returns(uint256 out)
func (_Uniswap *UniswapCaller) GetTokenToEthOutputPrice(opts *bind.CallOpts, eth_bought *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "getTokenToEthOutputPrice", eth_bought)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) returns(uint256 out)
func (_Uniswap *UniswapSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetTokenToEthOutputPrice(&_Uniswap.CallOpts, eth_bought)
}

// GetTokenToEthOutputPrice is a free data retrieval call binding the contract method 0x2640f62c.
//
// Solidity: function getTokenToEthOutputPrice(uint256 eth_bought) returns(uint256 out)
func (_Uniswap *UniswapCallerSession) GetTokenToEthOutputPrice(eth_bought *big.Int) (*big.Int, error) {
	return _Uniswap.Contract.GetTokenToEthOutputPrice(&_Uniswap.CallOpts, eth_bought)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32 out)
func (_Uniswap *UniswapCaller) Name(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "name")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32 out)
func (_Uniswap *UniswapSession) Name() ([32]byte, error) {
	return _Uniswap.Contract.Name(&_Uniswap.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() returns(bytes32 out)
func (_Uniswap *UniswapCallerSession) Name() ([32]byte, error) {
	return _Uniswap.Contract.Name(&_Uniswap.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32 out)
func (_Uniswap *UniswapCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32 out)
func (_Uniswap *UniswapSession) Symbol() ([32]byte, error) {
	return _Uniswap.Contract.Symbol(&_Uniswap.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() returns(bytes32 out)
func (_Uniswap *UniswapCallerSession) Symbol() ([32]byte, error) {
	return _Uniswap.Contract.Symbol(&_Uniswap.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() returns(address out)
func (_Uniswap *UniswapCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "tokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() returns(address out)
func (_Uniswap *UniswapSession) TokenAddress() (common.Address, error) {
	return _Uniswap.Contract.TokenAddress(&_Uniswap.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() returns(address out)
func (_Uniswap *UniswapCallerSession) TokenAddress() (common.Address, error) {
	return _Uniswap.Contract.TokenAddress(&_Uniswap.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_Uniswap *UniswapCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswap.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_Uniswap *UniswapSession) TotalSupply() (*big.Int, error) {
	return _Uniswap.Contract.TotalSupply(&_Uniswap.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256 out)
func (_Uniswap *UniswapCallerSession) TotalSupply() (*big.Int, error) {
	return _Uniswap.Contract.TotalSupply(&_Uniswap.CallOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Uniswap *UniswapTransactor) Default(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "__default__")
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Uniswap *UniswapSession) Default() (*types.Transaction, error) {
	return _Uniswap.Contract.Default(&_Uniswap.TransactOpts)
}

// Default is a paid mutator transaction binding the contract method 0x89402a72.
//
// Solidity: function __default__() returns()
func (_Uniswap *UniswapTransactorSession) Default() (*types.Transaction, error) {
	return _Uniswap.Contract.Default(&_Uniswap.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactor) AddLiquidity(opts *bind.TransactOpts, min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "addLiquidity", min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.AddLiquidity(&_Uniswap.TransactOpts, min_liquidity, max_tokens, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x422f1043.
//
// Solidity: function addLiquidity(uint256 min_liquidity, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) AddLiquidity(min_liquidity *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.AddLiquidity(&_Uniswap.TransactOpts, min_liquidity, max_tokens, deadline)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Uniswap *UniswapTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Uniswap *UniswapSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Approve(&_Uniswap.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool out)
func (_Uniswap *UniswapTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Approve(&_Uniswap.TransactOpts, _spender, _value)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactor) EthToTokenSwapInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "ethToTokenSwapInput", min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenSwapInput(&_Uniswap.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapInput is a paid mutator transaction binding the contract method 0xf39b5b9b.
//
// Solidity: function ethToTokenSwapInput(uint256 min_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) EthToTokenSwapInput(min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenSwapInput(&_Uniswap.TransactOpts, min_tokens, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactor) EthToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "ethToTokenSwapOutput", tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenSwapOutput(&_Uniswap.TransactOpts, tokens_bought, deadline)
}

// EthToTokenSwapOutput is a paid mutator transaction binding the contract method 0x6b1d4db7.
//
// Solidity: function ethToTokenSwapOutput(uint256 tokens_bought, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) EthToTokenSwapOutput(tokens_bought *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenSwapOutput(&_Uniswap.TransactOpts, tokens_bought, deadline)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactor) EthToTokenTransferInput(opts *bind.TransactOpts, min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "ethToTokenTransferInput", min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenTransferInput(&_Uniswap.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferInput is a paid mutator transaction binding the contract method 0xad65d76d.
//
// Solidity: function ethToTokenTransferInput(uint256 min_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) EthToTokenTransferInput(min_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenTransferInput(&_Uniswap.TransactOpts, min_tokens, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactor) EthToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "ethToTokenTransferOutput", tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenTransferOutput(&_Uniswap.TransactOpts, tokens_bought, deadline, recipient)
}

// EthToTokenTransferOutput is a paid mutator transaction binding the contract method 0x0b573638.
//
// Solidity: function ethToTokenTransferOutput(uint256 tokens_bought, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) EthToTokenTransferOutput(tokens_bought *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.EthToTokenTransferOutput(&_Uniswap.TransactOpts, tokens_bought, deadline, recipient)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Uniswap *UniswapTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "removeLiquidity", amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Uniswap *UniswapSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.RemoveLiquidity(&_Uniswap.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xf88bf15a.
//
// Solidity: function removeLiquidity(uint256 amount, uint256 min_eth, uint256 min_tokens, uint256 deadline) returns(uint256 out, uint256 out)
func (_Uniswap *UniswapTransactorSession) RemoveLiquidity(amount *big.Int, min_eth *big.Int, min_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.RemoveLiquidity(&_Uniswap.TransactOpts, amount, min_eth, min_tokens, deadline)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Uniswap *UniswapTransactor) Setup(opts *bind.TransactOpts, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "setup", token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Uniswap *UniswapSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.Setup(&_Uniswap.TransactOpts, token_addr)
}

// Setup is a paid mutator transaction binding the contract method 0x66d38203.
//
// Solidity: function setup(address token_addr) returns()
func (_Uniswap *UniswapTransactorSession) Setup(token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.Setup(&_Uniswap.TransactOpts, token_addr)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToEthSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToEthSwapInput", tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthSwapInput(&_Uniswap.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapInput is a paid mutator transaction binding the contract method 0x95e3c50b.
//
// Solidity: function tokenToEthSwapInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToEthSwapInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthSwapInput(&_Uniswap.TransactOpts, tokens_sold, min_eth, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToEthSwapOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToEthSwapOutput", eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthSwapOutput(&_Uniswap.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthSwapOutput is a paid mutator transaction binding the contract method 0x013efd8b.
//
// Solidity: function tokenToEthSwapOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToEthSwapOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthSwapOutput(&_Uniswap.TransactOpts, eth_bought, max_tokens, deadline)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToEthTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToEthTransferInput", tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthTransferInput(&_Uniswap.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferInput is a paid mutator transaction binding the contract method 0x7237e031.
//
// Solidity: function tokenToEthTransferInput(uint256 tokens_sold, uint256 min_eth, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToEthTransferInput(tokens_sold *big.Int, min_eth *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthTransferInput(&_Uniswap.TransactOpts, tokens_sold, min_eth, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToEthTransferOutput(opts *bind.TransactOpts, eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToEthTransferOutput", eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthTransferOutput(&_Uniswap.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToEthTransferOutput is a paid mutator transaction binding the contract method 0xd4e4841d.
//
// Solidity: function tokenToEthTransferOutput(uint256 eth_bought, uint256 max_tokens, uint256 deadline, address recipient) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToEthTransferOutput(eth_bought *big.Int, max_tokens *big.Int, deadline *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToEthTransferOutput(&_Uniswap.TransactOpts, eth_bought, max_tokens, deadline, recipient)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToExchangeSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToExchangeSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeSwapInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapInput is a paid mutator transaction binding the contract method 0xb1cb43bf.
//
// Solidity: function tokenToExchangeSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToExchangeSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeSwapInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToExchangeSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToExchangeSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeSwapOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeSwapOutput is a paid mutator transaction binding the contract method 0xea650c7d.
//
// Solidity: function tokenToExchangeSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToExchangeSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeSwapOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToExchangeTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToExchangeTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeTransferInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferInput is a paid mutator transaction binding the contract method 0xec384a3e.
//
// Solidity: function tokenToExchangeTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToExchangeTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeTransferInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToExchangeTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToExchangeTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeTransferOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToExchangeTransferOutput is a paid mutator transaction binding the contract method 0x981a1327.
//
// Solidity: function tokenToExchangeTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address exchange_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToExchangeTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, exchange_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToExchangeTransferOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, exchange_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToTokenSwapInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToTokenSwapInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenSwapInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapInput is a paid mutator transaction binding the contract method 0xddf7e1a7.
//
// Solidity: function tokenToTokenSwapInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToTokenSwapInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenSwapInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToTokenSwapOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToTokenSwapOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenSwapOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenSwapOutput is a paid mutator transaction binding the contract method 0xb040d545.
//
// Solidity: function tokenToTokenSwapOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToTokenSwapOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenSwapOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToTokenTransferInput(opts *bind.TransactOpts, tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToTokenTransferInput", tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenTransferInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferInput is a paid mutator transaction binding the contract method 0xf552d91b.
//
// Solidity: function tokenToTokenTransferInput(uint256 tokens_sold, uint256 min_tokens_bought, uint256 min_eth_bought, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToTokenTransferInput(tokens_sold *big.Int, min_tokens_bought *big.Int, min_eth_bought *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenTransferInput(&_Uniswap.TransactOpts, tokens_sold, min_tokens_bought, min_eth_bought, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactor) TokenToTokenTransferOutput(opts *bind.TransactOpts, tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "tokenToTokenTransferOutput", tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenTransferOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// TokenToTokenTransferOutput is a paid mutator transaction binding the contract method 0xf3c0efe9.
//
// Solidity: function tokenToTokenTransferOutput(uint256 tokens_bought, uint256 max_tokens_sold, uint256 max_eth_sold, uint256 deadline, address recipient, address token_addr) returns(uint256 out)
func (_Uniswap *UniswapTransactorSession) TokenToTokenTransferOutput(tokens_bought *big.Int, max_tokens_sold *big.Int, max_eth_sold *big.Int, deadline *big.Int, recipient common.Address, token_addr common.Address) (*types.Transaction, error) {
	return _Uniswap.Contract.TokenToTokenTransferOutput(&_Uniswap.TransactOpts, tokens_bought, max_tokens_sold, max_eth_sold, deadline, recipient, token_addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Uniswap *UniswapTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Uniswap *UniswapSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Transfer(&_Uniswap.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool out)
func (_Uniswap *UniswapTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.Transfer(&_Uniswap.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Uniswap *UniswapTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Uniswap *UniswapSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.TransferFrom(&_Uniswap.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool out)
func (_Uniswap *UniswapTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Uniswap.Contract.TransferFrom(&_Uniswap.TransactOpts, _from, _to, _value)
}

// UniswapAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Uniswap contract.
type UniswapAddLiquidityIterator struct {
	Event *UniswapAddLiquidity // Event containing the contract specifics and raw log

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
func (it *UniswapAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapAddLiquidity)
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
		it.Event = new(UniswapAddLiquidity)
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
func (it *UniswapAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapAddLiquidity represents a AddLiquidity event raised by the Uniswap contract.
type UniswapAddLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswap *UniswapFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*UniswapAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &UniswapAddLiquidityIterator{contract: _Uniswap.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswap *UniswapFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *UniswapAddLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "AddLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapAddLiquidity)
				if err := _Uniswap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x06239653922ac7bea6aa2b19dc486b9361821d37712eb796adfd38d81de278ca.
//
// Solidity: event AddLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswap *UniswapFilterer) ParseAddLiquidity(log types.Log) (*UniswapAddLiquidity, error) {
	event := new(UniswapAddLiquidity)
	if err := _Uniswap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Uniswap contract.
type UniswapApprovalIterator struct {
	Event *UniswapApproval // Event containing the contract specifics and raw log

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
func (it *UniswapApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapApproval)
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
		it.Event = new(UniswapApproval)
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
func (it *UniswapApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapApproval represents a Approval event raised by the Uniswap contract.
type UniswapApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Uniswap *UniswapFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*UniswapApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &UniswapApprovalIterator{contract: _Uniswap.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Uniswap *UniswapFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *UniswapApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapApproval)
				if err := _Uniswap.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Uniswap *UniswapFilterer) ParseApproval(log types.Log) (*UniswapApproval, error) {
	event := new(UniswapApproval)
	if err := _Uniswap.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapEthPurchaseIterator is returned from FilterEthPurchase and is used to iterate over the raw logs and unpacked data for EthPurchase events raised by the Uniswap contract.
type UniswapEthPurchaseIterator struct {
	Event *UniswapEthPurchase // Event containing the contract specifics and raw log

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
func (it *UniswapEthPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapEthPurchase)
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
		it.Event = new(UniswapEthPurchase)
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
func (it *UniswapEthPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapEthPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapEthPurchase represents a EthPurchase event raised by the Uniswap contract.
type UniswapEthPurchase struct {
	Buyer      common.Address
	TokensSold *big.Int
	EthBought  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEthPurchase is a free log retrieval operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Uniswap *UniswapFilterer) FilterEthPurchase(opts *bind.FilterOpts, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (*UniswapEthPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return &UniswapEthPurchaseIterator{contract: _Uniswap.contract, event: "EthPurchase", logs: logs, sub: sub}, nil
}

// WatchEthPurchase is a free log subscription operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Uniswap *UniswapFilterer) WatchEthPurchase(opts *bind.WatchOpts, sink chan<- *UniswapEthPurchase, buyer []common.Address, tokens_sold []*big.Int, eth_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokens_soldRule []interface{}
	for _, tokens_soldItem := range tokens_sold {
		tokens_soldRule = append(tokens_soldRule, tokens_soldItem)
	}
	var eth_boughtRule []interface{}
	for _, eth_boughtItem := range eth_bought {
		eth_boughtRule = append(eth_boughtRule, eth_boughtItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "EthPurchase", buyerRule, tokens_soldRule, eth_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapEthPurchase)
				if err := _Uniswap.contract.UnpackLog(event, "EthPurchase", log); err != nil {
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

// ParseEthPurchase is a log parse operation binding the contract event 0x7f4091b46c33e918a0f3aa42307641d17bb67029427a5369e54b353984238705.
//
// Solidity: event EthPurchase(address indexed buyer, uint256 indexed tokens_sold, uint256 indexed eth_bought)
func (_Uniswap *UniswapFilterer) ParseEthPurchase(log types.Log) (*UniswapEthPurchase, error) {
	event := new(UniswapEthPurchase)
	if err := _Uniswap.contract.UnpackLog(event, "EthPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Uniswap contract.
type UniswapRemoveLiquidityIterator struct {
	Event *UniswapRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *UniswapRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapRemoveLiquidity)
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
		it.Event = new(UniswapRemoveLiquidity)
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
func (it *UniswapRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapRemoveLiquidity represents a RemoveLiquidity event raised by the Uniswap contract.
type UniswapRemoveLiquidity struct {
	Provider    common.Address
	EthAmount   *big.Int
	TokenAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswap *UniswapFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (*UniswapRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return &UniswapRemoveLiquidityIterator{contract: _Uniswap.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswap *UniswapFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *UniswapRemoveLiquidity, provider []common.Address, eth_amount []*big.Int, token_amount []*big.Int) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var eth_amountRule []interface{}
	for _, eth_amountItem := range eth_amount {
		eth_amountRule = append(eth_amountRule, eth_amountItem)
	}
	var token_amountRule []interface{}
	for _, token_amountItem := range token_amount {
		token_amountRule = append(token_amountRule, token_amountItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "RemoveLiquidity", providerRule, eth_amountRule, token_amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapRemoveLiquidity)
				if err := _Uniswap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x0fbf06c058b90cb038a618f8c2acbf6145f8b3570fd1fa56abb8f0f3f05b36e8.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256 indexed eth_amount, uint256 indexed token_amount)
func (_Uniswap *UniswapFilterer) ParseRemoveLiquidity(log types.Log) (*UniswapRemoveLiquidity, error) {
	event := new(UniswapRemoveLiquidity)
	if err := _Uniswap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the Uniswap contract.
type UniswapTokenPurchaseIterator struct {
	Event *UniswapTokenPurchase // Event containing the contract specifics and raw log

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
func (it *UniswapTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapTokenPurchase)
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
		it.Event = new(UniswapTokenPurchase)
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
func (it *UniswapTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapTokenPurchase represents a TokenPurchase event raised by the Uniswap contract.
type UniswapTokenPurchase struct {
	Buyer        common.Address
	EthSold      *big.Int
	TokensBought *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Uniswap *UniswapFilterer) FilterTokenPurchase(opts *bind.FilterOpts, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (*UniswapTokenPurchaseIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return &UniswapTokenPurchaseIterator{contract: _Uniswap.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Uniswap *UniswapFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *UniswapTokenPurchase, buyer []common.Address, eth_sold []*big.Int, tokens_bought []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var eth_soldRule []interface{}
	for _, eth_soldItem := range eth_sold {
		eth_soldRule = append(eth_soldRule, eth_soldItem)
	}
	var tokens_boughtRule []interface{}
	for _, tokens_boughtItem := range tokens_bought {
		tokens_boughtRule = append(tokens_boughtRule, tokens_boughtItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "TokenPurchase", buyerRule, eth_soldRule, tokens_boughtRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapTokenPurchase)
				if err := _Uniswap.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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

// ParseTokenPurchase is a log parse operation binding the contract event 0xcd60aa75dea3072fbc07ae6d7d856b5dc5f4eee88854f5b4abf7b680ef8bc50f.
//
// Solidity: event TokenPurchase(address indexed buyer, uint256 indexed eth_sold, uint256 indexed tokens_bought)
func (_Uniswap *UniswapFilterer) ParseTokenPurchase(log types.Log) (*UniswapTokenPurchase, error) {
	event := new(UniswapTokenPurchase)
	if err := _Uniswap.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
}

// UniswapTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Uniswap contract.
type UniswapTransferIterator struct {
	Event *UniswapTransfer // Event containing the contract specifics and raw log

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
func (it *UniswapTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniswapTransfer)
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
		it.Event = new(UniswapTransfer)
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
func (it *UniswapTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniswapTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniswapTransfer represents a Transfer event raised by the Uniswap contract.
type UniswapTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Uniswap *UniswapFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*UniswapTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Uniswap.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &UniswapTransferIterator{contract: _Uniswap.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Uniswap *UniswapFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *UniswapTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Uniswap.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniswapTransfer)
				if err := _Uniswap.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Uniswap *UniswapFilterer) ParseTransfer(log types.Log) (*UniswapTransfer, error) {
	event := new(UniswapTransfer)
	if err := _Uniswap.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
