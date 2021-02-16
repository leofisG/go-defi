# Go Defi Framework

This is a framework to combine multiple interactions with Ethereum Defi framework into one atomic transaction.

This library support flash loan, so you can use this library to do things like arbitrage. The things that you can do includes:
- Arbitrage yCRV ([here](https://furucombo.app/explore/combo_curve_00015))
- Leverage Short DAI ([here](https://furucombo.app/explore/combo_maker_00009))

### Supported Protocols
- Compound
    - Supply token: `client.Compound().SupplyActions()`
    - Redeem token: `client.Compound().RedeemActions()`
- Aave
    - Flash loan: `client.Aave().FlashLoanActions()`
- Uniswap
    - Swap: `client.Uniswap().SwapActions()`
- Kyberswap
    - Swap: `client.Kyberswap().SwapActions()`
- Yearn
	- Supply to Vault: `client.Yearn().AddLiquidityActions()`
	- Withdraw from Vault: `client.Yearn().RemoveLiquidityActions()`
- Curve
	- Exchange token: `client.Curve().ExchangeActions()`
	- Exchange underlying token `client.Curve().ExchangeUnderlyingActions`
	- Add Liquidity: `client.Curve().AddLiquidityActions()`
	- Remove Liquidity: `client.Curve().RemoveLiquidityActions()`
- Sushiswap
    - Swap: `client.Sushiswap().SwapActions()`
- MakerDao
	- Create New Vault: `client.Maker().GenerateDaiAction()`
	- Burn DAI and reduce debt: `client.Maker().WipeAction()`
	- Add more collateral to vault: `client.Maker().DepositCollateralActions()`

### APIs

The main API for this tool is the `ExecuteActions` API.

The tool sends the transaction into a proxy contract, and let the proxy
contract to actually interact with the underlying protocols. 

A graph illustration of the idea is the following:
![Proxy Contract Interaction](./images/illustration_with_compound.png)

The `Client`, i.e. this tool interact with the proxy contract, the proxy contract does the following:
1. check if the handler is valid through the `isValid` function of the `Registry` contract
2. If step 1 is successful, delegate call the compound handler contract to interact with the underlyng compound code.

In the client we create an empty `Actions` by doing:
```go
actions := new(client.Actions)
```
We then add more `Actions` by calling the `Add` function of the `Actions` struct:
```go
actions.Add(
	defiClient.Compound().SupplyActions(big.NewInt(1e18), client.DAI),
	defiClient.Uniswap().SwapActions(big.NewInt(1e18), client.DAI, client.ETH),
)
```
After that we send the `actions` by calling the `ExecuteActions` function:
```go
err = defiClient.ExecuteActions(actions)
```
And done we have executed a transaction

## Complete Working Example for flash loan
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
	defiClient := client.NewClient(bind.NewKeyedTransactor(key), ethClient)
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
Go Version: 1.13

## Documentation
A more thorough documentation can be found [here](https://godoc.org/github.com/524119574/go-defi/client).

## Contact
- Leonard Ge [Twitter](https://twitter.com/ge_leonard)
- Kaihua Qin
- Liyi Zhou
- Arthur Gervais

## Acknowledgement
This is project is inspired by [Furucombo](https://furucombo.app/).
