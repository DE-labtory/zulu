package types

type ListWalletResponse struct {
	Id string `json:"id"`
	Scheme Scheme `json:"scheme"`
	Addresses []Address `json:"addresses"`
}