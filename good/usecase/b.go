package usecase

import (
	"accept-interfaces-return-structs/good/repository"
)

type B struct {
	repository repository.B
}

func NewB(r repository.B) *B {
	return &B{
		repository: r,
	}
}

func (b *B) Execute() error {
	b.repository.Find(struct{}{})
	return nil
}
