package client

import (
	"fmt"
	"log"
	"testing"
)

// Test inbox contract gets deployed correctly
func TestNewClient(t *testing.T) {

	client, err := NewClient(MAIN_NET, "xyz", "abc")

	if err != nil {
		t.Errorf("Error creating client: %v.", err)
	}

	if client == nil {
		t.Errorf("Empty client.")
	}

	tx, err := client.Uniswap().Swap(1000, DAI, ETH)
	if err != nil {
		log.Fatalf("Failed to swap in Uniswap: %v", err)
	}

	fmt.Print("Test End.", tx)

}