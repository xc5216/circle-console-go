package devwallet

import "github.com/xc5216/circle-console-go/model"

type basicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type publicKeyData struct {
	PublicKey string `json:"publicKey"`
}

type publicKeyResponse struct {
	Data publicKeyData `json:"data"`
}

type walletSetCreateRequest struct {
	Name                   string `json:"name"`
	IdempotencyKey         string `json:"idempotencyKey"`
	EntitySecretCipherText string `json:"entitySecretCipherText"`
}
type walletSetData struct {
	WalletSet model.WalletSet `json:"walletSet"`
}
type walletSetCreateResponse struct {
	basicResponse
	Data walletSetData `json:"data"`
}

type walletSetsData struct {
	WalletSets []model.WalletSet `json:"walletSets"`
}

type walletSetsGetResponse struct {
	basicResponse
	Data walletSetsData `json:"data"`
}

type walletsCreateRequest struct {
	IdempotencyKey         string   `json:"idempotencyKey"`
	EntitySecretCipherText string   `json:"entitySecretCipherText"`
	WalletSetID            string   `json:"walletSetId"`
	Blockchains            []string `json:"blockchains"`
	Count                  int      `json:"count"`
}

type walletsCreateData struct {
	Wallets []model.Wallet `json:"wallets"`
}

type walletsCreateResponse struct {
	basicResponse
	Data walletsCreateData `json:"data"`
}
