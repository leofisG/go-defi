package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/524119574/go-defi/binding/compound/cETH"
	"github.com/524119574/go-defi/binding/compound/cToken"
	"github.com/524119574/go-defi/binding/erc20"
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
)

const (
	// uniswapAddr is UniswapV2Router, see here: https://uniswap.org/docs/v2/smart-contracts/router02/#address
	uniswapAddr   string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
	yRegistryAddr string = "0x3eE41C098f9666ed2eA246f4D2558010e59d63A0"
)

// CoinToAddressMap returns a mapping from coin to address
var CoinToAddressMap = map[coinType]common.Address{
	ETH:  common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
	BAT:  common.HexToAddress("0x0d8775f648430679a709e98d2b0cb6250d2887ef"),
	COMP: common.HexToAddress("0xc00e94cb662c3520282e6f5717214004a7f26888"),
	DAI:  common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
	USDC: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	USDT: common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
}

var coinToCompoundMap = map[coinType]common.Address{
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
		err := approve(c.client, quoteCurrency, common.HexToAddress(uniswapAddr), big.NewInt(size))
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
		cETHContract, err := cETH.NewCETH(cTokenAddr, c.client.conn)
		if err != nil {
			return err
		}

		tx, err = cETHContract.Mint(opts)
	case BAT, COMP, DAI, REP, SAI, UNI, USDC, USDT, WBTC, ZRX:
		err = approve(c.client, coin, cTokenAddr, big.NewInt(amount))
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
		cETHContract, err := cETH.NewCETH(cTokenAddr, c.client.conn)
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
		cETHContract, err := cETH.NewCETH(cTokenAddr, c.client.conn)
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
		cETHContract, err := cETH.NewCETH(cTokenAddr, c.client.conn)
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

func (c *CompoundClient) getPoolAddrFromCoin(coin coinType) (common.Address, error) {
	if val, ok := coinToCompoundMap[coin]; ok {
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
		weth, err := yweth.NewYweth(common.HexToAddress("0xe1237aA7f535b0CC33Fd973D66cBf830354D16c7"), c.client.conn)
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
		err = approve(c.client, coin, vaultAddr, size)
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
	return nil
}

// utility------------------------------------------------------------------------
func approve(client *ActualClient, coin coinType, addr common.Address, size *big.Int) error {
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
	tx, err := erc20Contract.Approve(opts, coinToCompoundMap[DAI], size)
	bind.WaitMined(context.Background(), client.conn, tx)
	return nil
}
