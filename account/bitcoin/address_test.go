package bitcoin_test

import (
	"testing"

	"github.com/DE-labtory/zulu/account/bitcoin/chaincfg"

	"github.com/DE-labtory/zulu/account/bitcoin"
	"github.com/DE-labtory/zulu/types"
)

func TestNewAddress(t *testing.T) {
	addr, err := bitcoin.NewAddress(testKey, types.Testnet)
	if err != nil {
		t.Fatalf("failed to create new address: %v", err)
	}
	if !addr.IsForNet(chaincfg.Supplier[types.Testnet].Spec) {
		t.Fatalf("error created address is not testnet format")
	}
}

func TestAddress_WhenPubKeySame(t *testing.T) {
	addr, _ := bitcoin.NewAddress(testKey, types.Testnet)
	addr2, _ := bitcoin.NewAddress(testKey, types.Testnet)
	if addr.EncodeAddress() != addr2.EncodeAddress() {
		t.Fatalf("two address with same pub key is different: %s, %s", addr.EncodeAddress(), addr2.EncodeAddress())
	}
}
