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

func (ctrl devWalletCtrl) GetWalletSets(idempotencyKey string) (requestID string, walletSets []model.WalletSet, err error) {
	response := walletSetsData{}
	requestID, err = util.SendRequest[any, any, walletSetsData](
		util.GetEndPointGetWalletSets(), ctrl.apiKey, idempotencyKey, nil, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.WalletSets, nil
}

func (ctrl devWalletCtrl) CreateWalletSet(name string, idempotencyKey string) (requestID string, walletSet model.WalletSet, err error) {
	request := &model.CreateWalletSetRequest{
		Name:                   name,
		EntitySecretCipherText: ctrl.getEntitySecretCipherText(),
	}
	response := walletSetData{}
	requestID, err = util.SendRequest[model.CreateWalletSetRequest, any, walletSetData](
		util.GetEndPointCreateWalletSet(), ctrl.apiKey, idempotencyKey, request, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.WalletSet, nil
}

func (ctrl devWalletCtrl) GetWalletSet(walletSetID string, idempotencyKey string) (requestID string, walletSet model.WalletSet, err error) {
	response := walletSetData{}
	requestID, err = util.SendRequest[any, any, walletSetData](
		util.GetEndPointGetWalletSetByID(walletSetID), ctrl.apiKey, idempotencyKey, nil, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.WalletSet, nil
}

func (ctrl devWalletCtrl) UpdateWalletSetName(walletSetID string, name string, idempotencyKey string) (requestID string, walletSet model.WalletSet, err error) {
	req := &walletSetUpdateRequest{
		Name: name,
	}
	response := walletSetData{}
	requestID, err = util.SendRequest[walletSetUpdateRequest, any, walletSetData](
		util.GetEndPointUpdateWalletSetNameByID(walletSetID), ctrl.apiKey, idempotencyKey, req, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.WalletSet, nil
}
