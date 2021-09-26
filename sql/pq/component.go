package pq

import (
	"context"
	"database/sql"
)

type component struct {
	db *sql.DB
}

func (c *component) Status(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *component) GetDB() *sql.DB {
	return c.db
}
