package klaytn

import (
	"math/big"

	"github.com/DE-labtory/zulu/types"
)

type Params struct {
	NodeUrl  string
	ChainId  *big.Int
	GasLimit uint64
}

// chainId from https://github.com/klaytn/caver-java/blob/ea087d85ca53f90d627e67a6816c78dc72d887d8/core/src/main/java/com/klaytn/caver/utils/ChainId.java
var (
	CypressParams = Params{
		NodeUrl:  "https://api.cypress.klaytn.net:8651",
		ChainId:  big.NewInt(8217),
		GasLimit: 23000,
	}
	BaobabParams = Params{
		NodeUrl:  "https://api.baobab.klaytn.net:8651",
		ChainId:  big.NewInt(1001),
		GasLimit: 23000,
	}
)

var Supplier = map[types.Network]Params{
	types.Cypress: CypressParams,
	types.Baobab:  BaobabParams,
}
