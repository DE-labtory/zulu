package bitcoin

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type bitcoinType struct {
	network   types.Network
	txService *TxService
}

func NewService(network types.Network) *bitcoinType {
	return &bitcoinType{
		network:   network,
		txService: NewTxService(network, NewAdapter(network)),
	}
}

func (b *bitcoinType) DeriveAccount(key keychain.Key) (types.Account, error) {
	addr, err := DeriveAddress(key, b.network)
	if err != nil {
		return types.Account{}, err
	}
	unspents, err := b.txService.ListUnspent(addr.EncodeAddress())
	if err != nil {
		return types.Account{}, err
	}
	return addr.ToAccount(unspents.Balance()), nil
}

// TODO: implement me
func (b *bitcoinType) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	return types.Transaction{}, nil
}

func (b *bitcoinType) GetInfo() types.Coin {
	return Coin(b.network)
}
