package client

import (
	"testing"
	"fmt"
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

	fmt.Print("Test End.")

}