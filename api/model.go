package api

import "github.com/xc5216/circle-console-go/model"

type walletSetData struct {
	WalletSet model.WalletSet `json:"walletSet"`
}

type walletsCreateRequest struct {
	EntitySecretCipherText string `json:"entitySecretCipherText"`
	model.CreateWalletRequest
}

type walletsData struct {
	Wallets []model.Wallet `json:"wallets"`
}
