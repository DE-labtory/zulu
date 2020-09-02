package ethereum

import (
	"math/big"
	"os"

	"github.com/DE-labtory/zulu/types"
	"github.com/joho/godotenv"
)

type Params struct {
	NodeUrl  string
	ChainId  *big.Int
	GasLimit uint64
}

var Supplier = map[types.Network]Params{}

func init() {
	_ = godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/DE-labtory/zulu/.env"))

	Supplier[types.Mainnet] = Params{
		NodeUrl:  os.Getenv("INFURA_MAINNET_URL"),
		ChainId:  big.NewInt(1),
		GasLimit: 23000,
	}
	Supplier[types.Ropsten] = Params{
		NodeUrl:  os.Getenv("INFURA_ROPSTEN_URL"),
		ChainId:  big.NewInt(3),
		GasLimit: 23000,
	}
}
