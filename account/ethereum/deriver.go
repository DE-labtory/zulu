package ethereum

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Deriver struct {
	network types.Network
	client  Client
	coin    types.Coin
}

func NewDeriver(coin types.Coin) *Deriver {
	network := coin.Blockchain.Network
	return &Deriver{
		network: network,
		client:  NewGethClient(network),
		coin:    coin,
	}
}

func (d *Deriver) DeriveAccount(key keychain.Key) (types.Account, error) {
	address := d.deriveAddress(key)

	balance, err := d.client.BalanceAt(address)
	if err != nil {
		return types.Account{}, err
	}

	return types.Account{
		Address: address,
		Coin:    d.coin,
		Balance: balance.String(),
	}, nil
}

func (d *Deriver) deriveAddress(key keychain.Key) string {
	pubkeyBytes := key.PublicKey
	keyHash := crypto.Keccak256(pubkeyBytes[1:])
	addressBytes := common.BytesToAddress(keyHash[12:])
	return addressBytes.String()
}
