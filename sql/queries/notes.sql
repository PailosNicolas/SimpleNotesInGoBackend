-- name: CreateNote :one
INSERT INTO notes (id, title, body, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
