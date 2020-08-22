package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	BalanceAt(address string, blockNumber *big.Int) (*big.Int, error)
}

type GethClient struct {
	client *ethclient.Client
}

func (c *GethClient) BalanceAt(address string, blockNumber *big.Int) (*big.Int, error) {
	return nil, nil
}
