package types

type Transaction struct {
	TxHash string `json:"txHash"`
	Meta Meta `json:"meta"`
}
