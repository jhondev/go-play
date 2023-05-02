-- name: GetGame :one
SELECT id, name FROM arena.games
WHERE id = $1 LIMIT 1;