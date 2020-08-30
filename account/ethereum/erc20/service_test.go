package erc20

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type MockClient struct {
}

func (c *MockClient) BalanceAt(address string) (*big.Int, error) {
	return big.NewInt(0), nil
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

func loadDefaultPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA("9ca9700d14db691586ace71b25fe9973f1d2e0dd874e02e3d2d994ea7594f3e6")
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func evt() types.Coin {
	return types.Coin{
		Id: "1",
		Blockchain: types.Blockchain{
			Platform: types.Ethereum,
			Network:  types.Ropsten,
		},
		Symbol:   "EVT",
		Decimals: 18,
		Meta:     nil,
	}
}

func newEvtService() *Service {
	return NewService(
		evt(),
		&MockClient{},
		"0x0d4c27c49906208fbd9a9f3a43c63ccbd089f3bf", // EVT
	)
}

func TestService_Transfer(t *testing.T) {
	service := newEvtService()

	privKey, err := loadDefaultPrivateKey()
	if err != nil {
		t.Fatal(err)
	}

	key, err := keychain.NewKey(*privKey)
	if err != nil {
		t.Fatal(err)
	}

	tx, err := service.Transfer(key, "0x33ffe564A61d48408b5b8Db0c112e7Cc79d023a5", "0.00001")
	assert.NoError(t, err)
	assert.Equal(t, "transactionHash", tx.TxHash)
	t.Log(tx.TxHash)
}

func TestService_GetInfo(t *testing.T) {
	// given
	service := newEvtService()

	// when
	info := service.GetInfo()

	// then
	assert.Equal(t, evt(), info)
}
