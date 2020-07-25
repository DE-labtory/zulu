package types

type Blockchain struct {
	Platform Platform `json:"platform"`
	Network  Network  `json:"network"`
}
