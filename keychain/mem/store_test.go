package mem

import (
	"crypto/ecdsa"
	"testing"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

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

func TestStore(t *testing.T) {
	ks := NewKeyStore()

	for _, k := range keys {
		err := ks.Store(k)
		assert.NoError(t, err)

		assert.Equal(t, k, ks.keys[k.ID])
	}
}

func TestGet(t *testing.T) {
	ks := NewKeyStore()

	for _, k := range keys {
		ks.keys[k.ID] = k

		storedKey, err := ks.Get(k.ID)
		assert.NoError(t, err)

		assert.Equal(t, storedKey, k)
	}
}

func TestGet_NotFound(t *testing.T) {
	ks := NewKeyStore()

	k, err := ks.Get("not found")
	assert.Error(t, err)
	assert.Equal(t, keychain.Key{}, k)
}

func TestGetAll(t *testing.T) {
	ks := NewKeyStore()
	var expected []keychain.Key

	// case: empty key store
	result, err := ks.GetAll()
	assert.NoError(t, err)
	assert.ElementsMatch(t, result, expected)

	// case: key store has 1 key-value pair
	err = ks.Store(keys[0])
	expected = append(expected, keys[0])
	assert.NoError(t, err)

	result, err = ks.GetAll()
	assert.NoError(t, err)
	assert.ElementsMatch(t, result, expected)

	// case: key store has 2 key-value pair
	err = ks.Store(keys[1])
	expected = append(expected, keys[1])
	assert.NoError(t, err)

	result, err = ks.GetAll()
	assert.NoError(t, err)
	assert.ElementsMatch(t, result, expected)
}
