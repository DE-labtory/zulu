package keychain

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	// when
	g := NewKeyGenerator()
	key := g.Generate()
	msg := crypto.Keccak256([]byte("foo"))

	// then
	priv, err := crypto.ToECDSA(key.PrivateKey)
	if err != nil {
		t.Errorf("private key is invalid: %v", err)
	}

	sig, err := crypto.Sign(msg, priv)
	if err != nil {
		t.Errorf("sign error: %v", err)
	}

	pub1, err := crypto.SigToPub(msg, sig)
	if err != nil {
		t.Errorf("derive public key from signatrue error: %v", err)
	}
	pub2 := crypto.FromECDSAPub(pub1)

	// result
	assert.Equal(t, pub2[1:65], key.PublicKey)

}

func TestDeriveID(t *testing.T) {
	tests := []struct {
		pub []byte
		ID  string
	}{
		{
			pub: []byte{
				0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
				0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
				0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
				0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
				0xa3, 0xe0, 0xf0, 0xc8, 0x14, 0xc8, 0xa1, 0x30,
				0xcb, 0x59, 0xb6, 0x59, 0xd9, 0x48, 0xd7, 0x95,
				0x4c, 0x33, 0xd3, 0xf2, 0x3a, 0x68, 0x59, 0x9a,
				0xad, 0x02, 0x3b, 0x83, 0xd6, 0x61, 0x45, 0xa6,
			},
			ID: "28762cb4592cbe44f993032bb1656b201aae9b9a9b730f8f1c37e5c597cbd8cc",
		},
	}

	for _, test := range tests {
		ID := DeriveID(test.pub)
		assert.Equal(t, ID, test.ID)
	}
}
