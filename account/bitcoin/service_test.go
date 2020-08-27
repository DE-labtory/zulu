package bitcoin_test

import (
	"testing"

	"github.com/DE-labtory/zulu/keychain"

	"github.com/btcsuite/btcutil"

	"github.com/DE-labtory/zulu/account/bitcoin"

	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/chaincfg"
)

// https://github.com/bitcoin/bips/blob/master/bip-0137.mediawiki
var testKey = keychain.Key{
	PublicKey: []byte{
		0x02,
		0x62, 0xf2, 0x32, 0x4e, 0x79, 0xb6, 0x78, 0x59,
		0x7b, 0xdd, 0x7b, 0x5a, 0x48, 0xb6, 0x46, 0xdb,
		0xdc, 0x8c, 0x7f, 0x67, 0x28, 0x33, 0xd0, 0xef,
		0x88, 0xc7, 0x5f, 0xad, 0x69, 0xde, 0x55, 0xf0,
	},
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
