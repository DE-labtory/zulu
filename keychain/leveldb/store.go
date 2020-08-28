package leveldb

import (
	"encoding/hex"
	"errors"

	"github.com/DE-labtory/zulu/db/leveldb"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

func NewKeyStore(handle *leveldb.DBHandle) *KeyStore {
	return &KeyStore{handle: handle}
}

type KeyStore struct {
	handle *leveldb.DBHandle
}

func (ks *KeyStore) Store(k keychain.Key) error {
	rawData := marshalKey(k)

	err := ks.handle.Put([]byte(k.ID), rawData, true)
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
	k, err = unmarshalKey(rawData)
	if err != nil {
		logrus.Error("error while unmarshal key :" + id + ", in level db key store")
		return k, err
	}

	// with no err
	return k, nil
}

func (ks *KeyStore) GetAll() ([]keychain.Key, error) {
	keys := make([]keychain.Key, 0)
	iter := ks.handle.GetIteratorWithPrefix()
	for iter.Next() {
		key := iter.Key()
		rawData := iter.Value()

		k, err := unmarshalKey(rawData)
		if err != nil {
			logrus.Error("error while unmarshal key :" + string(key) + ", in level db key store")
			continue
		}

		keys = append(keys, k)
	}

	return keys, nil
}

func marshalKey(k keychain.Key) []byte {
	return crypto.FromECDSA(&k.PrivateKey)
}

func unmarshalKey(b []byte) (k keychain.Key, err error) {
	privKey, err := crypto.ToECDSA(b)
	if err != nil {
		logrus.Error("error while parsing private key :" + hex.EncodeToString(b) + ", in level db key store")
		return k, err
	}
	id, _ := keychain.DeriveID(*privKey)

	return keychain.Key{
		ID:         id,
		PrivateKey: *privKey,
	}, nil
}
