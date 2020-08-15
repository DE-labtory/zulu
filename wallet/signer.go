package wallet

type Signer interface {
	PubKey() string
	PrivKey() string
}
