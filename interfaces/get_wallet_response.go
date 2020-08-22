package interfaces

import "github.com/DE-labtory/zulu/types"

type GetWalletResponse struct {
	Id       string          `json:"id"`
	Accounts []types.Account `json:"accounts"`
}
