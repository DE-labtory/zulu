package ethereum

import (
	"crypto/ecdsa"
	"github.com/DE-labtory/zulu/types"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type AccountBuilder struct {
}

func (a *AccountBuilder) build(pubKey []byte) (types.Account, error) {
	return types.Account{}, nil
}

type TransactionBuilder struct {
	signer ethTypes.Signer
}

func NewTransactionBuilder(signer ethTypes.Signer) *TransactionBuilder {
	return &TransactionBuilder{
		signer: signer,
	}
}

func (t *TransactionBuilder) build(
	nonce uint64,
	gasPrice uint64,
	gasLimit uint64,
	toAddress string,
	amount uint,
	payload []byte,
	privKey *ecdsa.PrivateKey,
) string {
	return ""
}
