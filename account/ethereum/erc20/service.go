package erc20

import (
	"github.com/DE-labtory/zulu/account/ethereum"
	"math/big"

	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type Service struct {
	decimal            int
	address            string
	network            types.Network
	transactionBuilder ethereum.TransactionBuilder
	client             ethereum.Client
}

func NewService(network types.Network, decimal int, address string) *Service {
	txBuilder := ethereum.NewTransactionBuilder(ethTypes.NewEIP155Signer(ethereum.Supplier[network].ChainId))
	client := ethereum.NewGethClient(network)
	return &Service{
		decimal:            decimal,
		address:            address,
		network:            network,
		transactionBuilder: *txBuilder,
		client:             client,
	}
}

func (s *Service) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	nonce, err := s.client.NonceAt(to)
	if err != nil {
		return types.Transaction{}, err
	}

	gasPrice, err := s.client.SuggestGasPrice()
	if err != nil {
		return types.Transaction{}, err
	}

	payload, err := s.buildErc20TransferData(to, amount)
	if err != nil {
		return types.Transaction{}, err
	}

	rawTx, err := s.transactionBuilder.Build(
		nonce,
		gasPrice,
		ethereum.Supplier[s.network].GasLimit,
		s.address,
		big.NewInt(0),
		payload,
		key.PrivateKey,
	)

	if err != nil {
		return types.Transaction{}, err
	}

	txHash, err := s.client.SendTransaction(rawTx)
	if err != nil {
		return types.Transaction{}, err
	}

	return types.Transaction{TxHash: txHash}, nil
}

func (s *Service) GetInfo() types.Coin {
	return types.Coin{}
}

func getTransferFunctionSignatrue() []byte {
	return []byte{0xa9, 0x05, 0x9c, 0xbb} // 0xa9059cbb
}

func (s *Service) buildErc20TransferData(to string, amount string) ([]byte, error) {
	address := common.HexToAddress(to)
	a := new(big.Int)
	a.SetString(amount, 10)

	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(a.Bytes(), 32)

	var data []byte
	data = append(data, getTransferFunctionSignatrue()...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	return data, nil
}
