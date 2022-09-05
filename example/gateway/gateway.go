package gateway

import (
	"accept-interfaces-return-structs/example/model"
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

type Gateway struct {
	*bun.DB
}

func NewGateway() *Gateway {
	dsn := os.Getenv("DATABASE_URL")

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	config.PreferSimpleProtocol = true

	db := stdlib.OpenDB(*config)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Gateway{
		bun.NewDB(db, pgdialect.New()),
	}
}

func (g *Gateway) Save(ctx context.Context, example model.Example) error {
	e := Example{
		Name: example.Name,
	}

	_, err := g.DB.NewInsert().
		Model(&e).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gateway) FindAll(ctx context.Context) ([]model.Example, error) {
	var entities []Example

	err := g.DB.NewSelect().
		Model(&entities).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	examples := make([]model.Example, 0, len(entities))

	for _, entity := range entities {
		example := model.Example{
			ID:   entity.ID,
			Name: entity.Name,
		}
		examples = append(examples, example)
	}

	return examples, nil
}

func (g *Gateway) CreateTable(ctx context.Context) error {
	_, err := g.DB.NewCreateTable().
		Model((*Example)(nil)).
		IfNotExists().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
