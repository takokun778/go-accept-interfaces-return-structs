package main

import (
	"accept-interfaces-return-structs/good/gateway"
	"accept-interfaces-return-structs/good/usecase"
)

func main() {
	g := gateway.NewGateway()

	// can call
	g.Do()

	a := usecase.NewA(g)

	b := usecase.NewB(g)

	a.Execute()

	b.Execute()
}
