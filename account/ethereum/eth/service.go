package eth

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type Service struct{}

func NewService() *Service {
	return nil
}

func (s *Service) DeriveAccount(signer keychain.Signer) (types.Account, error) {
	signer.PubKey()
	return types.Account{}, nil
}

func (s *Service) Transfer(signer keychain.Signer, to string, amount uint) (types.Transaction, error) {
	return types.Transaction{}, nil
}
