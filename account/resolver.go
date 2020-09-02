package account

import (
	"errors"

	"github.com/DE-labtory/zulu/account/bitcoin"
	"github.com/DE-labtory/zulu/account/ethereum"
	"github.com/DE-labtory/zulu/account/ethereum/eth"
	"github.com/DE-labtory/zulu/types"
)

var (
	ErrUnknownAccountService = errors.New("unknown account service error")
)

type Resolver struct {
	services map[string]Service
}

func NewResolver() *Resolver {
	return &Resolver{
		services: register(),
	}
}

func (r *Resolver) Resolve(id string) (Service, error) {
	service := r.services[id]
	if service == nil {
		return nil, ErrUnknownAccountService
	}
	return service, nil
}

func (r *Resolver) GetAllServices() []Service {
	var values []Service
	for _, value := range r.services {
		values = append(values, value)
	}
	return values
}

func register() map[string]Service {
	services := make(map[string]Service)

	// bitcoin service
	services["1"] = bitcoin.NewService(
		types.Coin{
			Id: "1",
			Blockchain: types.Blockchain{
				Platform: types.Bitcoin,
				Network:  types.Mainnet,
			},
			Symbol:   types.Btc,
			Decimals: bitcoin.Decimal.Int(),
		},
		types.Mainnet,
	)

	services["2"] = bitcoin.NewService(
		types.Coin{
			Id: "2",
			Blockchain: types.Blockchain{
				Platform: types.Bitcoin,
				Network:  types.Testnet,
			},
			Symbol:   types.Tbtc,
			Decimals: bitcoin.Decimal.Int(),
		},
		types.Testnet,
	)

	//
	services["3"] = eth.NewService(
		types.Coin{
			Id: "3",
			Blockchain: types.Blockchain{
				Platform: types.Ethereum,
				Network:  types.Mainnet,
			},
			Symbol:   "ETH",
			Decimals: 18,
			Meta:     nil,
		},
		ethereum.NewGethClient(types.Mainnet),
	)
	services["4"] = eth.NewService(
		types.Coin{
			Id: "4",
			Blockchain: types.Blockchain{
				Platform: types.Ethereum,
				Network:  types.Ropsten,
			},
			Symbol:   "ETH",
			Decimals: 18,
			Meta:     nil,
		},
		ethereum.NewGethClient(types.Ropsten),
	)
	return services
}
