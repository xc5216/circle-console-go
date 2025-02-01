package model

import "fmt"

type RequestInfo struct {
	IdempotencyKey string
	RequestID      string
}

type BasicResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CircleAPIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e CircleAPIError) Error() string {
	return fmt.Sprintf("%d:%s", e.Code, e.Message)
}

type CreateWalletSetRequest struct {
	Name                   string `json:"name"`
	EntitySecretCipherText string `json:"entitySecretCipherText"`
}

type CreateWalletRequest struct {
	WalletSetID string           `json:"walletSetId"`
	Blockchains []string         `json:"blockchains"`
	Count       int              `json:"count"`
	MetaData    []WalletMetadata `json:"metadata"`
	AccountType WalletType       `json:"accountType"`
}

type GetTestnetTokenRequest struct {
	Blockchain string `json:"blockchain"`
	Address    string `json:"address"`
	Native     bool   `json:"native"`
	USDC       bool   `json:"usdc"`
	EURC       bool   `json:"eurc"`
}
