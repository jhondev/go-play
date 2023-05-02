package games

import (
	"context"
	"sqlc-test/data"
)

func (g *Games) Create(ctx context.Context, game *data.CreateGameParams) error {
	// validations

	// db
	err := g.db.CreateGame(ctx, *game)
	return err
}
