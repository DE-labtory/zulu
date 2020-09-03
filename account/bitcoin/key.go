package bitcoin

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil"
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
	_, pub := btcec.PrivKeyFromBytes(btcec.S256(), kw.PrivateKey.Serialize())
	return btcutil.Hash160(pub.SerializeCompressed())
	//return elliptic.Marshal(kw.PublicKey.Curve, kw.PublicKey.X, kw.PublicKey.Y)
}
