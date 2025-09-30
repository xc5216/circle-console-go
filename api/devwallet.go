package api

import (
	"github.com/xc5216/circle-console-go/internal/util"
	"github.com/xc5216/circle-console-go/model"
)

func (ctrl devWalletCtrl) CreateWallets(request model.CreateWalletRequest, idempotencyKey string) (requestID string, wallets []model.Wallet, err error) {
	r := &walletsCreateRequest{
		EntitySecretCipherText: ctrl.getEntitySecretCipherText(),
		CreateWalletRequest:    request,
	}
	response := walletsData{}
	requestID, err = util.SendRequest[walletsCreateRequest, any, walletsData](
		util.GetEndPointCreateWallets(), ctrl.apiKey, idempotencyKey, r, nil, &response)
	if err != nil {
		return
	}
	return requestID, response.Wallets, nil
}

func (ctrl devWalletCtrl) GetWalletTokenBalance(walletID string, request model.GetWalletTokenBalanceRequest, idempotencyKey string) (requestID string, balance []model.WalletTokenBalance, err error) {
	response := WalletTokenBalanceData{}
	requestID, err = util.SendRequest[any, model.GetWalletTokenBalanceRequest, WalletTokenBalanceData](
		util.GetEndPointGetWalletTokenBalance(walletID), ctrl.apiKey, idempotencyKey, nil, &request, &response)
	if err != nil {
		return
	}
	return requestID, response.TokenBalances, nil
}

func (ctrl devWalletCtrl) Transfer(request model.TransferRequest, idempotencyKey string) (requestID string, err error) {
	r := &TransferRequest{
		TransferRequest:        request,
		EntitySecretCipherText: ctrl.getEntitySecretCipherText(),
		FeeLevel:               "MEDIUM",
	}
	response := TransferData{}
	requestID, err = util.SendRequest[TransferRequest, any, TransferData](
		util.GetEndPointCreateDeveloperTransfer(), ctrl.apiKey, idempotencyKey, r, nil, &response)
	if err != nil {
		return
	}
	return requestID, nil
}
