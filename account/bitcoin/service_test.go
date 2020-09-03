package bitcoin_test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec"

	"github.com/DE-labtory/zulu/keychain"

	"github.com/btcsuite/btcutil"

	"github.com/DE-labtory/zulu/account/bitcoin"

	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/chaincfg"
)

func loadKeychain() keychain.Key {
	// momXhvmA324DdWhZC9TqFqNd9C7qszBKyn
	keyBytes, _ := hex.DecodeString("eeb1c9cb82fa9d81008847259e7239fcae3031fea4cccc224eab3e4c009de161")
	pvt, _ := btcec.PrivKeyFromBytes(btcec.S256(), keyBytes)

	key, err := keychain.NewKey((ecdsa.PrivateKey)(*pvt))
	if err != nil {
		panic(err)
	}
	return key
}

func TestBitcoinType_DeriveAccount(t *testing.T) {
	w := bitcoin.NewService(types.Coin{
		Id: "2",
		Blockchain: types.Blockchain{
			Platform: types.Bitcoin,
			Network:  types.Testnet,
		},
		Symbol:   types.Tbtc,
		Decimals: bitcoin.Decimal.Int(),
	},
		types.Testnet)

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
	if result.Balance == "0" {
		t.Fatalf("expected balance is not 0 but got: %s", result.Balance)
	}
}

func TestBitcoinType_WhenPrivKeyIsSame_ThenTwoAccountIsEqual(t *testing.T) {
	w := bitcoin.NewService(types.Coin{
		Id: "2",
		Blockchain: types.Blockchain{
			Platform: types.Bitcoin,
			Network:  types.Testnet,
		},
		Symbol:   types.Tbtc,
		Decimals: bitcoin.Decimal.Int(),
	},
		types.Testnet)

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
	w := bitcoin.NewService(types.Coin{
		Id: "1",
		Blockchain: types.Blockchain{
			Platform: types.Bitcoin,
			Network:  types.Mainnet,
		},
		Symbol:   types.Btc,
		Decimals: bitcoin.Decimal.Int(),
	},
		types.Mainnet)
	info := w.GetInfo()

	if info.Id != "1" {
		t.Fatalf("expected ID is 1 but got: %s", info.Id)
	}
	if info.Blockchain.Platform != types.Bitcoin {
		t.Fatalf("expected Platform is Bitcoin but got: %s", info.Blockchain.Platform)
	}
	if info.Blockchain.Network != types.Mainnet {
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
	w := bitcoin.NewService(types.Coin{
		Id: "2",
		Blockchain: types.Blockchain{
			Platform: types.Bitcoin,
			Network:  types.Testnet,
		},
		Symbol:   types.Tbtc,
		Decimals: bitcoin.Decimal.Int(),
	},
		types.Testnet)
	toAddr := "muQqyVnEaUPLLco4rDtsKifE2AVyXsStFY"
	amount := "1"

	tx, err := w.Transfer(loadKeychain(), toAddr, amount)
	if err != nil {
		t.Fatalf("failed to transfer: %v", err)
	}
	fmt.Println(tx)
}
