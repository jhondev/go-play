-- name: CreateGame :exec
INSERT INTO games (
  key_name, name
) VALUES (
  $1, $2
);