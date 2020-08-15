package bitcoin_test

import (
	"testing"

	"github.com/DE-labtory/zulu/keychain"

	"github.com/btcsuite/btcutil"

	"github.com/DE-labtory/zulu/account/bitcoin"

	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/chaincfg"
)

var testKey = keychain.Key{
	PrivateKey: []byte("tprv8ZgxMBicQKsPeC1AToPuxY8zTgM26qLuqp3tTWwzZpqj5azR9NAoJiAqZeCNm3tudA5pzbAx3Jb4gLzJzCfsSvSnymLRmmUuk7ji1SQxAhs"),
}

func TestBitcoinType_DeriveAccount(t *testing.T) {
	w := bitcoin.NewService(types.Testnet)

	result, err := w.DeriveAccount(testKey)
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
	w := bitcoin.NewService(types.Testnet)

	account1, err := w.DeriveAccount(testKey)
	if err != nil {
		t.Fatalf("error when creating account1: %s", err)
	}
	account2, err := w.DeriveAccount(testKey)
	if err != nil {
		t.Fatalf("error when creating account1: %s", err)
	}
	if account1.Address != account2.Address {
		t.Fatalf("error two accounts address is not equal: %s, %s", account1.Address, account2.Address)
	}
}

func TestBitcoinType_GetInfo(t *testing.T) {
	w := bitcoin.NewService(types.Testnet)
	info := w.GetInfo()

	if info.Id != "1" {
		t.Fatalf("expected ID is 1 but got: %s", info.Id)
	}
	if info.Blockchain.Platform != types.Bitcoin {
		t.Fatalf("expected Platform is Bitcoin but got: %s", info.Blockchain.Platform)
	}
	if info.Blockchain.Network != types.Testnet {
		t.Fatalf("expected Platform is Testnet but got: %s", info.Blockchain.Network)
	}
	if info.Symbol != types.Btc {
		t.Fatalf("expected Platform is Btc but got: %s", info.Symbol)
	}
	if info.Decimals != bitcoin.Decimal.Int() {
		t.Fatalf("expected decimal is 8 but got: %d", info.Decimals)
	}
}
