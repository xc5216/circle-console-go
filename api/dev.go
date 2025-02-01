package api

import (
	"encoding/base64"

	"github.com/xc5216/circle-console-go/internal/util"
	"github.com/xc5216/circle-console-go/model"
)

type devWalletCtrl struct {
	*generalCtrl
	apiKey       string
	publicKey    string
	entitySecret []byte
}

func NewDevWalletCtrl(apiKey string, entitySecret []byte) (*devWalletCtrl, error) {
	ctrl := NewGeneralCtrl(apiKey)
	publicKey, err := ctrl.GetPublicKey(apiKey)
	if err != nil {
		return nil, err
	}

	return &devWalletCtrl{
		generalCtrl:  ctrl,
		apiKey:       apiKey,
		publicKey:    publicKey,
		entitySecret: entitySecret,
	}, nil
}

func (ctrl devWalletCtrl) getEntitySecretCipherText() string {
	entitySecretCipherText, err := util.EncryptEntitySecret(ctrl.entitySecret, ctrl.publicKey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(entitySecretCipherText)
}

func (ctrl devWalletCtrl) CreateWalletSet(name string, idempotencyKey string) (requestID string, walletSet model.WalletSet, err error) {
	request := model.CreateWalletSetRequest{
		Name:                   name,
		EntitySecretCipherText: ctrl.getEntitySecretCipherText(),
	}
	response := walletSetData{}
	requestID, err = util.SendRequest[model.CreateWalletSetRequest, any, walletSetData](
		util.EndPointCreateWalletSet, ctrl.apiKey, idempotencyKey, request, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.WalletSet, nil
}

func (ctrl devWalletCtrl) CreateWallets(request model.CreateWalletRequest, idempotencyKey string) (requestID string, wallets []model.Wallet, err error) {
	r := walletsCreateRequest{
		EntitySecretCipherText: ctrl.getEntitySecretCipherText(),
		CreateWalletRequest:    request,
	}
	response := walletsData{}
	requestID, err = util.SendRequest[walletsCreateRequest, any, walletsData](
		util.EndPointCreateWallets, ctrl.apiKey, idempotencyKey, r, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.Wallets, nil
}
