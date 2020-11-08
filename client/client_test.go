package client

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestInteractWithCompound(t *testing.T) {
	key, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	publicKeyECDSA, ok := (key.Public()).(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type")
	}
	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	beforeETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	auth := bind.NewKeyedTransactor(key)
	defiClient, err := NewClient(auth, ethClient, MainNet)
	if err != nil {
		t.Errorf("Error creating client: %v.", err)
	}

	err = defiClient.Compound().Supply(int64(1e18), ETH)
	if err != nil {
		log.Fatalf("Failed to supply in compound: %v", err)
	}

	cETH, err := defiClient.Compound().BalanceOf(ETH)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	if cETH == big.NewInt(0) {
		log.Fatal("CTH minting is not successful")
	}

	afterETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	if beforeETH.Cmp(afterETH) != 1 {
		log.Fatalf("ETH balance not decreasing.")
	}
}

func TestInteractWithUniswap(t *testing.T) {
	key, err := crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	publicKeyECDSA, ok := (key.Public()).(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type")
	}
	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	beforeETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	auth := bind.NewKeyedTransactor(key)
	defiClient, err := NewClient(auth, ethClient, MainNet)
	if err != nil {
		t.Errorf("Error creating client: %v.", err)
	}

	err = defiClient.Uniswap().Swap(1e18, DAI, ETH, fromAddr)
	if err != nil {
		log.Fatalf("Failed to swap in uniswap: %v", err)
	}

	afterETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)
	afterDAI, err := defiClient.BalanceOf(DAI)

	if beforeETH.Cmp(afterETH) != 1 {
		log.Fatalf("ETH balance not decreasing.")
	}
	if afterDAI.Cmp(big.NewInt(0)) != 1 {
		log.Fatalf("Dai hasn't increased!")
	}
}
