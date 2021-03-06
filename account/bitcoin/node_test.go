package bitcoin_test

import (
	"testing"

	"github.com/DE-labtory/zulu/account/bitcoin"
	"github.com/DE-labtory/zulu/types"
)

func TestAdapter_ListUTXO(t *testing.T) {
	a := bitcoin.NewAdapter(types.Testnet)
	utxos, err := a.ListUTXO("muQqyVnEaUPLLco4rDtsKifE2AVyXsStFY")
	if err != nil {
		t.Fatalf("error when list utxos: %s", err)
	}
	if len(utxos) == 0 {
		t.Fatalf("expected UTXO length is more than 0, but got %d", len(utxos))
	}
	if utxos[0].Value == 0 {
		t.Fatalf("expected UTXO value is not 0, but got %d", utxos[0].Value)
	}
	if utxos[0].Txid == "" {
		t.Fatalf("expected UTXO txid is not empty string, but got %s", utxos[0].Txid)
	}
}

func TestAdapter_EstimateFees(t *testing.T) {
	a := bitcoin.NewAdapter(types.Testnet)
	estimate, err := a.EstimateFeeRate()
	if err != nil {
		t.Fatalf("error when get estimates: %s", err)
	}
	if estimate == 0 {
		t.Fatalf("expected fee estimates is NOT 0, but got %f", estimate)
	}
}
