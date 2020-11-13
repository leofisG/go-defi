// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package yregistry

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

// YregistryABI is the input ABI used to generate the binding from.
const YregistryABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"addDelegatedVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"addWrappedVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"}],\"name\":\"getVaultInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strategy\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isDelegated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaults\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultsInfo\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"controllerArray\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokenArray\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"strategyArray\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"isWrappedArray\",\"type\":\"bool[]\"},{\"internalType\":\"bool[]\",\"name\":\"isDelegatedArray\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isDelegatedVault\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pendingGovernance\",\"type\":\"address\"}],\"name\":\"setPendingGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Yregistry is an auto generated Go binding around an Ethereum contract.
type Yregistry struct {
	YregistryCaller     // Read-only binding to the contract
	YregistryTransactor // Write-only binding to the contract
	YregistryFilterer   // Log filterer for contract events
}

// YregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type YregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YregistrySession struct {
	Contract     *Yregistry        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YregistryCallerSession struct {
	Contract *YregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// YregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YregistryTransactorSession struct {
	Contract     *YregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type YregistryRaw struct {
	Contract *Yregistry // Generic contract binding to access the raw methods on
}

// YregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YregistryCallerRaw struct {
	Contract *YregistryCaller // Generic read-only contract binding to access the raw methods on
}

// YregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YregistryTransactorRaw struct {
	Contract *YregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYregistry creates a new instance of Yregistry, bound to a specific deployed contract.
func NewYregistry(address common.Address, backend bind.ContractBackend) (*Yregistry, error) {
	contract, err := bindYregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Yregistry{YregistryCaller: YregistryCaller{contract: contract}, YregistryTransactor: YregistryTransactor{contract: contract}, YregistryFilterer: YregistryFilterer{contract: contract}}, nil
}

// NewYregistryCaller creates a new read-only instance of Yregistry, bound to a specific deployed contract.
func NewYregistryCaller(address common.Address, caller bind.ContractCaller) (*YregistryCaller, error) {
	contract, err := bindYregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YregistryCaller{contract: contract}, nil
}

// NewYregistryTransactor creates a new write-only instance of Yregistry, bound to a specific deployed contract.
func NewYregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*YregistryTransactor, error) {
	contract, err := bindYregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YregistryTransactor{contract: contract}, nil
}

// NewYregistryFilterer creates a new log filterer instance of Yregistry, bound to a specific deployed contract.
func NewYregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*YregistryFilterer, error) {
	contract, err := bindYregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YregistryFilterer{contract: contract}, nil
}

// bindYregistry binds a generic wrapper to an already deployed contract.
func bindYregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(YregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yregistry *YregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yregistry.Contract.YregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yregistry *YregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yregistry.Contract.YregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yregistry *YregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yregistry.Contract.YregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Yregistry *YregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Yregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Yregistry *YregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Yregistry *YregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Yregistry.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() pure returns(string)
func (_Yregistry *YregistryCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() pure returns(string)
func (_Yregistry *YregistrySession) GetName() (string, error) {
	return _Yregistry.Contract.GetName(&_Yregistry.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() pure returns(string)
func (_Yregistry *YregistryCallerSession) GetName() (string, error) {
	return _Yregistry.Contract.GetName(&_Yregistry.CallOpts)
}

// GetVault is a free data retrieval call binding the contract method 0x9403b634.
//
// Solidity: function getVault(uint256 index) view returns(address vault)
func (_Yregistry *YregistryCaller) GetVault(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "getVault", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVault is a free data retrieval call binding the contract method 0x9403b634.
//
// Solidity: function getVault(uint256 index) view returns(address vault)
func (_Yregistry *YregistrySession) GetVault(index *big.Int) (common.Address, error) {
	return _Yregistry.Contract.GetVault(&_Yregistry.CallOpts, index)
}

// GetVault is a free data retrieval call binding the contract method 0x9403b634.
//
// Solidity: function getVault(uint256 index) view returns(address vault)
func (_Yregistry *YregistryCallerSession) GetVault(index *big.Int) (common.Address, error) {
	return _Yregistry.Contract.GetVault(&_Yregistry.CallOpts, index)
}

// GetVaultInfo is a free data retrieval call binding the contract method 0x90229af7.
//
// Solidity: function getVaultInfo(address _vault) view returns(address controller, address token, address strategy, bool isWrapped, bool isDelegated)
func (_Yregistry *YregistryCaller) GetVaultInfo(opts *bind.CallOpts, _vault common.Address) (struct {
	Controller  common.Address
	Token       common.Address
	Strategy    common.Address
	IsWrapped   bool
	IsDelegated bool
}, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "getVaultInfo", _vault)

	outstruct := new(struct {
		Controller  common.Address
		Token       common.Address
		Strategy    common.Address
		IsWrapped   bool
		IsDelegated bool
	})

	outstruct.Controller = out[0].(common.Address)
	outstruct.Token = out[1].(common.Address)
	outstruct.Strategy = out[2].(common.Address)
	outstruct.IsWrapped = out[3].(bool)
	outstruct.IsDelegated = out[4].(bool)

	return *outstruct, err

}

// GetVaultInfo is a free data retrieval call binding the contract method 0x90229af7.
//
// Solidity: function getVaultInfo(address _vault) view returns(address controller, address token, address strategy, bool isWrapped, bool isDelegated)
func (_Yregistry *YregistrySession) GetVaultInfo(_vault common.Address) (struct {
	Controller  common.Address
	Token       common.Address
	Strategy    common.Address
	IsWrapped   bool
	IsDelegated bool
}, error) {
	return _Yregistry.Contract.GetVaultInfo(&_Yregistry.CallOpts, _vault)
}

// GetVaultInfo is a free data retrieval call binding the contract method 0x90229af7.
//
// Solidity: function getVaultInfo(address _vault) view returns(address controller, address token, address strategy, bool isWrapped, bool isDelegated)
func (_Yregistry *YregistryCallerSession) GetVaultInfo(_vault common.Address) (struct {
	Controller  common.Address
	Token       common.Address
	Strategy    common.Address
	IsWrapped   bool
	IsDelegated bool
}, error) {
	return _Yregistry.Contract.GetVaultInfo(&_Yregistry.CallOpts, _vault)
}

// GetVaults is a free data retrieval call binding the contract method 0x44d00f82.
//
// Solidity: function getVaults() view returns(address[])
func (_Yregistry *YregistryCaller) GetVaults(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "getVaults")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVaults is a free data retrieval call binding the contract method 0x44d00f82.
//
// Solidity: function getVaults() view returns(address[])
func (_Yregistry *YregistrySession) GetVaults() ([]common.Address, error) {
	return _Yregistry.Contract.GetVaults(&_Yregistry.CallOpts)
}

// GetVaults is a free data retrieval call binding the contract method 0x44d00f82.
//
// Solidity: function getVaults() view returns(address[])
func (_Yregistry *YregistryCallerSession) GetVaults() ([]common.Address, error) {
	return _Yregistry.Contract.GetVaults(&_Yregistry.CallOpts)
}

// GetVaultsInfo is a free data retrieval call binding the contract method 0x6b708788.
//
// Solidity: function getVaultsInfo() view returns(address[] controllerArray, address[] tokenArray, address[] strategyArray, bool[] isWrappedArray, bool[] isDelegatedArray)
func (_Yregistry *YregistryCaller) GetVaultsInfo(opts *bind.CallOpts) (struct {
	ControllerArray  []common.Address
	TokenArray       []common.Address
	StrategyArray    []common.Address
	IsWrappedArray   []bool
	IsDelegatedArray []bool
}, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "getVaultsInfo")

	outstruct := new(struct {
		ControllerArray  []common.Address
		TokenArray       []common.Address
		StrategyArray    []common.Address
		IsWrappedArray   []bool
		IsDelegatedArray []bool
	})

	outstruct.ControllerArray = out[0].([]common.Address)
	outstruct.TokenArray = out[1].([]common.Address)
	outstruct.StrategyArray = out[2].([]common.Address)
	outstruct.IsWrappedArray = out[3].([]bool)
	outstruct.IsDelegatedArray = out[4].([]bool)

	return *outstruct, err

}

// GetVaultsInfo is a free data retrieval call binding the contract method 0x6b708788.
//
// Solidity: function getVaultsInfo() view returns(address[] controllerArray, address[] tokenArray, address[] strategyArray, bool[] isWrappedArray, bool[] isDelegatedArray)
func (_Yregistry *YregistrySession) GetVaultsInfo() (struct {
	ControllerArray  []common.Address
	TokenArray       []common.Address
	StrategyArray    []common.Address
	IsWrappedArray   []bool
	IsDelegatedArray []bool
}, error) {
	return _Yregistry.Contract.GetVaultsInfo(&_Yregistry.CallOpts)
}

// GetVaultsInfo is a free data retrieval call binding the contract method 0x6b708788.
//
// Solidity: function getVaultsInfo() view returns(address[] controllerArray, address[] tokenArray, address[] strategyArray, bool[] isWrappedArray, bool[] isDelegatedArray)
func (_Yregistry *YregistryCallerSession) GetVaultsInfo() (struct {
	ControllerArray  []common.Address
	TokenArray       []common.Address
	StrategyArray    []common.Address
	IsWrappedArray   []bool
	IsDelegatedArray []bool
}, error) {
	return _Yregistry.Contract.GetVaultsInfo(&_Yregistry.CallOpts)
}

// GetVaultsLength is a free data retrieval call binding the contract method 0x44b19dfc.
//
// Solidity: function getVaultsLength() view returns(uint256)
func (_Yregistry *YregistryCaller) GetVaultsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "getVaultsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVaultsLength is a free data retrieval call binding the contract method 0x44b19dfc.
//
// Solidity: function getVaultsLength() view returns(uint256)
func (_Yregistry *YregistrySession) GetVaultsLength() (*big.Int, error) {
	return _Yregistry.Contract.GetVaultsLength(&_Yregistry.CallOpts)
}

// GetVaultsLength is a free data retrieval call binding the contract method 0x44b19dfc.
//
// Solidity: function getVaultsLength() view returns(uint256)
func (_Yregistry *YregistryCallerSession) GetVaultsLength() (*big.Int, error) {
	return _Yregistry.Contract.GetVaultsLength(&_Yregistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yregistry *YregistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yregistry *YregistrySession) Governance() (common.Address, error) {
	return _Yregistry.Contract.Governance(&_Yregistry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Yregistry *YregistryCallerSession) Governance() (common.Address, error) {
	return _Yregistry.Contract.Governance(&_Yregistry.CallOpts)
}

// IsDelegatedVault is a free data retrieval call binding the contract method 0x44064be7.
//
// Solidity: function isDelegatedVault(address ) view returns(bool)
func (_Yregistry *YregistryCaller) IsDelegatedVault(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "isDelegatedVault", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegatedVault is a free data retrieval call binding the contract method 0x44064be7.
//
// Solidity: function isDelegatedVault(address ) view returns(bool)
func (_Yregistry *YregistrySession) IsDelegatedVault(arg0 common.Address) (bool, error) {
	return _Yregistry.Contract.IsDelegatedVault(&_Yregistry.CallOpts, arg0)
}

// IsDelegatedVault is a free data retrieval call binding the contract method 0x44064be7.
//
// Solidity: function isDelegatedVault(address ) view returns(bool)
func (_Yregistry *YregistryCallerSession) IsDelegatedVault(arg0 common.Address) (bool, error) {
	return _Yregistry.Contract.IsDelegatedVault(&_Yregistry.CallOpts, arg0)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Yregistry *YregistryCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Yregistry.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Yregistry *YregistrySession) PendingGovernance() (common.Address, error) {
	return _Yregistry.Contract.PendingGovernance(&_Yregistry.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_Yregistry *YregistryCallerSession) PendingGovernance() (common.Address, error) {
	return _Yregistry.Contract.PendingGovernance(&_Yregistry.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yregistry *YregistryTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Yregistry.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yregistry *YregistrySession) AcceptGovernance() (*types.Transaction, error) {
	return _Yregistry.Contract.AcceptGovernance(&_Yregistry.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_Yregistry *YregistryTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _Yregistry.Contract.AcceptGovernance(&_Yregistry.TransactOpts)
}

// AddDelegatedVault is a paid mutator transaction binding the contract method 0x2019c75a.
//
// Solidity: function addDelegatedVault(address _vault) returns()
func (_Yregistry *YregistryTransactor) AddDelegatedVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Yregistry.contract.Transact(opts, "addDelegatedVault", _vault)
}

// AddDelegatedVault is a paid mutator transaction binding the contract method 0x2019c75a.
//
// Solidity: function addDelegatedVault(address _vault) returns()
func (_Yregistry *YregistrySession) AddDelegatedVault(_vault common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.AddDelegatedVault(&_Yregistry.TransactOpts, _vault)
}

// AddDelegatedVault is a paid mutator transaction binding the contract method 0x2019c75a.
//
// Solidity: function addDelegatedVault(address _vault) returns()
func (_Yregistry *YregistryTransactorSession) AddDelegatedVault(_vault common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.AddDelegatedVault(&_Yregistry.TransactOpts, _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_Yregistry *YregistryTransactor) AddVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Yregistry.contract.Transact(opts, "addVault", _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_Yregistry *YregistrySession) AddVault(_vault common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.AddVault(&_Yregistry.TransactOpts, _vault)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(address _vault) returns()
func (_Yregistry *YregistryTransactorSession) AddVault(_vault common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.AddVault(&_Yregistry.TransactOpts, _vault)
}

// AddWrappedVault is a paid mutator transaction binding the contract method 0x61de8389.
//
// Solidity: function addWrappedVault(address _vault) returns()
func (_Yregistry *YregistryTransactor) AddWrappedVault(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Yregistry.contract.Transact(opts, "addWrappedVault", _vault)
}

// AddWrappedVault is a paid mutator transaction binding the contract method 0x61de8389.
//
// Solidity: function addWrappedVault(address _vault) returns()
func (_Yregistry *YregistrySession) AddWrappedVault(_vault common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.AddWrappedVault(&_Yregistry.TransactOpts, _vault)
}

// AddWrappedVault is a paid mutator transaction binding the contract method 0x61de8389.
//
// Solidity: function addWrappedVault(address _vault) returns()
func (_Yregistry *YregistryTransactorSession) AddWrappedVault(_vault common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.AddWrappedVault(&_Yregistry.TransactOpts, _vault)
}

// SetPendingGovernance is a paid mutator transaction binding the contract method 0x0abb6035.
//
// Solidity: function setPendingGovernance(address _pendingGovernance) returns()
func (_Yregistry *YregistryTransactor) SetPendingGovernance(opts *bind.TransactOpts, _pendingGovernance common.Address) (*types.Transaction, error) {
	return _Yregistry.contract.Transact(opts, "setPendingGovernance", _pendingGovernance)
}

// SetPendingGovernance is a paid mutator transaction binding the contract method 0x0abb6035.
//
// Solidity: function setPendingGovernance(address _pendingGovernance) returns()
func (_Yregistry *YregistrySession) SetPendingGovernance(_pendingGovernance common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.SetPendingGovernance(&_Yregistry.TransactOpts, _pendingGovernance)
}

// SetPendingGovernance is a paid mutator transaction binding the contract method 0x0abb6035.
//
// Solidity: function setPendingGovernance(address _pendingGovernance) returns()
func (_Yregistry *YregistryTransactorSession) SetPendingGovernance(_pendingGovernance common.Address) (*types.Transaction, error) {
	return _Yregistry.Contract.SetPendingGovernance(&_Yregistry.TransactOpts, _pendingGovernance)
}
