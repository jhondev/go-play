// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: cmd_create_game.sql

package data

import (
	"context"
)

const createGame = `-- name: CreateGame :exec
INSERT INTO arena.games (
  key_name, name
) VALUES (
  $1, $2
)
`

type CreateGameParams struct {
	KeyName string
	Name    string
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) error {
	_, err := q.db.ExecContext(ctx, createGame, arg.KeyName, arg.Name)
	return err
}
