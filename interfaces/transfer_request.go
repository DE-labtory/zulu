package interfaces

import "github.com/DE-labtory/zulu/types"

type TransferRequest struct {
	To       string     `json:"to"`
	Amount   string     `json:"amount"`
	Password string     `json:"password"`
	CoinId   string     `json:"coinId"`
	Meta     types.Meta `json:"meta"`
}
