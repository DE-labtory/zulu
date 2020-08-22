package ethereum

import (
	"github.com/DE-labtory/zulu/types"
	"math/big"
)

type Params struct {
	NodeUrl string
	ChainId *big.Int
	GasLimit uint64
}

var (
	MainnetParams = Params{
		NodeUrl: "https://mainnet.infura.io",
		ChainId: big.NewInt(1),
		GasLimit: 23000,
	}
	RopstenParams = Params{
		NodeUrl: "https://ropsten.infura.io",
		ChainId: big.NewInt(3),
		GasLimit: 23000,
	}
)

var Supplier = map[types.Network]Params{
	types.Mainnet: MainnetParams,
	types.Ropsten: RopstenParams,
}
