package btc_test

import (
	"testing"

	"github.com/btcsuite/btcutil"

	"github.com/DE-labtory/zulu/coin/btc"

	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/chaincfg"

	"github.com/btcsuite/btcutil/hdkeychain"
)

type HDKeySigner struct {
	pvtKey *hdkeychain.ExtendedKey
}

func NewHDKeySigner(pvtKey *hdkeychain.ExtendedKey) *HDKeySigner {
	if !pvtKey.IsPrivate() {
		panic("provided private key is not private")
	}
	return &HDKeySigner{
		pvtKey: pvtKey,
	}
}

func (s *HDKeySigner) PrivKey() []byte {
	return []byte("tprv8ZgxMBicQKsPeC1AToPuxY8zTgM26qLuqp3tTWwzZpqj5azR9NAoJiAqZeCNm3tudA5pzbAx3Jb4gLzJzCfsSvSnymLRmmUuk7ji1SQxAhs")
}

func (s *HDKeySigner) PubKey() []byte {
	return []byte{}
}

func NewHDPrivateKey() *hdkeychain.ExtendedKey {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		panic(err)
	}
	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}
	return key
}

func TestBitcoinType_DeriveAccount(t *testing.T) {
	signer := NewHDKeySigner(NewHDPrivateKey())

	w, _ := btc.WalletService(types.Testnet)

	result, err := w.DeriveAccount(signer)
	if err != nil {
		t.Fatalf("error when creating account: %s", err)
	}
	if result.Address == "" {
		t.Fatalf("expected address is empty string but got: %v", result.Address)
	}
	addr, err := btcutil.DecodeAddress(result.Address, &chaincfg.TestNet3Params)
	if err != nil {
		t.Fatalf("error when decoding pubkeyhash from string: %s", err)
	}
	if !addr.IsForNet(&chaincfg.TestNet3Params) {
		t.Fatal("error address is not for testnet")
	}

	if result.Coin.Blockchain.Platform != types.Bitcoin {
		t.Fatalf("expected blockchain platform is Bitcoin but got: %v", result.Coin.Blockchain.Platform)
	}
	if result.Balance != "0" {
		t.Fatalf("expected balance is 0 but got: %s", result.Balance)
	}
}

func TestBitcoinType_WhenPrivKeyIsSame_ThenTwoAccountIsEqual(t *testing.T) {
	signer := NewHDKeySigner(NewHDPrivateKey())

	w, _ := btc.WalletService(types.Testnet)

	account1, err := w.DeriveAccount(signer)
	if err != nil {
		t.Fatalf("error when creating account1: %s", err)
	}
	account2, err := w.DeriveAccount(signer)
	if err != nil {
		t.Fatalf("error when creating account1: %s", err)
	}
	if account1.Address != account2.Address {
		t.Fatalf("error two accounts address is not equal: %s, %s", account1.Address, account2.Address)
	}
}
