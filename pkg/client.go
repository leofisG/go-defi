package client

import (
	"fmt"
	"math/big"

	"github.com/524119574/go_defi/pkg/uniswap"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)


type netType int

const (
	// MAIN_NET the main Ethereum network
	MAIN_NET netType = iota
	// TEST_NET the test net.
	TEST_NET netType = iota
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

var  coinToAddressMap = map[coinType]common.Address{
		DAI: common.HexToAddress("0x6b175474e89094c44da98b954eedeac495271d0f"),
		USDC: common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"),
		ETH: common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
	}

// Client is the new interface
type Client interface {
	Uniswap() UniswapClient
}

// NewClient Create a new client
func NewClient(net netType, content string, passphrase string, opts *bind.TransactOpts) (*ActualClient, error) {
	c := new(ActualClient)
	c.net = net
	c.content = content
	c.passphrase = passphrase
	c.opts = opts
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/1c3ca57d49fa4db8aff58645c99fcc30")
	if err != nil {
		return nil, err
	}
	
	c.conn = conn
	return c, nil
}

// ActualClient is the struct that stores the information.
type ActualClient struct {
	net netType
	content string
	passphrase string
	opts *bind.TransactOpts
	conn *ethclient.Client
}

// UniswapClient struct
type UniswapClient struct {
	net netType
	content string
	passphrase string
	opts *bind.TransactOpts
	conn *ethclient.Client
	uniswap *uniswap.Uniswap
}

// Uniswap returns a uniswap client
func (c *ActualClient) Uniswap() *UniswapClient {
	uniClient := new(UniswapClient)
	uniClient.net = c.net
	uniClient.content = c.content
	uniClient.passphrase = c.passphrase
	uniClient.conn = c.conn
	uniClient.opts = c.opts
	// UniswapV2Router, see here: https://uniswap.org/docs/v2/smart-contracts/router02/#address
	uniswap, err := uniswap.NewUniswap(
		common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"), c.conn)
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
		c.opts, big.NewInt(size), big.NewInt(0), path, receipient, big.NewInt(2145916800))
	if err != nil {
		return "", err
	}
	fmt.Print(tx)
	return "default_hash", nil
}