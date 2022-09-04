package usecase

import (
	"accept-interfaces-return-structs/good/repository"
)

type A struct {
	repository repository.A
}

func NewA(r repository.A) *A {
	return &A{
		repository: r,
	}
}

func (a *A) Execute() error {
	a.repository.Save(struct{}{})
	return nil
}
