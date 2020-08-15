package keychain

type Store interface {
	Store(k Key)
	Get(id string) Key
}
