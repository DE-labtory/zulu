package ethereum

import (
	"github.com/DE-labtory/zulu/types"
	"math/big"
)

type Params struct {
	NodeUrl string
	ChainId *big.Int
}

var (
	MainnetParams = Params{
		NodeUrl: "https://mainnet.infura.io",
		ChainId: big.NewInt(1),
	}
	RopstenParams = Params{
		NodeUrl: "https://ropsten.infura.io",
		ChainId: big.NewInt(3),
	}
)

var Supplier = map[types.Network]Params{
	types.Mainnet: MainnetParams,
	types.Ropsten: RopstenParams,
}
