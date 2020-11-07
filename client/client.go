package client

import (
	"context"
	"fmt"
	"math/big"

	"github.com/524119574/go_defi/binding/compound/cETH"
	"github.com/524119574/go_defi/binding/uniswap"
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
	// DAI weiwu
	DAI coinType = iota
	// USDC weiwu
	USDC coinType = iota
	// ETH weiwu
	ETH coinType = iota
)

const (
	// uniswapAddr is UniswapV2Router, see here: https://uniswap.org/docs/v2/smart-contracts/router02/#address
	uniswapAddr string = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
)

var coinToAddressMap = map[coinType]common.Address{
	DAI:  common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
	USDC: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
	ETH:  common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
}

var coinToCompoundMap = map[coinType]common.Address{
	DAI:  common.HexToAddress("0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"),
	USDC: common.HexToAddress("0x39aa39c021dfbae8fac545936693ac917d5e7563"),
	ETH:  common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"),
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
func (c *UniswapClient) Swap(size int64, baseCurrency coinType, quoteCurrency coinType, receipient common.Address) (TxHash, error) {
	path := []common.Address{coinToAddressMap[quoteCurrency], coinToAddressMap[ETH], coinToAddressMap[baseCurrency]}
	tx, err := c.uniswap.SwapExactTokensForTokens(
		// TODO: there is basically no minimum output amount set, so this could cause huge slippage, need to fix.
		// Also the time stamp is set to 2038 January 1, it's better to set it dynamically.
		c.client.opts, big.NewInt(size), big.NewInt(0), path, receipient, big.NewInt(2145916800))
	if err != nil {
		return "", err
	}
	fmt.Print(tx)
	return "default_hash", nil
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
		Value:    big.NewInt(amount),
	}

	switch coin {
	case ETH:
		cETHContract, err := cETH.NewCETH(cTokenAddr, c.client.conn)
		if err != nil {
			fmt.Printf("Error getting cETH contract")
		}

		tx, err = cETHContract.Mint(opts)
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
			fmt.Printf("Error getting cETH contract")
		}

		val, err = cETHContract.BalanceOf(nil, c.client.opts.From)
	}

	if err != nil {
		fmt.Printf("Error getting balance of cToken: %v", err)
		return big.NewInt(0), err
	}
	return val, nil
}

// BalanceOfUnderlying return the balance of given cToken
func (c *CompoundClient) BalanceOfUnderlying(coin coinType) (*big.Int, error) {
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
			fmt.Printf("Error getting cETH contract")
		}

		val, err = cETHContract.BalanceOf(nil, c.client.opts.From)
	}

	if err != nil {
		fmt.Printf("Error getting balance of cToken: %v", err)
		return big.NewInt(0), err
	}
	return val, nil
}

func (c *CompoundClient) getPoolAddrFromCoin(coin coinType) (common.Address, error) {
	if val, ok := coinToCompoundMap[coin]; ok {
		return val, nil
	}
	return common.Address{}, fmt.Errorf("No corresponding compound pool for token: %v", coin)
}
