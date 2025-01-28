package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/xc5216/circle-console-go/devwallet"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	entitySecretStr := os.Getenv("ENTITY_SECRET")
	entitySecret, err := hex.DecodeString(entitySecretStr)
	if err != nil {
		panic(err)
	}

	controller, err := devwallet.New(apiKey, entitySecret)
	if err != nil {
		panic(err)
	}
	_, walletSet, err := controller.CreateWalletSet("test wallet")
	if err != nil {
		panic(err)
	}
	fmt.Println("Wallet set: ", walletSet)

	_, wallets, err := controller.CreateWallets(walletSet.ID, []string{"ETH-SEPOLIA"}, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wallets: ", wallets)
}
