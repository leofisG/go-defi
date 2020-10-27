package client

type net_type int

const (
	MAIN_NET net_type = iota
	TEST_NET net_type = iota
)

type coin_type int

const (
	DAI coin_type = iota
	USDC coin_type = iota
	ETH coin_type = iota
)

// type UniswapClient interface {
// 	Swap(size int, base_currency coin_type, quote_currency coin_type)
// 	AddLiquidity(size int, currency coin_type)
// 	RemoveLiquidity(size int, currency coin_type)
// }

type Client interface {
	Uniswap() UniswapClient
}

func NewClient(net net_type, content string, passphrase string) (*ActualClient, error) {
	c := new(ActualClient)
	c.net = net
	c.content = content
	c.passphrase = passphrase
	return c, nil
}

type ActualClient struct {
	net net_type
	content string
	passphrase string
}

func (c *ActualClient) Uniswap() *UniswapClient {
	uni_client := new(UniswapClient)
	uni_client.net = c.net
	uni_client.content = c.content
	uni_client.passphrase = c.passphrase
	return uni_client
}

type UniswapClient struct {
	net net_type
	content string
	passphrase string
}

func (c *UniswapClient) Swap(size int, base_currency coin_type, quote_currency coin_type) {

}