package interfaces

import "github.com/DE-labtory/zulu/types"

type CreateWalletRequest struct {
	Scheme   types.Scheme `json:"scheme"`
	Password string       `json:"password"`
	Meta     types.Meta   `json:"meta"`
}
