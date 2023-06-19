package games

import (
	"context"
	"sqlc-test/data"
)

func (g *Games) CreateAll(ctx context.Context, game *data.Game) error {
	err := g.dbx.Model(game).Insert()
	return err
}
