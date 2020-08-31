package klaytn

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/client"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/ser/rlp"
)

type Client interface {
	BalanceAt(address string) (*big.Int, error)
	NonceAt(address string) (uint64, error)
	SendTransaction(rawTransaction string) (string, error)
	GasPrice() (*big.Int, error)
}

type KenClient struct {
	client *client.Client
}

func (k *KenClient) BalanceAt(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := k.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return &big.Int{}, err
	}

	return balance, nil
}

func (k *KenClient) NonceAt(address string) (uint64, error) {
	account := common.HexToAddress(address)
	return k.client.NonceAt(context.Background(), account, nil)
}

func (k *KenClient) SendTransaction(rawTransaction string) (string, error) {
	rawTxBytes, err := hex.DecodeString(rawTransaction)

	tx := new(types.Transaction)
	err = rlp.DecodeBytes(rawTxBytes, &tx)
	if err != nil {
		return "", err
	}

	err = k.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (k *KenClient) GasPrice() (*big.Int, error) {
	return k.client.SuggestGasPrice(context.Background())
}
