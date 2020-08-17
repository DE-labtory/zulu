package eth

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type Service struct{}

func NewService() *Service {
	return nil
}

func (s *Service) DeriveAccount(key keychain.Key) (types.Account, error) {
	return types.Account{}, nil
}

func (s *Service) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	return types.Transaction{}, nil
}

func (s *Service) GetInfo() types.Coin {
	return types.Coin{}
}
