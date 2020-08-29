package bitcoin

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type bitcoinType struct {
	network types.Network
	lister  *TxLister
}

func NewService(network types.Network) *bitcoinType {
	return &bitcoinType{
		network: network,
		lister:  NewTxLister(NewAdapter(network)),
	}
}

func (b *bitcoinType) DeriveAccount(key keychain.Key) (types.Account, error) {
	addr, err := NewAddress(key, b.network)
	if err != nil {
		return types.Account{}, err
	}
	bal, err := b.GetBalance(addr.EncodeAddress())
	if err != nil {
		return types.Account{}, err
	}
	return addr.ToAccount(bal), nil
}

// TODO: implement me
func (b *bitcoinType) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	//addr, err := NewAddress(key, b.network)
	//if err != nil {
	//	return types.Transaction{}, err
	//}

	return types.Transaction{}, nil
}

func (b *bitcoinType) GetInfo() types.Coin {
	return Coin(b.network)
}

// TODO: 중복 삭제
func (b *bitcoinType) GetBalance(addr string) (Amount, error) {
	utxos, err := b.lister.ListUnspent(addr)
	if err != nil {
		return Amount{}, err
	}
	var value int64
	for _, u := range utxos {
		value += u.Value
	}
	return NewAmount(value), nil
}
