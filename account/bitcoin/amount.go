package bitcoin

import (
	"math"
	"strconv"

	"github.com/DE-labtory/zulu/types"
)

// AmountUnit describes a method of converting an Amount to something
// other than the base unit of a bitcoin.  The value of the AmountUnit
// is the exponent component of the decadic multiple to convert from
// an amount in bitcoin to an amount counted in units.
type AmountUnit int

const (
	Decimal AmountUnit = 8
)

func (au AmountUnit) Int() int {
	return int(au)
}

type Amount int64

const (
	AmountZero = Amount(0)
	AmountOne  = Amount(1)
)

func NewAmount(i int) Amount {
	return Amount(int64(i) * int64(math.Pow10(int(Decimal))))
}

// TODO: implement me
func (a Amount) Add(i Amount) Amount {
	return 0
}

// TODO: implement me
func (a Amount) Sub(i Amount) Amount {
	return 0
}

func (a Amount) ToDecimal() string {
	return strconv.FormatInt(int64(a), 10)
}

func (a Amount) ToHex() string {
	return strconv.FormatInt(int64(a), 16)
}

func Coin(network types.Network) types.Coin {
	return types.Coin{
		Id: "1", // TODO: decide how to create
		Blockchain: types.Blockchain{
			Platform: types.Bitcoin,
			Network:  network,
		},
		Symbol:   types.Btc,
		Decimals: Decimal.Int(),
	}
}
