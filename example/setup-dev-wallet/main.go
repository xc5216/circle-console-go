package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/xc5216/circle-console-go/api"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	ctrl := api.NewGeneralCtrl(apiKey)

	entitySecret := api.GenerateRandomEntitySecret()
	publicKey, err := ctrl.GetPublicKey(apiKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Public key: ", publicKey)
	encryptedEntitySecret, err := api.EncryptEntitySecret(entitySecret, publicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println("Entity secret: ", hex.EncodeToString(entitySecret))
	fmt.Println("Encrypted entity secret: ", base64.StdEncoding.EncodeToString(encryptedEntitySecret))
}
