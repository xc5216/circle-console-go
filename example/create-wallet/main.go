package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/xc5216/circle-console-go/api"
	"github.com/xc5216/circle-console-go/model"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	entitySecretStr := os.Getenv("ENTITY_SECRET")
	entitySecret, err := hex.DecodeString(entitySecretStr)
	if err != nil {
		panic(err)
	}

	ctrl, err := api.NewDevWalletCtrl(apiKey, entitySecret)
	if err != nil {
		panic(err)
	}
	name := fmt.Sprintf("walletSet-%d", time.Now().Unix())
	requestID, walletSet, err := ctrl.CreateWalletSet(name, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("CreateWalletSet Request ID: ", requestID)
	fmt.Println("Wallet set: ", walletSet)

	requestID, wallets, err := ctrl.CreateWallets(model.CreateWalletRequest{
		WalletSetID: walletSet.ID.String(),
		Blockchains: []string{"MATIC-AMOY"},
		Count:       1,
		MetaData:    nil,
		AccountType: model.WalletTypeExternallyOwnedAccount,
	}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("CreateWallets Request ID: ", requestID)
	fmt.Println("Wallets: ", wallets)

	requestID, err = ctrl.RequestTestnetToken(model.GetTestnetTokenRequest{
		Blockchain: "MATIC-AMOY",
		Address:    wallets[0].Address,
		Native:     true,
	}, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("RequestTestnetToken Request ID: ", requestID)
}
