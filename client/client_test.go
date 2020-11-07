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
func TestInteractWithCompound(t *testing.T) {
	key, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	auth := bind.NewKeyedTransactor(key)
	defiClient, err := NewClient(auth, ethClient, MainNet)
	if err != nil {
		t.Errorf("Error creating client: %v.", err)
	}

	err = defiClient.Compound().Supply(int64(1e18), ETH)
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
