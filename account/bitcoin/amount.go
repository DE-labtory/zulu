package bitcoin

import (
	"fmt"
	"math/big"
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

type Amount struct {
	value *big.Int
}

func NewAmount(i int64) Amount {
	return Amount{
		value: big.NewInt(i),
	}
}

func ParseAmount(i string) (Amount, error) {
	v, ok := new(big.Int).SetString(i, 10)
	if !ok {
		return Amount{}, fmt.Errorf("invalid amount: %s", i)
	}
	return Amount{
		value: v,
	}, nil
}

func (a Amount) Add(x Amount) Amount {
	return NewAmount(a.value.Add(a.value, x.value).Int64())
}

func (a Amount) Sub(x Amount) Amount {
	return NewAmount(a.value.Sub(a.value, x.value).Int64())
}

func (a Amount) ToDecimal() string {
	return strconv.FormatInt(a.value.Int64(), 10)
}

func (a Amount) ToHex() string {
	return "0x" + strconv.FormatInt(a.value.Int64(), 16)
}

func (a Amount) Int64() int64 {
	return a.value.Int64()
}

func (a Amount) Compare(x Amount) int {
	return a.value.Cmp(x.value)
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
