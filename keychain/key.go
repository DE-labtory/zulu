package keychain

type Key struct {
	ID         string
	PrivateKey []byte
	PublicKey  []byte
}

type KeyGenerator struct {
}

func (k *KeyGenerator) Generate() Key {
	return Key{}
}
