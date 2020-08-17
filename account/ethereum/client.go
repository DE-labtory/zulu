package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Client interface {
	BalanceAt() (*big.Int, error)
}

type GethClient struct {
	client *ethclient.Client
}

func (c *GethClient) BalanceAt(address string, blockNumber *big.Int) (*big.Int, error) {
	return c.client.BalanceAt(context.Background(), common.HexToAddress(address), blockNumber)
}