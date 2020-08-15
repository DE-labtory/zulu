package wallet

import (
	"github.com/DE-labtory/zulu/types"
)

type Service interface {
	DeriveAccount(signer Signer) (types.Account, error)
	Transfer(signer Signer, to string, amount string) (types.Transaction, error)
}
