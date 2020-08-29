package bitcoin

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/btcsuite/btcd/btcec"
)

type KeyWrapper struct {
	*btcec.PrivateKey
}

func NewKeyWrapperFromKeychain(key keychain.Key) *KeyWrapper {
	return nil
}

func NewKeyWrapper(bytes []byte) *KeyWrapper {
	pvt, _ := btcec.PrivKeyFromBytes(btcec.S256(), bytes)
	return &KeyWrapper{
		PrivateKey: pvt,
	}
}
