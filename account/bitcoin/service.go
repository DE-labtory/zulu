package bitcoin

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type bitcoinType struct {
	network types.Network
	node    Adapter
}

func NewService(network types.Network) *bitcoinType {
	return &bitcoinType{
		network: network,
		node:    NewAdapter(network),
	}
}

func (b *bitcoinType) DeriveAccount(key keychain.Key) (types.Account, error) {
	addr, err := NewAddress(key, b.network)
	if err != nil {
		return types.Account{}, err
	}
	bal, err := b.Balance(*addr)
	if err != nil {
		return types.Account{}, err
	}
	return addr.ToAccount(bal), nil
}

// TODO: implement me
func (b *bitcoinType) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	return types.Transaction{}, nil
}

func (b *bitcoinType) GetInfo() types.Coin {
	return Coin(b.network)
}

func (b *bitcoinType) Balance(addr Address) (Amount, error) {
	utxos, err := b.node.ListUTXO(addr)
	if err != nil {
		return Amount{}, err
	}
	var value int64
	for _, u := range utxos {
		value += u.Value
	}
	return NewAmount(value), nil
}
