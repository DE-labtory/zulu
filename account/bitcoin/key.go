package bitcoin

import (
	"crypto/elliptic"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/btcsuite/btcd/btcec"
)

type KeyWrapper struct {
	*btcec.PrivateKey
}

func NewKeyWrapper(key keychain.Key) *KeyWrapper {
	return &KeyWrapper{
		PrivateKey: (*btcec.PrivateKey)(&key.PrivateKey),
	}
}

func (kw *KeyWrapper) MarshalPubKey() []byte {
	return elliptic.Marshal(kw.PublicKey.Curve, kw.PublicKey.X, kw.PublicKey.Y)
}
