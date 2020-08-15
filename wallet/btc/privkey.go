package btc

import (
	"fmt"

	"github.com/DE-labtory/zulu/types"
	"github.com/DE-labtory/zulu/wallet"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

type PrivKey struct {
	*hdkeychain.ExtendedKey
}

func GetPrivKey(signer wallet.Signer) (*PrivKey, error) {
	extKey, err := hdkeychain.NewKeyFromString(signer.PrivKey())
	if err != nil {
		return nil, err
	}
	return &PrivKey{
		extKey,
	}, nil
}

func (k *PrivKey) GetAddress(network types.Network) (*Address, error) {
	params, err := getNetworkParams(network)
	if err != nil {
		return nil, err
	}
	pubkeyHash, err := k.Address(params)
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

func getNetworkParams(network types.Network) (*chaincfg.Params, error) {
	var net *chaincfg.Params
	switch network {
	case types.Mainnet:
		net = &chaincfg.MainNetParams
	case types.Testnet:
		net = &chaincfg.TestNet3Params
	default:
		return nil, fmt.Errorf("invalid network type got: %s", network)
	}
	return net, nil
}
