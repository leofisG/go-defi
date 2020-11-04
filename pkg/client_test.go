package client

import (
	"fmt"
	"log"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// Test inbox contract gets deployed correctly
func TestNewClient(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)
	client, err := NewClient(MAIN_NET, "xyz", "abc", auth)

	if err != nil {
		t.Errorf("Error creating client: %v.", err)
	}

	if client == nil {
		t.Errorf("Empty client.")
	}

	tx, err := client.Uniswap().Swap(int64(1000), DAI, ETH, common.HexToAddress("0"))
	if err != nil {
		log.Fatalf("Failed to swap in Uniswap: %v", err)
	}

	fmt.Print("Test End.", tx)

}