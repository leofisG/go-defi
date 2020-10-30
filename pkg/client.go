package client

import (
	uniswapfactory "github.com/524119574/go_defi/pkg/uniswapfactory"
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

var  coinToAddressMap = map[coinType]string{
		DAI: "0x6b175474e89094c44da98b954eedeac495271d0f",
		USDC: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		ETH: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
	}

// Client is the new interface
type Client interface {
	Uniswap() UniswapClient
}

// NewClient Create a new client
func NewClient(net netType, content string, passphrase string) (*ActualClient, error) {
	c := new(ActualClient)
	c.net = net
	c.content = content
	c.passphrase = passphrase
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
	conn *ethclient.Client
}

// UniswapClient struct
type UniswapClient struct {
	net netType
	content string
	passphrase string
	uniswap *uniswapfactory.Uniswap
}

// Uniswap returns a uniswap client
func (c *ActualClient) Uniswap() *UniswapClient {
	uniClient := new(UniswapClient)
	uniClient.net = c.net
	uniClient.content = c.content
	uniClient.passphrase = c.passphrase
	uniswap, err := uniswapfactory.NewUniswap(common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"), c.conn)
	if err != nil {
		return nil
	}
	
	uniClient.uniswap = uniswap
	return uniClient
}

// TxHash represents a transaction hash
type TxHash string

// Swap in the Uniswap Exchange.
func (c *UniswapClient) Swap(size int, baseCurrency coinType, quoteCurrency coinType) (TxHash, error) {
	return "default_hash", nil
}