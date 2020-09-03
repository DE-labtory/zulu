package bitcoin

import (
	"fmt"

	"github.com/DE-labtory/zulu/account/bitcoin/chaincfg"
	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcutil"
)

type Address struct {
	*btcutil.AddressPubKeyHash
	Network types.Network
}

func DeriveAddress(key *KeyWrapper, network types.Network) (*Address, error) {
	addrPk, err := btcutil.NewAddressPubKeyHash(
		key.MarshalPubKey(),
		chaincfg.Supplier[network].Spec)
	if err != nil {
		return nil, err
	}
	return &Address{
		AddressPubKeyHash: addrPk,
		Network:           network,
	}, nil
}

func ParseAddressStr(addr string, network types.Network) (*Address, error) {
	a, err := btcutil.DecodeAddress(addr, chaincfg.Supplier[network].Spec)
	if err != nil {
		return nil, err
	}
	addrPkHash, ok := a.(*btcutil.AddressPubKeyHash)
	if !ok {
		return nil, fmt.Errorf("'%s' is not AddressPubKeyHash format", addr)
	}
	return &Address{
		AddressPubKeyHash: addrPkHash,
		Network:           network,
	}, nil
}

func (a *Address) PayToAddrScript() ([]byte, error) {
	pkScript, err := txscript.PayToAddrScript(a.AddressPubKeyHash)
	if err != nil {
		return nil, err
	}
	return pkScript, nil
}

func (a *Address) ToAccount(balance Amount, coin types.Coin) types.Account {
	return types.Account{
		Address: a.EncodeAddress(),
		Coin:    coin,
		Balance: balance.ToDecimal(),
	}
}
