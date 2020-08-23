package mem

import (
	"errors"

	"github.com/DE-labtory/zulu/keychain"
)

type KeyStore struct {
	keys map[string]keychain.Key
}

func NewKeyStore() KeyStore {
	return KeyStore{
		keys: make(map[string]keychain.Key),
	}
}

func (ks *KeyStore) Store(k keychain.Key) error {
	if err := k.ValidateKey(); err != nil {
		return err
	}

	ks.keys[k.ID] = k
	return nil
}

func (ks *KeyStore) Get(id string) (keychain.Key, error) {
	k, ok := ks.keys[id]
	if !ok {
		return k, errors.New("key not found error")
	}
	return k, nil
}

func (ks *KeyStore) GetAll() ([]keychain.Key, error) {
	var keys []keychain.Key

	for _, v := range ks.keys {
		keys = append(keys, v)
	}
	return keys, nil
}
