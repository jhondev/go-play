-- name: GetMatch :one
SELECT id, name FROM arena.matches
WHERE id = $1 LIMIT 1;