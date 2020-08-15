package mem

import "github.com/DE-labtory/zulu/keychain"

type MemStore struct {
	keys map[string]keychain.Key
}

func (ks *MemStore) Store(k keychain.Key) {

}

func (ks *MemStore) Get(id string) keychain.Key {
	return keychain.Key{}
}
