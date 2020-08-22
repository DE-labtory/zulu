package keychain

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// Key is essentially a public and private key pair based on ECDSA which uses the secp256k1 curve
type Key struct {
	ID         string
	PrivateKey []byte
	PublicKey  []byte
}

type KeyGenerator struct {
}

// NewKey generates ECDSA key pair with its unique id which derived from its public key
func NewKey(priv *ecdsa.PrivateKey) (key Key, err error) {
	pubWithPrefix := crypto.FromECDSAPub(&priv.PublicKey)
	// Exclude the first byte which indicates if this public key is compressed or uncompressed
	pub := pubWithPrefix[1:65]

	id, err := DeriveID(pub)
	if err != nil {
		return key, err
	}

	return Key{
		ID:         id,
		PrivateKey: crypto.FromECDSA(priv),
		PublicKey:  pub,
	}, nil
}

func NewKeyGenerator() KeyGenerator {
	return KeyGenerator{}
}

// Generate creates new Key instance
func (g *KeyGenerator) Generate() (key Key, err error) {
	keyPair, err := crypto.GenerateKey()
	if err != nil {
		return key, err
	}

	k, err := NewKey(keyPair)
	if err != nil {
		return key, err
	}
	return k, nil
}

// ValidateKey checks if a Key satisfies cryptographic binding
func (k *Key) ValidateKey() error {
	// Validate if id contains only hex string
	if _, err := hex.DecodeString(k.ID); err != nil {
		return err
	}
	if len(k.ID) != 64 {
		return errors.New("length of key id should be 64")
	}
	if len(k.PrivateKey) != 32 {
		return errors.New("length of private key should be 32")
	}
	if len(k.PublicKey) != 64 {
		return errors.New("length of public key should be 64")
	}

	raw := k.PrivateKey
	priv, _ := crypto.ToECDSA(raw)
	pubWithPrefix := crypto.FromECDSAPub(&priv.PublicKey)
	pub := pubWithPrefix[1:65]
	if !bytes.Equal(k.PublicKey, pub) {
		return errors.New("invalid key pair")
	}

	id, err := DeriveID(pub)
	if err != nil {
		return err
	}
	if strings.Compare(k.ID, id) != 0 {
		return errors.New("invalid key id")
	}
	return nil
}

// DeriveID generates unique Key's id from its public key
func DeriveID(publicKey []byte) (id string, err error) {
	if publicKey == nil {
		return id, errors.New("public key is empty")
	}
	id = hex.EncodeToString(crypto.Keccak256(publicKey))
	return id, nil
}
