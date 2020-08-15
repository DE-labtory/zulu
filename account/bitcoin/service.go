package bitcoin

import (
	"github.com/DE-labtory/zulu/account/bitcoin/node"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type bitcoinType struct {
	network types.Network
	adapter node.Adapter
}

func WalletService(network types.Network) *bitcoinType {
	return &bitcoinType{
		network: network,
	}
}

func (b *bitcoinType) DeriveAccount(key keychain.Key) (types.Account, error) {
	privkey, err := GetPrivKey(key)
	if err != nil {
		return types.Account{}, err
	}
	addr, err := privkey.GetAddress(b.network)
	if err != nil {
		return types.Account{}, err
	}
	return types.Account{
		Address: addr.EncodeAddress(),
		Coin:    Coin(b.network),
		Balance: AmountZero.ToDecimal(),
	}, nil
}

// TODO: implement me
func (b *bitcoinType) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	acc, err := b.DeriveAccount(key)
	if err != nil {
		return types.Transaction{}, nil
	}
	// utxo
	_, err = b.adapter.ListUTXO(acc.Address)
	if err != nil {
		return types.Transaction{}, nil
	}
	return types.Transaction{}, nil
}

func (b *bitcoinType) GetInfo() types.Coin {
	return Coin(b.network)
}
