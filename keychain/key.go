package keychain

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

// Key is essentially a public and private key pair based on ECDSA which uses the secp256k1 curve
type Key struct {
	ID         string
	PrivateKey []byte
	PublicKey  []byte
}

type KeyGenerator struct {
}

func NewKeyGenerator() KeyGenerator {
	return KeyGenerator{}
}

func (k *KeyGenerator) Generate() Key {
	keyPair, _ := crypto.GenerateKey()
	pub := crypto.FromECDSAPub(&keyPair.PublicKey)
	// Exclude the first byte which indicates if this public key is compressed or uncompressed
	pub = pub[1:65]

	return Key{ID: DeriveID(pub),
		PrivateKey: crypto.FromECDSA(keyPair),
		PublicKey:  pub,
	}
}

func DeriveID(pub []byte) string {
	return hex.EncodeToString(crypto.Keccak256(pub))
}

func ValidateKey(k Key) error {
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
	pub := crypto.FromECDSAPub(&priv.PublicKey)
	if !bytes.Equal(k.PublicKey, pub[1:65]) {
		return errors.New("invalid key pair")
	}

	id := DeriveID(k.PublicKey)
	if strings.Compare(k.ID, id) != 0 {
		return errors.New("invalid key id")
	}
	return nil
}
