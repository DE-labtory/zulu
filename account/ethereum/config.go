package ethereum

import "github.com/DE-labtory/zulu/types"

type Params struct {
	NodeUrl string
}

var (
	MainnetParams = Params{
		NodeUrl: "https://mainnet.infura.io",
	}
	RopstenParams = Params{
		NodeUrl: "https://ropsten.infura.io",
	}
)

var Supplier = map[types.Network]Params{
	types.Mainnet: MainnetParams,
	types.Ropsten: RopstenParams,
}
