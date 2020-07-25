package types

type Coin struct {
	Id         string     `json:"id"`
	Blockchain Blockchain `json:"blockchain"`
	Symbol     Symbol     `json:"symbol"`
	Decimals   int        `json:"decimals"`
	Meta       Meta       `json:"meta"`
}
