package api

import (
	"github.com/xc5216/circle-console-go/internal/util"
	"github.com/xc5216/circle-console-go/model"
)

type generalCtrl struct {
	apiKey string
}

func NewGeneralCtrl(apiKey string) *generalCtrl {
	return &generalCtrl{
		apiKey: apiKey,
	}
}

// GetPublicKey will get public key from circle server
func (ctrl generalCtrl) GetPublicKey(apiKey string) (string, error) {
	result := publicKeyData{}
	_, err := util.SendRequest[any, any](util.EndPointGetPublicKey, apiKey, "", nil, nil, &result)
	if err != nil {
		return "", err
	}
	return result.PublicKey, nil
}

// RequestTestnetToken requests testnet tokens for the specified address.
func (ctrl generalCtrl) RequestTestnetToken(request model.GetTestnetTokenRequest, idempotencyKey string) (requestID string, err error) {
	return util.SendRequest[any, any, any](util.EndPointFaucet, ctrl.apiKey, idempotencyKey, request, nil, nil)
}
