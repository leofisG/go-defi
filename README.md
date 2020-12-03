# Go Defi Framework

This is a framework to interact with various Ethereum Defi Protocol. In particular, it supports combine multiple actions with different protocols into one transaction which makes all these actions atomic. Furthermore the library support Aave Flash Loan.

## Getting Started
Initialize a flash loan
```
import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
    "github.com/524119574/go-defi/client"
)

	key, err = crypto.HexToECDSA("... your private key here ...")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err = ethclient.Dial("... ETH gateway, could be Infura ...")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	publicKeyECDSA, ok := (key.Public()).(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to get public key")
	}
	fromAddr = crypto.PubkeyToAddress(*publicKeyECDSA)

    // Create a New Defi Client
    defiClient, err = client.NewClient(bind.NewKeyedTransactor(key), ethClient, MainNet)
	if err != nil {
		log.Fatalf("Error creating client: %v.", err)
	}

    // Approve the Transaction.
    client.Approve(defiClient, client.DAI, common.HexToAddress(client.furucomboAddr), big.NewInt(2e18))
	if err != nil {
		t.Errorf("Error getting DAI balance")
	}

	actions := new(Actions)

    // The action that you want to take before returning the loaned fund
	flashLoanActions := new(client.Actions)

    // Supply 1 Dai to Compound
	flashLoanActions.Add(
		defiClient.Compound().SupplyActions(big.NewInt(1e18), DAI),
	)
    // Transfer some fund so that we have enough to pay back the principal and interest
	flashLoanActions.Add(
		defiClient.SupplyFundActions(big.NewInt(2e18), DAI),
	)
    // Redeem some cDAI
	flashLoanActions.Add(
		defiClient.Compound().RedeemActions(big.NewInt(1), DAI),
	)

    // Add the flash loan action
	actions.Add(
		defiClient.Aave().FlashLoanActions(
			big.NewInt(1e18),
			DAI,
			flashLoanActions,
		),
	)

	err = defiClient.executeActions(actions)

	if err != nil {
		log.Errorf("Failed to interact with Furucombo: %v", err)
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
- Yearn (TODO)
- Curve (TODO)
- 1inch (TODO)

## Contact
Leonard Ge [Twitter](https://twitter.com/ge_leonard)
