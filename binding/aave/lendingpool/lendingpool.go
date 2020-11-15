// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lendingpool

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

// LendingpoolABI is the input ABI used to generate the binding from.
const LendingpoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_originationFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"_referral\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_totalFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_accruedBorrowInterest\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_feeLiquidated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_liquidatedCollateralForFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"OriginationFeeLiquidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"RedeemUnderlying\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amountMinusFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_borrowBalanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"LENDINGPOOL_REVISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UINT_MAX_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"addressesProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"core\",\"outputs\":[{\"internalType\":\"contractLendingPoolCore\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolDataProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parametersProvider\",\"outputs\":[{\"internalType\":\"contractLendingPoolParametersProvider\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractLendingPoolAddressesProvider\",\"name\":\"_addressesProvider\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_aTokenBalanceAfterRedeem\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"_referralCode\",\"type\":\"uint16\"}],\"name\":\"borrow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_collateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_purchaseAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsStable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsVariable\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"utilizationRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalLiquidityETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalFeesETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalBorrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originationFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Lendingpool is an auto generated Go binding around an Ethereum contract.
type Lendingpool struct {
	LendingpoolCaller     // Read-only binding to the contract
	LendingpoolTransactor // Write-only binding to the contract
	LendingpoolFilterer   // Log filterer for contract events
}

// LendingpoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingpoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingpoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingpoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingpoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingpoolSession struct {
	Contract     *Lendingpool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingpoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingpoolCallerSession struct {
	Contract *LendingpoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// LendingpoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingpoolTransactorSession struct {
	Contract     *LendingpoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LendingpoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingpoolRaw struct {
	Contract *Lendingpool // Generic contract binding to access the raw methods on
}

// LendingpoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingpoolCallerRaw struct {
	Contract *LendingpoolCaller // Generic read-only contract binding to access the raw methods on
}

// LendingpoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingpoolTransactorRaw struct {
	Contract *LendingpoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLendingpool creates a new instance of Lendingpool, bound to a specific deployed contract.
func NewLendingpool(address common.Address, backend bind.ContractBackend) (*Lendingpool, error) {
	contract, err := bindLendingpool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lendingpool{LendingpoolCaller: LendingpoolCaller{contract: contract}, LendingpoolTransactor: LendingpoolTransactor{contract: contract}, LendingpoolFilterer: LendingpoolFilterer{contract: contract}}, nil
}

// NewLendingpoolCaller creates a new read-only instance of Lendingpool, bound to a specific deployed contract.
func NewLendingpoolCaller(address common.Address, caller bind.ContractCaller) (*LendingpoolCaller, error) {
	contract, err := bindLendingpool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingpoolCaller{contract: contract}, nil
}

// NewLendingpoolTransactor creates a new write-only instance of Lendingpool, bound to a specific deployed contract.
func NewLendingpoolTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingpoolTransactor, error) {
	contract, err := bindLendingpool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingpoolTransactor{contract: contract}, nil
}

// NewLendingpoolFilterer creates a new log filterer instance of Lendingpool, bound to a specific deployed contract.
func NewLendingpoolFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingpoolFilterer, error) {
	contract, err := bindLendingpool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingpoolFilterer{contract: contract}, nil
}

// bindLendingpool binds a generic wrapper to an already deployed contract.
func bindLendingpool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingpoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lendingpool *LendingpoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lendingpool.Contract.LendingpoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lendingpool *LendingpoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lendingpool.Contract.LendingpoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lendingpool *LendingpoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lendingpool.Contract.LendingpoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lendingpool *LendingpoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lendingpool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lendingpool *LendingpoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lendingpool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lendingpool *LendingpoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lendingpool.Contract.contract.Transact(opts, method, params...)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_Lendingpool *LendingpoolCaller) LENDINGPOOLREVISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "LENDINGPOOL_REVISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_Lendingpool *LendingpoolSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _Lendingpool.Contract.LENDINGPOOLREVISION(&_Lendingpool.CallOpts)
}

// LENDINGPOOLREVISION is a free data retrieval call binding the contract method 0x8afaff02.
//
// Solidity: function LENDINGPOOL_REVISION() view returns(uint256)
func (_Lendingpool *LendingpoolCallerSession) LENDINGPOOLREVISION() (*big.Int, error) {
	return _Lendingpool.Contract.LENDINGPOOLREVISION(&_Lendingpool.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Lendingpool *LendingpoolCaller) UINTMAXVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "UINT_MAX_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Lendingpool *LendingpoolSession) UINTMAXVALUE() (*big.Int, error) {
	return _Lendingpool.Contract.UINTMAXVALUE(&_Lendingpool.CallOpts)
}

// UINTMAXVALUE is a free data retrieval call binding the contract method 0xd0fc81d2.
//
// Solidity: function UINT_MAX_VALUE() view returns(uint256)
func (_Lendingpool *LendingpoolCallerSession) UINTMAXVALUE() (*big.Int, error) {
	return _Lendingpool.Contract.UINTMAXVALUE(&_Lendingpool.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Lendingpool *LendingpoolCaller) AddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "addressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Lendingpool *LendingpoolSession) AddressesProvider() (common.Address, error) {
	return _Lendingpool.Contract.AddressesProvider(&_Lendingpool.CallOpts)
}

// AddressesProvider is a free data retrieval call binding the contract method 0xc72c4d10.
//
// Solidity: function addressesProvider() view returns(address)
func (_Lendingpool *LendingpoolCallerSession) AddressesProvider() (common.Address, error) {
	return _Lendingpool.Contract.AddressesProvider(&_Lendingpool.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Lendingpool *LendingpoolCaller) Core(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "core")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Lendingpool *LendingpoolSession) Core() (common.Address, error) {
	return _Lendingpool.Contract.Core(&_Lendingpool.CallOpts)
}

// Core is a free data retrieval call binding the contract method 0xf2f4eb26.
//
// Solidity: function core() view returns(address)
func (_Lendingpool *LendingpoolCallerSession) Core() (common.Address, error) {
	return _Lendingpool.Contract.Core(&_Lendingpool.CallOpts)
}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_Lendingpool *LendingpoolCaller) DataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "dataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_Lendingpool *LendingpoolSession) DataProvider() (common.Address, error) {
	return _Lendingpool.Contract.DataProvider(&_Lendingpool.CallOpts)
}

// DataProvider is a free data retrieval call binding the contract method 0xb334ed86.
//
// Solidity: function dataProvider() view returns(address)
func (_Lendingpool *LendingpoolCallerSession) DataProvider() (common.Address, error) {
	return _Lendingpool.Contract.DataProvider(&_Lendingpool.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_Lendingpool *LendingpoolCaller) GetReserveConfigurationData(opts *bind.CallOpts, _reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "getReserveConfigurationData", _reserve)

	outstruct := new(struct {
		Ltv                         *big.Int
		LiquidationThreshold        *big.Int
		LiquidationBonus            *big.Int
		InterestRateStrategyAddress common.Address
		UsageAsCollateralEnabled    bool
		BorrowingEnabled            bool
		StableBorrowRateEnabled     bool
		IsActive                    bool
	})

	outstruct.Ltv = out[0].(*big.Int)
	outstruct.LiquidationThreshold = out[1].(*big.Int)
	outstruct.LiquidationBonus = out[2].(*big.Int)
	outstruct.InterestRateStrategyAddress = out[3].(common.Address)
	outstruct.UsageAsCollateralEnabled = out[4].(bool)
	outstruct.BorrowingEnabled = out[5].(bool)
	outstruct.StableBorrowRateEnabled = out[6].(bool)
	outstruct.IsActive = out[7].(bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_Lendingpool *LendingpoolSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	return _Lendingpool.Contract.GetReserveConfigurationData(&_Lendingpool.CallOpts, _reserve)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address _reserve) view returns(uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, address interestRateStrategyAddress, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive)
func (_Lendingpool *LendingpoolCallerSession) GetReserveConfigurationData(_reserve common.Address) (struct {
	Ltv                         *big.Int
	LiquidationThreshold        *big.Int
	LiquidationBonus            *big.Int
	InterestRateStrategyAddress common.Address
	UsageAsCollateralEnabled    bool
	BorrowingEnabled            bool
	StableBorrowRateEnabled     bool
	IsActive                    bool
}, error) {
	return _Lendingpool.Contract.GetReserveConfigurationData(&_Lendingpool.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_Lendingpool *LendingpoolCaller) GetReserveData(opts *bind.CallOpts, _reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "getReserveData", _reserve)

	outstruct := new(struct {
		TotalLiquidity          *big.Int
		AvailableLiquidity      *big.Int
		TotalBorrowsStable      *big.Int
		TotalBorrowsVariable    *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		UtilizationRate         *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		ATokenAddress           common.Address
		LastUpdateTimestamp     *big.Int
	})

	outstruct.TotalLiquidity = out[0].(*big.Int)
	outstruct.AvailableLiquidity = out[1].(*big.Int)
	outstruct.TotalBorrowsStable = out[2].(*big.Int)
	outstruct.TotalBorrowsVariable = out[3].(*big.Int)
	outstruct.LiquidityRate = out[4].(*big.Int)
	outstruct.VariableBorrowRate = out[5].(*big.Int)
	outstruct.StableBorrowRate = out[6].(*big.Int)
	outstruct.AverageStableBorrowRate = out[7].(*big.Int)
	outstruct.UtilizationRate = out[8].(*big.Int)
	outstruct.LiquidityIndex = out[9].(*big.Int)
	outstruct.VariableBorrowIndex = out[10].(*big.Int)
	outstruct.ATokenAddress = out[11].(common.Address)
	outstruct.LastUpdateTimestamp = out[12].(*big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_Lendingpool *LendingpoolSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	return _Lendingpool.Contract.GetReserveData(&_Lendingpool.CallOpts, _reserve)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address _reserve) view returns(uint256 totalLiquidity, uint256 availableLiquidity, uint256 totalBorrowsStable, uint256 totalBorrowsVariable, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 utilizationRate, uint256 liquidityIndex, uint256 variableBorrowIndex, address aTokenAddress, uint40 lastUpdateTimestamp)
func (_Lendingpool *LendingpoolCallerSession) GetReserveData(_reserve common.Address) (struct {
	TotalLiquidity          *big.Int
	AvailableLiquidity      *big.Int
	TotalBorrowsStable      *big.Int
	TotalBorrowsVariable    *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	UtilizationRate         *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	ATokenAddress           common.Address
	LastUpdateTimestamp     *big.Int
}, error) {
	return _Lendingpool.Contract.GetReserveData(&_Lendingpool.CallOpts, _reserve)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_Lendingpool *LendingpoolCaller) GetReserves(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_Lendingpool *LendingpoolSession) GetReserves() ([]common.Address, error) {
	return _Lendingpool.Contract.GetReserves(&_Lendingpool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(address[])
func (_Lendingpool *LendingpoolCallerSession) GetReserves() ([]common.Address, error) {
	return _Lendingpool.Contract.GetReserves(&_Lendingpool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Lendingpool *LendingpoolCaller) GetUserAccountData(opts *bind.CallOpts, _user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "getUserAccountData", _user)

	outstruct := new(struct {
		TotalLiquidityETH           *big.Int
		TotalCollateralETH          *big.Int
		TotalBorrowsETH             *big.Int
		TotalFeesETH                *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})

	outstruct.TotalLiquidityETH = out[0].(*big.Int)
	outstruct.TotalCollateralETH = out[1].(*big.Int)
	outstruct.TotalBorrowsETH = out[2].(*big.Int)
	outstruct.TotalFeesETH = out[3].(*big.Int)
	outstruct.AvailableBorrowsETH = out[4].(*big.Int)
	outstruct.CurrentLiquidationThreshold = out[5].(*big.Int)
	outstruct.Ltv = out[6].(*big.Int)
	outstruct.HealthFactor = out[7].(*big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Lendingpool *LendingpoolSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _Lendingpool.Contract.GetUserAccountData(&_Lendingpool.CallOpts, _user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address _user) view returns(uint256 totalLiquidityETH, uint256 totalCollateralETH, uint256 totalBorrowsETH, uint256 totalFeesETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_Lendingpool *LendingpoolCallerSession) GetUserAccountData(_user common.Address) (struct {
	TotalLiquidityETH           *big.Int
	TotalCollateralETH          *big.Int
	TotalBorrowsETH             *big.Int
	TotalFeesETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _Lendingpool.Contract.GetUserAccountData(&_Lendingpool.CallOpts, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_Lendingpool *LendingpoolCaller) GetUserReserveData(opts *bind.CallOpts, _reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "getUserReserveData", _reserve, _user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentBorrowBalance     *big.Int
		PrincipalBorrowBalance   *big.Int
		BorrowRateMode           *big.Int
		BorrowRate               *big.Int
		LiquidityRate            *big.Int
		OriginationFee           *big.Int
		VariableBorrowIndex      *big.Int
		LastUpdateTimestamp      *big.Int
		UsageAsCollateralEnabled bool
	})

	outstruct.CurrentATokenBalance = out[0].(*big.Int)
	outstruct.CurrentBorrowBalance = out[1].(*big.Int)
	outstruct.PrincipalBorrowBalance = out[2].(*big.Int)
	outstruct.BorrowRateMode = out[3].(*big.Int)
	outstruct.BorrowRate = out[4].(*big.Int)
	outstruct.LiquidityRate = out[5].(*big.Int)
	outstruct.OriginationFee = out[6].(*big.Int)
	outstruct.VariableBorrowIndex = out[7].(*big.Int)
	outstruct.LastUpdateTimestamp = out[8].(*big.Int)
	outstruct.UsageAsCollateralEnabled = out[9].(bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_Lendingpool *LendingpoolSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _Lendingpool.Contract.GetUserReserveData(&_Lendingpool.CallOpts, _reserve, _user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address _reserve, address _user) view returns(uint256 currentATokenBalance, uint256 currentBorrowBalance, uint256 principalBorrowBalance, uint256 borrowRateMode, uint256 borrowRate, uint256 liquidityRate, uint256 originationFee, uint256 variableBorrowIndex, uint256 lastUpdateTimestamp, bool usageAsCollateralEnabled)
func (_Lendingpool *LendingpoolCallerSession) GetUserReserveData(_reserve common.Address, _user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentBorrowBalance     *big.Int
	PrincipalBorrowBalance   *big.Int
	BorrowRateMode           *big.Int
	BorrowRate               *big.Int
	LiquidityRate            *big.Int
	OriginationFee           *big.Int
	VariableBorrowIndex      *big.Int
	LastUpdateTimestamp      *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _Lendingpool.Contract.GetUserReserveData(&_Lendingpool.CallOpts, _reserve, _user)
}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_Lendingpool *LendingpoolCaller) ParametersProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lendingpool.contract.Call(opts, &out, "parametersProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_Lendingpool *LendingpoolSession) ParametersProvider() (common.Address, error) {
	return _Lendingpool.Contract.ParametersProvider(&_Lendingpool.CallOpts)
}

// ParametersProvider is a free data retrieval call binding the contract method 0x58707e06.
//
// Solidity: function parametersProvider() view returns(address)
func (_Lendingpool *LendingpoolCallerSession) ParametersProvider() (common.Address, error) {
	return _Lendingpool.Contract.ParametersProvider(&_Lendingpool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_Lendingpool *LendingpoolTransactor) Borrow(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "borrow", _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_Lendingpool *LendingpoolSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Lendingpool.Contract.Borrow(&_Lendingpool.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Borrow is a paid mutator transaction binding the contract method 0xc858f5f9.
//
// Solidity: function borrow(address _reserve, uint256 _amount, uint256 _interestRateMode, uint16 _referralCode) returns()
func (_Lendingpool *LendingpoolTransactorSession) Borrow(_reserve common.Address, _amount *big.Int, _interestRateMode *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Lendingpool.Contract.Borrow(&_Lendingpool.TransactOpts, _reserve, _amount, _interestRateMode, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_Lendingpool *LendingpoolTransactor) Deposit(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "deposit", _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_Lendingpool *LendingpoolSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Lendingpool.Contract.Deposit(&_Lendingpool.TransactOpts, _reserve, _amount, _referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xd2d0e066.
//
// Solidity: function deposit(address _reserve, uint256 _amount, uint16 _referralCode) payable returns()
func (_Lendingpool *LendingpoolTransactorSession) Deposit(_reserve common.Address, _amount *big.Int, _referralCode uint16) (*types.Transaction, error) {
	return _Lendingpool.Contract.Deposit(&_Lendingpool.TransactOpts, _reserve, _amount, _referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_Lendingpool *LendingpoolTransactor) FlashLoan(opts *bind.TransactOpts, _receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "flashLoan", _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_Lendingpool *LendingpoolSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Lendingpool.Contract.FlashLoan(&_Lendingpool.TransactOpts, _receiver, _reserve, _amount, _params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address _receiver, address _reserve, uint256 _amount, bytes _params) returns()
func (_Lendingpool *LendingpoolTransactorSession) FlashLoan(_receiver common.Address, _reserve common.Address, _amount *big.Int, _params []byte) (*types.Transaction, error) {
	return _Lendingpool.Contract.FlashLoan(&_Lendingpool.TransactOpts, _receiver, _reserve, _amount, _params)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_Lendingpool *LendingpoolTransactor) Initialize(opts *bind.TransactOpts, _addressesProvider common.Address) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "initialize", _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_Lendingpool *LendingpoolSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.Initialize(&_Lendingpool.TransactOpts, _addressesProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _addressesProvider) returns()
func (_Lendingpool *LendingpoolTransactorSession) Initialize(_addressesProvider common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.Initialize(&_Lendingpool.TransactOpts, _addressesProvider)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_Lendingpool *LendingpoolTransactor) LiquidationCall(opts *bind.TransactOpts, _collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "liquidationCall", _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_Lendingpool *LendingpoolSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _Lendingpool.Contract.LiquidationCall(&_Lendingpool.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address _collateral, address _reserve, address _user, uint256 _purchaseAmount, bool _receiveAToken) payable returns()
func (_Lendingpool *LendingpoolTransactorSession) LiquidationCall(_collateral common.Address, _reserve common.Address, _user common.Address, _purchaseAmount *big.Int, _receiveAToken bool) (*types.Transaction, error) {
	return _Lendingpool.Contract.LiquidationCall(&_Lendingpool.TransactOpts, _collateral, _reserve, _user, _purchaseAmount, _receiveAToken)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_Lendingpool *LendingpoolTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, _reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "rebalanceStableBorrowRate", _reserve, _user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_Lendingpool *LendingpoolSession) RebalanceStableBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.RebalanceStableBorrowRate(&_Lendingpool.TransactOpts, _reserve, _user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address _reserve, address _user) returns()
func (_Lendingpool *LendingpoolTransactorSession) RebalanceStableBorrowRate(_reserve common.Address, _user common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.RebalanceStableBorrowRate(&_Lendingpool.TransactOpts, _reserve, _user)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_Lendingpool *LendingpoolTransactor) RedeemUnderlying(opts *bind.TransactOpts, _reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "redeemUnderlying", _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_Lendingpool *LendingpoolSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _Lendingpool.Contract.RedeemUnderlying(&_Lendingpool.TransactOpts, _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x9895e3d8.
//
// Solidity: function redeemUnderlying(address _reserve, address _user, uint256 _amount, uint256 _aTokenBalanceAfterRedeem) returns()
func (_Lendingpool *LendingpoolTransactorSession) RedeemUnderlying(_reserve common.Address, _user common.Address, _amount *big.Int, _aTokenBalanceAfterRedeem *big.Int) (*types.Transaction, error) {
	return _Lendingpool.Contract.RedeemUnderlying(&_Lendingpool.TransactOpts, _reserve, _user, _amount, _aTokenBalanceAfterRedeem)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_Lendingpool *LendingpoolTransactor) Repay(opts *bind.TransactOpts, _reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "repay", _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_Lendingpool *LendingpoolSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.Repay(&_Lendingpool.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x5ceae9c4.
//
// Solidity: function repay(address _reserve, uint256 _amount, address _onBehalfOf) payable returns()
func (_Lendingpool *LendingpoolTransactorSession) Repay(_reserve common.Address, _amount *big.Int, _onBehalfOf common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.Repay(&_Lendingpool.TransactOpts, _reserve, _amount, _onBehalfOf)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_Lendingpool *LendingpoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, _reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "setUserUseReserveAsCollateral", _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_Lendingpool *LendingpoolSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Lendingpool.Contract.SetUserUseReserveAsCollateral(&_Lendingpool.TransactOpts, _reserve, _useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address _reserve, bool _useAsCollateral) returns()
func (_Lendingpool *LendingpoolTransactorSession) SetUserUseReserveAsCollateral(_reserve common.Address, _useAsCollateral bool) (*types.Transaction, error) {
	return _Lendingpool.Contract.SetUserUseReserveAsCollateral(&_Lendingpool.TransactOpts, _reserve, _useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_Lendingpool *LendingpoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, _reserve common.Address) (*types.Transaction, error) {
	return _Lendingpool.contract.Transact(opts, "swapBorrowRateMode", _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_Lendingpool *LendingpoolSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.SwapBorrowRateMode(&_Lendingpool.TransactOpts, _reserve)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x48ca1300.
//
// Solidity: function swapBorrowRateMode(address _reserve) returns()
func (_Lendingpool *LendingpoolTransactorSession) SwapBorrowRateMode(_reserve common.Address) (*types.Transaction, error) {
	return _Lendingpool.Contract.SwapBorrowRateMode(&_Lendingpool.TransactOpts, _reserve)
}

// LendingpoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Lendingpool contract.
type LendingpoolBorrowIterator struct {
	Event *LendingpoolBorrow // Event containing the contract specifics and raw log

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
func (it *LendingpoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolBorrow)
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
		it.Event = new(LendingpoolBorrow)
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
func (it *LendingpoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolBorrow represents a Borrow event raised by the Lendingpool contract.
type LendingpoolBorrow struct {
	Reserve               common.Address
	User                  common.Address
	Amount                *big.Int
	BorrowRateMode        *big.Int
	BorrowRate            *big.Int
	OriginationFee        *big.Int
	BorrowBalanceIncrease *big.Int
	Referral              uint16
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterBorrow(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*LendingpoolBorrowIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "Borrow", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolBorrowIterator{contract: _Lendingpool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LendingpoolBorrow, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "Borrow", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolBorrow)
				if err := _Lendingpool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x1e77446728e5558aa1b7e81e0cdab9cc1b075ba893b740600c76a315c2caa553.
//
// Solidity: event Borrow(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _borrowRateMode, uint256 _borrowRate, uint256 _originationFee, uint256 _borrowBalanceIncrease, uint16 indexed _referral, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseBorrow(log types.Log) (*LendingpoolBorrow, error) {
	event := new(LendingpoolBorrow)
	if err := _Lendingpool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Lendingpool contract.
type LendingpoolDepositIterator struct {
	Event *LendingpoolDeposit // Event containing the contract specifics and raw log

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
func (it *LendingpoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolDeposit)
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
		it.Event = new(LendingpoolDeposit)
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
func (it *LendingpoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolDeposit represents a Deposit event raised by the Lendingpool contract.
type LendingpoolDeposit struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Referral  uint16
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterDeposit(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _referral []uint16) (*LendingpoolDepositIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolDepositIterator{contract: _Lendingpool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *LendingpoolDeposit, _reserve []common.Address, _user []common.Address, _referral []uint16) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	var _referralRule []interface{}
	for _, _referralItem := range _referral {
		_referralRule = append(_referralRule, _referralItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "Deposit", _reserveRule, _userRule, _referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolDeposit)
				if err := _Lendingpool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xc12c57b1c73a2c3a2ea4613e9476abb3d8d146857aab7329e24243fb59710c82.
//
// Solidity: event Deposit(address indexed _reserve, address indexed _user, uint256 _amount, uint16 indexed _referral, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseDeposit(log types.Log) (*LendingpoolDeposit, error) {
	event := new(LendingpoolDeposit)
	if err := _Lendingpool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Lendingpool contract.
type LendingpoolFlashLoanIterator struct {
	Event *LendingpoolFlashLoan // Event containing the contract specifics and raw log

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
func (it *LendingpoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolFlashLoan)
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
		it.Event = new(LendingpoolFlashLoan)
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
func (it *LendingpoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolFlashLoan represents a FlashLoan event raised by the Lendingpool contract.
type LendingpoolFlashLoan struct {
	Target      common.Address
	Reserve     common.Address
	Amount      *big.Int
	TotalFee    *big.Int
	ProtocolFee *big.Int
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, _target []common.Address, _reserve []common.Address) (*LendingpoolFlashLoanIterator, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolFlashLoanIterator{contract: _Lendingpool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *LendingpoolFlashLoan, _target []common.Address, _reserve []common.Address) (event.Subscription, error) {

	var _targetRule []interface{}
	for _, _targetItem := range _target {
		_targetRule = append(_targetRule, _targetItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "FlashLoan", _targetRule, _reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolFlashLoan)
				if err := _Lendingpool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x5b8f46461c1dd69fb968f1a003acee221ea3e19540e350233b612ddb43433b55.
//
// Solidity: event FlashLoan(address indexed _target, address indexed _reserve, uint256 _amount, uint256 _totalFee, uint256 _protocolFee, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseFlashLoan(log types.Log) (*LendingpoolFlashLoan, error) {
	event := new(LendingpoolFlashLoan)
	if err := _Lendingpool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the Lendingpool contract.
type LendingpoolLiquidationCallIterator struct {
	Event *LendingpoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *LendingpoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolLiquidationCall)
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
		it.Event = new(LendingpoolLiquidationCall)
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
func (it *LendingpoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolLiquidationCall represents a LiquidationCall event raised by the Lendingpool contract.
type LendingpoolLiquidationCall struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	PurchaseAmount             *big.Int
	LiquidatedCollateralAmount *big.Int
	AccruedBorrowInterest      *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*LendingpoolLiquidationCallIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolLiquidationCallIterator{contract: _Lendingpool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *LendingpoolLiquidationCall, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "LiquidationCall", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolLiquidationCall)
				if err := _Lendingpool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0x56864757fd5b1fc9f38f5f3a981cd8ae512ce41b902cf73fc506ee369c6bc237.
//
// Solidity: event LiquidationCall(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _purchaseAmount, uint256 _liquidatedCollateralAmount, uint256 _accruedBorrowInterest, address _liquidator, bool _receiveAToken, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseLiquidationCall(log types.Log) (*LendingpoolLiquidationCall, error) {
	event := new(LendingpoolLiquidationCall)
	if err := _Lendingpool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolOriginationFeeLiquidatedIterator is returned from FilterOriginationFeeLiquidated and is used to iterate over the raw logs and unpacked data for OriginationFeeLiquidated events raised by the Lendingpool contract.
type LendingpoolOriginationFeeLiquidatedIterator struct {
	Event *LendingpoolOriginationFeeLiquidated // Event containing the contract specifics and raw log

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
func (it *LendingpoolOriginationFeeLiquidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolOriginationFeeLiquidated)
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
		it.Event = new(LendingpoolOriginationFeeLiquidated)
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
func (it *LendingpoolOriginationFeeLiquidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolOriginationFeeLiquidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolOriginationFeeLiquidated represents a OriginationFeeLiquidated event raised by the Lendingpool contract.
type LendingpoolOriginationFeeLiquidated struct {
	Collateral                 common.Address
	Reserve                    common.Address
	User                       common.Address
	FeeLiquidated              *big.Int
	LiquidatedCollateralForFee *big.Int
	Timestamp                  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterOriginationFeeLiquidated is a free log retrieval operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterOriginationFeeLiquidated(opts *bind.FilterOpts, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (*LendingpoolOriginationFeeLiquidatedIterator, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolOriginationFeeLiquidatedIterator{contract: _Lendingpool.contract, event: "OriginationFeeLiquidated", logs: logs, sub: sub}, nil
}

// WatchOriginationFeeLiquidated is a free log subscription operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchOriginationFeeLiquidated(opts *bind.WatchOpts, sink chan<- *LendingpoolOriginationFeeLiquidated, _collateral []common.Address, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _collateralRule []interface{}
	for _, _collateralItem := range _collateral {
		_collateralRule = append(_collateralRule, _collateralItem)
	}
	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "OriginationFeeLiquidated", _collateralRule, _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolOriginationFeeLiquidated)
				if err := _Lendingpool.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
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

// ParseOriginationFeeLiquidated is a log parse operation binding the contract event 0x36ca8b16d61dc13b1062adff83e3778ab92d14f9e35bfe9fd1283e02b13fb0a1.
//
// Solidity: event OriginationFeeLiquidated(address indexed _collateral, address indexed _reserve, address indexed _user, uint256 _feeLiquidated, uint256 _liquidatedCollateralForFee, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseOriginationFeeLiquidated(log types.Log) (*LendingpoolOriginationFeeLiquidated, error) {
	event := new(LendingpoolOriginationFeeLiquidated)
	if err := _Lendingpool.contract.UnpackLog(event, "OriginationFeeLiquidated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the Lendingpool contract.
type LendingpoolRebalanceStableBorrowRateIterator struct {
	Event *LendingpoolRebalanceStableBorrowRate // Event containing the contract specifics and raw log

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
func (it *LendingpoolRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolRebalanceStableBorrowRate)
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
		it.Event = new(LendingpoolRebalanceStableBorrowRate)
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
func (it *LendingpoolRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the Lendingpool contract.
type LendingpoolRebalanceStableBorrowRate struct {
	Reserve               common.Address
	User                  common.Address
	NewStableRate         *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*LendingpoolRebalanceStableBorrowRateIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "RebalanceStableBorrowRate", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolRebalanceStableBorrowRateIterator{contract: _Lendingpool.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *LendingpoolRebalanceStableBorrowRate, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "RebalanceStableBorrowRate", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolRebalanceStableBorrowRate)
				if err := _Lendingpool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
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

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x5050ad184862424ee0852d1838d355ad65bed1e5e6da67ac9a2dac1922677f60.
//
// Solidity: event RebalanceStableBorrowRate(address indexed _reserve, address indexed _user, uint256 _newStableRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*LendingpoolRebalanceStableBorrowRate, error) {
	event := new(LendingpoolRebalanceStableBorrowRate)
	if err := _Lendingpool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolRedeemUnderlyingIterator is returned from FilterRedeemUnderlying and is used to iterate over the raw logs and unpacked data for RedeemUnderlying events raised by the Lendingpool contract.
type LendingpoolRedeemUnderlyingIterator struct {
	Event *LendingpoolRedeemUnderlying // Event containing the contract specifics and raw log

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
func (it *LendingpoolRedeemUnderlyingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolRedeemUnderlying)
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
		it.Event = new(LendingpoolRedeemUnderlying)
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
func (it *LendingpoolRedeemUnderlyingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolRedeemUnderlyingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolRedeemUnderlying represents a RedeemUnderlying event raised by the Lendingpool contract.
type LendingpoolRedeemUnderlying struct {
	Reserve   common.Address
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedeemUnderlying is a free log retrieval operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterRedeemUnderlying(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*LendingpoolRedeemUnderlyingIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "RedeemUnderlying", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolRedeemUnderlyingIterator{contract: _Lendingpool.contract, event: "RedeemUnderlying", logs: logs, sub: sub}, nil
}

// WatchRedeemUnderlying is a free log subscription operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchRedeemUnderlying(opts *bind.WatchOpts, sink chan<- *LendingpoolRedeemUnderlying, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "RedeemUnderlying", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolRedeemUnderlying)
				if err := _Lendingpool.contract.UnpackLog(event, "RedeemUnderlying", log); err != nil {
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

// ParseRedeemUnderlying is a log parse operation binding the contract event 0x9c4ed599cd8555b9c1e8cd7643240d7d71eb76b792948c49fcb4d411f7b6b3c6.
//
// Solidity: event RedeemUnderlying(address indexed _reserve, address indexed _user, uint256 _amount, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseRedeemUnderlying(log types.Log) (*LendingpoolRedeemUnderlying, error) {
	event := new(LendingpoolRedeemUnderlying)
	if err := _Lendingpool.contract.UnpackLog(event, "RedeemUnderlying", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the Lendingpool contract.
type LendingpoolRepayIterator struct {
	Event *LendingpoolRepay // Event containing the contract specifics and raw log

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
func (it *LendingpoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolRepay)
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
		it.Event = new(LendingpoolRepay)
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
func (it *LendingpoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolRepay represents a Repay event raised by the Lendingpool contract.
type LendingpoolRepay struct {
	Reserve               common.Address
	User                  common.Address
	Repayer               common.Address
	AmountMinusFees       *big.Int
	Fees                  *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterRepay(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address, _repayer []common.Address) (*LendingpoolRepayIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _repayerRule []interface{}
	for _, _repayerItem := range _repayer {
		_repayerRule = append(_repayerRule, _repayerItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "Repay", _reserveRule, _userRule, _repayerRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolRepayIterator{contract: _Lendingpool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *LendingpoolRepay, _reserve []common.Address, _user []common.Address, _repayer []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _repayerRule []interface{}
	for _, _repayerItem := range _repayer {
		_repayerRule = append(_repayerRule, _repayerItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "Repay", _reserveRule, _userRule, _repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolRepay)
				if err := _Lendingpool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0xb718f0b14f03d8c3adf35b15e3da52421b042ac879e5a689011a8b1e0036773d.
//
// Solidity: event Repay(address indexed _reserve, address indexed _user, address indexed _repayer, uint256 _amountMinusFees, uint256 _fees, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseRepay(log types.Log) (*LendingpoolRepay, error) {
	event := new(LendingpoolRepay)
	if err := _Lendingpool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the Lendingpool contract.
type LendingpoolReserveUsedAsCollateralDisabledIterator struct {
	Event *LendingpoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *LendingpoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolReserveUsedAsCollateralDisabled)
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
		it.Event = new(LendingpoolReserveUsedAsCollateralDisabled)
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
func (it *LendingpoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the Lendingpool contract.
type LendingpoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_Lendingpool *LendingpoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*LendingpoolReserveUsedAsCollateralDisabledIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolReserveUsedAsCollateralDisabledIterator{contract: _Lendingpool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_Lendingpool *LendingpoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *LendingpoolReserveUsedAsCollateralDisabled, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolReserveUsedAsCollateralDisabled)
				if err := _Lendingpool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed _reserve, address indexed _user)
func (_Lendingpool *LendingpoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*LendingpoolReserveUsedAsCollateralDisabled, error) {
	event := new(LendingpoolReserveUsedAsCollateralDisabled)
	if err := _Lendingpool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the Lendingpool contract.
type LendingpoolReserveUsedAsCollateralEnabledIterator struct {
	Event *LendingpoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *LendingpoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolReserveUsedAsCollateralEnabled)
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
		it.Event = new(LendingpoolReserveUsedAsCollateralEnabled)
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
func (it *LendingpoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the Lendingpool contract.
type LendingpoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_Lendingpool *LendingpoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*LendingpoolReserveUsedAsCollateralEnabledIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolReserveUsedAsCollateralEnabledIterator{contract: _Lendingpool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_Lendingpool *LendingpoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *LendingpoolReserveUsedAsCollateralEnabled, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolReserveUsedAsCollateralEnabled)
				if err := _Lendingpool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed _reserve, address indexed _user)
func (_Lendingpool *LendingpoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*LendingpoolReserveUsedAsCollateralEnabled, error) {
	event := new(LendingpoolReserveUsedAsCollateralEnabled)
	if err := _Lendingpool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LendingpoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Lendingpool contract.
type LendingpoolSwapIterator struct {
	Event *LendingpoolSwap // Event containing the contract specifics and raw log

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
func (it *LendingpoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LendingpoolSwap)
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
		it.Event = new(LendingpoolSwap)
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
func (it *LendingpoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LendingpoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LendingpoolSwap represents a Swap event raised by the Lendingpool contract.
type LendingpoolSwap struct {
	Reserve               common.Address
	User                  common.Address
	NewRateMode           *big.Int
	NewRate               *big.Int
	BorrowBalanceIncrease *big.Int
	Timestamp             *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) FilterSwap(opts *bind.FilterOpts, _reserve []common.Address, _user []common.Address) (*LendingpoolSwapIterator, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.FilterLogs(opts, "Swap", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return &LendingpoolSwapIterator{contract: _Lendingpool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *LendingpoolSwap, _reserve []common.Address, _user []common.Address) (event.Subscription, error) {

	var _reserveRule []interface{}
	for _, _reserveItem := range _reserve {
		_reserveRule = append(_reserveRule, _reserveItem)
	}
	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Lendingpool.contract.WatchLogs(opts, "Swap", _reserveRule, _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LendingpoolSwap)
				if err := _Lendingpool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xb3e2773606abfd36b5bd91394b3a54d1398336c65005baf7bf7a05efeffaf75b.
//
// Solidity: event Swap(address indexed _reserve, address indexed _user, uint256 _newRateMode, uint256 _newRate, uint256 _borrowBalanceIncrease, uint256 _timestamp)
func (_Lendingpool *LendingpoolFilterer) ParseSwap(log types.Log) (*LendingpoolSwap, error) {
	event := new(LendingpoolSwap)
	if err := _Lendingpool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	return event, nil
}
