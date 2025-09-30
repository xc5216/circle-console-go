package api

import "github.com/xc5216/circle-console-go/model"

type walletsCreateRequest struct {
	EntitySecretCipherText string `json:"entitySecretCipherText"`
	model.CreateWalletRequest
}

type walletSetUpdateRequest struct {
	Name string `json:"name"`
}

type walletSetData struct {
	WalletSet model.WalletSet `json:"walletSet"`
}

type walletsData struct {
	Wallets []model.Wallet `json:"wallets"`
}

type publicKeyData struct {
	PublicKey string `json:"publicKey"`
}

type walletSetsData struct {
	WalletSets []model.WalletSet `json:"walletSets"`
}

type WalletTokenBalanceData struct {
	TokenBalances []model.WalletTokenBalance `json:"tokenBalances"`
}

type TransferRequest struct {
	model.TransferRequest
	EntitySecretCipherText string `json:"entitySecretCipherText"`
	FeeLevel               string `json:"feeLevel"`
}

type TransferData struct {
	ID    string `json:"id"`
	State string `json:"state"`
}
