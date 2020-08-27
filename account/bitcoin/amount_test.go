package bitcoin_test

import (
	"testing"

	"github.com/DE-labtory/zulu/account/bitcoin"
)

func TestAmount(t *testing.T) {
	amountTests := []struct {
		str     string
		decimal string
		hex     string
	}{
		//{float: 1.23456789, decimal: "123456789", hex: "0x75bcd15"},
		{str: "123456000", decimal: "123456000", hex: "0x75bca00"},
		{str: "123456", decimal: "123456", hex: "0x1e240"},
		{str: "0123", decimal: "123", hex: "0x7b"},
	}
	for _, tt := range amountTests {
		amount, err := bitcoin.ParseAmount(tt.str)
		if err != nil {
			t.Fatalf("error create amount: %v", err)
		}
		if amount.ToDecimal() != tt.decimal {
			t.Fatalf("expected decimal string result is '%s', but got %s", tt.decimal, amount.ToDecimal())
		}
		if amount.ToHex() != tt.hex {
			t.Fatalf("expected hex string result is '%s', but got %s", tt.hex, amount.ToHex())
		}
	}
}
