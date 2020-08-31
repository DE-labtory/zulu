package klaytn

import (
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/crypto"
	"github.com/klaytn/klaytn/ser/rlp"
)

var (
	errNoFromAddr = errors.New("from addr must be set")
	errNoToAddr   = errors.New("to addr must be set")
)

type TransactionBuilder struct {
	signer      types.Signer
	values      map[types.TxValueKeyType]interface{}
	fromPrvKey  []byte
	feePayerKey []byte
}

func NewTransactionBuilder(signer types.Signer) *TransactionBuilder {
	return &TransactionBuilder{
		signer: signer,
	}
}

func (tb *TransactionBuilder) SetNonce(n uint64) *TransactionBuilder {
	tb.values[types.TxValueKeyNonce] = n
	return tb
}

func (tb *TransactionBuilder) SetAmount(amount int64) *TransactionBuilder {
	tb.values[types.TxValueKeyAmount] = big.NewInt(amount)
	return tb
}
func (tb *TransactionBuilder) SetFrom(addr string, prvKey []byte) *TransactionBuilder {
	tb.values[types.TxValueKeyFrom] = common.HexToAddress(addr)
	tb.fromPrvKey = prvKey
	return tb
}

func (tb *TransactionBuilder) SetTo(addr string) *TransactionBuilder {
	tb.values[types.TxValueKeyTo] = common.HexToAddress(addr)
	return tb
}

func (tb *TransactionBuilder) SetGasLimit(gasLimit uint64) *TransactionBuilder {
	tb.values[types.TxValueKeyGasLimit] = gasLimit
	return tb
}

func (tb *TransactionBuilder) SetGasPrice(gasPrice int64) *TransactionBuilder {
	tb.values[types.TxValueKeyGasPrice] = big.NewInt(gasPrice)
	return tb
}

func (tb *TransactionBuilder) SetFeePayer(addr string, prvKey []byte) *TransactionBuilder {
	tb.values[types.TxValueKeyTo] = common.HexToAddress(addr)
	tb.feePayerKey = prvKey
	return tb
}

func (tb *TransactionBuilder) Build() (string, error) {
	if len(tb.fromPrvKey) == 0 {
		return "", errNoFromAddr
	}

	if _, ok := tb.values[types.TxValueKeyTo]; !ok {
		return "", errNoToAddr
	}

	txType := types.TxTypeValueTransfer
	if _, ok := tb.values[types.TxValueKeyTo]; ok {
		txType = types.TxTypeFeeDelegatedValueTransfer
	}

	tx, err := types.NewTransactionWithMap(txType, tb.values)
	if err != nil {
		return "", err
	}

	fromPrvKey, err := crypto.ToECDSA(tb.fromPrvKey)
	if err != nil {
		return "", err
	}

	err = tx.Sign(tb.signer, fromPrvKey)
	if err != nil {
		return "", err
	}

	if txType == types.TxTypeFeeDelegatedValueTransfer {
		feePayerPrvKey, err := crypto.ToECDSA(tb.feePayerKey)
		if err != nil {
			return "", err
		}
		err = tx.SignFeePayer(tb.signer, feePayerPrvKey)
		if err != nil {
			return "", err
		}
	}

	b, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}
