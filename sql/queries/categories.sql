-- name: CreateCategory :one
INSERT INTO categories (id, name, user_id)
VALUES ($1, $2, $3)
RETURNING *;
