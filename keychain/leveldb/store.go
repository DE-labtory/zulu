package leveldb

import (
	"encoding/json"
	"errors"

	"github.com/DE-labtory/zulu/db/leveldb"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/sirupsen/logrus"
)

func NewKeyStore(handle *leveldb.DBHandle) *KeyStore {
	return &KeyStore{handle: handle}
}

type KeyStore struct {
	handle *leveldb.DBHandle
}

func (ks *KeyStore) Store(k keychain.Key) error {
	rawData, err := json.Marshal(k)
	if err != nil {
		return err
	}
	err = ks.handle.Put([]byte(k.ID), rawData, true)
	if err != nil {
		return err
	}
	return nil
}

func (ks *KeyStore) Get(id string) (keychain.Key, error) {
	var k keychain.Key
	rawData, err := ks.handle.Get([]byte(id))

	// check leveldb error
	if err != nil {
		logrus.Error("error while find key :" + id + ", in level db key store")
		return k, err
	}

	// check empty data
	if rawData == nil {
		return k, errors.New("key not found")
	}

	// check unmarshal data
	err = json.Unmarshal(rawData, &k)
	if err != nil {
		logrus.Error("error while unmarshal key :" + id + ", in level db key store")
		return k, err
	}

	// with no err
	return k, nil
}
