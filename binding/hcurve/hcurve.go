// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hcurve

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

// HcurveABI is the input ABI used to generate the binding from.
const HcurveABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ONE_SPLIT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minMintAmount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenI\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenJ\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"}],\"name\":\"exchange\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenI\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenJ\",\"type\":\"address\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"j\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"}],\"name\":\"exchangeUnderlying\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenI\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneCoin\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"handler\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenI\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"int128\",\"name\":\"i\",\"type\":\"int128\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneCoinDust\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReturn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"distribution\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"featureFlags\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Hcurve is an auto generated Go binding around an Ethereum contract.
type Hcurve struct {
	HcurveCaller     // Read-only binding to the contract
	HcurveTransactor // Write-only binding to the contract
	HcurveFilterer   // Log filterer for contract events
}

// HcurveCaller is an auto generated read-only Go binding around an Ethereum contract.
type HcurveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HcurveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HcurveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HcurveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HcurveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HcurveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HcurveSession struct {
	Contract     *Hcurve           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HcurveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HcurveCallerSession struct {
	Contract *HcurveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HcurveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HcurveTransactorSession struct {
	Contract     *HcurveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HcurveRaw is an auto generated low-level Go binding around an Ethereum contract.
type HcurveRaw struct {
	Contract *Hcurve // Generic contract binding to access the raw methods on
}

// HcurveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HcurveCallerRaw struct {
	Contract *HcurveCaller // Generic read-only contract binding to access the raw methods on
}

// HcurveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HcurveTransactorRaw struct {
	Contract *HcurveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHcurve creates a new instance of Hcurve, bound to a specific deployed contract.
func NewHcurve(address common.Address, backend bind.ContractBackend) (*Hcurve, error) {
	contract, err := bindHcurve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hcurve{HcurveCaller: HcurveCaller{contract: contract}, HcurveTransactor: HcurveTransactor{contract: contract}, HcurveFilterer: HcurveFilterer{contract: contract}}, nil
}

// NewHcurveCaller creates a new read-only instance of Hcurve, bound to a specific deployed contract.
func NewHcurveCaller(address common.Address, caller bind.ContractCaller) (*HcurveCaller, error) {
	contract, err := bindHcurve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HcurveCaller{contract: contract}, nil
}

// NewHcurveTransactor creates a new write-only instance of Hcurve, bound to a specific deployed contract.
func NewHcurveTransactor(address common.Address, transactor bind.ContractTransactor) (*HcurveTransactor, error) {
	contract, err := bindHcurve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HcurveTransactor{contract: contract}, nil
}

// NewHcurveFilterer creates a new log filterer instance of Hcurve, bound to a specific deployed contract.
func NewHcurveFilterer(address common.Address, filterer bind.ContractFilterer) (*HcurveFilterer, error) {
	contract, err := bindHcurve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HcurveFilterer{contract: contract}, nil
}

// bindHcurve binds a generic wrapper to an already deployed contract.
func bindHcurve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HcurveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hcurve *HcurveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hcurve.Contract.HcurveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hcurve *HcurveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hcurve.Contract.HcurveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hcurve *HcurveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hcurve.Contract.HcurveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hcurve *HcurveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hcurve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hcurve *HcurveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hcurve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hcurve *HcurveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hcurve.Contract.contract.Transact(opts, method, params...)
}

// ONESPLIT is a free data retrieval call binding the contract method 0x33a469e6.
//
// Solidity: function ONE_SPLIT() view returns(address)
func (_Hcurve *HcurveCaller) ONESPLIT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hcurve.contract.Call(opts, &out, "ONE_SPLIT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ONESPLIT is a free data retrieval call binding the contract method 0x33a469e6.
//
// Solidity: function ONE_SPLIT() view returns(address)
func (_Hcurve *HcurveSession) ONESPLIT() (common.Address, error) {
	return _Hcurve.Contract.ONESPLIT(&_Hcurve.CallOpts)
}

// ONESPLIT is a free data retrieval call binding the contract method 0x33a469e6.
//
// Solidity: function ONE_SPLIT() view returns(address)
func (_Hcurve *HcurveCallerSession) ONESPLIT() (common.Address, error) {
	return _Hcurve.Contract.ONESPLIT(&_Hcurve.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51a34ca5.
//
// Solidity: function addLiquidity(address handler, address pool, address[] tokens, uint256[] amounts, uint256 minMintAmount) payable returns()
func (_Hcurve *HcurveTransactor) AddLiquidity(opts *bind.TransactOpts, handler common.Address, pool common.Address, tokens []common.Address, amounts []*big.Int, minMintAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "addLiquidity", handler, pool, tokens, amounts, minMintAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51a34ca5.
//
// Solidity: function addLiquidity(address handler, address pool, address[] tokens, uint256[] amounts, uint256 minMintAmount) payable returns()
func (_Hcurve *HcurveSession) AddLiquidity(handler common.Address, pool common.Address, tokens []common.Address, amounts []*big.Int, minMintAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.AddLiquidity(&_Hcurve.TransactOpts, handler, pool, tokens, amounts, minMintAmount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51a34ca5.
//
// Solidity: function addLiquidity(address handler, address pool, address[] tokens, uint256[] amounts, uint256 minMintAmount) payable returns()
func (_Hcurve *HcurveTransactorSession) AddLiquidity(handler common.Address, pool common.Address, tokens []common.Address, amounts []*big.Int, minMintAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.AddLiquidity(&_Hcurve.TransactOpts, handler, pool, tokens, amounts, minMintAmount)
}

// Exchange is a paid mutator transaction binding the contract method 0x8337782d.
//
// Solidity: function exchange(address handler, address tokenI, address tokenJ, int128 i, int128 j, uint256 dx, uint256 minDy) payable returns()
func (_Hcurve *HcurveTransactor) Exchange(opts *bind.TransactOpts, handler common.Address, tokenI common.Address, tokenJ common.Address, i *big.Int, j *big.Int, dx *big.Int, minDy *big.Int) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "exchange", handler, tokenI, tokenJ, i, j, dx, minDy)
}

// Exchange is a paid mutator transaction binding the contract method 0x8337782d.
//
// Solidity: function exchange(address handler, address tokenI, address tokenJ, int128 i, int128 j, uint256 dx, uint256 minDy) payable returns()
func (_Hcurve *HcurveSession) Exchange(handler common.Address, tokenI common.Address, tokenJ common.Address, i *big.Int, j *big.Int, dx *big.Int, minDy *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.Exchange(&_Hcurve.TransactOpts, handler, tokenI, tokenJ, i, j, dx, minDy)
}

// Exchange is a paid mutator transaction binding the contract method 0x8337782d.
//
// Solidity: function exchange(address handler, address tokenI, address tokenJ, int128 i, int128 j, uint256 dx, uint256 minDy) payable returns()
func (_Hcurve *HcurveTransactorSession) Exchange(handler common.Address, tokenI common.Address, tokenJ common.Address, i *big.Int, j *big.Int, dx *big.Int, minDy *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.Exchange(&_Hcurve.TransactOpts, handler, tokenI, tokenJ, i, j, dx, minDy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xfef6074e.
//
// Solidity: function exchangeUnderlying(address handler, address tokenI, address tokenJ, int128 i, int128 j, uint256 dx, uint256 minDy) payable returns()
func (_Hcurve *HcurveTransactor) ExchangeUnderlying(opts *bind.TransactOpts, handler common.Address, tokenI common.Address, tokenJ common.Address, i *big.Int, j *big.Int, dx *big.Int, minDy *big.Int) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "exchangeUnderlying", handler, tokenI, tokenJ, i, j, dx, minDy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xfef6074e.
//
// Solidity: function exchangeUnderlying(address handler, address tokenI, address tokenJ, int128 i, int128 j, uint256 dx, uint256 minDy) payable returns()
func (_Hcurve *HcurveSession) ExchangeUnderlying(handler common.Address, tokenI common.Address, tokenJ common.Address, i *big.Int, j *big.Int, dx *big.Int, minDy *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.ExchangeUnderlying(&_Hcurve.TransactOpts, handler, tokenI, tokenJ, i, j, dx, minDy)
}

// ExchangeUnderlying is a paid mutator transaction binding the contract method 0xfef6074e.
//
// Solidity: function exchangeUnderlying(address handler, address tokenI, address tokenJ, int128 i, int128 j, uint256 dx, uint256 minDy) payable returns()
func (_Hcurve *HcurveTransactorSession) ExchangeUnderlying(handler common.Address, tokenI common.Address, tokenJ common.Address, i *big.Int, j *big.Int, dx *big.Int, minDy *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.ExchangeUnderlying(&_Hcurve.TransactOpts, handler, tokenI, tokenJ, i, j, dx, minDy)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hcurve *HcurveTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hcurve *HcurveSession) PostProcess() (*types.Transaction, error) {
	return _Hcurve.Contract.PostProcess(&_Hcurve.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hcurve *HcurveTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Hcurve.Contract.PostProcess(&_Hcurve.TransactOpts)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x301ee996.
//
// Solidity: function removeLiquidityOneCoin(address handler, address pool, address tokenI, uint256 tokenAmount, int128 i, uint256 minAmount) payable returns()
func (_Hcurve *HcurveTransactor) RemoveLiquidityOneCoin(opts *bind.TransactOpts, handler common.Address, pool common.Address, tokenI common.Address, tokenAmount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "removeLiquidityOneCoin", handler, pool, tokenI, tokenAmount, i, minAmount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x301ee996.
//
// Solidity: function removeLiquidityOneCoin(address handler, address pool, address tokenI, uint256 tokenAmount, int128 i, uint256 minAmount) payable returns()
func (_Hcurve *HcurveSession) RemoveLiquidityOneCoin(handler common.Address, pool common.Address, tokenI common.Address, tokenAmount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.RemoveLiquidityOneCoin(&_Hcurve.TransactOpts, handler, pool, tokenI, tokenAmount, i, minAmount)
}

// RemoveLiquidityOneCoin is a paid mutator transaction binding the contract method 0x301ee996.
//
// Solidity: function removeLiquidityOneCoin(address handler, address pool, address tokenI, uint256 tokenAmount, int128 i, uint256 minAmount) payable returns()
func (_Hcurve *HcurveTransactorSession) RemoveLiquidityOneCoin(handler common.Address, pool common.Address, tokenI common.Address, tokenAmount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.RemoveLiquidityOneCoin(&_Hcurve.TransactOpts, handler, pool, tokenI, tokenAmount, i, minAmount)
}

// RemoveLiquidityOneCoinDust is a paid mutator transaction binding the contract method 0xe3878667.
//
// Solidity: function removeLiquidityOneCoinDust(address handler, address pool, address tokenI, uint256 tokenAmount, int128 i, uint256 minAmount) payable returns()
func (_Hcurve *HcurveTransactor) RemoveLiquidityOneCoinDust(opts *bind.TransactOpts, handler common.Address, pool common.Address, tokenI common.Address, tokenAmount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "removeLiquidityOneCoinDust", handler, pool, tokenI, tokenAmount, i, minAmount)
}

// RemoveLiquidityOneCoinDust is a paid mutator transaction binding the contract method 0xe3878667.
//
// Solidity: function removeLiquidityOneCoinDust(address handler, address pool, address tokenI, uint256 tokenAmount, int128 i, uint256 minAmount) payable returns()
func (_Hcurve *HcurveSession) RemoveLiquidityOneCoinDust(handler common.Address, pool common.Address, tokenI common.Address, tokenAmount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.RemoveLiquidityOneCoinDust(&_Hcurve.TransactOpts, handler, pool, tokenI, tokenAmount, i, minAmount)
}

// RemoveLiquidityOneCoinDust is a paid mutator transaction binding the contract method 0xe3878667.
//
// Solidity: function removeLiquidityOneCoinDust(address handler, address pool, address tokenI, uint256 tokenAmount, int128 i, uint256 minAmount) payable returns()
func (_Hcurve *HcurveTransactorSession) RemoveLiquidityOneCoinDust(handler common.Address, pool common.Address, tokenI common.Address, tokenAmount *big.Int, i *big.Int, minAmount *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.RemoveLiquidityOneCoinDust(&_Hcurve.TransactOpts, handler, pool, tokenI, tokenAmount, i, minAmount)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Hcurve *HcurveTransactor) Swap(opts *bind.TransactOpts, fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Hcurve.contract.Transact(opts, "swap", fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Hcurve *HcurveSession) Swap(fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.Swap(&_Hcurve.TransactOpts, fromToken, toToken, amount, minReturn, distribution, featureFlags)
}

// Swap is a paid mutator transaction binding the contract method 0xe2a7515e.
//
// Solidity: function swap(address fromToken, address toToken, uint256 amount, uint256 minReturn, uint256[] distribution, uint256 featureFlags) payable returns()
func (_Hcurve *HcurveTransactorSession) Swap(fromToken common.Address, toToken common.Address, amount *big.Int, minReturn *big.Int, distribution []*big.Int, featureFlags *big.Int) (*types.Transaction, error) {
	return _Hcurve.Contract.Swap(&_Hcurve.TransactOpts, fromToken, toToken, amount, minReturn, distribution, featureFlags)
}
