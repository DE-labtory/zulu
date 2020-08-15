package btc

import (
	"github.com/DE-labtory/zulu/coin"
	"github.com/DE-labtory/zulu/types"
)

type bitcoinType struct {
	network types.Network
}

func WalletService(network types.Network) (*bitcoinType, error) {
	return &bitcoinType{
		network: network,
	}, nil
}

func (b *bitcoinType) DeriveAccount(signer coin.Signer) (types.Account, error) {
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

func (b *bitcoinType) Transfer(signer coin.Signer, to string, amount string) (types.Transaction, error) {
	return types.Transaction{}, nil
}
