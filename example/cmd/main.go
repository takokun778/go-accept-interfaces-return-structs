package main

import (
	"accept-interfaces-return-structs/example/gateway"
	"context"
	"log"
)

func main() {
	log.Println("migrate")

	g := gateway.NewGateway()

	if err := g.CreateTable(context.Background()); err != nil {
		log.Print(err)
	}
}
