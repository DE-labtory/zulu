package bitcoin

import (
	"github.com/DE-labtory/zulu/account/bitcoin/chaincfg"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcutil"
)

type Address struct {
	*btcutil.AddressPubKeyHash
	Network types.Network
}

func NewAddress(key keychain.Key, network types.Network) (*Address, error) {
	addrPk, err := btcutil.NewAddressPubKey(
		key.PublicKey,
		chaincfg.Supplier[network].Spec)
	if err != nil {
		return nil, err
	}
	return &Address{
		AddressPubKeyHash: addrPk.AddressPubKeyHash(),
		Network:           network,
	}, nil
}

func (a *Address) ToAccount(balance Amount) types.Account {
	return types.Account{
		Address: a.EncodeAddress(),
		Coin:    Coin(a.Network),
		Balance: balance.ToDecimal(),
	}
}
