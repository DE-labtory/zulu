package ethereum

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type Deriver struct {

}

func (s *Deriver) DeriveAccount(key keychain.Key) (types.Account, error) {
	return types.Account{}, nil
}