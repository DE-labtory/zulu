package coin

type Signer interface {
	PubKey() []byte
	PrivKey() []byte
}
