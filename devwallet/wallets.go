package devwallet

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/xc5216/circle-console-go/internal/setting"
	"github.com/xc5216/circle-console-go/internal/util"
	"github.com/xc5216/circle-console-go/model"
)

func (c *controller) CreateWallets(walletSetId uuid.UUID, blockchains []string, count int, accountType model.WalletType, metadata []model.WalletMetadata) (info *model.RequestInfo, wallets []model.Wallet, err error) {
	if metadata != nil && len(metadata) != count {
		err = fmt.Errorf("meta length must be equal to count: %w", model.ErrInvalidParameter)
		return
	}

	url := fmt.Sprintf("%s/w3s/developer/wallets", setting.GetServerURL())
	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	info = &model.RequestInfo{
		IdempotencyKey: id.String(),
		RequestID:      util.GenerateRequestID(),
	}
	request := walletsCreateRequest{
		IdempotencyKey:         info.IdempotencyKey,
		EntitySecretCipherText: c.getEntitySecretCipherText(),
		WalletSetID:            walletSetId.String(),
		Blockchains:            blockchains,
		Count:                  count,
		MetaData:               metadata,
		AccountType:            accountType,
	}
	req, err := util.GenerateJsonPostRequest(url, request, c.apiKey)
	if err != nil {
		return
	}
	util.SetRequestID(req, info.RequestID)
	response, err := util.DoRequestAndParseResultAs[walletsCreateResponse](req)
	if err != nil {
		return
	}
	return info, response.Data.Wallets, nil
}

func (c *controller) GetWallets(filter WalletFilter) (wallets []model.Wallet, err error) {
	walletsGetRequest := filter.ToRequest()
	url := fmt.Sprintf("%s/w3s/wallets", setting.GetServerURL())
	req, err := util.GenerateGetRequest(url, walletsGetRequest, c.apiKey)
	if err != nil {
		return
	}
	response, err := util.DoRequestAndParseResultAs[walletsGetResponse](req)
	if err != nil {
		return
	}
	return response.Data.Wallets, nil
}

func (c *controller) GetWallet(id uuid.UUID) (wallet model.Wallet, err error) {
	url := fmt.Sprintf("%s/w3s/wallets/%s", setting.GetServerURL(), id.String())
	req, err := util.GenerateGetRequest(url, nil, c.apiKey)
	if err != nil {
		return
	}
	response, err := util.DoRequestAndParseResultAs[walletGetResponse](req)
	if err != nil {
		return
	}
	return response.Data.Wallet, nil
}
