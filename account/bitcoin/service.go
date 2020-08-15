package bitcoin

import (
	"github.com/DE-labtory/zulu/account"
	"github.com/DE-labtory/zulu/account/bitcoin/node"
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

func (b *bitcoinType) DeriveAccount(signer account.Signer) (types.Account, error) {
	privkey, err := GetPrivKey(signer)
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

func (b *bitcoinType) Transfer(signer account.Signer, to string, amount string) (types.Transaction, error) {
	acc, err := b.DeriveAccount(signer)
	if err != nil {
		return types.Transaction{}, nil
	}
	// utxo
	_, err = b.adapter.ListUTXO(acc.Address)
	if err != nil {
		return types.Transaction{}, nil
	}

	// fee calculate

	// utxo create 2

	// sign ~ send

	return types.Transaction{}, nil
}
