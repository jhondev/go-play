package games

import (
	"sqlc-test/data"

	"github.com/pocketbase/dbx"
)

type Games struct {
	db  *data.Queries
	dbx *dbx.DB
}

func New(db *dbx.DB) *Games {
	return &Games{
		db:  data.New(db.DB()),
		dbx: db,
	}
}
