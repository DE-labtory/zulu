package bitcoin

import (
	"github.com/DE-labtory/zulu/account"
	"github.com/DE-labtory/zulu/account/bitcoin/chaincfg"
	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

type PrivKey struct {
	*hdkeychain.ExtendedKey
}

func GetPrivKey(signer account.Signer) (*PrivKey, error) {
	extKey, err := hdkeychain.NewKeyFromString(string(signer.PrivKey()))
	if err != nil {
		return nil, err
	}
	return &PrivKey{
		extKey,
	}, nil
}

func (k *PrivKey) GetAddress(network types.Network) (*Address, error) {
	pubkeyHash, err := k.Address(chaincfg.Supplier[network].Spec)
	if err != nil {
		return nil, err
	}
	return &Address{
		pubkeyHash,
	}, nil
}

type Address struct {
	*btcutil.AddressPubKeyHash
}
