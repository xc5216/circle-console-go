package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type CustodyType string

const (
	CustodyTypeDeveloper CustodyType = "DEVELOPER"
	CustodyTypeUser      CustodyType = "USER"
)

type WalletSet struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	CustodyType CustodyType `json:"custodyType"`
	UpdateDate  time.Time   `json:"updateDate"`
	CreateDate  time.Time   `json:"createDate"`
}
