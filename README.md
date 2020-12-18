# Go Defi Framework

This is a framework to combine multiple interactions with Ethereum Defi framework into one atomic transaction.

This library support flash loan, so you can use this library to do things like arbitrage.

## Getting Started
Initialize a flash loan

```go
package main

import (
	"log"
	"math/big"

	"github.com/524119574/go-defi/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	key, err := crypto.HexToECDSA("... your private key here ...")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("... ETH gateway, could be Infura ...")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	// Create a New Defi Client
	defiClient, err := client.NewClient(bind.NewKeyedTransactor(key), ethClient, client.MainNet)
	if err != nil {
		log.Fatalf("Error creating client: %v.", err)
	}

	// Approve the Transaction.
	client.Approve(defiClient, client.DAI, common.HexToAddress(client.FurucomboAddr), big.NewInt(2e18))
	if err != nil {
		log.Fatal("Error getting DAI balance")
	}

	actions := new(client.Actions)

	// The action that you want to take before returning the loaned fund
	flashLoanActions := new(client.Actions)

	// Supply 1 Dai to Compound
	flashLoanActions.Add(
		defiClient.Compound().SupplyActions(big.NewInt(1e18), client.DAI),
	)
	// Transfer some fund so that we have enough to pay back the principal and interest
	flashLoanActions.Add(
		defiClient.SupplyFundActions(big.NewInt(2e18), client.DAI),
	)
	// Redeem some cDAI
	flashLoanActions.Add(
		defiClient.Compound().RedeemActions(big.NewInt(1), client.DAI),
	)

	// Add the flash loan action
	actions.Add(
		defiClient.Aave().FlashLoanActions(
			big.NewInt(1e18),
			client.DAI,
			flashLoanActions,
		),
	)

	err = defiClient.ExecuteActions(actions)

	if err != nil {
		log.Fatalf("Failed to interact with Furucombo: %v", err)
	}
}

```

### APIs
- Compound
    - Supply token
    - Redeem token
- Aave (TODO)
    - Flash loan
- Uniswap
    - Swap
- Kyberswap
    - Swap
- Yearn
	- Supply to Vault
	- Withdraw from Vault
- Curve
	- Exchange token
	- Exchange underlying token 
- 1inch (TODO)
- MakerDao (TODO)

## Contact
- Leonard Ge [Twitter](https://twitter.com/ge_leonard)
- Kaihua Qin
- Liyi Zhou
- Arthur Gervais

