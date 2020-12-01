// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package huniswap

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

// HuniswapABI is the input ABI used to generate the binding from.
const HuniswapABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Huniswap is an auto generated Go binding around an Ethereum contract.
type Huniswap struct {
	HuniswapCaller     // Read-only binding to the contract
	HuniswapTransactor // Write-only binding to the contract
	HuniswapFilterer   // Log filterer for contract events
}

// HuniswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type HuniswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HuniswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HuniswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HuniswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HuniswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HuniswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HuniswapSession struct {
	Contract     *Huniswap         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HuniswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HuniswapCallerSession struct {
	Contract *HuniswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HuniswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HuniswapTransactorSession struct {
	Contract     *HuniswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HuniswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type HuniswapRaw struct {
	Contract *Huniswap // Generic contract binding to access the raw methods on
}

// HuniswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HuniswapCallerRaw struct {
	Contract *HuniswapCaller // Generic read-only contract binding to access the raw methods on
}

// HuniswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HuniswapTransactorRaw struct {
	Contract *HuniswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHuniswap creates a new instance of Huniswap, bound to a specific deployed contract.
func NewHuniswap(address common.Address, backend bind.ContractBackend) (*Huniswap, error) {
	contract, err := bindHuniswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Huniswap{HuniswapCaller: HuniswapCaller{contract: contract}, HuniswapTransactor: HuniswapTransactor{contract: contract}, HuniswapFilterer: HuniswapFilterer{contract: contract}}, nil
}

// NewHuniswapCaller creates a new read-only instance of Huniswap, bound to a specific deployed contract.
func NewHuniswapCaller(address common.Address, caller bind.ContractCaller) (*HuniswapCaller, error) {
	contract, err := bindHuniswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HuniswapCaller{contract: contract}, nil
}

// NewHuniswapTransactor creates a new write-only instance of Huniswap, bound to a specific deployed contract.
func NewHuniswapTransactor(address common.Address, transactor bind.ContractTransactor) (*HuniswapTransactor, error) {
	contract, err := bindHuniswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HuniswapTransactor{contract: contract}, nil
}

// NewHuniswapFilterer creates a new log filterer instance of Huniswap, bound to a specific deployed contract.
func NewHuniswapFilterer(address common.Address, filterer bind.ContractFilterer) (*HuniswapFilterer, error) {
	contract, err := bindHuniswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HuniswapFilterer{contract: contract}, nil
}

// bindHuniswap binds a generic wrapper to an already deployed contract.
func bindHuniswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HuniswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Huniswap *HuniswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Huniswap.Contract.HuniswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Huniswap *HuniswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Huniswap.Contract.HuniswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Huniswap *HuniswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Huniswap.Contract.HuniswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Huniswap *HuniswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Huniswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Huniswap *HuniswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Huniswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Huniswap *HuniswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Huniswap.Contract.contract.Transact(opts, method, params...)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x3351733f.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin) payable returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Huniswap *HuniswapTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x3351733f.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin) payable returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Huniswap *HuniswapSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.AddLiquidity(&_Huniswap.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x3351733f.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin) payable returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_Huniswap *HuniswapTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.AddLiquidity(&_Huniswap.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x58871c81.
//
// Solidity: function addLiquidityETH(uint256 value, address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Huniswap *HuniswapTransactor) AddLiquidityETH(opts *bind.TransactOpts, value *big.Int, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "addLiquidityETH", value, token, amountTokenDesired, amountTokenMin, amountETHMin)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x58871c81.
//
// Solidity: function addLiquidityETH(uint256 value, address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Huniswap *HuniswapSession) AddLiquidityETH(value *big.Int, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.AddLiquidityETH(&_Huniswap.TransactOpts, value, token, amountTokenDesired, amountTokenMin, amountETHMin)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x58871c81.
//
// Solidity: function addLiquidityETH(uint256 value, address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_Huniswap *HuniswapTransactorSession) AddLiquidityETH(value *big.Int, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.AddLiquidityETH(&_Huniswap.TransactOpts, value, token, amountTokenDesired, amountTokenMin, amountETHMin)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Huniswap *HuniswapTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Huniswap *HuniswapSession) PostProcess() (*types.Transaction, error) {
	return _Huniswap.Contract.PostProcess(&_Huniswap.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Huniswap *HuniswapTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Huniswap.Contract.PostProcess(&_Huniswap.TransactOpts)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin) payable returns(uint256 amountA, uint256 amountB)
func (_Huniswap *HuniswapTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin) payable returns(uint256 amountA, uint256 amountB)
func (_Huniswap *HuniswapSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.RemoveLiquidity(&_Huniswap.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0xe2dc85dc.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin) payable returns(uint256 amountA, uint256 amountB)
func (_Huniswap *HuniswapTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.RemoveLiquidity(&_Huniswap.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xa1cfacde.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin) payable returns(uint256 amountToken, uint256 amountETH)
func (_Huniswap *HuniswapTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xa1cfacde.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin) payable returns(uint256 amountToken, uint256 amountETH)
func (_Huniswap *HuniswapSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.RemoveLiquidityETH(&_Huniswap.TransactOpts, token, liquidity, amountTokenMin, amountETHMin)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xa1cfacde.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin) payable returns(uint256 amountToken, uint256 amountETH)
func (_Huniswap *HuniswapTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int) (*types.Transaction, error) {
	return _Huniswap.Contract.RemoveLiquidityETH(&_Huniswap.TransactOpts, token, liquidity, amountTokenMin, amountETHMin)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x87151a79.
//
// Solidity: function swapETHForExactTokens(uint256 value, uint256 amountOut, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, value *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "swapETHForExactTokens", value, amountOut, path)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x87151a79.
//
// Solidity: function swapETHForExactTokens(uint256 value, uint256 amountOut, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapSession) SwapETHForExactTokens(value *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapETHForExactTokens(&_Huniswap.TransactOpts, value, amountOut, path)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x87151a79.
//
// Solidity: function swapETHForExactTokens(uint256 value, uint256 amountOut, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactorSession) SwapETHForExactTokens(value *big.Int, amountOut *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapETHForExactTokens(&_Huniswap.TransactOpts, value, amountOut, path)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0xd0241dac.
//
// Solidity: function swapExactETHForTokens(uint256 value, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, value *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "swapExactETHForTokens", value, amountOutMin, path)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0xd0241dac.
//
// Solidity: function swapExactETHForTokens(uint256 value, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapSession) SwapExactETHForTokens(value *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapExactETHForTokens(&_Huniswap.TransactOpts, value, amountOutMin, path)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0xd0241dac.
//
// Solidity: function swapExactETHForTokens(uint256 value, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactorSession) SwapExactETHForTokens(value *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapExactETHForTokens(&_Huniswap.TransactOpts, value, amountOutMin, path)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xef66f725.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xef66f725.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapExactTokensForETH(&_Huniswap.TransactOpts, amountIn, amountOutMin, path)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xef66f725.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapExactTokensForETH(&_Huniswap.TransactOpts, amountIn, amountOutMin, path)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x86818f26.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x86818f26.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapExactTokensForTokens(&_Huniswap.TransactOpts, amountIn, amountOutMin, path)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x86818f26.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapExactTokensForTokens(&_Huniswap.TransactOpts, amountIn, amountOutMin, path)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x18a22c40.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "swapTokensForExactETH", amountOut, amountInMax, path)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x18a22c40.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapTokensForExactETH(&_Huniswap.TransactOpts, amountOut, amountInMax, path)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x18a22c40.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut, uint256 amountInMax, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactorSession) SwapTokensForExactETH(amountOut *big.Int, amountInMax *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapTokensForExactETH(&_Huniswap.TransactOpts, amountOut, amountInMax, path)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x397d4b4a.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut *big.Int, amountInMax *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.contract.Transact(opts, "swapTokensForExactTokens", amountOut, amountInMax, path)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x397d4b4a.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapTokensForExactTokens(&_Huniswap.TransactOpts, amountOut, amountInMax, path)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x397d4b4a.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut, uint256 amountInMax, address[] path) payable returns(uint256[] amounts)
func (_Huniswap *HuniswapTransactorSession) SwapTokensForExactTokens(amountOut *big.Int, amountInMax *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Huniswap.Contract.SwapTokensForExactTokens(&_Huniswap.TransactOpts, amountOut, amountInMax, path)
}
