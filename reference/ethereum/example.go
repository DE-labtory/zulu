package main

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

var transferMethodID = []byte{0xa9, 0x05, 0x9c, 0xbb} // 0xa9059cbb

type Transaction struct {
	Nonce     uint64          `json:"nonce"    gencodec:"required"`
	GasPrice  *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit  uint64          `json:"gas"      gencodec:"required"`
	Recipient *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount    *big.Int        `json:"value"    gencodec:"required"`
	Payload   []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`
}

func BuildERC20TransferData(toAddress *common.Address, amount *big.Int) []byte {
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, transferMethodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	return data
}

func Sign(digestHash []byte, chainId *big.Int, prv *ecdsa.PrivateKey) (signature []byte, V *big.Int, R *big.Int, S *big.Int, err error) {
	signature, err = crypto.Sign(digestHash[:], prv)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])
	v := new(big.Int).SetBytes([]byte{signature[64] + 35})

	v = v.Add(v, big.NewInt(chainId.Int64()*2))

	return signature, v, r, s, nil
}

func BuildTransaction(transaction interface{}, metadata interface{}) ([]byte, error) {
	tx, ok := transaction.(*Transaction)
	if !ok {
		return nil, errors.New("invalid ethereum transaction format")
	}

	chainId, ok := metadata.(*big.Int)
	if !ok {
		return nil, errors.New("invalid ethereum chainId format")
	}

	var encodedTransaction bytes.Buffer
	rlp.Encode(&encodedTransaction, []interface{}{
		tx.Nonce,
		tx.GasPrice,
		tx.GasLimit,
		tx.Recipient,
		tx.Amount,
		tx.Payload,
		chainId, uint(0), uint(0)})
	return crypto.Keccak256(encodedTransaction.Bytes()), nil
}

func BuildMessage(data string) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func RLPEncode(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	err := rlp.Encode(&buffer, data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
