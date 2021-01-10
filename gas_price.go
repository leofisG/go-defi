package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/524119574/go-defi/client"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func measureGasPrice() {
	var (
		blockNum  *big.Int
		diffTotal *big.Int
		exceedCnt int
	)
	key, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	// Create a New Defi Client
	defiClient := client.NewClient(bind.NewKeyedTransactor(key), ethClient)
	if err != nil {
		log.Fatalf("Error creating client: %v.", err)
	}

	blockNum = big.NewInt(0)
	blockNum.SetString("11500000", 10)
	exceedCnt = 0
	diffTotal = big.NewInt(0)
	cnt := 0

	for cnt < 100 {
		gasPrice, err := defiClient.SuggestGasPrice(blockNum)
		if err != nil {
			log.Printf("Error getting suggested gas price: %v. ", err)
			blockNum = blockNum.Add(blockNum, big.NewInt(1))
			continue
		}
		fmt.Printf("block %v gas price is: %v \n", blockNum, gasPrice)

		blockNum = blockNum.Add(blockNum, big.NewInt(1))
		block, err := ethClient.BlockByNumber(context.Background(), blockNum)
		if err != nil {
			log.Printf("Error getting next block transaction: %v. ", err)
			continue
		}

		txs := block.Transactions()
		if len(txs) == 0 {
			log.Printf("jump: %v\n\n", blockNum)
			continue
		}

		minGas := txs[0].GasPrice()
		maxGas := txs[0].GasPrice()
		for _, tx := range txs {
			if minGas.Cmp(tx.GasPrice()) == 1 {
				minGas = tx.GasPrice()
			}
			if maxGas.Cmp(tx.GasPrice()) == -1 {
				maxGas = tx.GasPrice()
			}
		}

		t := new(big.Int)
		if gasPrice.Cmp(minGas) != -1 {
			exceedCnt++
			t.Sub(gasPrice, minGas)
			diffTotal.Add(diffTotal, t)
		}
		cnt++
	}

	fmt.Printf("Diff total is: %v \n", diffTotal)
	fmt.Printf("Exceed Cnt is: %v \n", exceedCnt)
	fmt.Printf("Diff average is: %v \n", diffTotal.Div(diffTotal, big.NewInt(int64(exceedCnt))))
}

func measureExecutionTime() {
	key, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	// Create a New Defi Client
	defiClient := client.NewClient(bind.NewKeyedTransactor(key), ethClient)
	if err != nil {
		log.Fatalf("Error creating client: %v.", err)
	}

	start := time.Now()
	for i := 0; i < 1000; i++ {
		actions := new(client.Actions)

		// Add the flash loan action
		actions.Add(
			defiClient.Compound().SupplyActions(big.NewInt(1e18), client.ETH),
		)
		defiClient.CombineActions(actions)
	}
	duration := time.Since(start)
	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println("Time take in seconds", duration.Seconds()/1000.0)

}

func main() {
	// measureExecutionTime()
	measureGasPrice()
}
