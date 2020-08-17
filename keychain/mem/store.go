package mem

import (
	"errors"
	"github.com/DE-labtory/zulu/keychain"
)

type MemStore struct {
	keys map[string]keychain.Key
}

func NewMemStore() MemStore {
	return MemStore{
		keys: make(map[string]keychain.Key),
	}
}

func (ks *MemStore) Store(k keychain.Key) error {
	if err := keychain.ValidateKey(k); err != nil {
		return err
	}

	ks.keys[k.ID] = k
	return nil
}

func (ks *MemStore) Get(id string) (keychain.Key, error) {
	k, ok := ks.keys[id]
	if !ok {
		return k, errors.New("key not found error")
	}
	return k, nil
}
