package types

type TransferRequest struct {
	To string `json:"to"`
	Amount string `json:"amount"`
	Coin Coin `json:"coin"`
	Meta Meta `json:"meta"`
}