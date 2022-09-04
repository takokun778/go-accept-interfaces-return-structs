package main

import (
	"accept-interfaces-return-structs/bad/gateway"
	"accept-interfaces-return-structs/good/usecase"
)

func main() {
	g := gateway.NewGateway()

	// can't call
	// g.Do()

	a := usecase.NewA(g)

	b := usecase.NewB(g)

	a.Execute()

	b.Execute()
}
