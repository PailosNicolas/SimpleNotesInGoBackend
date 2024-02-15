-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: UpdateTokens :one
UPDATE users
SET token = $1, refresh_token = $2
WHERE id = $3
RETURNING *;