package devwallet

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/xc5216/circle-console-go/model"
)

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
	model.BasicResponse
	Data walletSetData `json:"data"`
}

type walletSetsData struct {
	WalletSets []model.WalletSet `json:"walletSets"`
}

type walletSetsGetResponse struct {
	model.BasicResponse
	Data walletSetsData `json:"data"`
}

type walletsCreateRequest struct {
	WalletSetID            string                 `json:"walletSetId"`
	EntitySecretCipherText string                 `json:"entitySecretCipherText"`
	Blockchains            []string               `json:"blockchains"`
	IdempotencyKey         string                 `json:"idempotencyKey"`
	Count                  int                    `json:"count"`
	MetaData               []model.WalletMetadata `json:"metadata"`
	AccountType            model.WalletType       `json:"accountType"`
}

type walletsData struct {
	Wallets []model.Wallet `json:"wallets"`
}

type walletsCreateResponse struct {
	model.BasicResponse
	Data walletsData `json:"data"`
}

type WalletFilter struct {
	Address     string     `url:"address"`
	Blockchain  string     `url:"blockchain"`
	WalletSetID uuid.UUID  `url:"walletSetId"`
	RefID       string     `url:"refId"`
	From        *time.Time `url:"from"`
	To          *time.Time `url:"to"`
	PageBefore  *uuid.UUID `url:"pageBefore"`
	PageAfter   *uuid.UUID `url:"pageAfter"`
	PageSize    int        `url:"pageSize"`
}

type walletsGetRequest struct {
	Address     string `url:"address"`
	Blockchain  string `url:"blockchain"`
	WalletSetID string `url:"walletSetId"`
	RefID       string `url:"refId"`
	From        string `url:"from"`
	To          string `url:"to"`
	PageBefore  string `url:"pageBefore"`
	PageAfter   string `url:"pageAfter"`
	PageSize    int    `url:"pageSize"`
}

type walletsGetResponse struct {
	model.BasicResponse
	Data walletsData `json:"data"`
}

func (f WalletFilter) ToRequest() walletsGetRequest {
	ret := walletsGetRequest{
		Address:     f.Address,
		Blockchain:  f.Blockchain,
		WalletSetID: f.WalletSetID.String(),
		RefID:       f.RefID,
		PageSize:    f.PageSize,
	}
	if f.From != nil {
		ret.From = f.From.Format(time.RFC3339)
	}
	if f.To != nil {
		ret.To = f.To.Format(time.RFC3339)
	}
	if f.PageBefore != nil {
		ret.PageBefore = f.PageBefore.String()
	} else if f.PageAfter != nil {
		ret.PageAfter = f.PageAfter.String()
	}
	return ret
}

type walletData struct {
	Wallet model.Wallet `json:"wallet"`
}

type walletGetResponse struct {
	model.BasicResponse
	Data walletData `json:"data"`
}
