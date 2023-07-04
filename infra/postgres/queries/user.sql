-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserByID :one
UPDATE users SET updated_at = $1 WHERE id = $2 RETURNING *;

-- name: RemoveUserByID :exec
DELETE FROM users WHERE id = $1;
