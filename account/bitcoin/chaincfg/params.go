package chaincfg

import (
	"github.com/DE-labtory/zulu/types"
	"github.com/btcsuite/btcd/chaincfg"
)

type Params struct {
	NodeUrl string
	Spec    *chaincfg.Params
}

var (
	TestnetParams = Params{
		NodeUrl: "https://blockstream.info/testnet/api",
		Spec:    &chaincfg.TestNet3Params,
	}
	MainnetParams = Params{
		NodeUrl: "https://blockstream.info/api",
		Spec:    &chaincfg.MainNetParams,
	}
)

var Supplier = map[types.Network]Params{
	types.Testnet: TestnetParams,
	types.Mainnet: MainnetParams,
}
