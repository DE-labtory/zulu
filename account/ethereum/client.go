package ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type Client interface {
	BalanceAt(address string, blockNumber *big.Int) (*big.Int, error)
}

type GethClient struct {
	client *ethclient.Client
}

func (c *GethClient) BalanceAt(address string, blockNumber *big.Int) (*big.Int, error) {
	return 0, nil
}
