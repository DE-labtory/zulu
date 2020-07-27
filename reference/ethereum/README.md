## Ethereum Signing
The signing process for transaction is as follows.
> Sig = sign(keccak256(rlp.encode(tx.raw)), privateKey)

### Transaction Signing Process
**Transaction Structure (Reference - [Go Ethereum](https://github.com/ethereum/go-ethereum/blob/master/core/types/transaction.go#L46))**
```go
type txdata struct {
    AccountNonce uint64          `json:"nonce"    gencodec:"required"`
    Price        *big.Int        `json:"gasPrice" gencodec:"required"`
    GasLimit     uint64          `json:"gas"      gencodec:"required"`
    Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
    Amount       *big.Int        `json:"value"    gencodec:"required"`
    Payload      []byte          `json:"input"    gencodec:"required"`
    
    // Signature values
    V *big.Int `json:"v" gencodec:"required"`
    R *big.Int `json:"r" gencodec:"required"`
    S *big.Int `json:"s" gencodec:"required"`
}
```

1. Serialize the transaction by RLP Encoding.
    - Required data and order: nonce, gas price, gas limit, to address, amount, payload, v(chain id), r, s

2. Keccak256 Hash the result of step 1

3. Sign the result of step 2 with a private key.

4. Extract v, r, s with the value of the signing result. It proves that these are signed.
    - The value of v follows [EIP155](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-155.md).
    
5. Insert V, R and S into the transaction used in 1 and proceed with RLP encoding.

6. Send the result to the node.

### Message Signing Process
The signing process for message is as follows.
In Ethereum, a special prefix is attached to the message to distinguish it is Ethereum message.
> Sig = sign(keccak256("\x19Ethereum Signed Message:\n" + len(message) + message)))
