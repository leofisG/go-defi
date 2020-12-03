package client

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/524119574/go-defi/binding/haave"
	"github.com/524119574/go-defi/binding/hcether"
	"github.com/524119574/go-defi/binding/hctoken"
	"github.com/524119574/go-defi/binding/hkyber"
	"github.com/524119574/go-defi/binding/huniswap"

	"github.com/524119574/go-defi/binding/herc20tokenin"
	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/524119574/go-defi/binding/aave/lendingpool"
	ceth_binding "github.com/524119574/go-defi/binding/compound/cETH"
	"github.com/524119574/go-defi/binding/compound/cToken"
	"github.com/524119574/go-defi/binding/erc20"
	"github.com/524119574/go-defi/binding/furucombo"
	"github.com/524119574/go-defi/binding/uniswap"
	"github.com/524119574/go-defi/binding/yearn/yregistry"
	"github.com/524119574/go-defi/binding/yearn/yvault"
	"github.com/524119574/go-defi/binding/yearn/yweth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type netType int

const (
	// MainNet the main Ethereum network
	MainNet netType = iota
	// TestNet the test net.
	TestNet netType = iota
)

type rateModel int64

const (
	// StableRate https://medium.com/aave/aave-borrowing-rates-upgraded-f6c8b27973a7
	StableRate rateModel = 1
	// VariableRate https://medium.com/aave/aave-borrowing-rates-upgraded-f6c8b27973a7
	VariableRate rateModel = 2
)

type coinType int

const (
	// ETH weiwu
	ETH coinType = iota
	// BAT weiwu
	BAT coinType = iota
	// COMP weiwu
	COMP coinType = iota
	// DAI weiwu
	DAI coinType = iota
	// REP weiwu
	REP coinType = iota
	// SAI weiwu
	SAI coinType = iota
	// UNI weiwu
	UNI coinType = iota
	// USDC weiwu
	USDC coinType = iota
	// USDT weiwu
	USDT coinType = iota
	// WBTC weiwu
	WBTC coinType = iota
	// ZRX weiwu
	ZRX coinType = iota
	// BUSD weiwu
	BUSD coinType = iota

	// cToken
	cETH = iota

	cDAI = iota

	cUSDC = iota
)

const (
	// uniswapAddr is UniswapV2Router, see here: https://uniswap.org/docs/v2/smart-contracts/router02/#address
	uniswapAddr             string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	yRegistryAddr           string = "0x3eE41C098f9666ed2eA246f4D2558010e59d63A0"
	yETHVaultAddr           string = "0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7"
	aaveLendingPoolAddr     string = "0x398eC7346DcD622eDc5ae82352F02bE94C62d119"
	aaveLendingPoolCoreAddr string = "0x3dfd23A6c5E8BbcFc9581d2E864a68feb6a076d3"
	// Furucombo related addresses

	// FurucomboAddr is the address of the Furucombo proxy
	FurucomboAddr string = "0x57805e5a227937bac2b0fdacaa30413ddac6b8e1"
	hCEtherAddr   string = "0x9A1049f7f87Dbb0468C745d9B3952e23d5d6CE5e"
	hErcInAddr    string = "0x914490a362f4507058403a99e28bdf685c5c767f"
	hCTokenAddr   string = "0x8973D623d883c5641Dd3906625Aac31cdC8790c5"
	hMakerDaoAddr string = "0x294fbca49c8a855e04d7d82b28256b086d39afea"
	hUniswapAddr  string = "0x58a21cfcee675d65d577b251668f7dc46ea9c3a0"
	hCurveAddr    string = "0xa36dfb057010c419c5917f3d68b4520db3671cdb"
	hYearnAddr    string = "0xC50C8F34c9955217a6b3e385a069184DCE17fD2A"
	hAaveAddr     string = "0xf579b009748a62b1978639d6b54259f8dc915229"
	hOneInch      string = "0x783f5c56e3c8b23d90e4a271d7acbe914bfcd319"
	hFunds        string = "0xf9b03e9ea64b2311b0221b2854edd6df97669c09"
	hKyberAddr    string = "0xe2a3431508cd8e72d53a0e4b57c24af2899322a0"
)

// CoinToAddressMap returns a mapping from coin to address
var CoinToAddressMap = map[coinType]common.Address{
	ETH:   common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
	BAT:   common.HexToAddress("0x0d8775f648430679a709e98d2b0cb6250d2887ef"),
	COMP:  common.HexToAddress("0xc00e94cb662c3520282e6f5717214004a7f26888"),
	DAI:   common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
	USDC:  common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	USDT:  common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
	cETH:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
	cDAI:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	cUSDC: common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563"),
	BUSD:  common.HexToAddress("0x4Fabb145d64652a948d72533023f6E7A623C7C53"),
}

// CoinToCompoundMap returns a mapping from coin to compound address
var CoinToCompoundMap = map[coinType]common.Address{
	ETH:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
	DAI:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	USDC: common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563"),
}

// Client is the new interface
type Client interface {
	Uniswap() UniswapClient
}

// NewClient Create a new client
func NewClient(opts *bind.TransactOpts, ethClient *ethclient.Client, net netType) (*ActualClient, error) {
	c := new(ActualClient)
	c.conn = ethClient
	c.opts = opts
	c.net = net
	return c, nil
}

// ActualClient is the struct that stores the information.
type ActualClient struct {
	opts *bind.TransactOpts
	conn *ethclient.Client
	net  netType
}

// BalanceOf returns the balance of a given coin.
func (c *ActualClient) BalanceOf(coin coinType) (*big.Int, error) {
	erc20, err := erc20.NewErc20(CoinToAddressMap[coin], c.conn)
	if err != nil {
		return nil, err
	}
	balance, err := erc20.BalanceOf(nil, c.opts.From)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (c *ActualClient) ExecuteActions(actions *Actions) error {
	handlers := []common.Address{}
	datas := make([][]byte, 0)
	totalEthers := big.NewInt(0)
	approvalTokens := make([]common.Address, 0)
	approvalAmounts := make([]*big.Int, 0)
	for i := 0; i < len(actions.Actions); i++ {
		handlers = append(handlers, actions.Actions[i].HandlerAddr)
		datas = append(datas, actions.Actions[i].Data)
		totalEthers.Add(totalEthers, actions.Actions[i].EthersNeeded)
		if actions.Actions[i].ApprovalTokenAmount != nil {
			approvalTokens = append(approvalTokens, actions.Actions[i].ApprovalToken)
			approvalAmounts = append(approvalAmounts, actions.Actions[i].ApprovalTokenAmount)
		}
	}

	if len(approvalTokens) > 0 {
		parsed, err := abi.JSON(strings.NewReader(herc20tokenin.Herc20tokeninABI))
		if err != nil {
			return err
		}
		injectData, err := parsed.Pack("inject", approvalTokens, approvalAmounts)

		if err != nil {
			return err
		}

		handlers = append([]common.Address{common.HexToAddress(hErcInAddr)}, handlers...)
		datas = append([][]byte{injectData}, datas...)
	}

	proxy, err := furucombo.NewFurucombo(common.HexToAddress(FurucomboAddr), c.conn)
	if err != nil {
		return nil
	}

	gasPrice, err := c.conn.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	opts := &bind.TransactOpts{
		Value:    totalEthers,
		Signer:   c.opts.Signer,
		From:     c.opts.From,
		GasLimit: 5000000,
		GasPrice: gasPrice,
	}
	tx, err := proxy.BatchExec(opts, handlers, datas)
	if err != nil {
		return nil
	}
	receipt, err := bind.WaitMined(context.Background(), c.conn, tx)
	if err != nil {
		return err
	}
	if receipt.Status != 1 {
		return fmt.Errorf("tx receipt status is not 1, indicating a failure occurred")
	}
	return nil

}

// SupplyFundActions transfer a certain amount of fund to the proxy
func (c *ActualClient) SupplyFundActions(size *big.Int, coin coinType) *Actions {
	parsed, err := abi.JSON(strings.NewReader(herc20tokenin.Herc20tokeninABI))
	if err != nil {
		return nil
	}
	injectData, err := parsed.Pack(
		"inject", []common.Address{CoinToAddressMap[coin]}, []*big.Int{size})

	if err != nil {
		return nil
	}

	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:  common.HexToAddress(hFunds),
				Data:         injectData,
				EthersNeeded: big.NewInt(0),
			},
		},
	}
}

// Action represents one action, e.g. supply to Compound, swap on Uniswap
type Action struct {
	HandlerAddr         common.Address
	Data                []byte
	EthersNeeded        *big.Int
	ApprovalToken       common.Address
	ApprovalTokenAmount *big.Int
}

type Actions struct {
	Actions []Action
}

// Add adds actions together
func (actions *Actions) Add(newActions *Actions) {
	actions.Actions = append(actions.Actions, newActions.Actions...)
}

// Uniswap---------------------------------------------------------------------

// UniswapClient struct
type UniswapClient struct {
	client  *ActualClient
	uniswap *uniswap.Uniswap
}

// Uniswap returns a uniswap client
func (c *ActualClient) Uniswap() *UniswapClient {
	uniClient := new(UniswapClient)
	uniClient.client = c
	uniswap, err := uniswap.NewUniswap(common.HexToAddress(uniswapAddr), c.conn)

	if err != nil {
		return nil
	}

	uniClient.uniswap = uniswap
	return uniClient
}

// TxHash represents a transaction hash
type TxHash string

// Swap in the Uniswap Exchange.
func (c *UniswapClient) Swap(size int64, baseCurrency coinType, quoteCurrency coinType, receipient common.Address) error {
	if quoteCurrency == ETH {
		return c.swapETHToToken(size, baseCurrency, receipient)
	} else {
		err := Approve(c.client, quoteCurrency, common.HexToAddress(uniswapAddr), big.NewInt(size))
		if err != nil {
			return err
		}
		if baseCurrency == ETH {
			return c.swapTokenToETH(size, quoteCurrency, receipient)
		} else {
			return c.swapTokenToToken(size, baseCurrency, quoteCurrency, receipient)
		}
	}
}

func (c *UniswapClient) swapETHToToken(size int64, baseCurrency coinType, receipient common.Address) error {
	path := []common.Address{CoinToAddressMap[ETH], CoinToAddressMap[baseCurrency]}
	tx, err := c.uniswap.SwapExactETHForTokens(
		// TODO: there is basically no minimum output amount set, so this could cause huge slippage, need to fix.
		// Also the time stamp is set to 2038 January 1, it's better to set it dynamically.
		&bind.TransactOpts{
			Value:    big.NewInt(size),
			Signer:   c.client.opts.Signer,
			From:     c.client.opts.From,
			GasLimit: 500000,
			GasPrice: big.NewInt(20000000000),
		},
		big.NewInt(0),
		path, receipient,
		big.NewInt(2145916800),
	)
	if err != nil {
		return err
	}
	bind.WaitMined(context.Background(), c.client.conn, tx)
	return nil
}

func (c *UniswapClient) swapTokenToToken(size int64, baseCurrency coinType, quoteCurrency coinType, receipient common.Address) error {
	path := []common.Address{CoinToAddressMap[quoteCurrency], CoinToAddressMap[ETH], CoinToAddressMap[baseCurrency]}

	tx, err := c.uniswap.SwapExactTokensForTokens(
		// TODO: there is basically no minimum output amount set, so this could cause huge slippage, need to fix.
		// Also the time stamp is set to 2038 January 1, it's better to set it dynamically.
		c.client.opts, big.NewInt(size), big.NewInt(0), path, receipient, big.NewInt(2145916800))
	if err != nil {
		return err
	}
	bind.WaitMined(context.Background(), c.client.conn, tx)
	return nil
}

func (c *UniswapClient) swapTokenToETH(size int64, quoteCurrency coinType, receipient common.Address) error {
	path := []common.Address{CoinToAddressMap[quoteCurrency], CoinToAddressMap[ETH]}

	tx, err := c.uniswap.SwapExactTokensForETH(
		// TODO: there is basically no minimum output amount set, so this could cause huge slippage, need to fix.
		// Also the time stamp is set to 2038 January 1, it's better to set it dynamically.
		c.client.opts, big.NewInt(size), big.NewInt(0), path, receipient, big.NewInt(2145916800))
	if err != nil {
		return err
	}
	bind.WaitMined(context.Background(), c.client.conn, tx)
	return nil
}

// SwapActions create a new swap action
func (c *UniswapClient) SwapActions(size *big.Int, baseCurrency coinType, quoteCurrency coinType) *Actions {
	var callData []byte
	var ethersNeeded = big.NewInt(0)
	if quoteCurrency == ETH {
		ethersNeeded = size
		callData = c.swapETHToTokenData(size, baseCurrency)
	} else {
		if baseCurrency == ETH {
			callData = c.swapTokenToETHData(size, quoteCurrency)
		} else {
			callData = c.swapTokenToTokenData(size, baseCurrency, quoteCurrency)
		}
	}

	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:  common.HexToAddress(hUniswapAddr),
				Data:         callData,
				EthersNeeded: ethersNeeded,
			},
		},
	}
}

func (c *UniswapClient) swapETHToTokenData(size *big.Int, baseCurrency coinType) []byte {
	parsed, err := abi.JSON(strings.NewReader(huniswap.HuniswapABI))
	if err != nil {
		return nil
	}
	data, err := parsed.Pack(
		"swapExactETHForTokens", size, big.NewInt(0), []common.Address{CoinToAddressMap[ETH], CoinToAddressMap[baseCurrency]})
	if err != nil {
		return nil
	}
	return data
}

func (c *UniswapClient) swapTokenToETHData(size *big.Int, quoteCurrency coinType) []byte {
	parsed, err := abi.JSON(strings.NewReader(huniswap.HuniswapABI))
	if err != nil {
		return nil
	}
	data, err := parsed.Pack(
		"swapExactTokensForETH", size, big.NewInt(0), []common.Address{CoinToAddressMap[quoteCurrency], CoinToAddressMap[ETH]})
	if err != nil {
		return nil
	}
	return data
}

func (c *UniswapClient) swapTokenToTokenData(size *big.Int, baseCurrency coinType, quoteCurrency coinType) []byte {
	parsed, err := abi.JSON(strings.NewReader(huniswap.HuniswapABI))
	if err != nil {
		return nil
	}
	data, err := parsed.Pack(
		"swapExactTokensForTokens", size, big.NewInt(0), []common.Address{
			CoinToAddressMap[quoteCurrency], CoinToAddressMap[ETH], CoinToAddressMap[baseCurrency]})
	if err != nil {
		return nil
	}
	return data
}

// Compound---------------------------------------------------------------------

// CompoundClient is an instance of Compound protocol.
type CompoundClient struct {
	client *ActualClient
}

// Compound returns a compound client
func (c *ActualClient) Compound() *CompoundClient {
	compoundClient := new(CompoundClient)
	compoundClient.client = c

	return compoundClient
}

// Supply supplies token to compound
// amoutn decimal is 1e18
func (c *CompoundClient) Supply(amount int64, coin coinType) error {
	var (
		tx  *types.Transaction
		err error
	)

	cTokenAddr, err := c.getPoolAddrFromCoin(coin)
	opts := &bind.TransactOpts{
		From:     c.client.opts.From,
		Signer:   c.client.opts.Signer,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}

	switch coin {
	case ETH:
		opts.Value = big.NewInt(amount)
		cETHContract, err := ceth_binding.NewCETH(cTokenAddr, c.client.conn)
		if err != nil {
			return err
		}

		tx, err = cETHContract.Mint(opts)
	case BAT, COMP, DAI, REP, SAI, UNI, USDC, USDT, WBTC, ZRX:
		err = Approve(c.client, coin, cTokenAddr, big.NewInt(amount))
		if err != nil {
			return err
		}
		cTokenContract, err := cToken.NewCToken(cTokenAddr, c.client.conn)
		if err != nil {
			return err
		}
		tx, err = cTokenContract.Mint(opts, big.NewInt(amount))
	default:
		return fmt.Errorf("Not supported")
	}

	if err != nil {
		fmt.Printf("Error mint ctoken: %v", err)
		return err
	}

	bind.WaitMined(context.Background(), c.client.conn, tx)

	return nil
}

// Redeem supplies token to compound
// amount decimal is 1e8
func (c *CompoundClient) Redeem(amount int64, coin coinType) error {
	var (
		tx  *types.Transaction
		err error
	)

	cTokenAddr, err := c.getPoolAddrFromCoin(coin)
	if err != nil {
		return err
	}

	opts := &bind.TransactOpts{
		From:     c.client.opts.From,
		Signer:   c.client.opts.Signer,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}

	switch coin {
	case ETH:
		cETHContract, err := ceth_binding.NewCETH(cTokenAddr, c.client.conn)
		if err != nil {
			return fmt.Errorf("Error getting cETH contract: %v", err)
		}

		tx, err = cETHContract.Redeem(opts, big.NewInt(amount))
	case BAT, COMP, DAI, REP, SAI, UNI, USDC, USDT, WBTC, ZRX:
		cTokenContract, err := cToken.NewCToken(cTokenAddr, c.client.conn)
		if err != nil {
			return fmt.Errorf("Error getting cToken contract: %v", err)
		}

		tx, err = cTokenContract.Redeem(opts, big.NewInt(amount))
	}

	if err != nil {
		return err
	}

	bind.WaitMined(context.Background(), c.client.conn, tx)

	return nil
}

// BalanceOf return the balance of given cToken
func (c *CompoundClient) BalanceOf(coin coinType) (*big.Int, error) {
	var (
		val *big.Int
		err error
	)

	cTokenAddr, err := c.getPoolAddrFromCoin(coin)
	if err != nil {
		return nil, err
	}

	switch coin {
	case ETH:
		cETHContract, err := ceth_binding.NewCETH(cTokenAddr, c.client.conn)
		if err != nil {
			return nil, fmt.Errorf("Error getting cETH contract")
		}

		val, err = cETHContract.BalanceOf(nil, c.client.opts.From)
	case BAT, COMP, DAI, REP, SAI, UNI, USDC, USDT, WBTC, ZRX:
		cTokenContract, err := cToken.NewCToken(cTokenAddr, c.client.conn)
		if err != nil {
			return nil, fmt.Errorf("Error getting cDai contract")
		}

		val, err = cTokenContract.BalanceOf(nil, c.client.opts.From)
	default:
		return nil, fmt.Errorf("Not support token in balanceOf: %v", coin)
	}

	if err != nil {
		return big.NewInt(0), fmt.Errorf("Error getting balance of cToken: %v", err)
	}
	return val, nil
}

// BalanceOfUnderlying return the balance of given cToken
// TODO: figure out why is this API so strange? Why is it different from BalanceOf?
func (c *CompoundClient) BalanceOfUnderlying(coin coinType) (*types.Transaction, error) {
	var (
		tx  *types.Transaction
		err error
	)

	cTokenAddr, err := c.getPoolAddrFromCoin(coin)
	if err != nil {
		return nil, err
	}

	switch coin {
	case ETH:
		cETHContract, err := ceth_binding.NewCETH(cTokenAddr, c.client.conn)
		if err != nil {
			fmt.Printf("Error getting cETH contract")
		}

		tx, err = cETHContract.BalanceOfUnderlying(nil, c.client.opts.From)
	case BAT, COMP, DAI, REP, SAI, UNI, USDC, USDT, WBTC, ZRX:
		cTokenContract, err := cToken.NewCToken(cTokenAddr, c.client.conn)
		if err != nil {
			return nil, fmt.Errorf("Error getting cDai contract")
		}

		tx, err = cTokenContract.BalanceOfUnderlying(nil, c.client.opts.From)
	}

	if err != nil {
		fmt.Printf("Error getting balance of cToken: %v", err)
		return nil, err
	}
	return tx, nil
}

// SupplyActions create a supply action
func (c *CompoundClient) SupplyActions(size *big.Int, coin coinType) *Actions {
	if coin == ETH {
		return c.supplyActionsETH(size, coin)
	} else {
		return c.supplyActionsERC20(size, coin)
	}
}

func (c *CompoundClient) supplyActionsETH(size *big.Int, coin coinType) *Actions {
	parsed, err := abi.JSON(strings.NewReader(hcether.HcetherABI))
	if err != nil {
		return nil
	}
	data, err := parsed.Pack("mint", size)
	if err != nil {
		return nil
	}
	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:  common.HexToAddress(hCEtherAddr),
				Data:         data,
				EthersNeeded: size,
			},
		},
	}
}

func (c *CompoundClient) supplyActionsERC20(size *big.Int, coin coinType) *Actions {
	parsed, err := abi.JSON(strings.NewReader(hctoken.HctokenABI))
	if err != nil {
		return nil
	}
	mintData, err := parsed.Pack("mint", CoinToCompoundMap[DAI], size)
	if err != nil {
		return nil
	}
	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:         common.HexToAddress(hCTokenAddr),
				Data:                mintData,
				EthersNeeded:        big.NewInt(0),
				ApprovalToken:       CoinToAddressMap[coin],
				ApprovalTokenAmount: size,
			},
		},
	}
}

// RedeemActions create a Compound redeem action to be executed
func (c *CompoundClient) RedeemActions(size *big.Int, coin coinType) *Actions {
	if coin == ETH {
		return c.redeemActionsETH(size, coin)
	} else {
		return c.redeemActionsERC20(size, coin)
	}
}

func (c *CompoundClient) redeemActionsETH(size *big.Int, coin coinType) *Actions {
	parsed, err := abi.JSON(strings.NewReader(hcether.HcetherABI))
	if err != nil {
		return nil
	}
	data, err := parsed.Pack("redeem", size)
	if err != nil {
		return nil
	}
	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:  common.HexToAddress(hCEtherAddr),
				Data:         data,
				EthersNeeded: size,
			},
		},
	}
}

func (c *CompoundClient) redeemActionsERC20(size *big.Int, coin coinType) *Actions {
	parsed, err := abi.JSON(strings.NewReader(hctoken.HctokenABI))
	if err != nil {
		return nil
	}
	redeemData, err := parsed.Pack("redeem", CoinToCompoundMap[coin], size)
	if err != nil {
		return nil
	}
	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:         common.HexToAddress(hCTokenAddr),
				Data:                redeemData,
				EthersNeeded:        big.NewInt(0),
				ApprovalToken:       CoinToCompoundMap[coin],
				ApprovalTokenAmount: size,
			},
		},
	}
}

// FlashLoanActions create an action to get flashloan
func (c *AaveClient) FlashLoanActions(size *big.Int, coin coinType, actions *Actions) *Actions {
	handlers := []common.Address{}
	datas := make([][]byte, 0)
	totalEthers := big.NewInt(0)
	for i := 0; i < len(actions.Actions); i++ {
		handlers = append(handlers, actions.Actions[i].HandlerAddr)
		datas = append(datas, actions.Actions[i].Data)
		totalEthers.Add(totalEthers, actions.Actions[i].EthersNeeded)
	}

	proxy, err := abi.JSON(strings.NewReader(furucombo.FurucomboABI))
	if err != nil {
		return nil
	}
	payloadData, err := proxy.Pack("execs", handlers, datas)
	if err != nil {
		return nil
	}
	haave, err := abi.JSON(strings.NewReader(haave.HaaveABI))
	if err != nil {
		return nil
	}
	// skip the first 4 bytes to omit the function selector
	flashLoanData, err := haave.Pack("flashLoan", CoinToAddressMap[coin], size, payloadData[4:])
	fmt.Printf("flash loan data: %v", hex.EncodeToString(flashLoanData))
	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:  common.HexToAddress(hAaveAddr),
				Data:         flashLoanData,
				EthersNeeded: totalEthers,
			},
		},
	}
}

func (c *CompoundClient) getPoolAddrFromCoin(coin coinType) (common.Address, error) {
	if val, ok := CoinToCompoundMap[coin]; ok {
		return val, nil
	}
	return common.Address{}, fmt.Errorf("No corresponding compound pool for token: %v", coin)
}

// yearn-----------------------------------------------------------------------------------------------------

// YearnClient is an instance of Compound protocol.
type YearnClient struct {
	client       *ActualClient
	tokenToVault map[common.Address]common.Address
}

// Yearn returns a Yearn client
func (c *ActualClient) Yearn() *YearnClient {
	yearnClient := new(YearnClient)
	yearnClient.client = c

	yregistry, err := yregistry.NewYregistry(common.HexToAddress(yRegistryAddr), c.conn)
	if err != nil {
		return nil
	}

	vaults, err := yregistry.GetVaults(nil)
	if err != nil {
		return nil
	}

	vaultInfos, err := yregistry.GetVaultsInfo(nil)
	if err != nil {
		return nil
	}

	yearnClient.tokenToVault = make(map[common.Address]common.Address)
	for i := 0; i < len(vaults); i++ {
		yearnClient.tokenToVault[vaultInfos.TokenArray[i]] = vaults[i]
	}

	return yearnClient
}

func (c *YearnClient) addLiquidity(size *big.Int, coin coinType) error {
	var (
		tx  *types.Transaction
		err error
	)
	opts := &bind.TransactOpts{
		From:     c.client.opts.From,
		Signer:   c.client.opts.Signer,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}

	if coin == ETH {
		weth, err := yweth.NewYweth(common.HexToAddress(yETHVaultAddr), c.client.conn)
		if err != nil {
			return fmt.Errorf("Error getting weth contract")
		}
		opts.Value = size
		tx, err = weth.DepositETH(opts)
	} else if coin != ETH {
		tokenAddr := CoinToAddressMap[coin]
		vaultAddr, ok := c.tokenToVault[tokenAddr]
		if !ok {
			return fmt.Errorf("No corresponding vault found for: %v ", coin)
		}
		err = Approve(c.client, coin, vaultAddr, size)
		yvault, err := yvault.NewYvault(vaultAddr, c.client.conn)
		if err != nil {
			return fmt.Errorf("Error getting weth contract")
		}
		opts.Value = size
		tx, err = yvault.Deposit(opts, size)
	}

	if err != nil {
		fmt.Printf("Error deposit into vault: %v", err)
		return err
	}

	bind.WaitMined(context.Background(), c.client.conn, tx)

	return nil
}

func (c *YearnClient) removeLiquidity(size *big.Int, coin coinType) error {
	var (
		tx  *types.Transaction
		err error
	)
	opts := &bind.TransactOpts{
		From:     c.client.opts.From,
		Signer:   c.client.opts.Signer,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}

	if coin == ETH {
		weth, err := yweth.NewYweth(common.HexToAddress(yETHVaultAddr), c.client.conn)
		if err != nil {
			return fmt.Errorf("Error getting weth contract")
		}
		tx, err = weth.WithdrawETH(opts, size)
	} else if coin != ETH {
		tokenAddr := CoinToAddressMap[coin]
		vaultAddr, ok := c.tokenToVault[tokenAddr]
		if !ok {
			return fmt.Errorf("No corresponding vault found for: %v ", coin)
		}
		yvault, err := yvault.NewYvault(vaultAddr, c.client.conn)
		if err != nil {
			return fmt.Errorf("Error getting weth contract")
		}
		opts.Value = size
		tx, err = yvault.Withdraw(opts, size)
	}

	if err != nil {
		fmt.Printf("Error withdraw from vault: %v", err)
		return err
	}

	bind.WaitMined(context.Background(), c.client.conn, tx)

	return nil
}

// Aave----------------------------------------------------------------------------

// AaveClient is an instance of Aave protocol.
type AaveClient struct {
	client      *ActualClient
	lendingPool *lendingpool.Lendingpool
}

// Aave returns a Aave client
func (c *ActualClient) Aave() *AaveClient {
	aaveClient := new(AaveClient)
	aaveClient.client = c

	lendingpool, err := lendingpool.NewLendingpool(common.HexToAddress(aaveLendingPoolAddr), c.conn)
	if err != nil {
		return nil
	}
	aaveClient.lendingPool = lendingpool
	return aaveClient
}

// Lend lend to the lending pool
func (c *AaveClient) Lend(size *big.Int, coin coinType) error {
	opts := &bind.TransactOpts{
		From:     c.client.opts.From,
		Signer:   c.client.opts.Signer,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}

	if coin != ETH {
		Approve(c.client, coin, common.HexToAddress(aaveLendingPoolCoreAddr), size)
	}

	tx, err := c.lendingPool.Deposit(opts, CoinToAddressMap[coin], size, 0)
	if err != nil {
		return err
	}
	bind.WaitMined(context.Background(), c.client.conn, tx)
	return nil
}

// Borrow borrow money from lending pool
func (c *AaveClient) Borrow(size *big.Int, coin coinType, interestRate rateModel) error {
	return nil
}

// ReserveData is a struct described the status of Aave lending pool
type ReserveData struct {
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
}

// GetUserReserveData get the reserve data
func (c *AaveClient) GetUserReserveData(addr common.Address, user common.Address) (ReserveData, error) {
	data, err := c.lendingPool.GetUserReserveData(nil, addr, user)
	if err != nil {
		return ReserveData{}, err
	}
	return data, nil
}

// Kyberswap----------------------------------------------------------------------

// KyberswapClient struct
type KyberswapClient struct {
	client *ActualClient
}

// Kyberswap returns a Kyberswap client
func (c *ActualClient) Kyberswap() *KyberswapClient {
	kyberClient := new(KyberswapClient)
	kyberClient.client = c
	return kyberClient
}

// SwapActions creates a swap action
func (c *KyberswapClient) SwapActions(size *big.Int, baseCurrency coinType, quoteCurrency coinType) *Actions {
	var (
		data         []byte
		err          error
		ethersNeeded *big.Int = big.NewInt(0)
	)

	parsed, err := abi.JSON(strings.NewReader(hkyber.HkyberABI))
	if err != nil {
		return nil
	}

	if quoteCurrency == ETH {
		ethersNeeded = size
		data, err = parsed.Pack("swapEtherToToken", size, CoinToAddressMap[baseCurrency], big.NewInt(0))
	} else {
		if baseCurrency == ETH {
			data, err = parsed.Pack("swapTokenToEther", CoinToAddressMap[baseCurrency], size, big.NewInt(0))
		} else {
			data, err = parsed.Pack("swapTokenToToken", CoinToAddressMap[baseCurrency], size, CoinToAddressMap[quoteCurrency], big.NewInt(0))
		}
	}

	if err != nil {
		return nil
	}

	return &Actions{
		Actions: []Action{
			{
				HandlerAddr:  common.HexToAddress(hKyberAddr),
				Data:         data,
				EthersNeeded: ethersNeeded,
			},
		},
	}

}

// utility------------------------------------------------------------------------
func Approve(client *ActualClient, coin coinType, addr common.Address, size *big.Int) error {
	erc20Contract, err := erc20.NewErc20(CoinToAddressMap[coin], client.conn)
	if err != nil {
		return err
	}
	opts := &bind.TransactOpts{
		Signer:   client.opts.Signer,
		From:     client.opts.From,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}
	tx, err := erc20Contract.Approve(opts, addr, size)
	bind.WaitMined(context.Background(), client.conn, tx)
	return nil
}
