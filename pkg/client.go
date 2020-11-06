package client

import (
	"fmt"
	"math/big"

	"github.com/524119574/go_defi/pkg/compound/cETH"
	"github.com/524119574/go_defi/pkg/uniswap"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
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
	// Compound
	cETHAddr string = "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"

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
func NewClient(net netType, content string, passphrase string, opts *bind.TransactOpts, addr string) (*ActualClient, error) {
	c := new(ActualClient)
	c.net = net
	c.content = content
	c.passphrase = passphrase
	c.opts = opts
	conn, err := ethclient.Dial(addr)
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


// Uniswap---------------------------------------------------------------------

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
		c.opts, big.NewInt(size), big.NewInt(0), path, receipient, big.NewInt(2145916800))
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
func (c *CompoundClient) Supply(amount int64, coin coinType) error {
	cETHContract, err := cETH.NewCETH(common.HexToAddress(cETHAddr), c.client.conn)
	if err != nil {
		fmt.Printf("Error getting cETH contract")
	}

	_, err = cETHContract.Mint(&bind.TransactOpts{
		From: c.client.opts.From,
		Signer: c.client.opts.Signer,
		Value: big.NewInt(amount),
		GasLimit: 150000,
		GasPrice: big.NewInt(20000000000),
	})

	if err != nil {
		fmt.Printf("Error mint ctoken: %v", err)
		return err
	}

	return nil
}

// Redeem supplies token to compound
func (c *CompoundClient) Redeem(amount int64, coin coinType) error {
	cETHContract, err := cETH.NewCETH(common.HexToAddress(cETHAddr), c.client.conn)
	if err != nil {
		fmt.Printf("Error getting cETH contract")
	}

	_, err = cETHContract.Redeem(&bind.TransactOpts{
		From: c.client.opts.From,
		Signer: c.client.opts.Signer,
		GasLimit: 500000,
		GasPrice: big.NewInt(20000000000),
	}, big.NewInt(1e8))

	if err != nil {
		fmt.Printf("Error mint ctoken: %v", err)
		return err
	}

	return nil
}

// BalanceOf return the balance of given cToken
func (c *CompoundClient) BalanceOf(coin coinType) (*big.Int, error) {
	cETHContract, err := cETH.NewCETH(common.HexToAddress(cETHAddr), c.client.conn)
	if err != nil {
		fmt.Printf("Error getting cETH contract")
	}
	
	val, err := cETHContract.BalanceOf(nil, c.client.opts.From)
	if err != nil {
		fmt.Printf("Error getting balance of cToken: %v", err)
		return big.NewInt(0), err
	}
	return val, nil
}



