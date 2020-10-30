package client

import (
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
		DAI: "xyz",
		USDC: "abc",
		ETH: "123",
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
	
	c.connection = conn
	return c, nil
}

// ActualClient is the struct that stores the information.
type ActualClient struct {
	net netType
	content string
	passphrase string
	connection *ethclient.Client
}

// UniswapClient struct
type UniswapClient struct {
	net netType
	content string
	passphrase string
}

// Uniswap returns a uniswap client
func (c *ActualClient) Uniswap() *UniswapClient {
	uniClient := new(UniswapClient)
	uniClient.net = c.net
	uniClient.content = c.content
	uniClient.passphrase = c.passphrase
	return uniClient
}

// TxHash represents a transaction hash
type TxHash string

// Swap in the Uniswap Exchange.
func (c *UniswapClient) Swap(size int, baseCurrency coinType, quoteCurrency coinType) (TxHash, error) {
	return "default_hash", nil
}