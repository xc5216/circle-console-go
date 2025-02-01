package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type WalletType string

const (
	WalletTypeSmartContractAccount   = "SCA"
	WalletTypeExternallyOwnedAccount = "EOA"
)

type WalletMetadata struct {
	Name  string `json:"name"`
	RefID string `json:"refId"`
}

type Wallet struct {
	WalletMetadata
	ID          uuid.UUID   `json:"id"`
	Address     string      `json:"address"`
	Blockchain  string      `json:"blockchain"`
	CreateDate  time.Time   `json:"createDate"`
	UpdateDate  time.Time   `json:"updateDate"`
	CustodyType CustodyType `json:"custodyType"`
	State       string      `json:"state"`
	WalletSetID uuid.UUID   `json:"walletSetId"`
	AccountType string      `json:"accountType"`
}
