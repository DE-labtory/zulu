package bitcoin

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/btcsuite/btcd/txscript"

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
	lock  sync.RWMutex
	items map[UnspentIndex]bool
}

func NewTxLock() *memUnspentLock {
	return &memUnspentLock{
		lock:  sync.RWMutex{},
		items: make(map[UnspentIndex]bool),
	}
}

func (l *memUnspentLock) Lock(idxList []UnspentIndex) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, idx := range idxList {
		l.items[idx] = true
	}
}

func (l *memUnspentLock) Unlock(idxList []UnspentIndex) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, idx := range idxList {
		l.items[idx] = false
	}
}

func (l *memUnspentLock) Locked(idx UnspentIndex) (item bool, ok bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	_, ok = l.items[idx]
	item = l.items[idx]
	return
}

type TxData struct {
	*wire.MsgTx
	UnspentList *UnspentList
	Rollback    func(*UnspentList, error)
}

func NewTxData(utxoList *UnspentList, rollback func(*UnspentList, error)) (*TxData, error) {
	txData := &TxData{
		MsgTx:       wire.NewMsgTx(wire.TxVersion),
		UnspentList: utxoList,
		Rollback:    rollback,
	}
	if err := txData.addInputsFromUTXO(); err != nil {
		return nil, err
	}
	return txData, nil
}

func (data *TxData) addInputsFromUTXO() error {
	for _, utxo := range data.UnspentList.Items {
		txHash, err := chainhash.NewHashFromStr(utxo.Txid)
		if err != nil {
			data.Rollback(data.UnspentList, err)
			return err
		}
		outpoint := wire.NewOutPoint(txHash, uint32(utxo.Vout))
		txIn := wire.NewTxIn(outpoint, nil, nil)
		data.AddTxIn(txIn)
	}
	return nil
}

func (data *TxData) AddOutputs(outputs [2]struct {
	*Address
	Amount
}) error {
	for _, out := range outputs {
		if err := data.AddOutput(out.Address, out.Amount); err != nil {
			return err
		}
	}
	return nil
}

func (data *TxData) AddOutput(addr *Address, amount Amount) error {
	pkScript, err := addr.PayToAddrScript()
	if err != nil {
		return err
	}
	fmt.Println(pkScript)
	txOut := wire.NewTxOut(amount.Int64(), pkScript)
	data.AddTxOut(txOut)
	return nil
}

func (data *TxData) SignFrom(wrapper *KeyWrapper, addr *Address) (string, error) {
	pkScript, err := addr.PayToAddrScript()
	if err != nil {
		data.Rollback(data.UnspentList, err)
		return "", err
	}
	for i, txIn := range data.TxIn {
		sig, err := txscript.SignatureScript(data.MsgTx, i, pkScript, txscript.SigHashAll, wrapper.PrivateKey, true)
		if err != nil {
			data.Rollback(data.UnspentList, err)
			return "", err
		}
		txIn.SignatureScript = sig
	}

	var buf bytes.Buffer
	if err := data.MsgTx.BtcEncode(&buf, 70002, wire.WitnessEncoding); err != nil {
		data.Rollback(data.UnspentList, err)
		return "", err
	}
	return hex.EncodeToString(buf.Bytes()), nil
}

type TxResult struct {
	TxId string
}

func (r TxResult) Transaction() types.Transaction {
	return types.Transaction{
		TxHash: r.TxId,
	}
}

type TxService struct {
	Network types.Network
	TxLock  TxLock
	node    Adapter
}

func NewTxService(network types.Network, node Adapter) *TxService {
	return &TxService{
		Network: network,
		TxLock:  NewTxLock(),
		node:    node,
	}
}

func (b *TxService) Create(addr, to *Address, amount Amount) (*TxData, error) {
	unspents, err := b.ListUnspent(addr.EncodeAddress())
	if err != nil {
		return nil, err
	}
	fee, err := b.calcFee(len(unspents.Items), 2)
	if err != nil {
		return nil, err
	}
	if ok := unspents.EnoughThan(amount, fee); !ok {
		return nil, fmt.Errorf("'%s' have not enough balance: %s", addr, amount.ToDecimal())
	}

	b.TxLock.Lock(unspents.IdxList())

	txData, err := NewTxData(unspents, func(utxoList *UnspentList, err error) {
		b.TxLock.Unlock(utxoList.IdxList())
	})
	if err != nil {
		return nil, err
	}
	txData.AddOutputs([2]struct {
		*Address
		Amount
	}{
		{addr, unspents.Balance().Sub(amount).Sub(fee)},
		{to, amount},
	})
	return txData, nil
}

func (b *TxService) SendRaw(data string) (TxResult, error) {
	txResult, err := b.node.SendRawTx(data)
	if err != nil {
		return TxResult{}, err
	}
	return txResult, nil
}

func (b *TxService) ListUnspent(addr string) (*UnspentList, error) {
	all, err := b.node.ListUTXO(addr)
	if err != nil {
		return nil, err
	}
	result := make([]Unspent, 0)
	for _, utxo := range all {
		locked, ok := b.TxLock.Locked(utxo.UnspentIndex)
		if !ok {
			result = append(result, utxo)
			continue
		}
		if locked {
			continue
		}
		result = append(result, utxo)
	}
	return NewUnspentList(result), nil
}

func (b *TxService) calcFee(input, output int) (Amount, error) {
	feeRate, err := b.node.EstimateFeeRate()
	if err != nil {
		return Amount{}, err
	}
	return NewAmount(int64(1.5 * feeRate *
			(bytesPerInput*float64(input) + bytesPerOutput*float64(output)))),
		nil
}
