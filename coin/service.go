package coin

import (
	"github.com/DE-labtory/zulu/signer"
	"github.com/DE-labtory/zulu/types"
)

type Service interface {
	DeriveAccount(signer signer.Signer) (types.Account, error)
	Transfer(signer signer.Signer, to string, amount string) (types.Transaction, error)
}
