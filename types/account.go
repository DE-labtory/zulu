package types

type Account struct {
	Address string  `json:"address"`
	Coin    Coin    `json:"coin"`
	Balance Balance `json:"balance"`
}
