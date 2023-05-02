-- name: CreateGame :exec
INSERT INTO arena.games (
  key_name, name
) VALUES (
  $1, $2
);