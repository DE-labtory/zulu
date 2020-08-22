package eth

import (
	"crypto/ecdsa"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func loadDefaultPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA("9ca9700d14db691586ace71b25fe9973f1d2e0dd874e02e3d2d994ea7594f3e6")
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func TestService_Transfer(t *testing.T) {
	service := NewService(types.Ropsten)

	privKey, err := loadDefaultPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	key, err := keychain.NewKey(privKey)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := service.Transfer(key, "0x33ffe564A61d48408b5b8Db0c112e7Cc79d023a5", "0.000001")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("transaction hash %s", tx.TxHash)
}