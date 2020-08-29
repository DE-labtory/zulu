package bitcoin

import (
	"fmt"

	"github.com/DE-labtory/zulu/types"

	"github.com/btcsuite/btcd/chaincfg/chainhash"

	"github.com/btcsuite/btcd/wire"
)

const (
	bytesPerInput  = 148
	bytesPerOutput = 34

	defaultFeeRate = 1.0
)

type UnspentIndex struct {
	Txid string `json:"txid"`
	Vout int    `json:"vout"`
}

type Unspent struct {
	UnspentIndex
	Value int64 `json:"value"`
}

type UnspentList struct {
	Items []Unspent
}

func NewUnspentList(items []Unspent) *UnspentList {
	return &UnspentList{
		Items: items,
	}
}

func (l *UnspentList) Balance() Amount {
	balance := NewAmount(0)
	for _, u := range l.Items {
		balance = balance.Add(NewAmount(u.Value))
	}
	return balance
}

func (l *UnspentList) EnoughThan(amounts ...Amount) bool {
	totalAmount := NewAmount(0)
	for _, amount := range amounts {
		totalAmount.Add(amount)
	}

	balance := l.Balance()
	if balance.Compare(totalAmount) < 0 {
		return false
	}
	return true
}

func (l *UnspentList) IdxList() []UnspentIndex {
	idxList := make([]UnspentIndex, 0)
	for _, u := range l.Items {
		idxList = append(idxList, UnspentIndex{Txid: u.Txid, Vout: u.Vout})
	}
	return idxList
}

type TxLock interface {
	Lock(idxList []UnspentIndex)
	Unlock(idxList []UnspentIndex)
	Locked(idx UnspentIndex) (bool, bool)
}

type memUnspentLock struct {
	items map[UnspentIndex]bool
}

func (l *memUnspentLock) Lock(idxList []UnspentIndex) {
	for _, idx := range idxList {
		l.items[idx] = true
	}
}

func (l *memUnspentLock) Unlock(idxList []UnspentIndex) {
	for _, idx := range idxList {
		l.items[idx] = false
	}
}

func (l *memUnspentLock) Locked(idx UnspentIndex) (item bool, ok bool) {
	_, ok = l.items[idx]
	item = l.items[idx]
	return
}

type TxLister struct {
	node   Adapter
	txLock TxLock
}

func NewTxLister(node Adapter) *TxLister {
	return &TxLister{
		node: node,
		txLock: &memUnspentLock{
			items: make(map[UnspentIndex]bool),
		},
	}
}

func (l *TxLister) ListUnspent(addr string) ([]Unspent, error) {
	all, err := l.node.ListUTXO(addr)
	if err != nil {
		return nil, err
	}
	result := make([]Unspent, 0)
	for _, utxo := range all {
		if locked, ok := l.txLock.Locked(utxo.UnspentIndex); !locked && ok {
			result = append(result, utxo)
		}
	}
	return result, nil
}

type Tx struct{}

type TxData struct {
	*wire.MsgTx
	OnErrAddingUTXO func(*UnspentList, error) error
}

func EmptyTxData() *TxData {
	return &TxData{
		MsgTx: wire.NewMsgTx(wire.TxVersion),
		OnErrAddingUTXO: func(*UnspentList, error) error {
			return nil
		},
	}
}

func (data *TxData) AddInputsFromUTXO(utxoList *UnspentList) error {
	for _, utxo := range utxoList.Items {
		txHash, err := chainhash.NewHashFromStr(utxo.Txid)
		if err != nil {
			return data.OnErrAddingUTXO(utxoList, err)
		}
		outpoint := wire.NewOutPoint(txHash, uint32(utxo.Vout))
		txIn := wire.NewTxIn(outpoint, nil, nil)
		data.AddTxIn(txIn)
	}
	return nil
}

func (data *TxData) AddOutputs(outputs []struct {
	Address
	Amount
}) error {
	for _, out := range outputs {
		if err := data.AddOutput(out.Address, out.Amount); err != nil {
			return err
		}
	}
	return nil
}

func (data *TxData) AddOutput(addr Address, amount Amount) error {
	pkScript, err := addr.PayToAddrScript()
	if err != nil {
		return err
	}
	txOut := wire.NewTxOut(amount.Int64(), pkScript)
	data.AddTxOut(txOut)
	return nil
}

type TxService struct {
	network  types.Network
	txLister TxLister
	txLock   TxLock
	node     Adapter
}

func (b *TxService) Create(addr, to Address, amount Amount) (*TxData, error) {
	utxos, err := b.txLister.ListUnspent(addr.EncodeAddress())
	if err != nil {
		return nil, err
	}
	fee, err := b.calcFee(len(utxos), 2)
	if err != nil {
		return nil, err
	}
	utxoList := NewUnspentList(utxos)
	if ok := utxoList.EnoughThan(amount, fee); !ok {
		return nil, fmt.Errorf("'%s' have not enough balance: %s", addr, amount.ToDecimal())
	}

	b.txLock.Lock(utxoList.IdxList())

	txData := EmptyTxData()
	txData.OnErrAddingUTXO = func(utxoList *UnspentList, err error) error {
		b.txLock.Unlock(utxoList.IdxList())
		return err
	}
	if err := txData.AddInputsFromUTXO(utxoList); err != nil {
		return nil, err
	}
	txData.AddOutputs([]struct {
		Address
		Amount
	}{
		{to, amount},
		{addr, utxoList.Balance().Sub(amount).Sub(fee)},
	})
	return txData, nil
}

func (b *TxService) Sign() (string, error) {
	return "", nil
}

func (b *TxService) calcFee(input, output int) (Amount, error) {
	feeRate, err := b.node.EstimateFeeRate()
	if err != nil {
		return Amount{}, err
	}
	return NewAmount(int64(feeRate *
			(bytesPerInput*float64(input) + bytesPerOutput*float64(output)))),
		nil
}
