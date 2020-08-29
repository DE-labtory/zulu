package keychain

import (
	"crypto/ecdsa"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	// when
	g := NewKeyGenerator()
	key, err := g.Generate()
	assert.NoError(t, err)

	err = key.ValidateKey()
	assert.NoError(t, err)
}

func TestGetPrivateKey(t *testing.T) {
	// when
	g := NewKeyGenerator()
	key, err := g.Generate()
	assert.NoError(t, err)

	binaryDump := key.GetPrivateKey()

	priv, err := crypto.ToECDSA(binaryDump)
	assert.NoError(t, err)
	assert.Equal(t, key.PrivateKey, *priv)
}

func TestDeriveID(t *testing.T) {
	tests := []struct {
		rawPriv []byte
		ID      string
	}{
		{
			rawPriv: []byte{
				0x4c, 0x0c, 0x68, 0x33, 0x88, 0x35, 0x94, 0xab,
				0x03, 0x45, 0xeb, 0x13, 0x53, 0x93, 0xb1, 0x64,
				0x9e, 0x80, 0xfb, 0x8a, 0xa6, 0xfe, 0x96, 0x72,
				0x1e, 0xbf, 0x6f, 0x24, 0xf4, 0x54, 0xe0, 0xcc,
			},
			ID: "41201150a0e7e9fbc97c83c5f16fa5c121ec84d65b47e9a5a5594017c2b60fbe",
		},
	}

	for _, test := range tests {
		priv, err := crypto.ToECDSA(test.rawPriv)
		assert.NoError(t, err)

		ID, err := DeriveID(*priv)
		assert.NoError(t, err)
		assert.Equal(t, ID, test.ID)
	}
}

func TestNewKey_EmptyPrivateKey(t *testing.T) {
	_, err := NewKey(ecdsa.PrivateKey{})
	assert.Error(t, err)
}
