package bitcoin_test

import (
	"testing"

	"github.com/DE-labtory/zulu/account/bitcoin"
	"github.com/DE-labtory/zulu/types"
)

const testAddr = "mruZNSFmfGwkwueYsMTWwEtHzUvLMseoVu"

func TestTxService_Create(t *testing.T) {
	svc := bitcoin.NewTxService(types.Testnet, bitcoin.NewAdapter(types.Testnet))
	addr, _ := bitcoin.ParseAddressStr(testAddr, types.Testnet)

	txData, err := svc.Create(addr, addr, bitcoin.NewAmount(3000))
	if err != nil {
		t.Fatalf("failed to create txData: %v", err)
	}
	if len(txData.UnspentList.Items) == 0 {
		t.Fatalf("expected UTXO is more than 0")
	}
}
