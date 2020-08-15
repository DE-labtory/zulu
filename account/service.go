package account

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type Service interface {
	DeriveAccount(key keychain.Key) (types.Account, error)
	Transfer(key keychain.Key, to string, amount string) (types.Transaction, error)
	GetInfo() types.Coin
}