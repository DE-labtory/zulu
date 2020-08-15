package interfaces

import "github.com/DE-labtory/zulu/types"

type TransferRequest struct {
	To       string     `json:"to"`
	Amount   string     `json:"amount"`
	Password string     `json:"password"`
	Coin     types.Coin `json:"coin"`
	Meta     types.Meta `json:"meta"`
}
