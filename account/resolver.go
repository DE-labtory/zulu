package account

import (
	"errors"
)

var (
	ErrUnknownCoin = errors.New("unknown account error")
)

type Resolver struct {
	services map[string]Service
}

func NewResolver() *Resolver {
	return &Resolver{
		services: make(map[string]Service),
	}
}

func (r *Resolver) Resolve(id string) (Service, error) {
	service := r.services[id]
	if service == nil {
		return nil, ErrUnknownCoin
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
