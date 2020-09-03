package eth

import (
	"github.com/DE-labtory/zulu/account/ethereum"
	"github.com/DE-labtory/zulu/keychain"
	"github.com/DE-labtory/zulu/types"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

const defaultDecimal = 18

type Service struct {
	coin               types.Coin
	transactionBuilder ethereum.TransactionBuilder
	client             ethereum.Client
	*ethereum.Deriver
}

func NewService(coin types.Coin, client ethereum.Client) *Service {
	txBuilder := ethereum.NewTransactionBuilder(ethTypes.NewEIP155Signer(ethereum.Supplier[coin.Blockchain.Network].ChainId))
	return &Service{
		coin:               coin,
		transactionBuilder: *txBuilder,
		client:             client,
		Deriver:            ethereum.NewDeriver(coin, client),
	}
}

func (s *Service) Transfer(key keychain.Key, to string, amount string) (types.Transaction, error) {
	account, err := s.DeriveAccount(key)
	if err != nil {
		return types.Transaction{}, err
	}

	nonce, err := s.client.NonceAt(account.Address)
	if err != nil {
		return types.Transaction{}, err
	}

	gasPrice, err := s.client.SuggestGasPrice()
	if err != nil {
		return types.Transaction{}, err
	}

	value, err := ethereum.ConvertWithDecimal(amount, defaultDecimal)
	if err != nil {
		return types.Transaction{}, err
	}

	rawTx, err := s.transactionBuilder.Build(
		nonce,
		gasPrice,
		ethereum.Supplier[s.coin.Blockchain.Network].GasLimit,
		to,
		value,
		nil,
		key.GetPrivateKey(),
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
	return s.coin
}
