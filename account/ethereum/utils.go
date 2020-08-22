package ethereum

import (
	"math"
	"math/big"
	"strconv"
)

func ConvertWithDecimal(value string, decimal int) (*big.Int, error) {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}

	a := floatValue * math.Pow10(decimal)
	b := int64(a)
	return big.NewInt(b), nil
}
