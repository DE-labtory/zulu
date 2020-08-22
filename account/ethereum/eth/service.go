package eth

import (
	"errors"
	"github.com/DE-labtory/zulu/account/ethereum"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type Service struct {
	network            types.Network
	transactionBuilder ethereum.TransactionBuilder
	client             ethereum.Client
}

func NewService(network types.Network) *Service {
	txBuilder := ethereum.NewTransactionBuilder(ethTypes.NewEIP155Signer(ethereum.Supplier[network].ChainId))
	client := ethereum.NewGethClient(network)
	return &Service{
		network:            network,
		transactionBuilder: *txBuilder,
		client:             client,
	}
}

func (s *Service) DeriveAccount(key keychain.Key) (types.Account, error) {
	return types.Account{}, nil
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

	value := new(big.Int)
	value, ok := value.SetString(amount, 10)
	if !ok {
		return types.Transaction{}, errors.New("invalid amount format")
	}

	rawTx, err := s.transactionBuilder.Build(
		nonce,
		gasPrice,
		ethereum.Supplier[s.network].GasLimit,
		to,
		value,
		nil,
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
