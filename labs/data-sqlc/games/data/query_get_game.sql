-- name: GetGame :one
SELECT id, name FROM games
WHERE id = $1 LIMIT 1;