package leveldb

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/DE-labtory/zulu/db/leveldb"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/stretchr/testify/assert"
)

const TestDbPath = "./test_db_path"

var gracefulTestDown = func(t *testing.T, provider *leveldb.DBProvider) {
	provider.Close()
	err := os.RemoveAll(TestDbPath)
	if err != nil {
		t.Log(err.Error())
	}
}

func TestKeyStore_Get(t *testing.T) {
	// set
	dbProvider := leveldb.CreateNewDBProvider(TestDbPath)
	defer gracefulTestDown(t, dbProvider)
	keyStoreHandle := dbProvider.GetDBHandle("keystore")

	// given
	keyId := "testKeyId"
	keyData := keychain.Key{
		ID:         keyId,
		PrivateKey: []byte("testPrivateKey"),
		PublicKey:  []byte("testPublicKey"),
	}
	rawData, err := json.Marshal(keyData)
	assert.NoError(t, err)

	err = keyStoreHandle.Put([]byte(keyId), rawData, true)
	assert.NoError(t, err)

	keyStore := NewKeyStore(keyStoreHandle)

	// when
	retrieveKey, err := keyStore.Get(keyId)

	// then
	assert.NoError(t, err)
	assert.Equal(t, keyData, retrieveKey)
}

func TestKeyStore_GetNotFound(t *testing.T) {
	// set
	dbProvider := leveldb.CreateNewDBProvider(TestDbPath)
	defer gracefulTestDown(t, dbProvider)
	keyStoreHandle := dbProvider.GetDBHandle("keystore")

	// given
	keyId := "testKeyId"
	keyData := keychain.Key{
		ID:         keyId,
		PrivateKey: []byte("testPrivateKey"),
		PublicKey:  []byte("testPublicKey"),
	}
	rawData, err := json.Marshal(keyData)
	assert.NoError(t, err)

	err = keyStoreHandle.Put([]byte(keyId), rawData, true)
	assert.NoError(t, err)

	keyStore := NewKeyStore(keyStoreHandle)

	// when
	retrieveKey, err := keyStore.Get(keyId + "NOT_FOUND")

	// then
	assert.Error(t, err)
	assert.Equal(t, keychain.Key{}, retrieveKey)
}

func TestKeyStore_Store(t *testing.T) {
	// set
	dbProvider := leveldb.CreateNewDBProvider(TestDbPath)
	defer gracefulTestDown(t, dbProvider)
	keyStoreHandle := dbProvider.GetDBHandle("keystore")

	// given
	keyId := "testKeyId"
	keyData := keychain.Key{
		ID:         keyId,
		PrivateKey: []byte("testPrivateKey"),
		PublicKey:  []byte("testPublicKey"),
	}
	rawData, err := json.Marshal(keyData)
	assert.NoError(t, err)

	keyStore := NewKeyStore(keyStoreHandle)

	// when
	err = keyStore.Store(keyData)

	// then
	assert.NoError(t, err)

	d, err := keyStoreHandle.Get([]byte(keyId))
	assert.NoError(t, err)
	assert.Equal(t, rawData, d)

}

func TestKeyStore_GetAll(t *testing.T) {
	// set
	dbProvider := leveldb.CreateNewDBProvider(TestDbPath)
	defer gracefulTestDown(t, dbProvider)
	keyStoreHandle := dbProvider.GetDBHandle("keystore")

	// given
	keyId1 := "testKeyId1"
	keyData1 := keychain.Key{
		ID:         keyId1,
		PrivateKey: []byte("testPrivateKey1"),
		PublicKey:  []byte("testPublicKey1"),
	}

	keyId2 := "testKeyId2"
	keyData2 := keychain.Key{
		ID:         keyId2,
		PrivateKey: []byte("testPrivateKey2"),
		PublicKey:  []byte("testPublicKey2"),
	}

	keyStore := NewKeyStore(keyStoreHandle)
	err := keyStore.Store(keyData1)
	assert.NoError(t, err)

	err = keyStore.Store(keyData2)
	assert.NoError(t, err)

	// when
	loadKeys, err := keyStore.GetAll()

	// then
	assert.NoError(t, err)
	assert.Len(t, loadKeys, 2)

}
