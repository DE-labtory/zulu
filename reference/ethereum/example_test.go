package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"strings"
	"testing"
)

var defaultAddress = "0x33ffe564A61d48408b5b8Db0c112e7Cc79d023a5"

func loadDefaultPrivateKey() (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA("9ca9700d14db691586ace71b25fe9973f1d2e0dd874e02e3d2d994ea7594f3e6");
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func TestBuildETHTransaction(t *testing.T) {
	privateKey, err := loadDefaultPrivateKey()
	if err != nil {
		t.Fatal(err)
	}
	recipient := common.HexToAddress(defaultAddress)

	tx := &Transaction{
		Nonce:     35947,
		GasPrice:  big.NewInt(30000000000),
		GasLimit:  23000,
		Recipient: &recipient,
		Amount:    big.NewInt(0),
		Payload:   []byte("0x0"),
	}

	chainId := big.NewInt(3)
	encodedTransaction, err := BuildTransaction(tx, chainId)
	if err != nil {
		t.Error(err)
	}

	_, v, r, s, err := Sign(
		encodedTransaction,
		chainId,
		privateKey,
	)
	if err != nil {
		t.Error(err)
	}

	tx.R = r
	tx.S = s
	tx.V = v

	rawTransaction, err := RLPEncode(tx)

	// below is signing with geth guide code
	expectedTx := types.NewTransaction(35947, recipient, big.NewInt(0), 23000, big.NewInt(30000000000), []byte("0x0"))

	expectedSignedTx, err := types.SignTx(expectedTx, types.NewEIP155Signer(chainId), privateKey)
	if err != nil {
		t.Error(err)
	}

	ts := types.Transactions{expectedSignedTx}
	expectedRawTransaction := ts.GetRlp(0)

	if !bytes.Equal(rawTransaction, expectedRawTransaction) {
		t.Error("invalid signing result")
	}
}

func TestERC20Token(t *testing.T) {
	expectedResult := "a9059cbb00000000000000000000000033ffe564a61d48408b5b8db0c112e7cc79d023a5000000000000000000000000000000000000000000000000016345785d8a0000"
	recipient := common.HexToAddress(defaultAddress)
	amount := new(big.Int)
	amount.SetString("100000000000000000", 10)

	data := BuildERC20TransferData(&recipient, amount)
	hexData := hex.EncodeToString(data)

	if !strings.EqualFold(expectedResult, hexData) {
		t.Error("invalid erc20 transfer payload")
	}
}