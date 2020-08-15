package account

type Signer interface {
	PubKey() []byte
	PrivKey() []byte
}
