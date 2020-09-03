package main

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
)

const (
	//PvtKey          = "eeb1c9cb82fa9d81008847259e7239fcae3031fea4cccc224eab3e4c009de161"
	//PubKey          = "02d7fdaed2d5429370f50eed95bf43fa86e38cefd67eb7994725c66e3f983d14e3"
	//SenderAddress   = "muQqyVnEaUPLLco4rDtsKifE2AVyXsStFY"

	PvtKey = "9ca9700d14db691586ace71b25fe9973f1d2e0dd874e02e3d2d994ea7594f3e6"

	SenderAddress   = "momXhvmA324DdWhZC9TqFqNd9C7qszBKyn"
	ReceiverAddress = "mhVHWNGL8LaqbD5j6dL2KQ1xoewQCjAjsa"
)

func main() {
	// createAddressScenario()
	createTxScenario()
}

func createAddressScenario() {
	masterKey := createMasterNode()

	// Derive the extended key for account 0.  This gives the path:
	//   m/0H
	acct0, err := masterKey.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	acct0Keychain, _ := createKeyChains(acct0)

	acct0Addr := createAddress(acct0Keychain)

	pvtKey, err := acct0Keychain.ECPrivKey()
	if err != nil {
		panic(err)
	}
	pubKey, err := acct0Keychain.ECPubKey()
	if err != nil {
		panic(err)
	}

	fmt.Println("addr: ", acct0Addr)
	fmt.Println("private key: ", hex.EncodeToString(pvtKey.Serialize()))
	fmt.Println("public key: ", hex.EncodeToString(pubKey.SerializeCompressed()), len(pubKey.SerializeCompressed()))

	addrPubKey, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}

	//  the string encoding of the public key as a pay-to-pubkey-hash
	a1 := addrPubKey.EncodeAddress()
	a2 := addrPubKey.AddressPubKeyHash().EncodeAddress()
	fmt.Println(a1)
	fmt.Println(a2)
	if a1 != a2 {
		panic("a1, a2 is different")
	}
}

func getPubPvtKeyFromString() (*btcec.PrivateKey, *btcec.PublicKey) {
	hexPvtkey, err := hex.DecodeString(PvtKey)
	if err != nil {
		panic(err)
	}

	pvt, pub := btcec.PrivKeyFromBytes(btcec.S256(), hexPvtkey)
	return pvt, pub
}

func createTxScenario() {
	// https://www.blockchain.com/ko/btc-testnet/address/muQqyVnEaUPLLco4rDtsKifE2AVyXsStFY
	txId := "a813089ec3f71f8432978c52ae73f4eb8709c65f167d59bfe8aa64427e3e6042"
	vout := 1

	// 트랜잭션 정보를 담을 구조체 생성
	mtx := wire.NewMsgTx(wire.TxVersion)

	// UTXO를 가져오기 위해서 txId로부터 트랜잭션 구조체를 만듦
	txHash, err := chainhash.NewHashFromStr(txId)
	if err != nil {
		panic(err)
	}

	// 해당 트랜잭션의 몇 번째 UTXO를 사용할지 지정. 그리고 그것을 TxIn으로 만들어 트랜잭션 구조체에 저장
	// 만약 하나의 UTXO로 부족하다면 다른 UTXO를 추가적으로 가져와서 TxIn으로
	prevOutpoint := wire.NewOutPoint(txHash, uint32(vout))
	txIn := wire.NewTxIn(prevOutpoint, nil, nil)
	mtx.AddTxIn(txIn)

	// output
	addOutput(mtx, SenderAddress, 0.00002)
	addOutput(mtx, ReceiverAddress, 0.000004)

	// sender pubkey script
	sndrAddr, err := btcutil.DecodeAddress(SenderAddress, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}
	pkScript, err := txscript.PayToAddrScript(sndrAddr)
	if err != nil {
		panic(err)
	}

	// sign
	pvtKey, _ := getPubPvtKeyFromString()
	sig, err := txscript.SignatureScript(
		mtx,                 // The tx to be signed.
		0,                   // The index of the txin the signature is for.
		pkScript,            // The other half of the script from the PubKeyHash.
		txscript.SigHashAll, // The signature flags that indicate what the sig covers.
		pvtKey,              // The key to generate the signature with.
		true,                // The compress sig flag. This saves space on the blockchain.
	)
	if err != nil {
		panic(err)
	}
	mtx.TxIn[0].SignatureScript = sig

	// to hex
	mtxHex, err := messageToHex(mtx)
	if err != nil {
		panic(err)
	}
	fmt.Println(mtxHex)
}

func addOutput(mtx *wire.MsgTx, encodedAddr string, amount float64) {
	addr, err := btcutil.DecodeAddress(encodedAddr, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}

	// Create a new script which pays to the provided address.
	pkScript, err := txscript.PayToAddrScript(addr)
	if err != nil {
		// "Failed to generate pay-to-address script"
		panic(err)
	}

	// Convert the amount to satoshi.
	satoshi, err := btcutil.NewAmount(amount)
	if err != nil {
		// "Failed to convert amount"
		panic(err)
	}
	txOut := wire.NewTxOut(int64(satoshi), pkScript)
	mtx.AddTxOut(txOut)
}

// messageToHex serializes a message to the wire protocol encoding using the
// latest protocol version and returns a hex-encoded string of the result.
func messageToHex(msg wire.Message) (string, error) {
	var buf bytes.Buffer
	if err := msg.BtcEncode(&buf, 70002, wire.WitnessEncoding); err != nil {
		// context := fmt.Sprintf("Failed to encode msg of type %T", msg)
		panic(err)
	}

	return hex.EncodeToString(buf.Bytes()), nil
}

func createMasterNode() *hdkeychain.ExtendedKey {
	// Generate a random seed at the recommended length.
	seed, err := hdkeychain.GenerateSeed(hdkeychain.RecommendedSeedLen)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Generate a new master node using the seed.
	key, err := hdkeychain.NewMaster(seed, &chaincfg.TestNet3Params)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Show that the generated master node extended key is private.
	fmt.Println("Private Extended Key?:", key.IsPrivate())

	// Output:
	// Private Extended Key?: true
	return key
}

func createKeyChains(account *hdkeychain.ExtendedKey) (*hdkeychain.ExtendedKey, *hdkeychain.ExtendedKey) {
	// Derive the extended key for the account 0 external chain.  This
	// gives the path:
	//   m/0H/0
	external, err := account.Child(0)
	if err != nil {
		panic(err)
	}

	// Derive the extended key for the account 0 internal chain.  This gives
	// the path:
	//   m/0H/1
	internal, err := account.Child(1)
	if err != nil {
		panic(err)
	}
	return external, internal
}

func createAddress(keychain *hdkeychain.ExtendedKey) *btcutil.AddressPubKeyHash {
	// Get and show the address associated with the extended keys for the
	// main bitcoin	network.

	addr, err := keychain.Address(&chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}
	return addr
}
