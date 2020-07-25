package types

type Address struct {
	Address string `json:"address"`
	Coin Coin `json:"coin"`
	Balance Balance `json:"balance"`
}