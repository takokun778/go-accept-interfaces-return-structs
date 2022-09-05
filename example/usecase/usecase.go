package usecase

import (
	"accept-interfaces-return-structs/example/model"
	"accept-interfaces-return-structs/example/repository"
	"context"
	"log"
)

type Usecase struct {
	repository repository.Repository
}

func NewUsecase(r repository.Repository) *Usecase {
	return &Usecase{
		repository: r,
	}
}

func (u *Usecase) Execute(ctx context.Context, example model.Example) error {
	if err := u.repository.Save(ctx, example); err != nil {
		return err
	}

	examples, err := u.repository.FindAll(ctx)
	if err != nil {
		return err
	}

	for i, example := range examples {
		log.Printf("%d example %+v\n", i, example)
	}

	return nil
}
