package ethereum

import (
	"testing"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/stretchr/testify/assert"
)

// Test in Ethereum Ropsten testnet
// address: 0x8361a11De23658fc5b1C49B30858d6cd59D0E3C9
// balance: 0.04 ETH
func TestDeriver_DeriveAccount_eth(t *testing.T) {
	// given
	eth := types.Coin{
		Id: "1",
		Blockchain: types.Blockchain{
			Platform: types.Ethereum,
			Network:  types.Ropsten,
		},
		Symbol:   "ETH",
		Decimals: 18,
		Meta:     nil,
	}
	key := keychain.Key{
		ID: "02b86722365cfe5402076e14a8cbf5c1c0320f0f50cbb531caef2f4049d62195",
		PrivateKey: []byte("��5�}��IN���o!���ɑ|������G "),
		PublicKey: []byte("��u�������\f\n���yi�D�0*�sh4m7%yy�{44���(U�do7��S����������"),
	}

	// when
	deriver := NewDeriver(eth)
	account, err := deriver.DeriveAccount(key)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, "40000000000000000", account.Balance)
}
