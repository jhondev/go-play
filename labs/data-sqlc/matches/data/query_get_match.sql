-- name: GetMatch :one
SELECT id, name FROM matches
WHERE id = $1 LIMIT 1;