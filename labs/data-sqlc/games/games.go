package games

import (
	"database/sql"
	"sqlc-test/data"
)

type Games struct {
	db *data.Queries
}

func New(db *sql.DB) *Games {
	return &Games{db: data.New(db)}
}
