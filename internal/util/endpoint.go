package util

import (
	"fmt"
	"net/http"
)

func GetEndPointGetPublicKey() EndPoint {
	return EndPoint{
		URL:    "/w3s/config/entity/publicKey",
		Method: http.MethodGet,
	}
}
func GetEndPointGetWalletSets() EndPoint {
	return EndPoint{
		URL:    "/w3s/developer/walletSets",
		Method: http.MethodPost,
	}
}
func GetEndPointCreateWalletSet() EndPoint {
	return EndPoint{
		URL:    "/w3s/developer/walletSets",
		Method: http.MethodPost,
	}
}
func GetEndPointCreateWallets() EndPoint {
	return EndPoint{
		URL:    "/w3s/developer/wallets",
		Method: http.MethodPost,
	}
}
func GetEndPointFaucet() EndPoint {
	return EndPoint{
		URL:    "/faucet/drips",
		Method: http.MethodPost,
	}
}
func GetEndPointGetWalletSetByID(walletSetID string) EndPoint {
	return EndPoint{
		URL:    "/w3s/developer/walletSets/" + walletSetID,
		Method: http.MethodGet,
	}
}
func GetEndPointUpdateWalletSetNameByID(walletSetID string) EndPoint {
	return EndPoint{
		URL:    "/w3s/developer/walletSets/" + walletSetID,
		Method: http.MethodPut,
	}
}
func GetEndPointGetWalletTokenBalance(walletID string) EndPoint {
	return EndPoint{
		URL:    fmt.Sprintf("/w3s/wallets/%s/balances", walletID),
		Method: http.MethodGet,
	}
}
func GetEndPointCreateDeveloperTransfer() EndPoint {
	return EndPoint{
		URL:    "/w3s/developer/transactions/transfer",
		Method: http.MethodPost,
	}
}
