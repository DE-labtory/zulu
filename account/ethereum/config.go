package ethereum

import (
	"math/big"

	"github.com/DE-labtory/zulu/conf"
	"github.com/DE-labtory/zulu/types"
)

type Params struct {
	NodeUrl  string
	ChainId  *big.Int
	GasLimit uint64
}

var (
	MainnetParams = Params{
		NodeUrl:  conf.GetConfiguration().Endpoint.Ethereum.Mainnet,
		ChainId:  big.NewInt(1),
		GasLimit: 23000,
	}
	RopstenParams = Params{
		NodeUrl:  conf.GetConfiguration().Endpoint.Ethereum.Ropsten,
		ChainId:  big.NewInt(3),
		GasLimit: 23000,
	}
)

var Supplier = map[types.Network]Params{
	types.Mainnet: MainnetParams,
	types.Ropsten: RopstenParams,
}
