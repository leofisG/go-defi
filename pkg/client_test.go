package client

import (
	"fmt"
	"log"
	"testing"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Test inbox contract gets deployed correctly
func TestNewClient(t *testing.T) {
	// key, _ := crypto.GenerateKey()
	key, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	auth := bind.NewKeyedTransactor(key)
	defiClient, err := NewClient(auth, ethClient, MainNet)

	if err != nil {
		t.Errorf("Error creating client: %v.", err)
	}

	if defiClient == nil {
		t.Errorf("Empty client.")
	}

	err = defiClient.Compound().Supply(int64(10000000), ETH)
	// // tx, err := client.Uniswap().Swap(int64(1000), DAI, ETH, common.HexToAddress("0"))
	if err != nil {
		log.Fatalf("Failed to supply in compound: %v", err)
	}

	val, err := defiClient.Compound().BalanceOf(ETH)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	log.Fatalf("Get balance: %d", val.Int64())
	fmt.Print("Test End.")
	fmt.Printf("number %d", val.Int64())

}