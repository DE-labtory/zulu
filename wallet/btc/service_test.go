package btc_test

import (
	"testing"

	"github.com/btcsuite/btcutil"

	"github.com/DE-labtory/zulu/wallet/btc"

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

func (s *HDKeySigner) PrivKey() string {
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		panic(err)
	}
	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}
	return key.String()
}

func (s *HDKeySigner) PubKey() string {
	return ""
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

	w, err := btc.WalletService(types.Testnet)
	if err != nil {
		t.Fatalf("error when creating wallet service: %s", err)
	}
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
