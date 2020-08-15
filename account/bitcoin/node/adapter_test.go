package node_test

import (
	"fmt"
	"testing"

	"github.com/DE-labtory/zulu/account/bitcoin/node"
)

func TestHttpClient_ListUtxo(t *testing.T) {
	c := node.NewHttpClient("https://blockstream.info/testnet/api")
	utxos, err := c.ListUTXO("muQqyVnEaUPLLco4rDtsKifE2AVyXsStFY")
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

func TestHttpClient_GetFeeEstimates(t *testing.T) {
	c := node.NewHttpClient("https://blockstream.info/testnet/api")
	estimates, err := c.GetFeeEstimates()
	if err != nil {
		t.Fatalf("error when get estimates: %s", err)
	}
	fmt.Println(estimates)
}
