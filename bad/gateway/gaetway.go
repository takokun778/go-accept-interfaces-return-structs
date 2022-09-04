package gateway

import "accept-interfaces-return-structs/bad/repository"

type Gateway struct{}

func NewGateway() repository.Repository {
	return &Gateway{}
}

func (g *Gateway) Save(item any) error {
	return nil
}

func (g *Gateway) Find(item any) (any, error) {
	return nil, nil
}

func (g *Gateway) Do() error {
	return nil
}
