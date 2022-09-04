package usecase

import "accept-interfaces-return-structs/bad/repository"

type A struct {
	repository repository.Repository
}

func NewA(r repository.Repository) *A {
	return &A{
		repository: r,
	}
}

func (a *A) Execute() error {
	a.repository.Save(struct{}{})
	return nil
}
