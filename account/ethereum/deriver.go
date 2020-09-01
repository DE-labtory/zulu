package ethereum

import (
	"crypto/ecdsa"
	"errors"
	"strings"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Deriver struct {
	network types.Network
	client  Client
	coin    types.Coin
}

func NewDeriver(coin types.Coin, client Client) *Deriver {
	network := coin.Blockchain.Network
	return &Deriver{
		network: network,
		client:  client,
		coin:    coin,
	}
}

func (d *Deriver) DeriveAccount(key keychain.Key) (types.Account, error) {
	address, err := d.deriveAddress(key)
	if err != nil {
		return types.Account{}, err
	}

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

func (d *Deriver) deriveAddress(key keychain.Key) (string, error) {
	publicKey := key.PrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).String()
	return strings.ToLower(address), nil
}
