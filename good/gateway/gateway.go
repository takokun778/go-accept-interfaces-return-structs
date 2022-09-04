package gateway

import "accept-interfaces-return-structs/good/repository"

var (
	_ repository.A = (*Gateway)(nil)
	_ repository.B = (*Gateway)(nil)
)

type Gateway struct{}

func NewGateway() *Gateway {
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

func (g *Gateway) Update(item any) (any, error) {
	return nil, nil
}
