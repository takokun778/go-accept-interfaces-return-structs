package repository

import (
	"accept-interfaces-return-structs/example/model"
	"context"
)

type Repository interface {
	Save(context.Context, model.Example) error
	FindAll(context.Context) ([]model.Example, error)
}
