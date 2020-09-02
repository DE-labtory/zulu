package bitcoin_test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/DE-labtory/zulu/keychain"

	"github.com/btcsuite/btcutil"

	"github.com/DE-labtory/zulu/account/bitcoin"

	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/chaincfg"
)

func loadKeychain() keychain.Key {
	// mruZNSFmfGwkwueYsMTWwEtHzUvLMseoVu
	privateKey, err := crypto.HexToECDSA("9ca9700d14db691586ace71b25fe9973f1d2e0dd874e02e3d2d994ea7594f3e6")
	if err != nil {
		panic(err)
	}
	key, err := keychain.NewKey(*privateKey)
	if err != nil {
		panic(err)
	}
	return key
}

func TestBitcoinType_DeriveAccount(t *testing.T) {
	w := bitcoin.NewService(types.Testnet)

	result, err := w.DeriveAccount(loadKeychain())
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
	if result.Balance != "1719736" {
		t.Fatalf("expected balance is 0 but got: %s", result.Balance)
	}
}

func TestBitcoinType_WhenPrivKeyIsSame_ThenTwoAccountIsEqual(t *testing.T) {
	w := bitcoin.NewService(types.Testnet)

	account1, err := w.DeriveAccount(loadKeychain())
	if err != nil {
		t.Fatalf("error when creating account1: %s", err)
	}
	account2, err := w.DeriveAccount(loadKeychain())
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

func TestBitcoinType_Transfer(t *testing.T) {
	w := bitcoin.NewService(types.Testnet)
	toAddr := "muQqyVnEaUPLLco4rDtsKifE2AVyXsStFY"
	amount := "3000"

	tx, err := w.Transfer(loadKeychain(), toAddr, amount)
	if err != nil {
		t.Fatalf("failed to transfer: %v", err)
	}
	fmt.Println(tx)
}
