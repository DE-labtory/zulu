package types

type CreateWalletRequest struct {
	Scheme   Scheme `json:"scheme"`
	Password string `json:"password"`
	Meta     Meta   `json:"meta"`
}
