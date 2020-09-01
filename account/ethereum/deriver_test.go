package ethereum

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

type MockClient struct{}

func (c *MockClient) BalanceAt(address string) (*big.Int, error) {
	return big.NewInt(40000000000000000), nil
}

func (c *MockClient) NonceAt(address string) (uint64, error) {
	return 0, nil
}

func (c *MockClient) SendTransaction(rawTransaction string) (string, error) {
	return "transactionHash", nil
}

func (c *MockClient) SuggestGasPrice() (*big.Int, error) {
	return big.NewInt(0), nil
}

func loadDefaultPrivateKey(t *testing.T) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA("9ca9700d14db691586ace71b25fe9973f1d2e0dd874e02e3d2d994ea7594f3e6")
	if err != nil {
		t.Fatal(err)
	}
	return privateKey
}

func TestDeriver_DeriveAccount_Eth_Mock(t *testing.T) {
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
	client := &MockClient{}
	deriver := NewDeriver(eth, client)

	privateKey := loadDefaultPrivateKey(t)
	key, err := keychain.NewKey(*privateKey)
	if err != nil {
		t.Fatal(err)
	}

	// when
	account, err := deriver.DeriveAccount(key)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, eth, account.Coin)
	assert.Equal(t, "40000000000000000", account.Balance)
}

func TestDeriver_DeriveAccount_Eth_Geth(t *testing.T) {
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
	ecdsaPrivateKey, _ := crypto.HexToECDSA("fdfbe5c027ab7a0e095d716a47dd4f64742c081c8bd159934d8949a88687c1cc")
	key, _ := keychain.NewKey(*ecdsaPrivateKey)
	ethClient := NewGethClient(eth.Blockchain.Network)
	deriver := NewDeriver(eth, ethClient)

	// when
	account, err := deriver.DeriveAccount(key)

	// then
	assert.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, eth, account.Coin)
	assert.Equal(t, "0xcdf2d74c3eaab2b1a38f3f13fe7080e8a90eae0a", account.Address)
	assert.Equal(t, "12345670000000000", account.Balance)
}
