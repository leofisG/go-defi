// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hmaker

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

// HmakerABI is the input ABI used to generate the binding from.
const HmakerABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"draw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"freeETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"freeGem\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"openLockETHAndDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"ilk\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"wadC\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wadD\",\"type\":\"uint256\"}],\"name\":\"openLockGemAndDraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"postProcess\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ethJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"safeLockETH\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gemJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"safeLockGem\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"wipe\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"daiJoin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cdp\",\"type\":\"uint256\"}],\"name\":\"wipeAll\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// Hmaker is an auto generated Go binding around an Ethereum contract.
type Hmaker struct {
	HmakerCaller     // Read-only binding to the contract
	HmakerTransactor // Write-only binding to the contract
	HmakerFilterer   // Log filterer for contract events
}

// HmakerCaller is an auto generated read-only Go binding around an Ethereum contract.
type HmakerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HmakerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HmakerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HmakerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HmakerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HmakerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HmakerSession struct {
	Contract     *Hmaker           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HmakerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HmakerCallerSession struct {
	Contract *HmakerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HmakerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HmakerTransactorSession struct {
	Contract     *HmakerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HmakerRaw is an auto generated low-level Go binding around an Ethereum contract.
type HmakerRaw struct {
	Contract *Hmaker // Generic contract binding to access the raw methods on
}

// HmakerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HmakerCallerRaw struct {
	Contract *HmakerCaller // Generic read-only contract binding to access the raw methods on
}

// HmakerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HmakerTransactorRaw struct {
	Contract *HmakerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHmaker creates a new instance of Hmaker, bound to a specific deployed contract.
func NewHmaker(address common.Address, backend bind.ContractBackend) (*Hmaker, error) {
	contract, err := bindHmaker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hmaker{HmakerCaller: HmakerCaller{contract: contract}, HmakerTransactor: HmakerTransactor{contract: contract}, HmakerFilterer: HmakerFilterer{contract: contract}}, nil
}

// NewHmakerCaller creates a new read-only instance of Hmaker, bound to a specific deployed contract.
func NewHmakerCaller(address common.Address, caller bind.ContractCaller) (*HmakerCaller, error) {
	contract, err := bindHmaker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HmakerCaller{contract: contract}, nil
}

// NewHmakerTransactor creates a new write-only instance of Hmaker, bound to a specific deployed contract.
func NewHmakerTransactor(address common.Address, transactor bind.ContractTransactor) (*HmakerTransactor, error) {
	contract, err := bindHmaker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HmakerTransactor{contract: contract}, nil
}

// NewHmakerFilterer creates a new log filterer instance of Hmaker, bound to a specific deployed contract.
func NewHmakerFilterer(address common.Address, filterer bind.ContractFilterer) (*HmakerFilterer, error) {
	contract, err := bindHmaker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HmakerFilterer{contract: contract}, nil
}

// bindHmaker binds a generic wrapper to an already deployed contract.
func bindHmaker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HmakerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hmaker *HmakerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hmaker.Contract.HmakerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hmaker *HmakerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hmaker.Contract.HmakerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hmaker *HmakerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hmaker.Contract.HmakerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hmaker *HmakerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hmaker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hmaker *HmakerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hmaker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hmaker *HmakerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hmaker.Contract.contract.Transact(opts, method, params...)
}

// Draw is a paid mutator transaction binding the contract method 0xf07ab7be.
//
// Solidity: function draw(address daiJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactor) Draw(opts *bind.TransactOpts, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "draw", daiJoin, cdp, wad)
}

// Draw is a paid mutator transaction binding the contract method 0xf07ab7be.
//
// Solidity: function draw(address daiJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerSession) Draw(daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.Draw(&_Hmaker.TransactOpts, daiJoin, cdp, wad)
}

// Draw is a paid mutator transaction binding the contract method 0xf07ab7be.
//
// Solidity: function draw(address daiJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactorSession) Draw(daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.Draw(&_Hmaker.TransactOpts, daiJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x2537e4b5.
//
// Solidity: function freeETH(address ethJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactor) FreeETH(opts *bind.TransactOpts, ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "freeETH", ethJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x2537e4b5.
//
// Solidity: function freeETH(address ethJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerSession) FreeETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.FreeETH(&_Hmaker.TransactOpts, ethJoin, cdp, wad)
}

// FreeETH is a paid mutator transaction binding the contract method 0x2537e4b5.
//
// Solidity: function freeETH(address ethJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactorSession) FreeETH(ethJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.FreeETH(&_Hmaker.TransactOpts, ethJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x7031b517.
//
// Solidity: function freeGem(address gemJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactor) FreeGem(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "freeGem", gemJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x7031b517.
//
// Solidity: function freeGem(address gemJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerSession) FreeGem(gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.FreeGem(&_Hmaker.TransactOpts, gemJoin, cdp, wad)
}

// FreeGem is a paid mutator transaction binding the contract method 0x7031b517.
//
// Solidity: function freeGem(address gemJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactorSession) FreeGem(gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.FreeGem(&_Hmaker.TransactOpts, gemJoin, cdp, wad)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0x5481e4a4.
//
// Solidity: function openLockETHAndDraw(uint256 value, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_Hmaker *HmakerTransactor) OpenLockETHAndDraw(opts *bind.TransactOpts, value *big.Int, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "openLockETHAndDraw", value, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0x5481e4a4.
//
// Solidity: function openLockETHAndDraw(uint256 value, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_Hmaker *HmakerSession) OpenLockETHAndDraw(value *big.Int, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.OpenLockETHAndDraw(&_Hmaker.TransactOpts, value, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockETHAndDraw is a paid mutator transaction binding the contract method 0x5481e4a4.
//
// Solidity: function openLockETHAndDraw(uint256 value, address ethJoin, address daiJoin, bytes32 ilk, uint256 wadD) payable returns(uint256 cdp)
func (_Hmaker *HmakerTransactorSession) OpenLockETHAndDraw(value *big.Int, ethJoin common.Address, daiJoin common.Address, ilk [32]byte, wadD *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.OpenLockETHAndDraw(&_Hmaker.TransactOpts, value, ethJoin, daiJoin, ilk, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0x73af24e7.
//
// Solidity: function openLockGemAndDraw(address gemJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD) payable returns(uint256 cdp)
func (_Hmaker *HmakerTransactor) OpenLockGemAndDraw(opts *bind.TransactOpts, gemJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "openLockGemAndDraw", gemJoin, daiJoin, ilk, wadC, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0x73af24e7.
//
// Solidity: function openLockGemAndDraw(address gemJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD) payable returns(uint256 cdp)
func (_Hmaker *HmakerSession) OpenLockGemAndDraw(gemJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.OpenLockGemAndDraw(&_Hmaker.TransactOpts, gemJoin, daiJoin, ilk, wadC, wadD)
}

// OpenLockGemAndDraw is a paid mutator transaction binding the contract method 0x73af24e7.
//
// Solidity: function openLockGemAndDraw(address gemJoin, address daiJoin, bytes32 ilk, uint256 wadC, uint256 wadD) payable returns(uint256 cdp)
func (_Hmaker *HmakerTransactorSession) OpenLockGemAndDraw(gemJoin common.Address, daiJoin common.Address, ilk [32]byte, wadC *big.Int, wadD *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.OpenLockGemAndDraw(&_Hmaker.TransactOpts, gemJoin, daiJoin, ilk, wadC, wadD)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hmaker *HmakerTransactor) PostProcess(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "postProcess")
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hmaker *HmakerSession) PostProcess() (*types.Transaction, error) {
	return _Hmaker.Contract.PostProcess(&_Hmaker.TransactOpts)
}

// PostProcess is a paid mutator transaction binding the contract method 0xc2722916.
//
// Solidity: function postProcess() payable returns()
func (_Hmaker *HmakerTransactorSession) PostProcess() (*types.Transaction, error) {
	return _Hmaker.Contract.PostProcess(&_Hmaker.TransactOpts)
}

// SafeLockETH is a paid mutator transaction binding the contract method 0xf6596590.
//
// Solidity: function safeLockETH(uint256 value, address ethJoin, uint256 cdp) payable returns()
func (_Hmaker *HmakerTransactor) SafeLockETH(opts *bind.TransactOpts, value *big.Int, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "safeLockETH", value, ethJoin, cdp)
}

// SafeLockETH is a paid mutator transaction binding the contract method 0xf6596590.
//
// Solidity: function safeLockETH(uint256 value, address ethJoin, uint256 cdp) payable returns()
func (_Hmaker *HmakerSession) SafeLockETH(value *big.Int, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.SafeLockETH(&_Hmaker.TransactOpts, value, ethJoin, cdp)
}

// SafeLockETH is a paid mutator transaction binding the contract method 0xf6596590.
//
// Solidity: function safeLockETH(uint256 value, address ethJoin, uint256 cdp) payable returns()
func (_Hmaker *HmakerTransactorSession) SafeLockETH(value *big.Int, ethJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.SafeLockETH(&_Hmaker.TransactOpts, value, ethJoin, cdp)
}

// SafeLockGem is a paid mutator transaction binding the contract method 0x63d070a5.
//
// Solidity: function safeLockGem(address gemJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactor) SafeLockGem(opts *bind.TransactOpts, gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "safeLockGem", gemJoin, cdp, wad)
}

// SafeLockGem is a paid mutator transaction binding the contract method 0x63d070a5.
//
// Solidity: function safeLockGem(address gemJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerSession) SafeLockGem(gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.SafeLockGem(&_Hmaker.TransactOpts, gemJoin, cdp, wad)
}

// SafeLockGem is a paid mutator transaction binding the contract method 0x63d070a5.
//
// Solidity: function safeLockGem(address gemJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactorSession) SafeLockGem(gemJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.SafeLockGem(&_Hmaker.TransactOpts, gemJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xc3b6cb4b.
//
// Solidity: function wipe(address daiJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactor) Wipe(opts *bind.TransactOpts, daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "wipe", daiJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xc3b6cb4b.
//
// Solidity: function wipe(address daiJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerSession) Wipe(daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.Wipe(&_Hmaker.TransactOpts, daiJoin, cdp, wad)
}

// Wipe is a paid mutator transaction binding the contract method 0xc3b6cb4b.
//
// Solidity: function wipe(address daiJoin, uint256 cdp, uint256 wad) payable returns()
func (_Hmaker *HmakerTransactorSession) Wipe(daiJoin common.Address, cdp *big.Int, wad *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.Wipe(&_Hmaker.TransactOpts, daiJoin, cdp, wad)
}

// WipeAll is a paid mutator transaction binding the contract method 0x6ddb4566.
//
// Solidity: function wipeAll(address daiJoin, uint256 cdp) payable returns()
func (_Hmaker *HmakerTransactor) WipeAll(opts *bind.TransactOpts, daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Hmaker.contract.Transact(opts, "wipeAll", daiJoin, cdp)
}

// WipeAll is a paid mutator transaction binding the contract method 0x6ddb4566.
//
// Solidity: function wipeAll(address daiJoin, uint256 cdp) payable returns()
func (_Hmaker *HmakerSession) WipeAll(daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.WipeAll(&_Hmaker.TransactOpts, daiJoin, cdp)
}

// WipeAll is a paid mutator transaction binding the contract method 0x6ddb4566.
//
// Solidity: function wipeAll(address daiJoin, uint256 cdp) payable returns()
func (_Hmaker *HmakerTransactorSession) WipeAll(daiJoin common.Address, cdp *big.Int) (*types.Transaction, error) {
	return _Hmaker.Contract.WipeAll(&_Hmaker.TransactOpts, daiJoin, cdp)
}
