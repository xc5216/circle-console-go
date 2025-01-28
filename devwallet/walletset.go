package devwallet

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/xc5216/circle-console-go/internal/setting"
	"github.com/xc5216/circle-console-go/internal/util"
	"github.com/xc5216/circle-console-go/model"
)

type controller struct {
	apiKey       string
	publicKey    string
	entitySecret []byte
}

func New(apiKey string, entitySecret []byte) (*controller, error) {
	publicKey, err := GetPublicKey(apiKey)
	if err != nil {
		return nil, err
	}

	return &controller{
		apiKey:       apiKey,
		publicKey:    publicKey,
		entitySecret: entitySecret,
	}, nil
}

func (c *controller) getEntitySecretCipherText() string {
	entitySecretCipherText, err := EncryptEntitySecret(c.entitySecret, c.publicKey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(entitySecretCipherText)
}

func (c *controller) CreateWalletSet(name string) (idempotencyKey string, walletSet model.WalletSet, err error) {
	url := fmt.Sprintf("%s/developer/walletSets", setting.GetServerURL())
	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	idempotencyKey = id.String()
	request := walletSetCreateRequest{
		Name:                   name,
		IdempotencyKey:         idempotencyKey,
		EntitySecretCipherText: c.getEntitySecretCipherText(),
	}

	req, err := util.GenerateJsonPostRequest(url, request, c.apiKey)
	if err != nil {
		return
	}

	response, err := util.DoRequestAndParseResultAs[walletSetCreateResponse](req)
	if err != nil {
		return
	}

	walletSet = response.Data.WalletSet
	return
}

func (c *controller) GetDevWalletSets() ([]model.WalletSet, error) {
	url := fmt.Sprintf("%s/walletSets", setting.GetServerURL())

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	util.SetApiKey(req, c.apiKey)

	response, err := util.DoRequestAndParseResultAs[walletSetsGetResponse](req)
	if err != nil {
		return nil, err
	}

	walletSets := util.Filter(response.Data.WalletSets, func(walletSet model.WalletSet) bool {
		return walletSet.CustodyType == model.CustodyTypeDeveloper
	})
	return walletSets, nil
}

func (c *controller) CreateWallets(walletSetId uuid.UUID, blockchains []string, count int) (idempotencyKey string, wallets []model.Wallet, err error) {
	url := fmt.Sprintf("%s/developer/wallets", setting.GetServerURL())
	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	idempotencyKey = id.String()
	request := walletsCreateRequest{
		IdempotencyKey:         idempotencyKey,
		EntitySecretCipherText: c.getEntitySecretCipherText(),
		WalletSetID:            walletSetId.String(),
		Blockchains:            blockchains,
		Count:                  count,
	}
	req, err := util.GenerateJsonPostRequest(url, request, c.apiKey)
	if err != nil {
		return
	}
	response, err := util.DoRequestAndParseResultAs[walletsCreateResponse](req)
	if err != nil {
		return
	}
	return idempotencyKey, response.Data.Wallets, nil
}
