package ethereum

import (
	"encoding/hex"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	"log"

	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type AccountBuilder struct {
}

func (a *AccountBuilder) Build(pubKey []byte) (types.Account, error) {
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

func (t *TransactionBuilder) Build(
	nonce uint64,
	gasPrice *big.Int,
	gasLimit uint64,
	toAddress string,
	amount *big.Int,
	payload []byte,
	privateKey []byte,
) (string, error) {
	tx := ethTypes.NewTransaction(
		nonce,
		common.HexToAddress(toAddress),
		amount,
		gasLimit,
		gasPrice,
		payload,
	)

	privKy, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return "", err
	}

	signedTx, err := ethTypes.SignTx(tx, t.signer, privKy)
	if err != nil {
		log.Fatal(err)
	}

	ts := ethTypes.Transactions{signedTx}
	rawTxBytes := ts.GetRlp(0)
	return hex.EncodeToString(rawTxBytes), nil
}
