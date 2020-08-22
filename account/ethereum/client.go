package ethereum

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

type Client interface {
	BalanceAt(address string) (*big.Int, error)
	NonceAt(address string) (uint64, error)
	SendTransaction(rawTransaction string) (string, error)
	SuggestGasPrice() (*big.Int, error)
}

type GethClient struct {
	client *ethclient.Client
}

func NewGethClient(network types.Network) *GethClient {
	client, err := ethclient.Dial(Supplier[network].NodeUrl)
	if err != nil {
		panic(err)
	}

	return &GethClient{
		client: client,
	}
}

func (c *GethClient) BalanceAt(address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return &big.Int{}, err
	}

	return balance, nil
}

func (c *GethClient) NonceAt(address string) (uint64, error) {
	account := common.HexToAddress(address)
	return c.client.NonceAt(context.Background(), account, nil)
}

func (c *GethClient) SendTransaction(rawTransaction string) (string, error) {
	rawTxBytes, err := hex.DecodeString(rawTransaction)

	tx := new(ethTypes.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = c.client.SendTransaction(context.Background(), tx)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (c *GethClient) SuggestGasPrice() (*big.Int, error) {
	return c.client.SuggestGasPrice(context.Background())
}
