package keychain

type Store interface {
	Store(k Key) error
	Get(id string) (Key, error)
}
