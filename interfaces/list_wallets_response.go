package interfaces

import "github.com/DE-labtory/zulu/types"

type ListWalletResponse struct {
	Id       string          `json:"id"`
	Scheme   types.Scheme    `json:"scheme"`
	Accounts []types.Account `json:"accounts"`
}
