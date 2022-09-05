package gateway

import (
	"github.com/uptrace/bun"
)

type Example struct {
	bun.BaseModel `bun:"table:examples"`

	ID   int    `bun:"id,pk,autoincrement"`
	Name string `bun:"name,notnull"`
}
