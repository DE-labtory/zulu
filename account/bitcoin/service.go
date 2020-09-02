package bitcoin

import (
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
)

type bitcoinType struct {
	coin      types.Coin
	network   types.Network
	txService *TxService
}

func NewService(coin types.Coin, network types.Network) *bitcoinType {
	return &bitcoinType{
		coin:      coin,
		network:   network,
		txService: NewTxService(network, NewAdapter(network)),
	}
}

func (b *bitcoinType) DeriveAccount(key keychain.Key) (types.Account, error) {
	addr, err := DeriveAddress(NewKeyWrapper(key), b.network)
	if err != nil {
		return types.Account{}, err
	}
	unspents, err := b.txService.ListUnspent(addr.EncodeAddress())
	if err != nil {
		return types.Account{}, err
	}
	return addr.ToAccount(unspents.Balance(), b.coin), nil
}

func (b *bitcoinType) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	kw := NewKeyWrapper(key)
	fromAddr, err := DeriveAddress(kw, b.network)
	if err != nil {
		return types.Transaction{}, err
	}
	toAddr, err := ParseAddressStr(to, b.network)
	if err != nil {
		return types.Transaction{}, err
	}
	amnt, err := ParseAmount(amount)
	if err != nil {
		return types.Transaction{}, err
	}
	txData, err := b.txService.Create(fromAddr, toAddr, amnt)
	if err != nil {
		return types.Transaction{}, err
	}
	rawData, err := txData.SignFrom(kw, fromAddr)
	if err != nil {
		return types.Transaction{}, err
	}
	result, err := b.txService.SendRaw(rawData)
	if err != nil {
		return types.Transaction{}, err
	}
	return result.Transaction(), nil
}

func (b *bitcoinType) GetInfo() types.Coin {
	return b.coin
}
