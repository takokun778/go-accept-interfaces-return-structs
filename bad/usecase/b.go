package usecase

import "accept-interfaces-return-structs/bad/repository"

type B struct {
	repository repository.Repository
}

func NewB(r repository.Repository) *B {
	return &B{
		repository: r,
	}
}

func (b *B) Execute() error {
	b.repository.Find(struct{}{})
	return nil
}
