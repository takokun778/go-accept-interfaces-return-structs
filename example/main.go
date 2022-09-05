package main

import (
	"accept-interfaces-return-structs/example/gateway"
	"accept-interfaces-return-structs/example/model"
	"accept-interfaces-return-structs/example/usecase"
	"context"
	"log"
)

func main() {
	log.Println("main")

	g := gateway.NewGateway()

	u := usecase.NewUsecase(g)

	e := model.Example{
		Name: "example",
	}

	if err := u.Execute(context.Background(), e); err != nil {
		log.Print(err)
	}
}
