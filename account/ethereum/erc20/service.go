package erc20

import (
	"encoding/hex"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Service struct {
	Decimal int
	address string
}

func NewService(
	decimal int,
	address string,
) *Service {

	return &Service{
		Decimal: decimal,
		address: address,
	}
}

func (s *Service) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	return types.Transaction{}, nil
}

func getTransferFunctionSignatrue() []byte {
	return []byte{0xa9, 0x05, 0x9c, 0xbb} // 0xa9059cbb
}

func (s *Service) buildErc20TransferData(to string, amount string) ([]byte, error) {
	toAddress, err := hex.DecodeString(to)
	if err != nil {
		return nil, err
	}
	a := new(big.Int)
	a.SetString(amount, 10)

	paddedAddress := common.LeftPadBytes(toAddress, 32)
	paddedAmount := common.LeftPadBytes(a.Bytes(), 32)

	var data []byte
	data = append(data, getTransferFunctionSignatrue()...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	return data, nil
}
