package keychain

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

// Key is essentially a public and private key pair based on ECDSA which uses the secp256k1 curve
type Key struct {
	// ID is 64 character long string which is keccack hash of Key's public key
	ID         string
	PrivateKey ecdsa.PrivateKey
}

type KeyGenerator struct {
}

// NewKey generates ECDSA key pair with its unique id which derived from its public key
func NewKey(priv ecdsa.PrivateKey) (key Key, err error) {
	var empty ecdsa.PrivateKey
	if priv == empty {
		return key, errors.New("`Key` generation failed because private key is empty")
	}

	id, _ := DeriveID(priv)

	return Key{
		ID:         id,
		PrivateKey: priv,
	}, nil
}

func (k *Key) GetPrivateKey() []byte {
	return crypto.FromECDSA(&k.PrivateKey)
}

func NewKeyGenerator() KeyGenerator {
	return KeyGenerator{}
}

// Generate creates new Key instance
func (g *KeyGenerator) Generate() (key Key, err error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return key, err
	}

	k, err := NewKey(*priv)
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

	id, err := DeriveID(k.PrivateKey)
	if err != nil {
		return err
	}
	if strings.Compare(k.ID, id) != 0 {
		return errors.New("invalid key id")
	}
	return nil
}

// DeriveID generates unique Key's id from its public key
func DeriveID(priv ecdsa.PrivateKey) (id string, err error) {
	pubKey := crypto.FromECDSAPub(&priv.PublicKey)
	if pubKey == nil {
		return id, errors.New("public key is empty")
	}
	id = hex.EncodeToString(crypto.Keccak256(pubKey))
	return id, nil
}
