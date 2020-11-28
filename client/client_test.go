package client

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"testing"

	"github.com/524119574/go-defi/binding/erc20"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var key *ecdsa.PrivateKey
var ethClient *ethclient.Client
var publicKey *ecdsa.PublicKey
var fromAddr common.Address
var defiClient *ActualClient

func init() {
	var err error
	key, err = crypto.HexToECDSA("b8c1b5c1d81f9475fdf2e334517d29f733bdfa40682207571b12fc1142cbf329")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	ethClient, err = ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to ETH: %v", err)
	}

	publicKeyECDSA, ok := (key.Public()).(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type")
	}
	fromAddr = crypto.PubkeyToAddress(*publicKeyECDSA)
	defiClient, err = NewClient(bind.NewKeyedTransactor(key), ethClient, MainNet)
	if err != nil {
		log.Fatalf("Error creating client: %v.", err)
	}
	_ = fromAddr
	_ = ethClient
	_ = defiClient
}

func TestInteractWithCompound(t *testing.T) {

	beforeETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	err = defiClient.Compound().Supply(int64(1e18), ETH)
	if err != nil {
		t.Errorf("Failed to supply in compound: %v", err)
	}

	cETH, err := defiClient.Compound().BalanceOf(ETH)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if cETH.Cmp(big.NewInt(0)) == 0 {
		t.Errorf("CETH minting is not successful")
	}

	afterETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	if beforeETH.Cmp(afterETH) != 1 {
		t.Errorf("ETH balance not decreasing.")
	}
}

func TestInteractWithUniswap(t *testing.T) {
	beforeETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	err = defiClient.Uniswap().Swap(1e18, DAI, ETH, fromAddr)
	if err != nil {
		t.Errorf("Failed to swap in uniswap: %v", err)
	}

	afterETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)
	afterDAI, err := defiClient.BalanceOf(DAI)

	if beforeETH.Cmp(afterETH) != 1 {
		t.Errorf("ETH balance not decreasing.")
	}
	if afterDAI.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("Dai hasn't increased!")
	}

	err = defiClient.Compound().Supply(int64(1e18), DAI)
	if err != nil {
		t.Errorf("Failed to supply Dai in compound: %v", err)
	}

	cDai, err := defiClient.Compound().BalanceOf(DAI)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if cDai.Cmp(big.NewInt(0)) != 1 {
		t.Errorf("cDAI minting is not successful: %v", cDai)
	}

}

func TestInteractWithCompoundInDai(t *testing.T) {

	beforeDAI, err := defiClient.BalanceOf(DAI)

	err = defiClient.Compound().Supply(int64(1e18), DAI)
	if err != nil {
		t.Errorf("Failed to supply in compound: %v", err)
	}

	cDAI, err := defiClient.Compound().BalanceOf(DAI)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if cDAI.Cmp(big.NewInt(0)) == 0 {
		t.Errorf("CDai minting is not successful")
	}

	afterDAI, err := defiClient.BalanceOf(DAI)

	if beforeDAI.Cmp(afterDAI) != 1 {
		t.Errorf("Dai balance not decreasing.")
	}
}

func TestMintSomeUSDC(t *testing.T) {
	_, err := erc20.NewErc20(CoinToAddressMap[USDC], ethClient)
	if err != nil {
		t.Errorf("Error getting USDC Contract")
	}
	beforeUSDC, err := defiClient.BalanceOf(USDC)
	if err != nil {
		t.Errorf("Error getting USDC balance")
	}

	if beforeUSDC.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("Before USDC not 0, %v", beforeUSDC)
	}
}

func TestInteractWithYearn(t *testing.T) {
	beforeETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)

	err = defiClient.Yearn().addLiquidity(big.NewInt(1e18), ETH)
	if err != nil {
		t.Errorf("Failed to add liquidity in yearn: %v", err)
	}

	afterETH, err := ethClient.BalanceAt(context.Background(), fromAddr, nil)
	if beforeETH.Cmp(afterETH) != 1 {
		t.Errorf("ETH balance not decreasing.")
	}
}

func TestInteractWithFurucomboWithCompoundNew(t *testing.T) {

	beforeCETH, err := defiClient.Compound().BalanceOf(ETH)

	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	actions := new(Actions)

	actions.Add(
		defiClient.Compound().SupplyActions(big.NewInt(1e18), ETH),
	)

	defiClient.executeActions(actions)

	if err != nil {
		t.Errorf("Failed to interact with Furucombo: %v", err)
	}

	afterCETH, err := defiClient.Compound().BalanceOf(ETH)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if afterCETH.Cmp(beforeCETH) != 1 {
		t.Errorf("cETH minting is not successful via Furucombo: %v %v", afterCETH, beforeCETH)
	}

}

func TestInteractWithFurucomboWithCompoundERC20New(t *testing.T) {
	approve(defiClient, DAI, common.HexToAddress(furucomboAddr), big.NewInt(1e18))
	beforeCDai, err := defiClient.Compound().BalanceOf(DAI)

	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	actions := new(Actions)

	actions.Add(
		defiClient.Compound().SupplyActions(big.NewInt(1e18), DAI),
	)

	defiClient.executeActions(actions)

	if err != nil {
		t.Errorf("Failed to interact with Furucombo: %v", err)
	}

	afterCDai, err := defiClient.Compound().BalanceOf(DAI)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if afterCDai.Cmp(beforeCDai) != 1 {
		t.Errorf("cDai minting is not successful via Furucombo: %v %v", afterCDai, beforeCDai)
	}

}

func TestInteractWithFurucomboWithCompoundERC20withRedeem(t *testing.T) {
	approve(defiClient, DAI, common.HexToAddress(furucomboAddr), big.NewInt(1e18))
	// approve(defiClient, cDAI, common.HexToAddress(furucomboAddr), big.NewInt(1e18))
	beforeCDai, err := defiClient.Compound().BalanceOf(DAI)

	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	actions := new(Actions)

	actions.Add(
		defiClient.Compound().SupplyActions(big.NewInt(1e18), DAI),
	)

	actions.Add(
		defiClient.Compound().RedeemActions(big.NewInt(100000), DAI),
	)

	defiClient.executeActions(actions)

	if err != nil {
		t.Errorf("Failed to interact with Furucombo: %v", err)
	}

	afterCDai, err := defiClient.Compound().BalanceOf(DAI)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if afterCDai.Cmp(beforeCDai) != 1 {
		t.Errorf("cDai minting is not successful via Furucombo: %v %v", afterCDai, beforeCDai)
	}

}

func TestInteractWithFurucomboFlashLoan(t *testing.T) {
	approve(defiClient, DAI, common.HexToAddress(furucomboAddr), big.NewInt(1e18))
	approve(defiClient, cDAI, common.HexToAddress(furucomboAddr), big.NewInt(1e18))
	beforeCDai, err := defiClient.Compound().BalanceOf(DAI)

	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}

	actions := new(Actions)
	flashLoanActions := new(Actions)

	// TODO: use spread operator
	flashLoanActions.Add(
		defiClient.Compound().SupplyActions(big.NewInt(10000000), DAI),
	)
	flashLoanActions.Add(
		defiClient.Compound().RedeemActions(big.NewInt(50), DAI),
	)

	actions.Add(
		defiClient.Aave().FlashLoanActions(
			big.NewInt(10000000),
			DAI,
			flashLoanActions,
		),
	)

	defiClient.executeActions(actions)

	if err != nil {
		t.Errorf("Failed to interact with Furucombo: %v", err)
	}

	afterCDai, err := defiClient.Compound().BalanceOf(DAI)
	if err != nil {
		t.Errorf("Failed to get balance: %v", err)
	}

	if afterCDai.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("cDai minting is not successful via Furucombo: %v %v", afterCDai, beforeCDai)
	}

}
