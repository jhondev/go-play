package grpc

import (
	"context"
	"sqlc-test/data"
	"sqlc-test/games"
)

type Server struct {
	games *games.Games
}

func New(games *games.Games) *Server {
	return &Server{games}
}

type CreateGameRequest struct {
	KeyName string
	Name    string
}

func (s *Server) CreateGame(ctx context.Context, req *CreateGameRequest) error {
	s.games.Create(ctx, &data.CreateGameParams{
		KeyName: req.KeyName,
		Name:    req.Name,
	})
	return nil
}
