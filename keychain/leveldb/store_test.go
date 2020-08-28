package leveldb

import (
	"crypto/ecdsa"
	"os"
	"testing"

	"github.com/DE-labtory/zulu/db/leveldb"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var keys = []keychain.Key{
	{
		ID:         "41201150a0e7e9fbc97c83c5f16fa5c121ec84d65b47e9a5a5594017c2b60fbe",
		PrivateKey: convert(priv[0]),
	},
	{
		ID:         "d69ef2141db922a51d45ccd9821eb282a3cbabc263878a960353a8a98451b1fe",
		PrivateKey: convert(priv[1]),
	},
}

var priv = [][]byte{
	{
		0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
		0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
		0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
		0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
	},
	{
		0x19, 0xe6, 0xa9, 0x0e, 0xcb, 0x53, 0xd9, 0x55,
		0x6e, 0xb0, 0x6e, 0x5e, 0x8a, 0x34, 0x64, 0x7f,
		0x2b, 0x03, 0x08, 0xcf, 0x52, 0x3d, 0x72, 0x59,
		0xe2, 0x5d, 0x4f, 0xd2, 0xc4, 0x2c, 0xa8, 0x2f,
	},
}

func convert(priv []byte) ecdsa.PrivateKey {
	p, err := crypto.ToECDSA(priv)
	if err != nil {
		panic("private key converting error")
	}
	return *p
}

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
	rawData := marshalKey(keys[0])

	err := keyStoreHandle.Put([]byte(keys[0].ID), rawData, true)
	assert.NoError(t, err)

	keyStore := NewKeyStore(keyStoreHandle)

	// when
	retrieveKey, err := keyStore.Get(keys[0].ID)

	// then
	assert.NoError(t, err)
	assert.Equal(t, keys[0], retrieveKey)
}

func TestKeyStore_GetNotFound(t *testing.T) {
	// set
	dbProvider := leveldb.CreateNewDBProvider(TestDbPath)
	defer gracefulTestDown(t, dbProvider)
	keyStoreHandle := dbProvider.GetDBHandle("keystore")

	// given
	rawData := marshalKey(keys[0])

	err := keyStoreHandle.Put([]byte(keys[0].ID), rawData, true)
	assert.NoError(t, err)

	keyStore := NewKeyStore(keyStoreHandle)

	// when
	retrieveKey, err := keyStore.Get(keys[0].ID + "NOT_FOUND")

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
	rawData := marshalKey(keys[0])

	keyStore := NewKeyStore(keyStoreHandle)

	// when
	err := keyStore.Store(keys[0])

	// then
	assert.NoError(t, err)

	d, err := keyStoreHandle.Get([]byte(keys[0].ID))
	assert.NoError(t, err)
	assert.Equal(t, rawData, d)

}

func TestKeyStore_GetAll(t *testing.T) {
	// set
	dbProvider := leveldb.CreateNewDBProvider(TestDbPath)
	defer gracefulTestDown(t, dbProvider)
	keyStoreHandle := dbProvider.GetDBHandle("keystore")

	keyStore := NewKeyStore(keyStoreHandle)
	err := keyStore.Store(keys[0])
	assert.NoError(t, err)

	err = keyStore.Store(keys[1])
	assert.NoError(t, err)

	// when
	loadKeys, err := keyStore.GetAll()

	// then
	assert.NoError(t, err)
	assert.Len(t, loadKeys, 2)

}
